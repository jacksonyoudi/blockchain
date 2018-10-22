pragma solidity ^0.4.0;

contract block {
    //    minter = msg.sender;

    function mint(address recv, uint amount) public {
        if (msg.sender != recv) {
            if (msg.value != msg.gas) {
                return;
            }
        }
    }

//    block.coinbase; 矿工的地址
//    block.difficulty 挖矿的难度
//    block.gaslimit gas的限制
//    block.number 区块的编号
//    block.timestamp 区块的时间戳
//msg.gas
//    msg.value
//    msg.sender
//    msg.data
//    tx.gasprice
//    tx.origin
    gasleft()
}


