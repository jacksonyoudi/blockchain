pragma solidity ^0.4.4;

/*
    public
    internal 合约属性默认的属性
    private

    public会有一个以attr的函数，但是自己函数优先级高于系统生成的属性函数

    internal、private的属性都不能被外部访问，，当属性是public的时候
    会生成一个和属性名相同并且返回值就是当前属性的get函数。

    当然，生成的这个get函数会覆盖掉public属性自动生成的函数。

*/


contract Person {
    uint  _age;
    uint public _weight;
    uint private _height;
    uint public _money;

    function _money() constant returns (uint) {
        return 120;
    }


}


/*

    在上述代码里面，_money 返回的数据是_money的函数。

 */
