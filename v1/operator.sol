pragma solidity ^0.4.4;


contract Test {
    uint _a;
    uint _b;

    function Test() {
        _a = 100;
        _b = 200;
    }

    function op1() constant public returns (bool) {
        return _a < _b;
    }



    function op2() constant public returns (bool) {
        return _a <= _b;
    }

    function op3() constant public returns (bool) {
        return _a != _b;
    }

    function op4() constant public returns (bool) {
        return _a == _b;
    }

    function op5() constant public returns (bool) {
        return _a > _b;
    }


    function op1() constant public returns (bool) {
        return _a >= _b;
    }
}
