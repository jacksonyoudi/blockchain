pragma solidity ^0.4.0;

contract D {
    uint x;
    function D(uint a) public payable {
        x = a;
    }

}

contract c {
    D d = new D(4); //将作为c构造的一部分被执行

    function createD(uint arg) public {
        D newD = new D(arg);
    }

    function createAndEndowD(uint arg, uint amount) public payable {
        //创建合约时发送以太币
        D newD = (new D).value(amount)(arg);
    }
}