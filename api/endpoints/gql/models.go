// Code generated by github.com/vektah/gqlgen, DO NOT EDIT.

package gql

import (
	fmt "fmt"
	io "io"
	strconv "strconv"
	msg "vega/msg"
)

type Market struct {
	Name   string          `json:"name"`
	Orders []msg.Order     `json:"orders"`
	Trades []msg.Trade     `json:"trades"`
	Depth  msg.MarketDepth `json:"depth"`
}
type Party struct {
	Name      string               `json:"name"`
	Orders    []msg.Order          `json:"orders"`
	Positions []msg.MarketPosition `json:"positions"`
}
type PreConsensus struct {
	Accepted bool `json:"accepted"`
}
type Vega struct {
	Markets []Market `json:"markets"`
	Parties []Party  `json:"parties"`
}

type OrderStatus string

const (
	// The order is active and not cancelled or expired, it could be unfilled, partially filled or fully filled.
	// Active does not necessarily mean it's still on the order book.
	OrderStatusActive    OrderStatus = "Active"    // The order is cancelled, the order could be partially filled or unfilled before it was cancelled. It is not possible to cancel an order with 0 remaining.
	OrderStatusCancelled OrderStatus = "Cancelled" // This order trades any amount and as much as possible and remains on the book until it either trades completely or is cancelled
	OrderStatusExpired   OrderStatus = "Expired"
)

func (e OrderStatus) IsValid() bool {
	switch e {
	case OrderStatusActive, OrderStatusCancelled, OrderStatusExpired:
		return true
	}
	return false
}

func (e OrderStatus) String() string {
	return string(e)
}

func (e *OrderStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderStatus", str)
	}
	return nil
}

func (e OrderStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type OrderType string

const (
	// The order either trades completely (remainingSize == 0 after adding) or not at all, does not remain on the book if it doesn't trade
	OrderTypeFok OrderType = "FOK" // The order trades any amount and as much as possible but does not remain on the book (whether it trades or not)
	OrderTypeEne OrderType = "ENE" // This order trades any amount and as much as possible and remains on the book until it either trades completely or is cancelled
	OrderTypeGtc OrderType = "GTC" // This order type trades any amount and as much as possible and remains on the book until they either trade completely, are cancelled, or expires at a set time
	// NOTE: this may in future be multiple types or have sub types for orders that provide different ways of specifying expiry
	OrderTypeGtt OrderType = "GTT"
)

func (e OrderType) IsValid() bool {
	switch e {
	case OrderTypeFok, OrderTypeEne, OrderTypeGtc, OrderTypeGtt:
		return true
	}
	return false
}

func (e OrderType) String() string {
	return string(e)
}

func (e *OrderType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderType", str)
	}
	return nil
}

func (e OrderType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Side string

const (
	SideBuy  Side = "Buy"
	SideSell Side = "Sell"
)

func (e Side) IsValid() bool {
	switch e {
	case SideBuy, SideSell:
		return true
	}
	return false
}

func (e Side) String() string {
	return string(e)
}

func (e *Side) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Side(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Side", str)
	}
	return nil
}

func (e Side) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ValueDirection string

const (
	ValueDirectionPositive ValueDirection = "Positive"
	ValueDirectionNegative ValueDirection = "Negative"
)

func (e ValueDirection) IsValid() bool {
	switch e {
	case ValueDirectionPositive, ValueDirectionNegative:
		return true
	}
	return false
}

func (e ValueDirection) String() string {
	return string(e)
}

func (e *ValueDirection) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ValueDirection(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ValueDirection", str)
	}
	return nil
}

func (e ValueDirection) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}