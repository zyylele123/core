// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package dosproxy

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// BN256ABI is the input ABI used to generate the binding from.
const BN256ABI = "[]"

// BN256Bin is the compiled bytecode used for deploying new contracts.
const BN256Bin = `0x604c602c600b82828239805160001a60731460008114601c57601e565bfe5b5030600052607381538281f30073000000000000000000000000000000000000000030146080604052600080fd00a165627a7a723058202e1858244ad6c497690860ca32afa6c06d0246a99a70a7a9f2edd5bf67760ae70029`

// DeployBN256 deploys a new Ethereum contract, binding an instance of BN256 to it.
func DeployBN256(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BN256, error) {
	parsed, err := abi.JSON(strings.NewReader(BN256ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BN256Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BN256{BN256Caller: BN256Caller{contract: contract}, BN256Transactor: BN256Transactor{contract: contract}, BN256Filterer: BN256Filterer{contract: contract}}, nil
}

// BN256 is an auto generated Go binding around an Ethereum contract.
type BN256 struct {
	BN256Caller     // Read-only binding to the contract
	BN256Transactor // Write-only binding to the contract
	BN256Filterer   // Log filterer for contract events
}

// BN256Caller is an auto generated read-only Go binding around an Ethereum contract.
type BN256Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BN256Transactor is an auto generated write-only Go binding around an Ethereum contract.
type BN256Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BN256Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BN256Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BN256Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BN256Session struct {
	Contract     *BN256            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BN256CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BN256CallerSession struct {
	Contract *BN256Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BN256TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BN256TransactorSession struct {
	Contract     *BN256Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BN256Raw is an auto generated low-level Go binding around an Ethereum contract.
type BN256Raw struct {
	Contract *BN256 // Generic contract binding to access the raw methods on
}

// BN256CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BN256CallerRaw struct {
	Contract *BN256Caller // Generic read-only contract binding to access the raw methods on
}

// BN256TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BN256TransactorRaw struct {
	Contract *BN256Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBN256 creates a new instance of BN256, bound to a specific deployed contract.
func NewBN256(address common.Address, backend bind.ContractBackend) (*BN256, error) {
	contract, err := bindBN256(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BN256{BN256Caller: BN256Caller{contract: contract}, BN256Transactor: BN256Transactor{contract: contract}, BN256Filterer: BN256Filterer{contract: contract}}, nil
}

// NewBN256Caller creates a new read-only instance of BN256, bound to a specific deployed contract.
func NewBN256Caller(address common.Address, caller bind.ContractCaller) (*BN256Caller, error) {
	contract, err := bindBN256(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BN256Caller{contract: contract}, nil
}

// NewBN256Transactor creates a new write-only instance of BN256, bound to a specific deployed contract.
func NewBN256Transactor(address common.Address, transactor bind.ContractTransactor) (*BN256Transactor, error) {
	contract, err := bindBN256(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BN256Transactor{contract: contract}, nil
}

// NewBN256Filterer creates a new log filterer instance of BN256, bound to a specific deployed contract.
func NewBN256Filterer(address common.Address, filterer bind.ContractFilterer) (*BN256Filterer, error) {
	contract, err := bindBN256(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BN256Filterer{contract: contract}, nil
}

// bindBN256 binds a generic wrapper to an already deployed contract.
func bindBN256(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BN256ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BN256 *BN256Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BN256.Contract.BN256Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BN256 *BN256Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BN256.Contract.BN256Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BN256 *BN256Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BN256.Contract.BN256Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BN256 *BN256CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BN256.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BN256 *BN256TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BN256.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BN256 *BN256TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BN256.Contract.contract.Transact(opts, method, params...)
}

// DOSProxyABI is the input ABI used to generate the binding from.
const DOSProxyABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"sig\",\"type\":\"uint256[2]\"}],\"name\":\"updateRandomness\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"whitelist_inited\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"blkNum\",\"type\":\"uint256\"},{\"name\":\"timeout\",\"type\":\"uint256\"},{\"name\":\"queryType\",\"type\":\"string\"},{\"name\":\"queryPath\",\"type\":\"string\"}],\"name\":\"query\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"x1\",\"type\":\"uint256\"},{\"name\":\"x2\",\"type\":\"uint256\"},{\"name\":\"y1\",\"type\":\"uint256\"},{\"name\":\"y2\",\"type\":\"uint256\"}],\"name\":\"setPublicKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getGroupPubKey\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256[4]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"queryId\",\"type\":\"uint256\"},{\"name\":\"result\",\"type\":\"bytes\"},{\"name\":\"sig\",\"type\":\"uint256[2]\"}],\"name\":\"triggerCallback\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"new_whitelisted_addr\",\"type\":\"address\"}],\"name\":\"resetWhitelistAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"resetContract\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"getWhitelistAddess\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"fireRandom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"size\",\"type\":\"uint256\"}],\"name\":\"grouping\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addresses\",\"type\":\"address[21]\"}],\"name\":\"initWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastRandomness\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"uploadNodeId\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastUpdatedBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"handleTimeout\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"queryId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"url\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"timeout\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"randomness\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"dispatchedGroup\",\"type\":\"uint256[4]\"}],\"name\":\"LogUrl\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"queryType\",\"type\":\"string\"}],\"name\":\"LogNonSupportedType\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"}],\"name\":\"LogNonContractCall\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"callbackAddr\",\"type\":\"address\"}],\"name\":\"LogCallbackTriggeredFor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"LogQueryFromNonExistentUC\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"lastRandomness\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"lastUpdatedBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"dispatchedGroup\",\"type\":\"uint256[4]\"}],\"name\":\"LogUpdateRandom\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"trafficType\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"trafficId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"message\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"signature\",\"type\":\"uint256[2]\"},{\"indexed\":false,\"name\":\"pubKey\",\"type\":\"uint256[4]\"},{\"indexed\":false,\"name\":\"pass\",\"type\":\"bool\"}],\"name\":\"LogValidationResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"LogInsufficientGroupNumber\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"NodeId\",\"type\":\"uint256[]\"}],\"name\":\"LogGrouping\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"x1\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"x2\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"y1\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"y2\",\"type\":\"uint256\"}],\"name\":\"LogPublicKeyAccepted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"previous\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"curr\",\"type\":\"address\"}],\"name\":\"WhitelistAddressReset\",\"type\":\"event\"}]"

// DOSProxyBin is the compiled bytecode used for deploying new contracts.
const DOSProxyBin = `0x60806040526023805460ff1916905534801561001a57600080fd5b506121778061002a6000396000f3006080604052600436106100e55763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166309ac86d381146100ea57806317dfa0d914610103578063482edfaa1461012c57806379a924c51461017f57806392021653146101a05780639261cc7e146101f05780639457d4e814610216578063b181a8fc14610237578063d82a24061461024c578063ea5cba1314610280578063eab14fe614610295578063f13a9626146102ad578063f2a3072d146102f2578063f89a15f714610307578063f90ce5ba1461031f578063fcfafeb614610334575b600080fd5b3480156100f657600080fd5b506101016004610349565b005b34801561010f57600080fd5b5061011861056b565b604080519115158252519081900360200190f35b34801561013857600080fd5b5061016d60048035600160a060020a031690602480359160443591606435808201929081013591608435908101910135610574565b60408051918252519081900360200190f35b34801561018b57600080fd5b50610101600435602435604435606435610951565b3480156101ac57600080fd5b506101b8600435610bf6565b6040518082608080838360005b838110156101dd5781810151838201526020016101c5565b5050505090500191505060405180910390f35b3480156101fc57600080fd5b506101016004803590602480359081019101356044610d03565b34801561022257600080fd5b50610101600160a060020a0360043516610f74565b34801561024357600080fd5b506101016110ab565b34801561025857600080fd5b5061026460043561113f565b60408051600160a060020a039092168252519081900360200190f35b34801561028c57600080fd5b506101016111c5565b3480156102a157600080fd5b50610101600435611348565b3480156102b957600080fd5b50604080516102a08181019092526101019136916004916102a49190839060159083908390808284375093965061150c95505050505050565b3480156102fe57600080fd5b5061016d611613565b34801561031357600080fd5b50610101600435611619565b34801561032b57600080fd5b5061016d6116c7565b34801561034057600080fd5b506101016116cd565b600754604080518082019091526006544081526020810182905260009161040391600191906103779061175e565b604080518082018252873581526020808901359082015281516080810180845291929091600891839190820190839060029082845b8154815260200190600101908083116103ac57505050918352505060408051808201918290526020909201919060028481019182845b8154815260200190600101908083116103e257505050505081525050611798565b151561040e57610567565b6040805183356020828101919091528085013582840152825180830384018152606090920192839052815191929182918401908083835b602083106104645780518252601f199092019160209182019101610445565b5181516020939093036101000a6000190180199091169216919091179052604051920182900390912060078190554360065560035493509150508115156104a757fe5b0690506003818154811015156104b957fe5b6000918252602090912060049091020160086104d781836002611f5f565b506104ea60028281019084810190611f5f565b509050507fb57d39816b5fafa35e9395c99afd19edb72160f491310ba0942ea8c58c8bc9f960075460065461051e84610bf6565b6040518084815260200183815260200182600460200280838360005b8381101561055257818101518382015260200161053a565b50505050905001935050505060405180910390a15b5050565b60235460ff1681565b6000806000806105838b611ae3565b1115610903576105f787878080601f016020809104026020016040519081016040528093929190818152602001838380828437505060408051808201909152600381527f415049000000000000000000000000000000000000000000000000000000000060208201529350611ae792505050565b156108aa57898989898989896040516020018088600160a060020a0316600160a060020a03166c01000000000000000000000000028152601401878152602001868152602001858580828437820191505083838082843782019150509750505050505050506040516020818303038152906040526040518082805190602001908083835b6020831061069a5780518252601f19909201916020918201910161067b565b5181516020939093036101000a6000190180199091169216919091179052604051920182900390912060035460075491965093509150508115156106da57fe5b0690506060604051908101604052808381526020016003838154811015156106fe57fe5b6000918252602090912060408051608081018083529093600402909201918391820190839060029082845b81548152602001906001019080831161072957505050918352505060408051808201918290526020909201919060028481019182845b81548152602001906001019080831161075f57505050919092525050508152600160a060020a038c1660209182015260008481526002808352604090912083518155918301518051909160018401916107ba91839190611f9a565b5060208201516107d09060028084019190611f9a565b50505060408201518160050160006101000a815481600160a060020a030219169083600160a060020a031602179055509050507f75b4fd18552b0d79989d6cc5de124544de30dc29b548e1f743a719ce5ad443298286868b60075461083487610bf6565b604051808781526020018060200185815260200184815260200183600460200280838360005b8381101561087257818101518382015260200161085a565b5050505090500182810382528787828181526020019250808284376040519201829003995090975050505050505050a1819250610944565b7f70714cf695ae953ee67221716a4b4dc9e944909fd2b66f07e790a49d9ac29b4187876040518080602001828103825284848281815260200192508082843760405192018290039550909350505050a160009250610944565b60408051600160a060020a038c16815290517f6cea43bb3db7220931a7c8ac633e65cbc8e7ba129f2ed84db2e71bc0adb73bb59181900360200190a1600092505b5050979650505050505050565b336000908152602260205260408120548015801590610986575033600c826016811061097957fe5b0154600160a060020a0316145b15156109ca576040805160e560020a62461bcd028152602060048201526010602482015260008051602061212c833981519152604482015290519081900360640190fd5b85858585604051602001808581526020018481526020018381526020018281526020019450505050506040516020818303038152906040526040518082805190602001908083835b60208310610a315780518252601f199092019160209182019101610a12565b51815160209384036101000a6000190180199092169116179052604080519290940182900390912060008181526004909252929020549195505060ff16159150610ac79050576040805160e560020a62461bcd02815260206004820152601c60248201527f67726f75702068617320616c7265616479207265676973746572656400000000604482015290519081900360640190fd5b60008281526005602052604081208054600101908190559054600290041015610bee5760408051608081018252808201888152606082018890528152815180830190925285825260208281018690528101919091526003805460018101808355600092909252825191929160049091027fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0190610b679082906002611f9a565b506020820151610b7d9060028084019190611f9a565b5050506000838152600460209081526040808320805460ff19166001179055600582528083209290925581518981529081018890528082018790526060810186905290517f5bf077aca7c3c2ae9ca6ebd7da84490edf523227621a171c3c2656b3d45e92fe92509081900360800190a15b505050505050565b610bfe611fc8565b6003548210610c57576040805160e560020a62461bcd02815260206004820152601860248201527f67726f757020696e646578206f7574206f662072616e67650000000000000000604482015290519081900360640190fd5b608060405190810160405280600384815481101515610c7257fe5b60009182526020822060049091020101548152602001600384815481101515610c9757fe5b60009182526020909120600490910201600101548152602001600384815481101515610cbf57fe5b60009182526020822060026004909202010101548152602001600384815481101515610ce757fe5b6000918252602090912060036004909202010154905292915050565b6000610deb60008686868080601f016020809104026020016040519081016040528093929190818152602001838380828437505060408051808201909152935083925089915060009050602090810291909101358252018760016020908102919091013590915260008b81526002918290526040908190208151608081018084526001830180549483019485529194919385939092859290918201606086018083116103ac575050509183525050604080518082019182905260028481018054835260209485019492939092600387019085018083116103e257505050505081525050611798565b1515610df657610f6d565b50600084815260026020526040902060050154600160a060020a0316801515610e47576040517f158bff16635ac24f3d1acce162f0626cc6751bd434047538d76421366edf590690600090a1610f6d565b60408051600160a060020a038316815290517f065d5d7c942a87321bf774d2780cfd4928766b9f04dc6728ab8b4490ef5edaf09181900360200190a1604080517f6d112977000000000000000000000000000000000000000000000000000000008152600481018781526024820192835260448201869052600160a060020a03841692636d112977928992899289926064018484808284378201915050945050505050600060405180830381600087803b158015610f0457600080fd5b505af1158015610f18573d6000803e3d6000fd5b505050600086815260026020526040812081815591506001820181610f3d8282611fe7565b610f4b600283016000611fe7565b505050600501805473ffffffffffffffffffffffffffffffffffffffff191690555b5050505050565b336000908152602260205260409020548015801590610fa9575033600c8260168110610f9c57fe5b0154600160a060020a0316145b1515610fed576040805160e560020a62461bcd028152602060048201526010602482015260008051602061212c833981519152604482015290519081900360640190fd5b600160a060020a0382161580159061100e5750600160a060020a0382163314155b151561101957600080fd5b60408051338152600160a060020a038416602082015281517f81212bcc8da3221a71cc77dc44d159e2a3aeeb6e1aaa17b1fffa5938540ab842929181900390910190a1336000908152602260205260409020548290600c906016811061107b57fe5b01805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03929092169190911790555050565b3360009081526022602052604090205480158015906110e0575033600c82601681106110d357fe5b0154600160a060020a0316145b1515611124576040805160e560020a62461bcd028152602060048201526010602482015260008051602061212c833981519152604482015290519081900360640190fd5b6000611131600182611ff5565b50600061056760038261201e565b60008082118015611151575060158211155b15156111a7576040805160e560020a62461bcd02815260206004820152601260248201527f496e646578206f7574206f662072616e67650000000000000000000000000000604482015290519081900360640190fd5b600c82601681106111b457fe5b0154600160a060020a031692915050565b3360009081526022602052604081205480158015906111fa575033600c82601681106111ed57fe5b0154600160a060020a0316145b151561123e576040805160e560020a62461bcd028152602060048201526010602482015260008051602061212c833981519152604482015290519081900360640190fd5b6040805160001943014060208083019190915282518083038201815291830192839052815191929182918401908083835b6020831061128e5780518252601f19909201916020918201910161126f565b5181516020939093036101000a6000190180199091169216919091179052604051920182900390912060078190554360065560035493509150508115156112d157fe5b0691506003828154811015156112e357fe5b60009182526020909120600490910201600861130181836002611f5f565b5061131460028281019084810190611f5f565b509050507fb57d39816b5fafa35e9395c99afd19edb72160f491310ba0942ea8c58c8bc9f960075460065461051e85610bf6565b33600090815260226020526040812054606091908015801590611381575033600c826016811061137457fe5b0154600160a060020a0316145b15156113c5576040805160e560020a62461bcd028152602060048201526010602482015260008051602061212c833981519152604482015290519081900360640190fd5b60008490556040805185815260208087028201019091528480156113f3578160200160208202803883390190505b5060015490935084111561142f576040517f08a70ba288e836bee6c9b4aea7482ee5ff8f63c5ad9d2533d9cf0ced64adc26290600090a1611506565b600091505b838210156114905760018054600019810190811061144e57fe5b9060005260206000200154838381518110151561146757fe5b602090810290910101526001805490611484906000198301611ff5565b50600190910190611434565b7f5f30b698cceb472bcb5a80c4acc8c52ea45ea704f5aeeb2527d2d4c95f793dd7836040518080602001828103825283818151815260200191508051906020019060200280838360005b838110156114f25781810151838201526020016114da565b505050509050019250505060405180910390a15b50505050565b60235460009060ff161561156a576040805160e560020a62461bcd02815260206004820152601e60248201527f57686974656c69737420616c726561647920696e697469616c697a6564210000604482015290519081900360640190fd5b5060005b60158110156116025781816015811061158357fe5b6020020151600c600183016016811061159857fe5b01805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039290921691909117905560018101602260008484601581106115d857fe5b60209081029190910151600160a060020a031682528101919091526040016000205560010161156e565b50506023805460ff19166001179055565b60075481565b33600090815260226020526040902054801580159061164e575033600c826016811061164157fe5b0154600160a060020a0316145b1515611692576040805160e560020a62461bcd028152602060048201526010602482015260008051602061212c833981519152604482015290519081900360640190fd5b506001805480820182556000919091527fb10e2d527612073b26eecdfd717e6a320cf44b4afac2b0732d9fcbe2b7fa0cf60155565b60065481565b336000908152602260205260408120548015801590611702575033600c82601681106116f557fe5b0154600160a060020a0316145b1515611746576040805160e560020a62461bcd028152602060048201526010602482015260008051602061212c833981519152604482015290519081900360640190fd5b439150600460065483031115610567576105676111c5565b6040805181815260608082018352918291906020820161080080388339505084516020808401919091529094015160408201529392505050565b3360009081526022602052604081205460609081908190849080158015906117d6575033600c82601681106117c957fe5b0154600160a060020a0316145b151561181a576040805160e560020a62461bcd028152602060048201526010602482015260008051602061212c833981519152604482015290519081900360640190fd5b8833600160a060020a03166001026040516020018083805190602001908083835b6020831061185a5780518252601f19909201916020918201910161183b565b51815160209384036101000a60001901801990921691161790529201938452506040805180850381526002928501838152608086018352909a50945090920190505b6118a461204a565b81526020019060019003908161189c57505060408051600280825260608201909252919550602082015b6118d6612061565b8152602001906001900390816118ce579050509250878460008151811015156118fb57fe5b6020908102909101015261190e85611bb6565b84600181518110151561191d57fe5b6020908102909101015261192f611c3b565b83600081518110151561193e57fe5b60209081029091010152825187908490600190811061195957fe5b6020908102909101015261196d8484611cfc565b6040805180820182528a5181526020808c01518183015282516080810184528b515181528b5182015181830152908b018051519382019390935291519294507fd33c44f7ce166bcd2616c8f4d811261d4d24e1af815f78614683e0c7682c42a5928e928e928a9290919060608201906001602002015181525087604051808760ff1660ff1681526020018681526020018060200185600260200280838360005b83811015611a25578181015183820152602001611a0d565b5050505090500184600460200280838360005b83811015611a50578181015183820152602001611a38565b5050505090500183151515158152602001828103825286818151815260200191508051906020019080838360005b83811015611a96578181015183820152602001611a7e565b50505050905090810190601f168015611ac35780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390a1509998505050505050505050565b3b90565b8051825160009184918491849114611b025760009350611bad565b5060005b8251811015611ba8578181815181101515611b1d57fe5b90602001015160f860020a900460f860020a027effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168382815181101515611b6057fe5b60209101015160f860020a90819004027fff000000000000000000000000000000000000000000000000000000000000001614611ba05760009350611bad565b600101611b06565b600193505b50505092915050565b611bbe61204a565b6000826040518082805190602001908083835b60208310611bf05780518252601f199092019160209182019101611bd1565b6001836020036101000a0380198251168184511680821785525050505050509050019150506040518091039020600190049050611c34611c2e611ef8565b82611f19565b9392505050565b611c43612061565b50604080516080810182527f198e9393920d483a7260bfb731fb5d25f1aa493335a9e71297e485b7aef312c28183019081527f1800deef121f1e76426a00665e5c4479674322d4f75edadd46debd5cd992f6ed6060830152815281518083019092527f090689d0585ff075ec9e99ad690c3395bc4b313370b38ef355acdadcd122975b82527f12c85ea5db8c6deb4aab71808dcb408fe3d1e7690c43d37b4ce6cc0166fa7daa6020838101919091528101919091525b90565b600080600060606000611d0d612087565b8651885160009114611d1e57600080fd5b8851955085600602945084604051908082528060200260200182016040528015611d52578160200160208202803883390190505b509350600092505b85831015611ec7578883815181101515611d7057fe5b60209081029091010151518451859060068602908110611d8c57fe5b602090810290910101528851899084908110611da457fe5b90602001906020020151602001518484600602600101815181101515611dc657fe5b602090810290910101528751889084908110611dde57fe5b602090810291909101015151518451859060026006870201908110611dff57fe5b602090810290910101528751889084908110611e1757fe5b6020908102919091018101515101518451859060036006870201908110611e3a57fe5b602090810290910101528751889084908110611e5257fe5b6020908102919091018101510151518451859060046006870201908110611e7557fe5b602090810290910101528751889084908110611e8d57fe5b60209081029190910181015181015101518451859060056006870201908110611eb257fe5b60209081029091010152600190920191611d5a565b6020826020870260208701600060086107d05a03f19050808015611eeb5750815115155b9998505050505050505050565b611f0061204a565b5060408051808201909152600181526002602082015290565b611f2161204a565b611f296120a6565b8351815260208085015190820152604080820184905282606083600060076107d05a03f11515611f5857600080fd5b5092915050565b8260028101928215611f8a579182015b82811115611f8a578254825591600101919060010190611f6f565b50611f969291506120c5565b5090565b8260028101928215611f8a579160200282015b82811115611f8a578251825591602001919060010190611fad565b6080604051908101604052806004906020820280388339509192915050565b506000815560010160009055565b815481835581811115612019576000838152602090206120199181019083016120c5565b505050565b8154818355818111156120195760040281600402836000526020600020918201910161201991906120df565b604080518082019091526000808252602082015290565b608060405190810160405280612075612110565b8152602001612082612110565b905290565b6020604051908101604052806001906020820280388339509192915050565b6060604051908101604052806003906020820280388339509192915050565b611cf991905b80821115611f9657600081556001016120cb565b611cf991905b80821115611f965760006120f98282611fe7565b612107600283016000611fe7565b506004016120e5565b6040805180820182529060029082908038833950919291505056004e6f742077686974656c69737465642100000000000000000000000000000000a165627a7a72305820e2df2fee7263f812f17470f88853ec1fb730a197accb135550c8c40da37b52af0029`

// DeployDOSProxy deploys a new Ethereum contract, binding an instance of DOSProxy to it.
func DeployDOSProxy(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DOSProxy, error) {
	parsed, err := abi.JSON(strings.NewReader(DOSProxyABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(DOSProxyBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DOSProxy{DOSProxyCaller: DOSProxyCaller{contract: contract}, DOSProxyTransactor: DOSProxyTransactor{contract: contract}, DOSProxyFilterer: DOSProxyFilterer{contract: contract}}, nil
}

// DOSProxy is an auto generated Go binding around an Ethereum contract.
type DOSProxy struct {
	DOSProxyCaller     // Read-only binding to the contract
	DOSProxyTransactor // Write-only binding to the contract
	DOSProxyFilterer   // Log filterer for contract events
}

// DOSProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type DOSProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DOSProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DOSProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DOSProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DOSProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DOSProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DOSProxySession struct {
	Contract     *DOSProxy         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DOSProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DOSProxyCallerSession struct {
	Contract *DOSProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// DOSProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DOSProxyTransactorSession struct {
	Contract     *DOSProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DOSProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type DOSProxyRaw struct {
	Contract *DOSProxy // Generic contract binding to access the raw methods on
}

// DOSProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DOSProxyCallerRaw struct {
	Contract *DOSProxyCaller // Generic read-only contract binding to access the raw methods on
}

// DOSProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DOSProxyTransactorRaw struct {
	Contract *DOSProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDOSProxy creates a new instance of DOSProxy, bound to a specific deployed contract.
func NewDOSProxy(address common.Address, backend bind.ContractBackend) (*DOSProxy, error) {
	contract, err := bindDOSProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DOSProxy{DOSProxyCaller: DOSProxyCaller{contract: contract}, DOSProxyTransactor: DOSProxyTransactor{contract: contract}, DOSProxyFilterer: DOSProxyFilterer{contract: contract}}, nil
}

// NewDOSProxyCaller creates a new read-only instance of DOSProxy, bound to a specific deployed contract.
func NewDOSProxyCaller(address common.Address, caller bind.ContractCaller) (*DOSProxyCaller, error) {
	contract, err := bindDOSProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DOSProxyCaller{contract: contract}, nil
}

// NewDOSProxyTransactor creates a new write-only instance of DOSProxy, bound to a specific deployed contract.
func NewDOSProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*DOSProxyTransactor, error) {
	contract, err := bindDOSProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DOSProxyTransactor{contract: contract}, nil
}

// NewDOSProxyFilterer creates a new log filterer instance of DOSProxy, bound to a specific deployed contract.
func NewDOSProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*DOSProxyFilterer, error) {
	contract, err := bindDOSProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DOSProxyFilterer{contract: contract}, nil
}

// bindDOSProxy binds a generic wrapper to an already deployed contract.
func bindDOSProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DOSProxyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DOSProxy *DOSProxyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DOSProxy.Contract.DOSProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DOSProxy *DOSProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DOSProxy.Contract.DOSProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DOSProxy *DOSProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DOSProxy.Contract.DOSProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DOSProxy *DOSProxyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _DOSProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DOSProxy *DOSProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DOSProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DOSProxy *DOSProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DOSProxy.Contract.contract.Transact(opts, method, params...)
}

// GetGroupPubKey is a free data retrieval call binding the contract method 0x92021653.
//
// Solidity: function getGroupPubKey(idx uint256) constant returns(uint256[4])
func (_DOSProxy *DOSProxyCaller) GetGroupPubKey(opts *bind.CallOpts, idx *big.Int) ([4]*big.Int, error) {
	var (
		ret0 = new([4]*big.Int)
	)
	out := ret0
	err := _DOSProxy.contract.Call(opts, out, "getGroupPubKey", idx)
	return *ret0, err
}

// GetGroupPubKey is a free data retrieval call binding the contract method 0x92021653.
//
// Solidity: function getGroupPubKey(idx uint256) constant returns(uint256[4])
func (_DOSProxy *DOSProxySession) GetGroupPubKey(idx *big.Int) ([4]*big.Int, error) {
	return _DOSProxy.Contract.GetGroupPubKey(&_DOSProxy.CallOpts, idx)
}

// GetGroupPubKey is a free data retrieval call binding the contract method 0x92021653.
//
// Solidity: function getGroupPubKey(idx uint256) constant returns(uint256[4])
func (_DOSProxy *DOSProxyCallerSession) GetGroupPubKey(idx *big.Int) ([4]*big.Int, error) {
	return _DOSProxy.Contract.GetGroupPubKey(&_DOSProxy.CallOpts, idx)
}

// GetWhitelistAddess is a free data retrieval call binding the contract method 0xd82a2406.
//
// Solidity: function getWhitelistAddess(idx uint256) constant returns(address)
func (_DOSProxy *DOSProxyCaller) GetWhitelistAddess(opts *bind.CallOpts, idx *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _DOSProxy.contract.Call(opts, out, "getWhitelistAddess", idx)
	return *ret0, err
}

// GetWhitelistAddess is a free data retrieval call binding the contract method 0xd82a2406.
//
// Solidity: function getWhitelistAddess(idx uint256) constant returns(address)
func (_DOSProxy *DOSProxySession) GetWhitelistAddess(idx *big.Int) (common.Address, error) {
	return _DOSProxy.Contract.GetWhitelistAddess(&_DOSProxy.CallOpts, idx)
}

// GetWhitelistAddess is a free data retrieval call binding the contract method 0xd82a2406.
//
// Solidity: function getWhitelistAddess(idx uint256) constant returns(address)
func (_DOSProxy *DOSProxyCallerSession) GetWhitelistAddess(idx *big.Int) (common.Address, error) {
	return _DOSProxy.Contract.GetWhitelistAddess(&_DOSProxy.CallOpts, idx)
}

// LastRandomness is a free data retrieval call binding the contract method 0xf2a3072d.
//
// Solidity: function lastRandomness() constant returns(uint256)
func (_DOSProxy *DOSProxyCaller) LastRandomness(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DOSProxy.contract.Call(opts, out, "lastRandomness")
	return *ret0, err
}

// LastRandomness is a free data retrieval call binding the contract method 0xf2a3072d.
//
// Solidity: function lastRandomness() constant returns(uint256)
func (_DOSProxy *DOSProxySession) LastRandomness() (*big.Int, error) {
	return _DOSProxy.Contract.LastRandomness(&_DOSProxy.CallOpts)
}

// LastRandomness is a free data retrieval call binding the contract method 0xf2a3072d.
//
// Solidity: function lastRandomness() constant returns(uint256)
func (_DOSProxy *DOSProxyCallerSession) LastRandomness() (*big.Int, error) {
	return _DOSProxy.Contract.LastRandomness(&_DOSProxy.CallOpts)
}

// LastUpdatedBlock is a free data retrieval call binding the contract method 0xf90ce5ba.
//
// Solidity: function lastUpdatedBlock() constant returns(uint256)
func (_DOSProxy *DOSProxyCaller) LastUpdatedBlock(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _DOSProxy.contract.Call(opts, out, "lastUpdatedBlock")
	return *ret0, err
}

// LastUpdatedBlock is a free data retrieval call binding the contract method 0xf90ce5ba.
//
// Solidity: function lastUpdatedBlock() constant returns(uint256)
func (_DOSProxy *DOSProxySession) LastUpdatedBlock() (*big.Int, error) {
	return _DOSProxy.Contract.LastUpdatedBlock(&_DOSProxy.CallOpts)
}

// LastUpdatedBlock is a free data retrieval call binding the contract method 0xf90ce5ba.
//
// Solidity: function lastUpdatedBlock() constant returns(uint256)
func (_DOSProxy *DOSProxyCallerSession) LastUpdatedBlock() (*big.Int, error) {
	return _DOSProxy.Contract.LastUpdatedBlock(&_DOSProxy.CallOpts)
}

// WhitelistInited is a free data retrieval call binding the contract method 0x17dfa0d9.
//
// Solidity: function whitelist_inited() constant returns(bool)
func (_DOSProxy *DOSProxyCaller) WhitelistInited(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _DOSProxy.contract.Call(opts, out, "whitelist_inited")
	return *ret0, err
}

// WhitelistInited is a free data retrieval call binding the contract method 0x17dfa0d9.
//
// Solidity: function whitelist_inited() constant returns(bool)
func (_DOSProxy *DOSProxySession) WhitelistInited() (bool, error) {
	return _DOSProxy.Contract.WhitelistInited(&_DOSProxy.CallOpts)
}

// WhitelistInited is a free data retrieval call binding the contract method 0x17dfa0d9.
//
// Solidity: function whitelist_inited() constant returns(bool)
func (_DOSProxy *DOSProxyCallerSession) WhitelistInited() (bool, error) {
	return _DOSProxy.Contract.WhitelistInited(&_DOSProxy.CallOpts)
}

// FireRandom is a paid mutator transaction binding the contract method 0xea5cba13.
//
// Solidity: function fireRandom() returns()
func (_DOSProxy *DOSProxyTransactor) FireRandom(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DOSProxy.contract.Transact(opts, "fireRandom")
}

// FireRandom is a paid mutator transaction binding the contract method 0xea5cba13.
//
// Solidity: function fireRandom() returns()
func (_DOSProxy *DOSProxySession) FireRandom() (*types.Transaction, error) {
	return _DOSProxy.Contract.FireRandom(&_DOSProxy.TransactOpts)
}

// FireRandom is a paid mutator transaction binding the contract method 0xea5cba13.
//
// Solidity: function fireRandom() returns()
func (_DOSProxy *DOSProxyTransactorSession) FireRandom() (*types.Transaction, error) {
	return _DOSProxy.Contract.FireRandom(&_DOSProxy.TransactOpts)
}

// Grouping is a paid mutator transaction binding the contract method 0xeab14fe6.
//
// Solidity: function grouping(size uint256) returns()
func (_DOSProxy *DOSProxyTransactor) Grouping(opts *bind.TransactOpts, size *big.Int) (*types.Transaction, error) {
	return _DOSProxy.contract.Transact(opts, "grouping", size)
}

// Grouping is a paid mutator transaction binding the contract method 0xeab14fe6.
//
// Solidity: function grouping(size uint256) returns()
func (_DOSProxy *DOSProxySession) Grouping(size *big.Int) (*types.Transaction, error) {
	return _DOSProxy.Contract.Grouping(&_DOSProxy.TransactOpts, size)
}

// Grouping is a paid mutator transaction binding the contract method 0xeab14fe6.
//
// Solidity: function grouping(size uint256) returns()
func (_DOSProxy *DOSProxyTransactorSession) Grouping(size *big.Int) (*types.Transaction, error) {
	return _DOSProxy.Contract.Grouping(&_DOSProxy.TransactOpts, size)
}

// HandleTimeout is a paid mutator transaction binding the contract method 0xfcfafeb6.
//
// Solidity: function handleTimeout() returns()
func (_DOSProxy *DOSProxyTransactor) HandleTimeout(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DOSProxy.contract.Transact(opts, "handleTimeout")
}

// HandleTimeout is a paid mutator transaction binding the contract method 0xfcfafeb6.
//
// Solidity: function handleTimeout() returns()
func (_DOSProxy *DOSProxySession) HandleTimeout() (*types.Transaction, error) {
	return _DOSProxy.Contract.HandleTimeout(&_DOSProxy.TransactOpts)
}

// HandleTimeout is a paid mutator transaction binding the contract method 0xfcfafeb6.
//
// Solidity: function handleTimeout() returns()
func (_DOSProxy *DOSProxyTransactorSession) HandleTimeout() (*types.Transaction, error) {
	return _DOSProxy.Contract.HandleTimeout(&_DOSProxy.TransactOpts)
}

// InitWhitelist is a paid mutator transaction binding the contract method 0xf13a9626.
//
// Solidity: function initWhitelist(addresses address[21]) returns()
func (_DOSProxy *DOSProxyTransactor) InitWhitelist(opts *bind.TransactOpts, addresses [21]common.Address) (*types.Transaction, error) {
	return _DOSProxy.contract.Transact(opts, "initWhitelist", addresses)
}

// InitWhitelist is a paid mutator transaction binding the contract method 0xf13a9626.
//
// Solidity: function initWhitelist(addresses address[21]) returns()
func (_DOSProxy *DOSProxySession) InitWhitelist(addresses [21]common.Address) (*types.Transaction, error) {
	return _DOSProxy.Contract.InitWhitelist(&_DOSProxy.TransactOpts, addresses)
}

// InitWhitelist is a paid mutator transaction binding the contract method 0xf13a9626.
//
// Solidity: function initWhitelist(addresses address[21]) returns()
func (_DOSProxy *DOSProxyTransactorSession) InitWhitelist(addresses [21]common.Address) (*types.Transaction, error) {
	return _DOSProxy.Contract.InitWhitelist(&_DOSProxy.TransactOpts, addresses)
}

// Query is a paid mutator transaction binding the contract method 0x482edfaa.
//
// Solidity: function query(from address, blkNum uint256, timeout uint256, queryType string, queryPath string) returns(uint256)
func (_DOSProxy *DOSProxyTransactor) Query(opts *bind.TransactOpts, from common.Address, blkNum *big.Int, timeout *big.Int, queryType string, queryPath string) (*types.Transaction, error) {
	return _DOSProxy.contract.Transact(opts, "query", from, blkNum, timeout, queryType, queryPath)
}

// Query is a paid mutator transaction binding the contract method 0x482edfaa.
//
// Solidity: function query(from address, blkNum uint256, timeout uint256, queryType string, queryPath string) returns(uint256)
func (_DOSProxy *DOSProxySession) Query(from common.Address, blkNum *big.Int, timeout *big.Int, queryType string, queryPath string) (*types.Transaction, error) {
	return _DOSProxy.Contract.Query(&_DOSProxy.TransactOpts, from, blkNum, timeout, queryType, queryPath)
}

// Query is a paid mutator transaction binding the contract method 0x482edfaa.
//
// Solidity: function query(from address, blkNum uint256, timeout uint256, queryType string, queryPath string) returns(uint256)
func (_DOSProxy *DOSProxyTransactorSession) Query(from common.Address, blkNum *big.Int, timeout *big.Int, queryType string, queryPath string) (*types.Transaction, error) {
	return _DOSProxy.Contract.Query(&_DOSProxy.TransactOpts, from, blkNum, timeout, queryType, queryPath)
}

// ResetContract is a paid mutator transaction binding the contract method 0xb181a8fc.
//
// Solidity: function resetContract() returns()
func (_DOSProxy *DOSProxyTransactor) ResetContract(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DOSProxy.contract.Transact(opts, "resetContract")
}

// ResetContract is a paid mutator transaction binding the contract method 0xb181a8fc.
//
// Solidity: function resetContract() returns()
func (_DOSProxy *DOSProxySession) ResetContract() (*types.Transaction, error) {
	return _DOSProxy.Contract.ResetContract(&_DOSProxy.TransactOpts)
}

// ResetContract is a paid mutator transaction binding the contract method 0xb181a8fc.
//
// Solidity: function resetContract() returns()
func (_DOSProxy *DOSProxyTransactorSession) ResetContract() (*types.Transaction, error) {
	return _DOSProxy.Contract.ResetContract(&_DOSProxy.TransactOpts)
}

// ResetWhitelistAddress is a paid mutator transaction binding the contract method 0x9457d4e8.
//
// Solidity: function resetWhitelistAddress(new_whitelisted_addr address) returns()
func (_DOSProxy *DOSProxyTransactor) ResetWhitelistAddress(opts *bind.TransactOpts, new_whitelisted_addr common.Address) (*types.Transaction, error) {
	return _DOSProxy.contract.Transact(opts, "resetWhitelistAddress", new_whitelisted_addr)
}

// ResetWhitelistAddress is a paid mutator transaction binding the contract method 0x9457d4e8.
//
// Solidity: function resetWhitelistAddress(new_whitelisted_addr address) returns()
func (_DOSProxy *DOSProxySession) ResetWhitelistAddress(new_whitelisted_addr common.Address) (*types.Transaction, error) {
	return _DOSProxy.Contract.ResetWhitelistAddress(&_DOSProxy.TransactOpts, new_whitelisted_addr)
}

// ResetWhitelistAddress is a paid mutator transaction binding the contract method 0x9457d4e8.
//
// Solidity: function resetWhitelistAddress(new_whitelisted_addr address) returns()
func (_DOSProxy *DOSProxyTransactorSession) ResetWhitelistAddress(new_whitelisted_addr common.Address) (*types.Transaction, error) {
	return _DOSProxy.Contract.ResetWhitelistAddress(&_DOSProxy.TransactOpts, new_whitelisted_addr)
}

// SetPublicKey is a paid mutator transaction binding the contract method 0x79a924c5.
//
// Solidity: function setPublicKey(x1 uint256, x2 uint256, y1 uint256, y2 uint256) returns()
func (_DOSProxy *DOSProxyTransactor) SetPublicKey(opts *bind.TransactOpts, x1 *big.Int, x2 *big.Int, y1 *big.Int, y2 *big.Int) (*types.Transaction, error) {
	return _DOSProxy.contract.Transact(opts, "setPublicKey", x1, x2, y1, y2)
}

// SetPublicKey is a paid mutator transaction binding the contract method 0x79a924c5.
//
// Solidity: function setPublicKey(x1 uint256, x2 uint256, y1 uint256, y2 uint256) returns()
func (_DOSProxy *DOSProxySession) SetPublicKey(x1 *big.Int, x2 *big.Int, y1 *big.Int, y2 *big.Int) (*types.Transaction, error) {
	return _DOSProxy.Contract.SetPublicKey(&_DOSProxy.TransactOpts, x1, x2, y1, y2)
}

// SetPublicKey is a paid mutator transaction binding the contract method 0x79a924c5.
//
// Solidity: function setPublicKey(x1 uint256, x2 uint256, y1 uint256, y2 uint256) returns()
func (_DOSProxy *DOSProxyTransactorSession) SetPublicKey(x1 *big.Int, x2 *big.Int, y1 *big.Int, y2 *big.Int) (*types.Transaction, error) {
	return _DOSProxy.Contract.SetPublicKey(&_DOSProxy.TransactOpts, x1, x2, y1, y2)
}

// TriggerCallback is a paid mutator transaction binding the contract method 0x9261cc7e.
//
// Solidity: function triggerCallback(queryId uint256, result bytes, sig uint256[2]) returns()
func (_DOSProxy *DOSProxyTransactor) TriggerCallback(opts *bind.TransactOpts, queryId *big.Int, result []byte, sig [2]*big.Int) (*types.Transaction, error) {
	return _DOSProxy.contract.Transact(opts, "triggerCallback", queryId, result, sig)
}

// TriggerCallback is a paid mutator transaction binding the contract method 0x9261cc7e.
//
// Solidity: function triggerCallback(queryId uint256, result bytes, sig uint256[2]) returns()
func (_DOSProxy *DOSProxySession) TriggerCallback(queryId *big.Int, result []byte, sig [2]*big.Int) (*types.Transaction, error) {
	return _DOSProxy.Contract.TriggerCallback(&_DOSProxy.TransactOpts, queryId, result, sig)
}

// TriggerCallback is a paid mutator transaction binding the contract method 0x9261cc7e.
//
// Solidity: function triggerCallback(queryId uint256, result bytes, sig uint256[2]) returns()
func (_DOSProxy *DOSProxyTransactorSession) TriggerCallback(queryId *big.Int, result []byte, sig [2]*big.Int) (*types.Transaction, error) {
	return _DOSProxy.Contract.TriggerCallback(&_DOSProxy.TransactOpts, queryId, result, sig)
}

// UpdateRandomness is a paid mutator transaction binding the contract method 0x09ac86d3.
//
// Solidity: function updateRandomness(sig uint256[2]) returns()
func (_DOSProxy *DOSProxyTransactor) UpdateRandomness(opts *bind.TransactOpts, sig [2]*big.Int) (*types.Transaction, error) {
	return _DOSProxy.contract.Transact(opts, "updateRandomness", sig)
}

// UpdateRandomness is a paid mutator transaction binding the contract method 0x09ac86d3.
//
// Solidity: function updateRandomness(sig uint256[2]) returns()
func (_DOSProxy *DOSProxySession) UpdateRandomness(sig [2]*big.Int) (*types.Transaction, error) {
	return _DOSProxy.Contract.UpdateRandomness(&_DOSProxy.TransactOpts, sig)
}

// UpdateRandomness is a paid mutator transaction binding the contract method 0x09ac86d3.
//
// Solidity: function updateRandomness(sig uint256[2]) returns()
func (_DOSProxy *DOSProxyTransactorSession) UpdateRandomness(sig [2]*big.Int) (*types.Transaction, error) {
	return _DOSProxy.Contract.UpdateRandomness(&_DOSProxy.TransactOpts, sig)
}

// UploadNodeId is a paid mutator transaction binding the contract method 0xf89a15f7.
//
// Solidity: function uploadNodeId(id uint256) returns()
func (_DOSProxy *DOSProxyTransactor) UploadNodeId(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _DOSProxy.contract.Transact(opts, "uploadNodeId", id)
}

// UploadNodeId is a paid mutator transaction binding the contract method 0xf89a15f7.
//
// Solidity: function uploadNodeId(id uint256) returns()
func (_DOSProxy *DOSProxySession) UploadNodeId(id *big.Int) (*types.Transaction, error) {
	return _DOSProxy.Contract.UploadNodeId(&_DOSProxy.TransactOpts, id)
}

// UploadNodeId is a paid mutator transaction binding the contract method 0xf89a15f7.
//
// Solidity: function uploadNodeId(id uint256) returns()
func (_DOSProxy *DOSProxyTransactorSession) UploadNodeId(id *big.Int) (*types.Transaction, error) {
	return _DOSProxy.Contract.UploadNodeId(&_DOSProxy.TransactOpts, id)
}

// DOSProxyLogCallbackTriggeredForIterator is returned from FilterLogCallbackTriggeredFor and is used to iterate over the raw logs and unpacked data for LogCallbackTriggeredFor events raised by the DOSProxy contract.
type DOSProxyLogCallbackTriggeredForIterator struct {
	Event *DOSProxyLogCallbackTriggeredFor // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DOSProxyLogCallbackTriggeredForIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DOSProxyLogCallbackTriggeredFor)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DOSProxyLogCallbackTriggeredFor)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DOSProxyLogCallbackTriggeredForIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DOSProxyLogCallbackTriggeredForIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DOSProxyLogCallbackTriggeredFor represents a LogCallbackTriggeredFor event raised by the DOSProxy contract.
type DOSProxyLogCallbackTriggeredFor struct {
	CallbackAddr common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterLogCallbackTriggeredFor is a free log retrieval operation binding the contract event 0x065d5d7c942a87321bf774d2780cfd4928766b9f04dc6728ab8b4490ef5edaf0.
//
// Solidity: e LogCallbackTriggeredFor(callbackAddr address)
func (_DOSProxy *DOSProxyFilterer) FilterLogCallbackTriggeredFor(opts *bind.FilterOpts) (*DOSProxyLogCallbackTriggeredForIterator, error) {

	logs, sub, err := _DOSProxy.contract.FilterLogs(opts, "LogCallbackTriggeredFor")
	if err != nil {
		return nil, err
	}
	return &DOSProxyLogCallbackTriggeredForIterator{contract: _DOSProxy.contract, event: "LogCallbackTriggeredFor", logs: logs, sub: sub}, nil
}

// WatchLogCallbackTriggeredFor is a free log subscription operation binding the contract event 0x065d5d7c942a87321bf774d2780cfd4928766b9f04dc6728ab8b4490ef5edaf0.
//
// Solidity: e LogCallbackTriggeredFor(callbackAddr address)
func (_DOSProxy *DOSProxyFilterer) WatchLogCallbackTriggeredFor(opts *bind.WatchOpts, sink chan<- *DOSProxyLogCallbackTriggeredFor) (event.Subscription, error) {

	logs, sub, err := _DOSProxy.contract.WatchLogs(opts, "LogCallbackTriggeredFor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DOSProxyLogCallbackTriggeredFor)
				if err := _DOSProxy.contract.UnpackLog(event, "LogCallbackTriggeredFor", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DOSProxyLogGroupingIterator is returned from FilterLogGrouping and is used to iterate over the raw logs and unpacked data for LogGrouping events raised by the DOSProxy contract.
type DOSProxyLogGroupingIterator struct {
	Event *DOSProxyLogGrouping // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DOSProxyLogGroupingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DOSProxyLogGrouping)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DOSProxyLogGrouping)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DOSProxyLogGroupingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DOSProxyLogGroupingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DOSProxyLogGrouping represents a LogGrouping event raised by the DOSProxy contract.
type DOSProxyLogGrouping struct {
	NodeId []*big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterLogGrouping is a free log retrieval operation binding the contract event 0x5f30b698cceb472bcb5a80c4acc8c52ea45ea704f5aeeb2527d2d4c95f793dd7.
//
// Solidity: e LogGrouping(NodeId uint256[])
func (_DOSProxy *DOSProxyFilterer) FilterLogGrouping(opts *bind.FilterOpts) (*DOSProxyLogGroupingIterator, error) {

	logs, sub, err := _DOSProxy.contract.FilterLogs(opts, "LogGrouping")
	if err != nil {
		return nil, err
	}
	return &DOSProxyLogGroupingIterator{contract: _DOSProxy.contract, event: "LogGrouping", logs: logs, sub: sub}, nil
}

// WatchLogGrouping is a free log subscription operation binding the contract event 0x5f30b698cceb472bcb5a80c4acc8c52ea45ea704f5aeeb2527d2d4c95f793dd7.
//
// Solidity: e LogGrouping(NodeId uint256[])
func (_DOSProxy *DOSProxyFilterer) WatchLogGrouping(opts *bind.WatchOpts, sink chan<- *DOSProxyLogGrouping) (event.Subscription, error) {

	logs, sub, err := _DOSProxy.contract.WatchLogs(opts, "LogGrouping")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DOSProxyLogGrouping)
				if err := _DOSProxy.contract.UnpackLog(event, "LogGrouping", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DOSProxyLogInsufficientGroupNumberIterator is returned from FilterLogInsufficientGroupNumber and is used to iterate over the raw logs and unpacked data for LogInsufficientGroupNumber events raised by the DOSProxy contract.
type DOSProxyLogInsufficientGroupNumberIterator struct {
	Event *DOSProxyLogInsufficientGroupNumber // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DOSProxyLogInsufficientGroupNumberIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DOSProxyLogInsufficientGroupNumber)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DOSProxyLogInsufficientGroupNumber)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DOSProxyLogInsufficientGroupNumberIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DOSProxyLogInsufficientGroupNumberIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DOSProxyLogInsufficientGroupNumber represents a LogInsufficientGroupNumber event raised by the DOSProxy contract.
type DOSProxyLogInsufficientGroupNumber struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogInsufficientGroupNumber is a free log retrieval operation binding the contract event 0x08a70ba288e836bee6c9b4aea7482ee5ff8f63c5ad9d2533d9cf0ced64adc262.
//
// Solidity: e LogInsufficientGroupNumber()
func (_DOSProxy *DOSProxyFilterer) FilterLogInsufficientGroupNumber(opts *bind.FilterOpts) (*DOSProxyLogInsufficientGroupNumberIterator, error) {

	logs, sub, err := _DOSProxy.contract.FilterLogs(opts, "LogInsufficientGroupNumber")
	if err != nil {
		return nil, err
	}
	return &DOSProxyLogInsufficientGroupNumberIterator{contract: _DOSProxy.contract, event: "LogInsufficientGroupNumber", logs: logs, sub: sub}, nil
}

// WatchLogInsufficientGroupNumber is a free log subscription operation binding the contract event 0x08a70ba288e836bee6c9b4aea7482ee5ff8f63c5ad9d2533d9cf0ced64adc262.
//
// Solidity: e LogInsufficientGroupNumber()
func (_DOSProxy *DOSProxyFilterer) WatchLogInsufficientGroupNumber(opts *bind.WatchOpts, sink chan<- *DOSProxyLogInsufficientGroupNumber) (event.Subscription, error) {

	logs, sub, err := _DOSProxy.contract.WatchLogs(opts, "LogInsufficientGroupNumber")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DOSProxyLogInsufficientGroupNumber)
				if err := _DOSProxy.contract.UnpackLog(event, "LogInsufficientGroupNumber", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DOSProxyLogNonContractCallIterator is returned from FilterLogNonContractCall and is used to iterate over the raw logs and unpacked data for LogNonContractCall events raised by the DOSProxy contract.
type DOSProxyLogNonContractCallIterator struct {
	Event *DOSProxyLogNonContractCall // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DOSProxyLogNonContractCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DOSProxyLogNonContractCall)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DOSProxyLogNonContractCall)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DOSProxyLogNonContractCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DOSProxyLogNonContractCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DOSProxyLogNonContractCall represents a LogNonContractCall event raised by the DOSProxy contract.
type DOSProxyLogNonContractCall struct {
	From common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogNonContractCall is a free log retrieval operation binding the contract event 0x6cea43bb3db7220931a7c8ac633e65cbc8e7ba129f2ed84db2e71bc0adb73bb5.
//
// Solidity: e LogNonContractCall(from address)
func (_DOSProxy *DOSProxyFilterer) FilterLogNonContractCall(opts *bind.FilterOpts) (*DOSProxyLogNonContractCallIterator, error) {

	logs, sub, err := _DOSProxy.contract.FilterLogs(opts, "LogNonContractCall")
	if err != nil {
		return nil, err
	}
	return &DOSProxyLogNonContractCallIterator{contract: _DOSProxy.contract, event: "LogNonContractCall", logs: logs, sub: sub}, nil
}

// WatchLogNonContractCall is a free log subscription operation binding the contract event 0x6cea43bb3db7220931a7c8ac633e65cbc8e7ba129f2ed84db2e71bc0adb73bb5.
//
// Solidity: e LogNonContractCall(from address)
func (_DOSProxy *DOSProxyFilterer) WatchLogNonContractCall(opts *bind.WatchOpts, sink chan<- *DOSProxyLogNonContractCall) (event.Subscription, error) {

	logs, sub, err := _DOSProxy.contract.WatchLogs(opts, "LogNonContractCall")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DOSProxyLogNonContractCall)
				if err := _DOSProxy.contract.UnpackLog(event, "LogNonContractCall", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DOSProxyLogNonSupportedTypeIterator is returned from FilterLogNonSupportedType and is used to iterate over the raw logs and unpacked data for LogNonSupportedType events raised by the DOSProxy contract.
type DOSProxyLogNonSupportedTypeIterator struct {
	Event *DOSProxyLogNonSupportedType // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DOSProxyLogNonSupportedTypeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DOSProxyLogNonSupportedType)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DOSProxyLogNonSupportedType)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DOSProxyLogNonSupportedTypeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DOSProxyLogNonSupportedTypeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DOSProxyLogNonSupportedType represents a LogNonSupportedType event raised by the DOSProxy contract.
type DOSProxyLogNonSupportedType struct {
	QueryType string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterLogNonSupportedType is a free log retrieval operation binding the contract event 0x70714cf695ae953ee67221716a4b4dc9e944909fd2b66f07e790a49d9ac29b41.
//
// Solidity: e LogNonSupportedType(queryType string)
func (_DOSProxy *DOSProxyFilterer) FilterLogNonSupportedType(opts *bind.FilterOpts) (*DOSProxyLogNonSupportedTypeIterator, error) {

	logs, sub, err := _DOSProxy.contract.FilterLogs(opts, "LogNonSupportedType")
	if err != nil {
		return nil, err
	}
	return &DOSProxyLogNonSupportedTypeIterator{contract: _DOSProxy.contract, event: "LogNonSupportedType", logs: logs, sub: sub}, nil
}

// WatchLogNonSupportedType is a free log subscription operation binding the contract event 0x70714cf695ae953ee67221716a4b4dc9e944909fd2b66f07e790a49d9ac29b41.
//
// Solidity: e LogNonSupportedType(queryType string)
func (_DOSProxy *DOSProxyFilterer) WatchLogNonSupportedType(opts *bind.WatchOpts, sink chan<- *DOSProxyLogNonSupportedType) (event.Subscription, error) {

	logs, sub, err := _DOSProxy.contract.WatchLogs(opts, "LogNonSupportedType")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DOSProxyLogNonSupportedType)
				if err := _DOSProxy.contract.UnpackLog(event, "LogNonSupportedType", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DOSProxyLogPublicKeyAcceptedIterator is returned from FilterLogPublicKeyAccepted and is used to iterate over the raw logs and unpacked data for LogPublicKeyAccepted events raised by the DOSProxy contract.
type DOSProxyLogPublicKeyAcceptedIterator struct {
	Event *DOSProxyLogPublicKeyAccepted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DOSProxyLogPublicKeyAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DOSProxyLogPublicKeyAccepted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DOSProxyLogPublicKeyAccepted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DOSProxyLogPublicKeyAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DOSProxyLogPublicKeyAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DOSProxyLogPublicKeyAccepted represents a LogPublicKeyAccepted event raised by the DOSProxy contract.
type DOSProxyLogPublicKeyAccepted struct {
	X1  *big.Int
	X2  *big.Int
	Y1  *big.Int
	Y2  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogPublicKeyAccepted is a free log retrieval operation binding the contract event 0x5bf077aca7c3c2ae9ca6ebd7da84490edf523227621a171c3c2656b3d45e92fe.
//
// Solidity: e LogPublicKeyAccepted(x1 uint256, x2 uint256, y1 uint256, y2 uint256)
func (_DOSProxy *DOSProxyFilterer) FilterLogPublicKeyAccepted(opts *bind.FilterOpts) (*DOSProxyLogPublicKeyAcceptedIterator, error) {

	logs, sub, err := _DOSProxy.contract.FilterLogs(opts, "LogPublicKeyAccepted")
	if err != nil {
		return nil, err
	}
	return &DOSProxyLogPublicKeyAcceptedIterator{contract: _DOSProxy.contract, event: "LogPublicKeyAccepted", logs: logs, sub: sub}, nil
}

// WatchLogPublicKeyAccepted is a free log subscription operation binding the contract event 0x5bf077aca7c3c2ae9ca6ebd7da84490edf523227621a171c3c2656b3d45e92fe.
//
// Solidity: e LogPublicKeyAccepted(x1 uint256, x2 uint256, y1 uint256, y2 uint256)
func (_DOSProxy *DOSProxyFilterer) WatchLogPublicKeyAccepted(opts *bind.WatchOpts, sink chan<- *DOSProxyLogPublicKeyAccepted) (event.Subscription, error) {

	logs, sub, err := _DOSProxy.contract.WatchLogs(opts, "LogPublicKeyAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DOSProxyLogPublicKeyAccepted)
				if err := _DOSProxy.contract.UnpackLog(event, "LogPublicKeyAccepted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DOSProxyLogQueryFromNonExistentUCIterator is returned from FilterLogQueryFromNonExistentUC and is used to iterate over the raw logs and unpacked data for LogQueryFromNonExistentUC events raised by the DOSProxy contract.
type DOSProxyLogQueryFromNonExistentUCIterator struct {
	Event *DOSProxyLogQueryFromNonExistentUC // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DOSProxyLogQueryFromNonExistentUCIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DOSProxyLogQueryFromNonExistentUC)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DOSProxyLogQueryFromNonExistentUC)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DOSProxyLogQueryFromNonExistentUCIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DOSProxyLogQueryFromNonExistentUCIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DOSProxyLogQueryFromNonExistentUC represents a LogQueryFromNonExistentUC event raised by the DOSProxy contract.
type DOSProxyLogQueryFromNonExistentUC struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterLogQueryFromNonExistentUC is a free log retrieval operation binding the contract event 0x158bff16635ac24f3d1acce162f0626cc6751bd434047538d76421366edf5906.
//
// Solidity: e LogQueryFromNonExistentUC()
func (_DOSProxy *DOSProxyFilterer) FilterLogQueryFromNonExistentUC(opts *bind.FilterOpts) (*DOSProxyLogQueryFromNonExistentUCIterator, error) {

	logs, sub, err := _DOSProxy.contract.FilterLogs(opts, "LogQueryFromNonExistentUC")
	if err != nil {
		return nil, err
	}
	return &DOSProxyLogQueryFromNonExistentUCIterator{contract: _DOSProxy.contract, event: "LogQueryFromNonExistentUC", logs: logs, sub: sub}, nil
}

// WatchLogQueryFromNonExistentUC is a free log subscription operation binding the contract event 0x158bff16635ac24f3d1acce162f0626cc6751bd434047538d76421366edf5906.
//
// Solidity: e LogQueryFromNonExistentUC()
func (_DOSProxy *DOSProxyFilterer) WatchLogQueryFromNonExistentUC(opts *bind.WatchOpts, sink chan<- *DOSProxyLogQueryFromNonExistentUC) (event.Subscription, error) {

	logs, sub, err := _DOSProxy.contract.WatchLogs(opts, "LogQueryFromNonExistentUC")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DOSProxyLogQueryFromNonExistentUC)
				if err := _DOSProxy.contract.UnpackLog(event, "LogQueryFromNonExistentUC", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DOSProxyLogUpdateRandomIterator is returned from FilterLogUpdateRandom and is used to iterate over the raw logs and unpacked data for LogUpdateRandom events raised by the DOSProxy contract.
type DOSProxyLogUpdateRandomIterator struct {
	Event *DOSProxyLogUpdateRandom // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DOSProxyLogUpdateRandomIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DOSProxyLogUpdateRandom)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DOSProxyLogUpdateRandom)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DOSProxyLogUpdateRandomIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DOSProxyLogUpdateRandomIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DOSProxyLogUpdateRandom represents a LogUpdateRandom event raised by the DOSProxy contract.
type DOSProxyLogUpdateRandom struct {
	LastRandomness   *big.Int
	LastUpdatedBlock *big.Int
	DispatchedGroup  [4]*big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLogUpdateRandom is a free log retrieval operation binding the contract event 0xb57d39816b5fafa35e9395c99afd19edb72160f491310ba0942ea8c58c8bc9f9.
//
// Solidity: e LogUpdateRandom(lastRandomness uint256, lastUpdatedBlock uint256, dispatchedGroup uint256[4])
func (_DOSProxy *DOSProxyFilterer) FilterLogUpdateRandom(opts *bind.FilterOpts) (*DOSProxyLogUpdateRandomIterator, error) {

	logs, sub, err := _DOSProxy.contract.FilterLogs(opts, "LogUpdateRandom")
	if err != nil {
		return nil, err
	}
	return &DOSProxyLogUpdateRandomIterator{contract: _DOSProxy.contract, event: "LogUpdateRandom", logs: logs, sub: sub}, nil
}

// WatchLogUpdateRandom is a free log subscription operation binding the contract event 0xb57d39816b5fafa35e9395c99afd19edb72160f491310ba0942ea8c58c8bc9f9.
//
// Solidity: e LogUpdateRandom(lastRandomness uint256, lastUpdatedBlock uint256, dispatchedGroup uint256[4])
func (_DOSProxy *DOSProxyFilterer) WatchLogUpdateRandom(opts *bind.WatchOpts, sink chan<- *DOSProxyLogUpdateRandom) (event.Subscription, error) {

	logs, sub, err := _DOSProxy.contract.WatchLogs(opts, "LogUpdateRandom")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DOSProxyLogUpdateRandom)
				if err := _DOSProxy.contract.UnpackLog(event, "LogUpdateRandom", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DOSProxyLogUrlIterator is returned from FilterLogUrl and is used to iterate over the raw logs and unpacked data for LogUrl events raised by the DOSProxy contract.
type DOSProxyLogUrlIterator struct {
	Event *DOSProxyLogUrl // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DOSProxyLogUrlIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DOSProxyLogUrl)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DOSProxyLogUrl)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DOSProxyLogUrlIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DOSProxyLogUrlIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DOSProxyLogUrl represents a LogUrl event raised by the DOSProxy contract.
type DOSProxyLogUrl struct {
	QueryId         *big.Int
	Url             string
	Timeout         *big.Int
	Randomness      *big.Int
	DispatchedGroup [4]*big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterLogUrl is a free log retrieval operation binding the contract event 0x75b4fd18552b0d79989d6cc5de124544de30dc29b548e1f743a719ce5ad44329.
//
// Solidity: e LogUrl(queryId uint256, url string, timeout uint256, randomness uint256, dispatchedGroup uint256[4])
func (_DOSProxy *DOSProxyFilterer) FilterLogUrl(opts *bind.FilterOpts) (*DOSProxyLogUrlIterator, error) {

	logs, sub, err := _DOSProxy.contract.FilterLogs(opts, "LogUrl")
	if err != nil {
		return nil, err
	}
	return &DOSProxyLogUrlIterator{contract: _DOSProxy.contract, event: "LogUrl", logs: logs, sub: sub}, nil
}

// WatchLogUrl is a free log subscription operation binding the contract event 0x75b4fd18552b0d79989d6cc5de124544de30dc29b548e1f743a719ce5ad44329.
//
// Solidity: e LogUrl(queryId uint256, url string, timeout uint256, randomness uint256, dispatchedGroup uint256[4])
func (_DOSProxy *DOSProxyFilterer) WatchLogUrl(opts *bind.WatchOpts, sink chan<- *DOSProxyLogUrl) (event.Subscription, error) {

	logs, sub, err := _DOSProxy.contract.WatchLogs(opts, "LogUrl")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DOSProxyLogUrl)
				if err := _DOSProxy.contract.UnpackLog(event, "LogUrl", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DOSProxyLogValidationResultIterator is returned from FilterLogValidationResult and is used to iterate over the raw logs and unpacked data for LogValidationResult events raised by the DOSProxy contract.
type DOSProxyLogValidationResultIterator struct {
	Event *DOSProxyLogValidationResult // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DOSProxyLogValidationResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DOSProxyLogValidationResult)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DOSProxyLogValidationResult)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DOSProxyLogValidationResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DOSProxyLogValidationResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DOSProxyLogValidationResult represents a LogValidationResult event raised by the DOSProxy contract.
type DOSProxyLogValidationResult struct {
	TrafficType uint8
	TrafficId   *big.Int
	Message     []byte
	Signature   [2]*big.Int
	PubKey      [4]*big.Int
	Pass        bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLogValidationResult is a free log retrieval operation binding the contract event 0xd33c44f7ce166bcd2616c8f4d811261d4d24e1af815f78614683e0c7682c42a5.
//
// Solidity: e LogValidationResult(trafficType uint8, trafficId uint256, message bytes, signature uint256[2], pubKey uint256[4], pass bool)
func (_DOSProxy *DOSProxyFilterer) FilterLogValidationResult(opts *bind.FilterOpts) (*DOSProxyLogValidationResultIterator, error) {

	logs, sub, err := _DOSProxy.contract.FilterLogs(opts, "LogValidationResult")
	if err != nil {
		return nil, err
	}
	return &DOSProxyLogValidationResultIterator{contract: _DOSProxy.contract, event: "LogValidationResult", logs: logs, sub: sub}, nil
}

// WatchLogValidationResult is a free log subscription operation binding the contract event 0xd33c44f7ce166bcd2616c8f4d811261d4d24e1af815f78614683e0c7682c42a5.
//
// Solidity: e LogValidationResult(trafficType uint8, trafficId uint256, message bytes, signature uint256[2], pubKey uint256[4], pass bool)
func (_DOSProxy *DOSProxyFilterer) WatchLogValidationResult(opts *bind.WatchOpts, sink chan<- *DOSProxyLogValidationResult) (event.Subscription, error) {

	logs, sub, err := _DOSProxy.contract.WatchLogs(opts, "LogValidationResult")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DOSProxyLogValidationResult)
				if err := _DOSProxy.contract.UnpackLog(event, "LogValidationResult", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// DOSProxyWhitelistAddressResetIterator is returned from FilterWhitelistAddressReset and is used to iterate over the raw logs and unpacked data for WhitelistAddressReset events raised by the DOSProxy contract.
type DOSProxyWhitelistAddressResetIterator struct {
	Event *DOSProxyWhitelistAddressReset // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *DOSProxyWhitelistAddressResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DOSProxyWhitelistAddressReset)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(DOSProxyWhitelistAddressReset)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *DOSProxyWhitelistAddressResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DOSProxyWhitelistAddressResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DOSProxyWhitelistAddressReset represents a WhitelistAddressReset event raised by the DOSProxy contract.
type DOSProxyWhitelistAddressReset struct {
	Previous common.Address
	Curr     common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWhitelistAddressReset is a free log retrieval operation binding the contract event 0x81212bcc8da3221a71cc77dc44d159e2a3aeeb6e1aaa17b1fffa5938540ab842.
//
// Solidity: e WhitelistAddressReset(previous address, curr address)
func (_DOSProxy *DOSProxyFilterer) FilterWhitelistAddressReset(opts *bind.FilterOpts) (*DOSProxyWhitelistAddressResetIterator, error) {

	logs, sub, err := _DOSProxy.contract.FilterLogs(opts, "WhitelistAddressReset")
	if err != nil {
		return nil, err
	}
	return &DOSProxyWhitelistAddressResetIterator{contract: _DOSProxy.contract, event: "WhitelistAddressReset", logs: logs, sub: sub}, nil
}

// WatchWhitelistAddressReset is a free log subscription operation binding the contract event 0x81212bcc8da3221a71cc77dc44d159e2a3aeeb6e1aaa17b1fffa5938540ab842.
//
// Solidity: e WhitelistAddressReset(previous address, curr address)
func (_DOSProxy *DOSProxyFilterer) WatchWhitelistAddressReset(opts *bind.WatchOpts, sink chan<- *DOSProxyWhitelistAddressReset) (event.Subscription, error) {

	logs, sub, err := _DOSProxy.contract.WatchLogs(opts, "WhitelistAddressReset")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DOSProxyWhitelistAddressReset)
				if err := _DOSProxy.contract.UnpackLog(event, "WhitelistAddressReset", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// UserContractInterfaceABI is the input ABI used to generate the binding from.
const UserContractInterfaceABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"__callback__\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// UserContractInterfaceBin is the compiled bytecode used for deploying new contracts.
const UserContractInterfaceBin = `0x`

// DeployUserContractInterface deploys a new Ethereum contract, binding an instance of UserContractInterface to it.
func DeployUserContractInterface(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UserContractInterface, error) {
	parsed, err := abi.JSON(strings.NewReader(UserContractInterfaceABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UserContractInterfaceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UserContractInterface{UserContractInterfaceCaller: UserContractInterfaceCaller{contract: contract}, UserContractInterfaceTransactor: UserContractInterfaceTransactor{contract: contract}, UserContractInterfaceFilterer: UserContractInterfaceFilterer{contract: contract}}, nil
}

// UserContractInterface is an auto generated Go binding around an Ethereum contract.
type UserContractInterface struct {
	UserContractInterfaceCaller     // Read-only binding to the contract
	UserContractInterfaceTransactor // Write-only binding to the contract
	UserContractInterfaceFilterer   // Log filterer for contract events
}

// UserContractInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type UserContractInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserContractInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UserContractInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserContractInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UserContractInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserContractInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UserContractInterfaceSession struct {
	Contract     *UserContractInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// UserContractInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UserContractInterfaceCallerSession struct {
	Contract *UserContractInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// UserContractInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UserContractInterfaceTransactorSession struct {
	Contract     *UserContractInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// UserContractInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type UserContractInterfaceRaw struct {
	Contract *UserContractInterface // Generic contract binding to access the raw methods on
}

// UserContractInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UserContractInterfaceCallerRaw struct {
	Contract *UserContractInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// UserContractInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UserContractInterfaceTransactorRaw struct {
	Contract *UserContractInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUserContractInterface creates a new instance of UserContractInterface, bound to a specific deployed contract.
func NewUserContractInterface(address common.Address, backend bind.ContractBackend) (*UserContractInterface, error) {
	contract, err := bindUserContractInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UserContractInterface{UserContractInterfaceCaller: UserContractInterfaceCaller{contract: contract}, UserContractInterfaceTransactor: UserContractInterfaceTransactor{contract: contract}, UserContractInterfaceFilterer: UserContractInterfaceFilterer{contract: contract}}, nil
}

// NewUserContractInterfaceCaller creates a new read-only instance of UserContractInterface, bound to a specific deployed contract.
func NewUserContractInterfaceCaller(address common.Address, caller bind.ContractCaller) (*UserContractInterfaceCaller, error) {
	contract, err := bindUserContractInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UserContractInterfaceCaller{contract: contract}, nil
}

// NewUserContractInterfaceTransactor creates a new write-only instance of UserContractInterface, bound to a specific deployed contract.
func NewUserContractInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*UserContractInterfaceTransactor, error) {
	contract, err := bindUserContractInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UserContractInterfaceTransactor{contract: contract}, nil
}

// NewUserContractInterfaceFilterer creates a new log filterer instance of UserContractInterface, bound to a specific deployed contract.
func NewUserContractInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*UserContractInterfaceFilterer, error) {
	contract, err := bindUserContractInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UserContractInterfaceFilterer{contract: contract}, nil
}

// bindUserContractInterface binds a generic wrapper to an already deployed contract.
func bindUserContractInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UserContractInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserContractInterface *UserContractInterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UserContractInterface.Contract.UserContractInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserContractInterface *UserContractInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserContractInterface.Contract.UserContractInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserContractInterface *UserContractInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserContractInterface.Contract.UserContractInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserContractInterface *UserContractInterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _UserContractInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserContractInterface *UserContractInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserContractInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserContractInterface *UserContractInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserContractInterface.Contract.contract.Transact(opts, method, params...)
}

// Callback_ is a paid mutator transaction binding the contract method 0x6d112977.
//
// Solidity: function __callback__( uint256,  bytes) returns()
func (_UserContractInterface *UserContractInterfaceTransactor) Callback_(opts *bind.TransactOpts, arg0 *big.Int, arg1 []byte) (*types.Transaction, error) {
	return _UserContractInterface.contract.Transact(opts, "__callback__", arg0, arg1)
}

// Callback_ is a paid mutator transaction binding the contract method 0x6d112977.
//
// Solidity: function __callback__( uint256,  bytes) returns()
func (_UserContractInterface *UserContractInterfaceSession) Callback_(arg0 *big.Int, arg1 []byte) (*types.Transaction, error) {
	return _UserContractInterface.Contract.Callback_(&_UserContractInterface.TransactOpts, arg0, arg1)
}

// Callback_ is a paid mutator transaction binding the contract method 0x6d112977.
//
// Solidity: function __callback__( uint256,  bytes) returns()
func (_UserContractInterface *UserContractInterfaceTransactorSession) Callback_(arg0 *big.Int, arg1 []byte) (*types.Transaction, error) {
	return _UserContractInterface.Contract.Callback_(&_UserContractInterface.TransactOpts, arg0, arg1)
}
