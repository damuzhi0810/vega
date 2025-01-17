// Copyright (C) 2023 Gobalsky Labs Limited
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package staking

import (
	"errors"
	"sort"
	"time"

	"code.vegaprotocol.io/vega/core/types"
	"code.vegaprotocol.io/vega/libs/num"
)

var (
	ErrEventAlreadyExists = errors.New("event already exists")
	ErrInvalidAmount      = errors.New("invalid amount")
	ErrInvalidEventKind   = errors.New("invalid event kind")
	ErrMissingEventID     = errors.New("missing event id")
	ErrMissingTimestamp   = errors.New("missing timestamp")
	ErrNegativeBalance    = errors.New("negative balance")
	ErrInvalidParty       = errors.New("invalid party")
)

type Account struct {
	Party   string
	Balance *num.Uint
	Events  []*types.StakeLinking
}

func NewStakingAccount(party string) *Account {
	return &Account{
		Party:   party,
		Balance: num.UintZero(),
		Events:  []*types.StakeLinking{},
	}
}

func (s *Account) validateEvent(evt *types.StakeLinking) error {
	if evt.Amount == nil || evt.Amount.IsZero() {
		return ErrInvalidAmount
	}
	if evt.Type != types.StakeLinkingTypeDeposited && evt.Type != types.StakeLinkingTypeRemoved {
		return ErrInvalidEventKind
	}
	if evt.TS <= 0 {
		return ErrMissingTimestamp
	}
	if len(evt.ID) <= 0 {
		return ErrMissingEventID
	}
	if evt.Party != s.Party {
		return ErrInvalidParty
	}

	for _, v := range s.Events {
		if evt.ID == v.ID {
			return ErrEventAlreadyExists
		}
	}

	return nil
}

// AddEvent will add a new event to the account.
func (s *Account) AddEvent(evt *types.StakeLinking) error {
	if err := s.validateEvent(evt); err != nil {
		return err
	}
	// save the new events
	s.insertSorted(evt)

	// now update the ongoing balance
	return s.computeOngoingBalance()
}

func (s *Account) GetAvailableBalance() *num.Uint {
	return s.Balance.Clone()
}

func (s *Account) GetAvailableBalanceAt(at time.Time) (*num.Uint, error) {
	atUnix := at.UnixNano()
	return s.calculateBalance(func(evt *types.StakeLinking) bool {
		return evt.TS <= atUnix
	})
}

// GetAvailableBalanceInRange could return a negative balance
// if some event are still expected to be received from the bridge.
func (s *Account) GetAvailableBalanceInRange(from, to time.Time) (*num.Uint, error) {
	// first compute the balance before the from time.
	balance, err := s.GetAvailableBalanceAt(from)
	if err != nil {
		return num.UintZero(), err
	}

	minBalance := balance.Clone()

	// now we have the balance at the from time.
	// we will want to check how much was added / removed
	// during the epoch, and make sure that the initial
	// balance is still covered
	var (
		fromUnix = from.UnixNano()
		toUnix   = to.UnixNano()
	)
	for i := 0; i < len(s.Events) && s.Events[i].TS <= toUnix; i++ {
		if s.Events[i].TS > fromUnix {
			evt := s.Events[i]
			switch evt.Type {
			case types.StakeLinkingTypeDeposited:
				balance.AddSum(evt.Amount)
			case types.StakeLinkingTypeRemoved:
				if balance.LT(evt.Amount) {
					return num.UintZero(), ErrNegativeBalance
				}
				balance.Sub(balance, evt.Amount)
				minBalance = num.Min(balance, minBalance)
			}
		}
	}

	return minBalance, nil
}

// computeOnGoingBalance can return only 1 error which would
// be ErrNegativeBalancem, while this sounds bad, it can happen
// because of event being processed out of order but we can't
// really prevent that, and would have to wait for the network
// to have seen all events before getting a positive balance.
func (s *Account) computeOngoingBalance() error {
	balance, err := s.calculateBalance(func(evt *types.StakeLinking) bool {
		return true
	})
	s.Balance.Set(balance)
	return err
}

func (s *Account) insertSorted(evt *types.StakeLinking) {
	s.Events = append(s.Events, evt)
	// sort anyway, but we would expect the events to come in a sorted manner
	sort.SliceStable(s.Events, func(i, j int) bool {
		// check if timestamps are the same
		if s.Events[i].TS == s.Events[j].TS {
			// now we want to put deposit first to avoid any remove
			// event before a withdraw
			if s.Events[i].Type == types.StakeLinkingTypeRemoved && s.Events[j].Type == types.StakeLinkingTypeDeposited {
				// we return false so they can switched
				return false
			}
			// any other case is find to be as they are
			return true
		}

		return s.Events[i].TS < s.Events[j].TS
	})
}

type timeFilter func(*types.StakeLinking) bool

func (s *Account) calculateBalance(f timeFilter) (*num.Uint, error) {
	balance := num.UintZero()
	for _, evt := range s.Events {
		if f(evt) {
			switch evt.Type {
			case types.StakeLinkingTypeDeposited:
				balance.Add(balance, evt.Amount)
			case types.StakeLinkingTypeRemoved:
				if balance.LT(evt.Amount) {
					return num.UintZero(), ErrNegativeBalance
				}
				balance.Sub(balance, evt.Amount)
			}
		}
	}
	return balance, nil
}
