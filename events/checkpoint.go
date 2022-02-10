package events

import (
	"context"
	"encoding/hex"

	eventspb "code.vegaprotocol.io/protos/vega/events/v1"
	vgcontext "code.vegaprotocol.io/vega/libs/context"
	"code.vegaprotocol.io/vega/types"
)

type Checkpoint struct {
	*Base
	data eventspb.CheckpointEvent
}

func NewCheckpointEvent(ctx context.Context, snap *types.CheckpointState) *Checkpoint {
	height, _ := vgcontext.BlockHeightFromContext(ctx)
	_, block := vgcontext.TraceIDFromContext(ctx)
	return &Checkpoint{
		Base: newBase(ctx, CheckpointEvent),
		data: eventspb.CheckpointEvent{
			Hash:        hex.EncodeToString(snap.Hash),
			BlockHash:   block,
			BlockHeight: uint64(height),
		},
	}
}

func (e Checkpoint) Proto() eventspb.CheckpointEvent {
	return e.data
}

func (e Checkpoint) StreamMessage() *eventspb.BusEvent {
	busEvent := newBusEventFromBase(e.Base)
	busEvent.Event = &eventspb.BusEvent_Checkpoint{
		Checkpoint: &e.data,
	}
	return busEvent
}

func CheckpointEventFromStream(ctx context.Context, be *eventspb.BusEvent) *Checkpoint {
	if event := be.GetCheckpoint(); event != nil {
		return &Checkpoint{
			Base: newBaseFromBusEvent(ctx, CheckpointEvent, be),
			data: *event,
		}
	}
	return nil
}