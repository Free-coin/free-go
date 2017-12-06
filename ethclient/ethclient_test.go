// Copyright 2016 The free-go Authors
// This file is part of the free-go library.
//
// The free-go library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The free-go library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the free-go library. If not, see <http://www.gnu.org/licenses/>.

package ethclient

import "github.com/freecoin/free-go"

// Verify that Client implements the freecoin interfaces.
var (
	_ = freecoin.ChainReader(&Client{})
	_ = freecoin.TransactionReader(&Client{})
	_ = freecoin.ChainStateReader(&Client{})
	_ = freecoin.ChainSyncReader(&Client{})
	_ = freecoin.ContractCaller(&Client{})
	_ = freecoin.GasEstimator(&Client{})
	_ = freecoin.GasPricer(&Client{})
	_ = freecoin.LogFilterer(&Client{})
	_ = freecoin.PendingStateReader(&Client{})
	// _ = freecoin.PendingStateEventer(&Client{})
	_ = freecoin.PendingContractCaller(&Client{})
)
