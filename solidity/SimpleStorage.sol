pragma solidity ^0.4.0;

contract SimpleStorage {
    uint storedData;

    function set(uint x) public {
        storedData = x;
    }

    function get() public constant returns (uint) {
        return storedData;
    }
}


// web3.eth.contract构造一个智能合约， 参数是ABI
var simplestorageContract = web3.eth.contract(
        [
        {
            "constant" : false,
            "inputs" : [
                            {"name" : "x", "type" : "uint256"}
                        ],
                        "name" : "set", "outputs" : [],
        "payable" : false,
        "stateMutability" : "nonpayable",
        "type": "function"
        },
        {"constant" : false,
        "inputs" : [],
        "name": "get",
        "outputs" : [{"name" : "", "type" : "uint256"}], "payable" : false, "stateMutability" : "nonpayable","type" : "function"}]);


//创建合约的交易，用于创建合约的EVM字节码
var simplestorage = simplestorageContract.new({
    from : web3.eth.accounts[0],
    data : '0x608060405234801561001057600080fd5b5060ec8061001f6000396000f3fe6080604052600436106049576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806360fe47b114604e5780636d4ce63c146085575b600080fd5b348015605957600080fd5b50608360048036036020811015606e57600080fd5b810190808035906020019092919050505060ad565b005b348015609057600080fd5b50609760b7565b6040518082815260200191505060405180910390f35b8060008190555050565b6000805490509056fea165627a7a723058204fb2b71aa998e5ee76e2d562580def7eba86cd9232c48a2e341670d09b94889e0029',
    gas : '4700000'
}, function (e, contract){
    console.log(e, contract);
    if (typeof contract .address != = 'undefined') {
        console.log('Contract mined! address: ' + contract .address + ' transactionHash: ' + contract .transactionHash);
    }
})