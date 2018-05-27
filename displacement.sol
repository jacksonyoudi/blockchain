pragma solidity ^0.4.4;

contract Test {
    int _a;

    function Test() {
        _a = 3;
    }

    function pow(uint8 b) constant returns (int) {
        return _a ** b;
    }

    function leftShift(uint8 b) constant returns (int) {
        return  _a << b;
    }

    function rightShift(uint8 b) constant returns (int) {
        return _a >> b;
    }

}
