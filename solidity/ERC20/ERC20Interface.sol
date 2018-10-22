contract ERC20Interface {

    // 代币名称
    string public constant name = "Token Name";
    // 符号
    string public constant symbol = "SYM";
    // 小数点位数
    uint8 public constant decimals = 18;  // 18 is the most common number of decimal places
    // 0.0000000000000000001  个代币

    // 法币总量
    function totalSupply() public constant returns (uint);

    // 查看对应账号的代币余额
    function balanceOf(address tokenOwner) public constant returns (uint balance);

    // 控制代币的交易
    function allowance(address tokenOwner, address spender) public constant returns (uint remaining);

    // 允许用户可花费的代币数
    function approve(address spender, uint tokens) public returns (bool success);

    // 实现代币交易,转账
    function transfer(address to, uint tokens) public returns (bool success);

    //实现代币之间的交易
    function transferFrom(address from, address to, uint tokens) public returns (bool success);

    event Transfer(address indexed from, address indexed to, uint tokens);


    event Approval(address indexed tokenOwner, address indexed spender, uint tokens);
}