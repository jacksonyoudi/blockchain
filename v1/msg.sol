pragma solidity ^0.4.4;


/*
msg gobal是当前运行的account
msg.sender当调用钱包的地址

通过 _owner == msg.sender 是否允许哪些操作
 */

contract Test {
    address public _owner; //合约的地址，第一个部署的地址
    uint _number;


    function Test() {
        _owner = msg.sender;
        _number = 100;
    }

    function msgSenderAddress() constant public returns (address) {
        return msg.sender;
    }

    function setNumber() {
        _number = _number + 1;
    }


    function setNumber() {
        _number = _number + 5;
    }

    function setNumber2() {
        if (_owner == msg.sender) {
            _number = _number + 1;
        }
    }
}
