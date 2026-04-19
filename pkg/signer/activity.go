package signer

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const DefaultLookbackBlocks = 15 * 7200

type Activity struct {
	Address    common.Address
	LastBlock  uint64
	TxCount    uint64
	DaysIdle   int
}

func GetActivity(ctx context.Context, client *ethclient.Client, addr common.Address) (*Activity, error) {
	nonce, err := client.NonceAt(ctx, addr, nil)
	if err != nil {
		return nil, err
	}
	current, _ := client.BlockNumber(ctx)
	return &Activity{
		Address:   addr,
		TxCount:   nonce,
		LastBlock: current,
	}, nil
}
