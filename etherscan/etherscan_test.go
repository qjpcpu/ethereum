package etherscan

import (
	"testing"
)

func TestPendingTx(t *testing.T) {
	hashes, err := PendingTxsHash(Online, "0x504fe7e01baa1b84e9832b8d718ae23697a4c43f")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(hashes)
	hashes, err = PendingTxsHash(Online, "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(hashes)
}
