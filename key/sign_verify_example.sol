pragma solidity ^ 0.4 .11;

/* generate signature in golang code then verify in solidity
import (
    crand "crypto/rand"
    "github.com/ethereum/go-ethereum/common/hexutil"
    "github.com/ethereum/go-ethereum/common/math"
    "github.com/ethereum/go-ethereum/crypto"
    "github.com/qjpcpu/ethereum/contracts"
    "math/big"
    "testing"
)

func TestSignature(t *testing.T) {
    pk, err := newKey(crand.Reader)
    if err != nil {
        t.Fatal(err)
    }
    msg := crypto.Keccak256(common.HexToAddress("0xaaa").Bytes(),contracts.PackNum(big.NewInt(522)), contracts.PackNum(big.NewInt(1000)), contracts.PackNum(big.NewInt(0)))
    sign, err := Sign(pk, msg)
    if err != nil {
        t.Fatal(err)
    }
    from := crypto.PubkeyToAddress(pk.PublicKey).Hex()
    signHex := hexutil.Encode(sign)
    if err := VerifySign(from, signHex, msg); err != nil {
        t.Fatal(err)
    }
    t.Logf("address:%v", from)
    t.Logf("sign:%v", signHex)
    // if SignVerifyExample's owner is "from", opCounter is 0
    // SignVerifyExample.func_need_verify_sign(522,1000,signHex)
}
*/

contract SignVerifyExample {
    address owner = msg.sender;
    mapping(address => uint256) userOpCounter;

    function func_need_verify_sign(uint256 amount1, uint256 amount2, bytes sig) public returns(bool) {

        // This recreates the message that was signed on the client.
        bytes32 message = prefixed(keccak256(msg.sender, amount1, amount2, userOpCounter[msg.sender]));

        require(recoverSigner(message, sig) == owner);
        userOpCounter[msg.sender]++;
        return true;
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
