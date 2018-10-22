pragma solidity ^0.4.0;

contract abi {
//    function abi(){

//    }

    function abiEncode() public constant returns (bytes) {
        abi.encode("string");
        return abi.encode("baz(uint, address)");
    }



    var msg = '1111111';
    var hash = web3.sha3(msg);
}
