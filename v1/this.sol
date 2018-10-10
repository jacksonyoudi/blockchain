pragma solidity ^0.4.4;

/*
this 指向当前的合约对象的指针

 */

contract adderssBalance {
    function getBalance(address addr) constant returns (uint) {
        return addr.balance;
    }

    function getCurrentAddressBalance() constant returns (uint) {
        return this.balance;
    }
}

/*

当前合约地址 也是合法的钱包地址

this.balance
*/
