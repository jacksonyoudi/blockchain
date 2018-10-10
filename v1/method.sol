pragma solidity ^0.4.4;


/*
    public 默认类型


*/

contract Person {
    function age() constant returns (uint) {
        return 180;
    }


    function weight() constant  public returns (uint) {
        return 100;
    }

    function money() constant  internal returns (uint) {
        return 20000;
    }

    function height() constant private returns (uint) {
        return 176;
    }


}

/**
    合约中的method默认是public类型，
    

 */
