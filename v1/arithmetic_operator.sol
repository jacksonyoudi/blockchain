pramga solidity ^0.4.4;

contract Test {
    int8 _a;
    int8 _b;

    function Test() {
        _a = 5;
        _b = 8;
    }

    function add() constant public returns (int) {
        return _a + _b;
    }

    function sub() constant public returns (int) {
        return _a - _b;
    }


    function mul() constant public returns (int) {
        return _a * _b;
    }

    function div() constant public returns (int) {
        return _a / _b;
    }

    function mod() constant public returns (int) {
        return _a % _b;
    }
}

/*
    注意类型的使用

*/
