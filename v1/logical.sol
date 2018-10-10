pragma solidity ^0.4.4;


contract Test {
    uint8 _a;
    uint8 _b;

    function Test() {
        _a = 5; // 0000 0101
        _b = 8; //0000 1000
    }

    function and() constant public returns (uint8) {
        return _a & _b;
    }

    function or() constant public returns (uint8) {
        return _a | _b;
    }

    //相同为0,相反为1
    function not_or() constant public returns (uint8) {
        return _a ^ _b;
    }


    //相同为0,相反为1
    function not() constant public returns (uint8) {
        return  ~ _a;
    }
}
