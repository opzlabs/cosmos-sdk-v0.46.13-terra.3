package posthandler

import (
	sdk "github.com/opzlabs/cosmos-sdk-v0.46.13-terra.3/types"
)

// HandlerOptions are the options required for constructing a default SDK PostHandler.
type HandlerOptions struct{}

// NewPostHandler returns an empty posthandler chain.
func NewPostHandler(options HandlerOptions) (sdk.AnteHandler, error) {
	postDecorators := []sdk.AnteDecorator{}

	return sdk.ChainAnteDecorators(postDecorators...), nil
}
