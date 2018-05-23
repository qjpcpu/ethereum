pragma solidity ^0.4.18;

import "./ownable.sol";

contract UnionPay is Ownable {
    address public platform;
    mapping(address => mapping(address => uint256)) public userReceipts;
    event UserPay(address _from,address _to,uint256 _amount, uint256 _amountIndeed,uint256 _nonce,uint256 _state);

    function payCash(uint256 _nonce,uint256 _feePercentage,address _to,uint256 _state, bytes _sig) payable public returns(bool) {
        require(_feePercentage>=0 && _feePercentage<=100);
        require(_to != address(0));
        require(_nonce == userReceipts[msg.sender][_to]+1);
        require(platform!=address(0));

        bytes32 message = prefixed(keccak256(msg.sender, _to, msg.value, _feePercentage,_nonce,_state));

        require(recoverSigner(message, _sig) == platform);
        userReceipts[msg.sender][_to]++;
        
        if (_feePercentage == 0){
            if (msg.value > 0){
                _to.transfer(msg.value);
            }
            emit UserPay(msg.sender,_to,msg.value,msg.value,_nonce,_state);
            return true;
        }        
        uint256 val = _feePercentage * msg.value;
        assert(val/_feePercentage == msg.value);
        val = val/100;
        if (msg.value-val>0){
            _to.transfer(msg.value - val);
        }
        emit UserPay(msg.sender,_to,msg.value,msg.value - val,_nonce,_state);
        return true;
    }
    
    function plainPay() public payable returns(bool){
        emit UserPay(msg.sender,address(this),msg.value,msg.value,0,0);
        return true;
    }
    
    function setPlatform(address _checker) public onlyOwner{
        require(_checker!=address(0));
        platform = _checker;
    }
    
    function withdraw() public onlyOwner{
        require(platform!=address(0));
        platform.transfer(address(this).balance);
    }
    
    function getBalance() public view returns(uint256){
        return address(this).balance;
    }


    // Signature methods

    function splitSignature(bytes sig)
    internal
    pure
    returns(uint8, bytes32, bytes32) {
        require(sig.length == 65);

        bytes32 r;
        bytes32 s;
        uint8 v;

        assembly {
            // first 32 bytes, after the length prefix
            r: = mload(add(sig, 32))
            // second 32 bytes
            s: = mload(add(sig, 64))
            // final byte (first byte of the next 32 bytes)
            v: = byte(0, mload(add(sig, 96)))
        }

        return (v, r, s);
    }

    function recoverSigner(bytes32 message, bytes sig)
    internal
    pure
    returns(address) {
        uint8 v;
        bytes32 r;
        bytes32 s;

        (v, r, s) = splitSignature(sig);

        return ecrecover(message, v, r, s);
    }

    // Builds a prefixed hash to mimic the behavior of eth_sign.
    function prefixed(bytes32 hash) internal pure returns(bytes32) {
        return keccak256("\x19Ethereum Signed Message:\n32", hash);
    }
}

