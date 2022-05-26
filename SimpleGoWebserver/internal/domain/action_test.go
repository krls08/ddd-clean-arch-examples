package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBTCAction(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name         string
		actionString string
		expectAction BTCAction
		err          error
	}{
		{
			name:         "valid buy action",
			actionString: "buy",
			expectAction: BuyBTCAction,
		},
		{
			name:         "valid sell action",
			actionString: "sell",
			expectAction: SellBTCAction,
		},
		{
			name:         "unknown action",
			actionString: "unknown",
			err:          ErrInvalidAction,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			action, err := NewBTCAction(tc.actionString)
			assert.Equal(t, tc.err, err)
			assert.Equal(t, tc.expectAction, action)
		})
	}
}

func TestUSDAction(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name         string
		actionString string
		expectAction USDAction
		err          error
	}{
		{
			name:         "valid deposit action",
			actionString: "deposit",
			expectAction: DepositUSDAction,
		},
		{
			name:         "valid withdraw action",
			actionString: "withdraw",
			expectAction: WithdrawUSDAction,
		},
		{
			name:         "unknown action",
			actionString: "unknown",
			err:          ErrInvalidAction,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			action, err := NewUSDAction(tc.actionString)

			assert.ErrorIs(t, err, tc.err)
			assert.Equal(t, tc.expectAction, action)
		})
	}
}
