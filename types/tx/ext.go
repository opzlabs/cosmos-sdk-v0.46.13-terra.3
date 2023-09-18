package tx

import (
	"github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/codec/types"
)

// TxExtensionOptionI defines the interface for tx extension options
type TxExtensionOptionI interface{}

// unpackTxExtensionOptionsI unpacks Any's to TxExtensionOptionI's.
func unpackTxExtensionOptionsI(unpacker types.AnyUnpacker, anys []*types.Any) error {
	for _, any := range anys {
		var opt TxExtensionOptionI
		err := unpacker.UnpackAny(any, &opt)
		if err != nil {
			return err
		}
	}

	return nil
}
