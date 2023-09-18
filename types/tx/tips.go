package tx

import (
	sdk "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/types"
)

// TipTx defines the interface to be implemented by Txs that handle Tips.
type TipTx interface {
	sdk.FeeTx
	GetTip() *Tip
}
