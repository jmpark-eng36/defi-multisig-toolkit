package signer

import "github.com/ethereum/go-ethereum/common"

type Activity struct { Address common.Address; TxCount uint64 }
const LookbackBlocks = 165600
