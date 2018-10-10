pragma solidity ^0.4.4;

contract Test {

    function test1() constant returns (uint) {
        uint8 a = 2 ** 800;
        return a;
    }
}
