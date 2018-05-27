pragma solidity ^0.4.4;

//uint8 uint16 uint24 ... uint256
//int8 int16 int24 int32 ...int256

/* int => int256 */
/* uint => uint256 */

// int8代表有符号 +1， -3
// uint代表无符号 3,8,10

/*
int8-> 对应8位 0-255
uint8 ->  -127~127
int16 -> 0~511
uint16 ->  -255~255
*/



contract Test {
    uint8 public _a;

    function Test(uint8 a) {
        _a = a;
    }

    function a() constant public returns (uint8) {
        return _a;
    }

    function setA(uint8 a) {
        _a = a;
    }



}


/*
超过多少位，就会溢出，就是最大值
uint8 和 uint16之间可以隐式转换的




 */
