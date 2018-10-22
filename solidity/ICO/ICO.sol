pragma solidity ^0.4.16;


// token接口
interface token {
    function transfer(address receiver, uint amount) external ;
}



contract Ico {
    // 受益人
    address public beneficiary;
    // 额度
    uint public fundingGoal;
    uint public amountRaised; // 参与数量

    uint public deadline;  // 截止日期
    uint public price; // 汇率
    token public tokenReward; // 要卖的token

    mapping(address => uint256) public balanceOf; //
    bool crowdsaleClosed = false; //

    /**
    事件 可以用来跟踪信息

    */
    event GoalReached(address recipient, uint totalAmountRaised);  //
    event FundTransfer(address backer, uint amount, bool isContribution); //


    // 构造器器
    constructor (
        uint fundingGoalInEthers,
        uint durationInMinutes,
        uint etherCostOfEachToken,
        address addressOfTokenUsedAsReward
    ) public {
        beneficiary = msg.sender;
        fundingGoal = fundingGoalInEthers * 1 ether;
        deadline = now + durationInMinutes * 1 minutes;
        price = etherCostOfEachToken * 1 ether;

        // 传入已发布的token合约的地址来创建实例
        tokenReward = token(addressOfTokenUsedAsReward);
    }


    // 回退函数
    // 在向合约转账时，这个函数会被调用
    function () public payable {
        require(!crowdsaleClosed);

        uint amount = msg.value;  // wei
        balanceOf[msg.sender] += amount;

        amountRaised += amount;

        tokenReward.transfer(msg.sender, amount / price);

        emit FundTransfer(msg.sender, amount, true);
    }


    /**
    定义了函数修改器
    用于在函数执行前检查某些前置条件
    _ 表示继续执行之后的代码

    */
    modifier afterDeadline() {
        if (now >= deadline) {
            _;
        }
    }


    // 判断众筹是否完成
    function checkGoalReached() public afterDeadline {
        if (amountRaised >= fundingGoal) {
            emit GoalReached(beneficiary, amountRaised);
        }
        crowdsaleClosed = true;
    }



    function safeWithdrawal() public afterDeadline {

        // 众筹失败，退款
        if (amountRaised < fundingGoal) {
            uint amount = balanceOf[msg.sender];
            balanceOf[msg.sender] = 0;
            if (amount > 0) {
                msg.sender.transfer(amount);
                emit FundTransfer(msg.sender, amount, false);
            }
        }

        // 众筹成功，发到beneficiary地址上
        if (fundingGoal <= amountRaised && beneficiary == msg.sender) {
            beneficiary.transfer(amountRaised);
            emit FundTransfer(beneficiary, amountRaised, false);
        }
    }
}
