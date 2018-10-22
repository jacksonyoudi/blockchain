pragma solidity ^0.4.0;

contract owned {
    function owned() public {
        owner = msg.sender;
    }

    address owner;


    //函数修饰器
    modifier onlyOwner {
        require(msg.sender == owner);
        _;
    }

}


//合约继承
contract mortal is owned {
    //函数修饰器继承
    function close() public onlyOwner {
        selfdestruct(owner);
    }
}


contract priced {
    //函数修改器可以接收参数
    modifier costs(uint price) {
        if (msg.value >= price) {
            _;
        }
    }
}



contract Register is priced, owned {
    mapping (address => bool ) registerAddresses;
    uint price;
    function Register(uint initialPrice) public {
        price = initialPrice;
    }


    function register() public payable costs(price) {
        registerAddresses[msg.sender] = true;
    }

    function changePrice(uint _price) public onlyOwner {
        price = _price;
    }
 }


contract Mutex {
    bool locked;
    modifier noReentrancy() {
        require(!locked);
        locked = true;
        _;
        locked = false;
    }


    function f() public noReentrancy returns (uint) {
        require(msg.sender.call());
        return 7;
    }

}

