package gql

import (
	"context"

	types "code.vegaprotocol.io/vega/proto"
)

type newAssetResolver VegaResolverRoot

func (r *newAssetResolver) Source(ctx context.Context, obj *types.NewAsset) (AssetSource, error) {
	return AssetSourceFromProto(obj.Changes)
}