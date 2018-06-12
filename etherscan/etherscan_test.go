package etherscan

import (
	"testing"
)

func TestPendingTx(t *testing.T) {
	hashes, err := PendingTxs(Online, "0x504fe7e01baa1b84e9832b8d718ae23697a4c43f")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(hashes)
	hashes, err = PendingTxs(Online, "")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(hashes)
}
func TestPendingDetail(t *testing.T) {
	detail, err := GetBlockedPendingTx(Online, `0xe35f3e2a93322b61e5d8931f806ff38f4a4f4d88`, 83)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v", detail)
}
