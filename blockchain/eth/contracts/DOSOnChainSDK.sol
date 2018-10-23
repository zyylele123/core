pragma solidity ^0.4.24;

interface DOSProxyInterface {
  function query(address, uint, uint, string, string) external returns (uint);
}

interface DOSAddressBridgeInterface {
  function getProxyAddress() external view returns (address);
}

contract DOSOnChainSDK {
  DOSProxyInterface dos_proxy;
  DOSAddressBridgeInterface dos_addr_bridge =
    DOSAddressBridgeInterface(0x87095a8115b8385E6A4852640eC9852cD9b6ad9E);

  modifier resolveAddress {
      dos_proxy = DOSProxyInterface(dos_addr_bridge.getProxyAddress());
      _;
  }

  function fromDOSProxyContract() internal view returns (address) {
      return dos_addr_bridge.getProxyAddress();
  }

  // TODO: Working on response parser.
  // @return a unique query_id for parallel requests. A zeroed (0x0) query_id
  // represents error.
  function DOSQuery(uint timeout, string query_type, string query_string)
    resolveAddress
    internal
    returns (uint)
    {
      return dos_proxy.query(this, block.number, timeout, query_type, query_string);
    }
}
