package util

import (
	"testing"
)

func TestGetBlockNumberByTransaction(t *testing.T) {

	idx := "23ca36ad-c549-4c91-a1c8-bf662c2e0339"
	blockNumber := GetBlockNumberByTransaction(idx)

	if blockNumber != 4 {
		t.Fatalf("blockNumber is incorrect")
	}

}
