contract Token {
  function name() constant returns (string name);
    // 可选方法，返回代币符号，如EOS
    function symbol() constant returns (string symbol);
    // 可选方法,返回代币小数位数，如8
    function decimals() constant returns (uint8 decimals);

    // 货币总发行量
    function totalSupply() constant returns (uint256 totalSupply);
    // 获取某个账户的代币余额
    function balanceOf(address _owner) constant returns (uint256 balance);
    // (本人)向某人转账
    function transfer(address _to, uint256 _value) returns (bool success);
    // (本人)批准只能合约可以向某人转账
    function approve(address _spender, uint256 _value) returns (bool success);
    // 合约代理from向to转账(须先经过from账户approve)
    function transferFrom(address _from, address _to, uint256 _value) returns (bool success);
    // 查询_owner允许合约代理向_spender转账的金额
    function allowance(address _owner, address _spender) constant returns (uint256 remaining);
}
