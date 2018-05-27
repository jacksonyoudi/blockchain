pragma solidity ^0.4.4;



/*
public internal private

只有public类型才能继承到子合约中
 */

contract Animal {
    uint _weight;
    uint private _height;
    uint internal _age;
    uint public _money;


    //public
    function test() constant returns (uint) {
        //this._money
        return _weight;
    }

    function test1() constant public returns (uint) {
        return _height;
    }

    function test2() constant internal returns (uint) {
        return _age;
    }

    function test3() constant private returns (uint) {
        return _money;
    }

}


contract Animal1 {
    uint _sex; // 1 male 2: female

    function Animal1() {
        _sex = 1;
    }

    function sex() constant returns (uint) {
        return _sex;
    }
}



contract Dog is Animal, Animal1{
    function testWeight() constant returns (uint) {
        return _weight;
    }
    function testAge() constant returns (uint) {
        return _age;
    }

    function testMoney() constant returns (uint) {
        return _money;
    }

    /* function testHeight() constant returns (uint) {
        return _height;
    } */

    function sex() constant returns (uint) {
        return 0;
    }

}







/*


*/
