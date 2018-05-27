pragma solidity ^0.4.4;

/*
pragma:版本声明
solidity: 开发语言
0.4.4:当前合约的版本，0.4代表主版本， .4代表修复bug是升级版本
^:代表向上兼容，0.4.4 ~ 0.4.9可以对我们当前的代码进行编译
 */



//类比： contract Person class Person extends Contract

/*
属性
行为
 */

contract Person {
    uint _height; //属性名一个下划线
    uint _age;
    address _owner; //代表合约的拥有者


    //初始化实例化 构造函数、析构函数

    /*
    方法名和合约相同时就属于析构函数，在创建对象的时，构造函数

     */
    function Person() {
        _height = 180;
        _age = 29;
        _owner = msg.sender;
    }

    function ownder() constant returns (address) {
        return _owner;
    }

    //set方法，可以修改_height属性
    function setHeight(uint height) {  //驼峰
        _height = height;
    }

    /*
        获取_height值, constant 代表方法只读
        function funName() constant return (unit)
     */
    function height() constant returns (uint){  //get方法
        return _height;
    }

    function setAge(uint age) {
        _age = age;
    }

    function age() constant returns (uint) {
        return  _age;
    }

    /*
    析构函数，调用后，实例就删除了
    */
    function kill() {
        if (_owner == msg.sender) {
            //析构函数
            selfdestruct(msg.sender);
        }
    }

}
