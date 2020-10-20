package tests

import (
	"context"
	"ecoin-api-client/api"
	fopt "ecoin-api-client/options"
	"github.com/stretchr/testify/require"
	"testing"
)

var DealsAPI = api.NewAPI("").Deals()

func TestDealsList(t *testing.T) {
	resp, err := DealsAPI.List(context.Background())
	require.NoError(t, err)
	t.Log("deals list:", resp)
}

func TestDealsRedeem(t *testing.T) {
	resp, err := DealsAPI.Redeem(context.Background(), "dealidxxxxx", fopt.Deals.GasPrice("0.1"), fopt.Deals.GasLimit("0"))
	require.NoError(t, err)
	t.Log("redeem vouchers for FIL:", resp)
}

func TestDealsShow(t *testing.T) {
	resp, err := DealsAPI.Show(context.Background(), "dealidxxxxx")
	require.NoError(t, err)
	t.Log("deals detail info:", resp)
}
