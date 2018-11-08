package onchain

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
	"sync"
	"time"

	"github.com/DOSNetwork/core/configuration"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EthCommon struct {
	id       *big.Int
	key      *keystore.Key
	Client   *ethclient.Client
	lock     *sync.Mutex
	ethNonce uint64
	config   *configuration.ChainConfig
}

func (e *EthCommon) DialToEth() (err error) {
	fmt.Println("dialing...")
	e.Client, err = ethclient.Dial(e.config.RemoteNodeAddress)
	for err != nil {
		fmt.Println(err)
		fmt.Println("Cannot connect to the network, retrying...", e.config.RemoteNodeAddress)
		e.Client, err = ethclient.Dial(e.config.RemoteNodeAddress)
	}
	return
}

func (e *EthCommon) Init(credentialPath string, config *configuration.ChainConfig) (err error) {

	e.config = config

	fmt.Println("start initial onChainConn...", config.DOSProxyAddress)

	e.DialToEth()
	e.setAccount(credentialPath)
	return
}

func (e *EthCommon) setAccount(credentialPath string) (err error) {
	fmt.Println("credentialPath: ", credentialPath)

	passPhraseBytes, err := ioutil.ReadFile(credentialPath + "/passPhrase")
	if err != nil {
		return
	}

	passPhrase := string(passPhraseBytes)

	newKeyStore := keystore.NewKeyStore(credentialPath, keystore.StandardScryptN, keystore.StandardScryptP)
	if len(newKeyStore.Accounts()) < 1 {
		_, err = newKeyStore.NewAccount(passPhrase)
		if err != nil {
			return
		}
	}

	usrKeyPath := newKeyStore.Accounts()[0].URL.Path
	usrKeyJson, err := ioutil.ReadFile(usrKeyPath)
	if err != nil {
		return
	}

	usrKey, err := keystore.DecryptKey(usrKeyJson, passPhrase)
	if err != nil {
		return
	}

	e.key = usrKey
	e.id = new(big.Int)
	e.id.SetBytes(e.key.Address.Bytes())
	e.ethNonce, err = e.Client.PendingNonceAt(context.Background(), e.key.Address)
	if err != nil {
		return
	}
	//for correctness of the first call of GetAuth, because GetAuth always ++,
	e.ethNonce--

	return
}

func (e *EthCommon) GetAuth() (auth *bind.TransactOpts, err error) {
	auth = bind.NewKeyedTransactor(e.key.PrivateKey)
	if err != nil {
		return
	}

	automatedNonce, err := e.Client.PendingNonceAt(context.Background(), e.key.Address)
	if err != nil {
		return
	}

	e.ethNonce++
	if automatedNonce > e.ethNonce {
		e.ethNonce = automatedNonce
	}
	auth.Nonce = big.NewInt(int64(e.ethNonce))

	_, err = e.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return
	}

	auth.GasLimit = uint64(6000000)

	return
}

func (e *EthCommon) getKey(keyPath, passPrase string) (key *keystore.Key, err error) {
	var keyJson []byte
	keyJson, err = ioutil.ReadFile(keyPath)
	if err != nil {
		return
	}

	key, err = keystore.DecryptKey(keyJson, passPrase)
	if err != nil {
		return
	}
	return
}

func (e *EthCommon) BalanceMaintain(rootKeyPath, usrKeyPath, rPassPhrase, uPassPhrase string) (err error) {
	fmt.Println("EthCommon BalanceMaintain")

	rootKey, err := e.getKey(rootKeyPath, rPassPhrase)
	if err != nil {
		return
	}
	usrKey, err := e.getKey(usrKeyPath, uPassPhrase)
	if err != nil {
		return
	}
	err = e.balanceMaintain(usrKey, rootKey)
	return
}

func (e *EthCommon) balanceMaintain(usrKey, rootKey *keystore.Key) (err error) {
	usrKeyBalance, err := e.getBalance(usrKey)
	if err != nil {
		return
	}
	fmt.Println("usrKeyBalance ", usrKeyBalance)

	rootKeyBalance, err := e.getBalance(rootKey)
	if err != nil {
		return
	}
	fmt.Println("rootKeyBalance ", rootKeyBalance)

	if usrKeyBalance.Cmp(big.NewFloat(0.7)) == -1 {
		fmt.Println("userKey account replenishing...")
		if err = e.transferEth(rootKey, usrKey); err == nil {
			fmt.Println("userKey account replenished.")
		}
	}

	return
}

func (e *EthCommon) getBalance(key *keystore.Key) (balance *big.Float, err error) {
	wei, err := e.Client.BalanceAt(context.Background(), key.Address, nil)
	if err != nil {
		return
	}

	balance = new(big.Float)
	balance.SetString(wei.String())
	balance = balance.Quo(balance, big.NewFloat(math.Pow10(18)))
	return
}

func (e *EthCommon) transferEth(from, to *keystore.Key) (err error) {
	nonce, err := e.Client.PendingNonceAt(context.Background(), from.Address)
	if err != nil {
		return
	}

	gasPrice, err := e.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return
	}

	value := big.NewInt(800000000000000000) //0.8 Eth
	gasLimit := uint64(1000000)

	tx := types.NewTransaction(nonce, to.Address, value, gasLimit, gasPrice.Mul(gasPrice, big.NewInt(3)), nil)

	chainId, err := e.Client.NetworkID(context.Background())
	if err != nil {
		return
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), from.PrivateKey)
	if err != nil {
		return
	}

	err = e.Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return
	}
	fmt.Println("replenishing tx sent: ", signedTx.Hash().Hex(), ", waiting for confirmation...")

	err = e.CheckTransaction(signedTx)

	return
}

func (e *EthCommon) CheckTransaction(tx *types.Transaction) (err error) {
	//start := time.Now()
	receipt, err := e.Client.TransactionReceipt(context.Background(), tx.Hash())
	for err == ethereum.NotFound {
		/*
			if time.Since(start).Seconds() > 30 {
				fmt.Println("no receipt yet, set to successful")
				return
			}
		*/
		time.Sleep(1 * time.Second)
		receipt, err = e.Client.TransactionReceipt(context.Background(), tx.Hash())
	}
	if err != nil {
		return
	}

	if receipt.Status == 1 {
		fmt.Println("transaction confirmed.")
	} else {
		err = errors.New("transaction failed")
	}

	return
}