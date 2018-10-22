pragma solidity ^0.4.0;

contract Selector {
    function f() public view returns (bytes4){
        return this.f.selector;
    }
}
