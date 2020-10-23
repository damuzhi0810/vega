// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/events.proto

package proto

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *MarketEvent) Validate() error {
	return nil
}
func (this *TimeUpdate) Validate() error {
	return nil
}
func (this *TransferResponses) Validate() error {
	for _, item := range this.Responses {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Responses", err)
			}
		}
	}
	return nil
}
func (this *PositionResolution) Validate() error {
	return nil
}
func (this *LossSocialization) Validate() error {
	return nil
}
func (this *TradeSettlement) Validate() error {
	return nil
}
func (this *SettlePosition) Validate() error {
	for _, item := range this.TradeSettlements {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("TradeSettlements", err)
			}
		}
	}
	return nil
}
func (this *SettleDistressed) Validate() error {
	return nil
}
func (this *MarketTick) Validate() error {
	return nil
}
func (this *AuctionEvent) Validate() error {
	return nil
}
func (this *BusEvent) Validate() error {
	if oneOfNester, ok := this.GetEvent().(*BusEvent_TimeUpdate); ok {
		if oneOfNester.TimeUpdate != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.TimeUpdate); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("TimeUpdate", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_TransferResponses); ok {
		if oneOfNester.TransferResponses != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.TransferResponses); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("TransferResponses", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_PositionResolution); ok {
		if oneOfNester.PositionResolution != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.PositionResolution); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("PositionResolution", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_Order); ok {
		if oneOfNester.Order != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Order); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Order", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_Account); ok {
		if oneOfNester.Account != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Account); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Account", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_Party); ok {
		if oneOfNester.Party != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Party); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Party", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_Trade); ok {
		if oneOfNester.Trade != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Trade); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Trade", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_MarginLevels); ok {
		if oneOfNester.MarginLevels != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.MarginLevels); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MarginLevels", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_Proposal); ok {
		if oneOfNester.Proposal != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Proposal); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Proposal", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_Vote); ok {
		if oneOfNester.Vote != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Vote); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Vote", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_MarketData); ok {
		if oneOfNester.MarketData != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.MarketData); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MarketData", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_NodeSignature); ok {
		if oneOfNester.NodeSignature != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.NodeSignature); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("NodeSignature", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_LossSocialization); ok {
		if oneOfNester.LossSocialization != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.LossSocialization); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("LossSocialization", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_SettlePosition); ok {
		if oneOfNester.SettlePosition != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.SettlePosition); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("SettlePosition", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_SettleDistressed); ok {
		if oneOfNester.SettleDistressed != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.SettleDistressed); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("SettleDistressed", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_MarketCreated); ok {
		if oneOfNester.MarketCreated != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.MarketCreated); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MarketCreated", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_Asset); ok {
		if oneOfNester.Asset != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Asset); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Asset", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_MarketTick); ok {
		if oneOfNester.MarketTick != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.MarketTick); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MarketTick", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_Withdrawal); ok {
		if oneOfNester.Withdrawal != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Withdrawal); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Withdrawal", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_Deposit); ok {
		if oneOfNester.Deposit != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Deposit); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Deposit", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_Auction); ok {
		if oneOfNester.Auction != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Auction); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Auction", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_RiskFactor); ok {
		if oneOfNester.RiskFactor != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.RiskFactor); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("RiskFactor", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_NetworkParameter); ok {
		if oneOfNester.NetworkParameter != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.NetworkParameter); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("NetworkParameter", err)
			}
		}
	}
	if oneOfNester, ok := this.GetEvent().(*BusEvent_Market); ok {
		if oneOfNester.Market != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(oneOfNester.Market); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Market", err)
			}
		}
	}
	return nil
}