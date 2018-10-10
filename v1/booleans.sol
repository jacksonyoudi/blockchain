pragma solidity ^0.4.4;

contract Test {
    uint _a;
    uint _b;
    bool _c;

    function Test() {
        _a = 100;
        _b = 200;
        _c = true;
    }

    // !
    function not() constant public returns (bool) {
        return !_c;
    }

    //&&
    function and() constant public returns (bool) {
        return _a && (_a == 100);
    }

    // ||
    function or() constant public returns (bool) {
        return _a || (_a == 100);
    }

    // noteq
    function noteq() constant public returns (bool) {
        return _a && (_a != 100);
    }

}
