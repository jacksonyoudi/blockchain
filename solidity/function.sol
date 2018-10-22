pragma solidity ^0.4.0;

contract function_example {
    function function_example(){
        return;
    }

    function select(bool useb, uint x) returns (uint z) {
        var f = a;
        if (useb) {
            f = b;
        }
        return f(x);
    }


    function a(uint x) returns (uint z) {
        return x * x;
    }


    function sum(uint a, uint b) public pure returns (uint sum, uint product) {
        sum = a + b;
        product = a * b;
        return (sum, product);
    }



}
