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

package steps

import (
	"fmt"

	"code.vegaprotocol.io/vega/core/integration/stubs"
	types "code.vegaprotocol.io/vega/protos/vega"
)

func TheGlobalInsuranceBalanceShouldBeForTheAsset(
	broker *stubs.BrokerStub,
	rawAmount, asset string,
) error {
	amount := parseExpectedInsurancePoolBalance(rawAmount)

	acc, err := broker.GetAssetGlobalInsuranceAccount(asset)
	if err != nil {
		return errCannotGetGlobalInsuranceAccountForAsset(asset, err)
	}

	if amount != stringToU64(acc.Balance) {
		return errInvalidAssetGlobalInsuranceBalance(amount, acc)
	}
	return nil
}

func errCannotGetGlobalInsuranceAccountForAsset(asset string, err error) error {
	return fmt.Errorf("couldn't get global insurance account for asset(%s): %s", asset, err.Error())
}

func errInvalidAssetGlobalInsuranceBalance(amount uint64, acc types.Account) error {
	return fmt.Errorf(
		"invalid balance for global insurance, expected %v, got %v",
		amount, acc.Balance,
	)
}