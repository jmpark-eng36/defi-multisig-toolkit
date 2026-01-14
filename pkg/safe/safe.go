package safe

import "github.com/ethereum/go-ethereum/common"

type Safe struct { Address common.Address }
type SafeInfo struct { Threshold uint64; Nonce uint64 }
const MaxRetries = 24
