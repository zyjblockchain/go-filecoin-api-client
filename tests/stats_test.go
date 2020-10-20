package tests

import (
	"context"
	"github.com/stretchr/testify/require"
	"github.com/zyjblockchain/go-filecoin-api-client/api"
	"testing"
)

var StatsAPI = api.NewAPI("").Stats()

func TestStatsBandwidth(t *testing.T) {
	resp, err := StatsAPI.BandWidth(context.Background())
	require.NoError(t, err)
	t.Log("View bandwidth usage metrics:", resp)
}
