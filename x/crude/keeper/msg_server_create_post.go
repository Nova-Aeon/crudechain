package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"crudechain/x/crude/types"
)

func (k msgServer) CreatePost(goCtx context.Context, msg *types.MsgCreatePost) (*types.MsgCreatePostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var post = types.Post{
		Creator: msg.Creator,
		Body:    msg.Body,
	}
	id := k.AppendPost(
		ctx,
		post,
	)
	return &types.MsgCreatePostResponse{
		Id: id,
	}, nil
}
