package validators_test

import (
	"bytes"
	"context"
	"encoding/hex"
	"testing"

	"code.vegaprotocol.io/vega/validators"

	abcitypes "github.com/tendermint/tendermint/abci/types"
	types1 "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtypes "github.com/tendermint/tendermint/types"

	commandspb "code.vegaprotocol.io/protos/vega/commands/v1"
	snapshot "code.vegaprotocol.io/protos/vega/snapshot/v1"
	"code.vegaprotocol.io/vega/types"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var topKey = (&types.PayloadTopology{}).Key()

func TestEmptySnapshot(t *testing.T) {
	top := getTestTopology(t)
	defer top.ctrl.Finish()

	h, err := top.GetHash(topKey)
	assert.Nil(t, err)
	assert.NotEmpty(t, h)

	s, p, err := top.GetState(topKey)
	assert.Nil(t, err)
	assert.Empty(t, p)
	assert.NotEmpty(t, s)

	assert.Equal(t, 1, len(top.Keys()))
}

func TestTopologySnapshot(t *testing.T) {
	top := getTestTopWithDefaultValidator(t)
	updateValidatorPerformanceToNonDefaultState(t, top.Topology)
	defer top.ctrl.Finish()

	h1, err := top.GetHash(topKey)
	require.Nil(t, err)

	tmPubKeys := []string{"tm-pubkey-1", "tm-pubkey-2"}
	top.UpdateValidatorSet(tmPubKeys)

	h2, err := top.GetHash(topKey)
	require.Nil(t, err)

	ctx := context.Background()

	nr1 := commandspb.NodeRegistration{
		Id:              "vega-master-pubkey",
		ChainPubKey:     tmPubKeys[0],
		VegaPubKey:      "vega-key",
		EthereumAddress: "eth-address",
	}
	err = top.AddNodeRegistration(ctx, &nr1)
	assert.NoError(t, err)

	nr2 := commandspb.NodeRegistration{
		Id:              "vega-master-pubkey-2",
		ChainPubKey:     tmPubKeys[1],
		VegaPubKey:      "vega-key-2",
		EthereumAddress: "eth-address-2",
	}
	err = top.AddNodeRegistration(ctx, &nr2)
	assert.NoError(t, err)

	kr1 := &commandspb.KeyRotateSubmission{
		NewPubKeyIndex:    1,
		TargetBlock:       10,
		NewPubKey:         "new-vega-key",
		CurrentPubKeyHash: hashKey(nr1.VegaPubKey),
	}
	err = top.AddKeyRotate(ctx, nr1.Id, 5, kr1)
	assert.NoError(t, err)

	kr2 := &commandspb.KeyRotateSubmission{
		NewPubKeyIndex:    1,
		TargetBlock:       11,
		NewPubKey:         "new-vega-key-2",
		CurrentPubKeyHash: hashKey(nr2.VegaPubKey),
	}
	err = top.AddKeyRotate(ctx, nr2.Id, 5, kr2)
	assert.NoError(t, err)

	// Check the hashes have changed after each state change
	h3, err := top.GetHash(topKey)
	require.Nil(t, err)
	require.False(t, bytes.Equal(h1, h2))
	require.False(t, bytes.Equal(h2, h3))
	require.False(t, bytes.Equal(h1, h3))

	// Get the state ready to load into a new instance of the engine
	state, _, _ := top.GetState(topKey)
	snap := &snapshot.Payload{}
	err = proto.Unmarshal(state, snap)
	require.Nil(t, err)

	snapTop := getTestTopWithDefaultValidator(t)
	defer snapTop.ctrl.Finish()

	_, err = snapTop.LoadState(context.Background(), types.PayloadFromProto(snap))
	require.Nil(t, err)

	// Check the new reloaded engine is the same as the original
	h4, err := snapTop.GetHash(topKey)
	require.Nil(t, err)
	require.True(t, bytes.Equal(h3, h4))
	assert.ElementsMatch(t, top.AllNodeIDs(), snapTop.AllNodeIDs())
	assert.ElementsMatch(t, top.AllVegaPubKeys(), snapTop.AllVegaPubKeys())
	assert.Equal(t, top.IsValidator(), snapTop.IsValidator())
	assert.Equal(t, top.GetPendingKeyRotation(kr1.TargetBlock, kr1.NewPubKey), snapTop.GetPendingKeyRotation(kr1.TargetBlock, kr1.NewPubKey))
	assert.Equal(t, top.GetPendingKeyRotation(kr2.TargetBlock, kr2.NewPubKey), snapTop.GetPendingKeyRotation(kr2.TargetBlock, kr2.NewPubKey))

	require.Equal(t, "0.5", snapTop.ValidatorPerformanceScore(hex.EncodeToString(address1)).String())
	require.Equal(t, "1", snapTop.ValidatorPerformanceScore(hex.EncodeToString(address2)).String())
	require.Equal(t, "1", snapTop.ValidatorPerformanceScore(hex.EncodeToString(address3)).String())
	require.Equal(t, "1", snapTop.ValidatorPerformanceScore(hex.EncodeToString(address4)).String())
	require.Equal(t, "1", snapTop.ValidatorPerformanceScore(hex.EncodeToString(address5)).String())
}

func updateValidatorPerformanceToNonDefaultState(t *testing.T, top *validators.Topology) {
	t.Helper()
	vd1 := []*tmtypes.Validator{
		{Address: address1, VotingPower: 3715, ProposerPriority: 5249},
		{Address: address2, VotingPower: 3351, ProposerPriority: 796},
		{Address: address3, VotingPower: 2793, ProposerPriority: -797},
		{Address: address4, VotingPower: 139, ProposerPriority: 1016},
		{Address: address5, VotingPower: 1, ProposerPriority: -6264},
	}
	req1 := abcitypes.RequestBeginBlock{Header: types1.Header{ProposerAddress: address1, Height: int64(1)}}
	top.BeginBlock(context.Background(), req1, vd1)

	vd2 := []*tmtypes.Validator{
		{Address: address1, VotingPower: 3715, ProposerPriority: 6433},
		{Address: address2, VotingPower: 3351, ProposerPriority: -1853},
		{Address: address3, VotingPower: 2793, ProposerPriority: 5347},
		{Address: address4, VotingPower: 139, ProposerPriority: -3701},
		{Address: address5, VotingPower: 1, ProposerPriority: -6226},
	}

	// expecting address1 to propose but got address3
	req2 := abcitypes.RequestBeginBlock{Header: types1.Header{ProposerAddress: address3, Height: int64(1)}}
	top.BeginBlock(context.Background(), req2, vd2)

	require.Equal(t, "0.5", top.ValidatorPerformanceScore(hex.EncodeToString(address1)).String())
	require.Equal(t, "1", top.ValidatorPerformanceScore(hex.EncodeToString(address2)).String())
	require.Equal(t, "1", top.ValidatorPerformanceScore(hex.EncodeToString(address3)).String())
	require.Equal(t, "1", top.ValidatorPerformanceScore(hex.EncodeToString(address4)).String())
	require.Equal(t, "1", top.ValidatorPerformanceScore(hex.EncodeToString(address5)).String())
}