package events

import (
	"context"

	eventspb "code.vegaprotocol.io/protos/vega/events/v1"

	"code.vegaprotocol.io/vega/types"
)

type StakeLinking struct {
	*Base
	evt eventspb.StakeLinking
}

func NewStakeLinking(ctx context.Context, evt types.StakeLinking) *StakeLinking {
	return &StakeLinking{
		Base: newBase(ctx, StakeLinkingEvent),
		evt:  *(evt.IntoProto()),
	}
}

func (s StakeLinking) StakeLinking() eventspb.StakeLinking {
	return s.evt
}

func (s StakeLinking) Proto() eventspb.StakeLinking {
	return s.evt
}

func (s StakeLinking) StreamMessage() *eventspb.BusEvent {
	busEvent := newBusEventFromBase(s.Base)
	busEvent.Event = &eventspb.BusEvent_StakeLinking{
		StakeLinking: &s.evt,
	}

	return busEvent
}

func StakeLinkingFromStream(ctx context.Context, be *eventspb.BusEvent) *StakeLinking {
	return &StakeLinking{
		Base: newBaseFromBusEvent(ctx, StakeLinkingEvent, be),
		evt:  *be.GetStakeLinking(),
	}
}