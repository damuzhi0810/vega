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

package query

import (
	"errors"
	"fmt"

	apipb "code.vegaprotocol.io/vega/protos/vega/api/v1"

	"github.com/golang/protobuf/jsonpb"
)

type NetworkParametersCmd struct {
	NodeAddress string `default:"0.0.0.0:3002" description:"The address of the vega node to use" long:"node-address"`
}

func (opts *NetworkParametersCmd) Execute(params []string) error {
	if len(params) > 1 {
		return errors.New("only one network parameter key can be to be specified")
	}

	var key string
	if len(params) == 1 {
		key = params[0]
	}

	req := apipb.ListNetworkParametersRequest{
		NetworkParameterKey: key,
	}
	return getPrintNetworkParameters(opts.NodeAddress, &req)
}

func getPrintNetworkParameters(nodeAddress string, req *apipb.ListNetworkParametersRequest) error {
	clt, err := getClient(nodeAddress)
	if err != nil {
		return fmt.Errorf("could not connect to the vega node: %w", err)
	}

	ctx, cancel := timeoutContext()
	defer cancel()
	res, err := clt.ListNetworkParameters(ctx, req)
	if err != nil {
		return fmt.Errorf("error querying the vega node: %w", err)
	}

	m := jsonpb.Marshaler{
		Indent: "  ",
	}
	buf, err := m.MarshalToString(res)
	if err != nil {
		return fmt.Errorf("invalid response from vega node: %w", err)
	}

	fmt.Printf("%v", buf)

	return nil
}