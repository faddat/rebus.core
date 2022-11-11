package claim

import (
	"fmt"

	"github.com/rebuschain/rebus.core/v1/x/claim/keeper"
	"github.com/rebuschain/rebus.core/v1/x/claim/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// NewHandler returns claim module messages
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		_ = ctx.WithEventManager(sdk.NewEventManager())

		errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
	}
}
