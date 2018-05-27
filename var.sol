pragma solidity ^0.4.4;

//uint8 0-255

contract Test {

    function testVar() constant returns (uint16) {
        uint8 a = 100;
        var b = a;
        return b;
    }

    function testVar1() constant public returns (uint) {
        uint a;
        for (var i=0;i<=255; i++) {
            a = i;
        }
        return a;
    }

}


/*
    var 是赋值类型是计算的类型 var b =a;
    类型不匹配，是错误的。

    uint8 和 uint16之间可以隐式转换的
*/
