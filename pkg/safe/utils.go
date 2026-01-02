package safe

import "math/big"

func WeiToEth(wei *big.Int) float64 {
	f, _ := new(big.Float).Quo(new(big.Float).SetInt(wei), new(big.Float).SetInt(big.NewInt(1e18))).Float64()
	return f
}
// v10
