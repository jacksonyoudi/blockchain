pragma solidity ^0.4.0;

contract addressDemo {
    function addressDemo(){
        address a = 0xddddddddd;
    }

    function getAddress() returns (address) {
        return this;
    }

    function getBaclance() returns (float) {
        return this.balance;
    }

}
