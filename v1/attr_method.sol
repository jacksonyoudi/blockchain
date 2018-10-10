pragma solidity ^0.4.4;


contract Person {
    uint  _weight;
    uint internal _age;
    uint private _height;
    uint public _money;


    function weight() constant returns (uint) {
            return _weight;
    }

    function height() constant public  returns (uint) {
            return _height;
    }

    function age() constant internal returns (uint) {
            return _age;
    }

    function money() constant private returns (uint) {
            return _money;
    }

    function test() constant returns (uint) {
        return this.weight();
    }

}
/*
    在public method中可以访问所有的属性

    属性默认权限是internal,只有public类型的属性才可能供外部访问
    internal和private类型的属性只能在合约内部使用

    函数的权限默认是public类型，public类型的函数可以供外部访问，
    而internal和private类型的函数不能够通过指针访问，哪怕是在内部通过
    this进行访问也会报错。

    this._money 属性 this._money()方法

    不管是属性还是方法，只有是public类型时，才可以通过合约地址进行访问，
    如果是合约内部的this就是当前合约地址，在合约内部要访问internal和
    private类型的属性和方法，直接访问即可，不要试图通过this去访问。

 */
