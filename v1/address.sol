pramga solidity ^0.4.4;


/*
以太坊中的地址的长度是20个字节，一个字节等于8位，一共160位
所以 address类型可以使用uint160来声明。
address用于存储

 */

contract Test {
    address _owner;
    uint160 _ownerUint;

    function Test() {
        _owner = 0xed57fcD7059a401bc083622B55915c4c6fC81E31;
        _ownerUint = 1354993002981732877008587219546772909048753298993;
    }

    function owner() constant returns (address) {
        return _owner;
    }

    function ownerUint160() constant public returns (uint) {
            return uint160(_owner);
    }

    function ownderUnitToAddress() constant returns (address) {
        return address(_ownerUint);
    }
}

//
