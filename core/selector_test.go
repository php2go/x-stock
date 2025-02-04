package core

import (
	"testing"

	"github.com/axiaoxin-com/logging"
	"github.com/axiaoxin-com/x-stock/datacenter/eastmoney"
	"github.com/stretchr/testify/require"
)

func TestAutoFilterStocks(t *testing.T) {
	logging.SetLevel("error")
	checker := NewChecker(_ctx, DefaultCheckerOptions)
	s := NewSelector(_ctx, eastmoney.DefaultFilter, checker)
	_, err := s.AutoFilterStocks(_ctx)
	require.Nil(t, err)
	// t.Log(result)
}
