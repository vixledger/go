package network

import (
	"testing"

	"github.com/vixledger/go/xdr"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHashTransaction(t *testing.T) {
	var txe xdr.TransactionEnvelope

	err := xdr.SafeUnmarshalBase64("AAAAAGL8HQvQkbK2HA3WVjRrKmjX00fG8sLI7m0ERwJW/AX3AAAACgAAAAAAAAABAAAAAAAAAAAAAAABAAAAAAAAAAAAAAAArqN6LeOagjxMaUP96Bzfs9e0corNZXzBWJkFoK7kvkwAAAAAO5rKAAAAAAAAAAABVvwF9wAAAEAKZ7IPj/46PuWU6ZOtyMosctNAkXRNX9WCAI5RnfRk+AyxDLoDZP/9l3NvsxQtWj9juQOuoBlFLnWu8intgxQA", &txe)

	require.NoError(t, err)

	expected := [32]uint8{
		0x43, 0xca, 0xd2, 0x9d, 0xf7, 0xad, 0x1a, 0x7e,
		0x98, 0xa5, 0x58, 0x3c, 0x50, 0x2b, 0xb4, 0x2b,
		0xa9, 0xbb, 0x2d, 0x7b, 0x1c, 0x15, 0x78, 0xb9,
		0x54, 0xcc, 0xcf, 0xec, 0xa1, 0x77, 0x6f, 0x4d}

	actual, err := HashTransaction(&txe.Tx, TestNetworkPassphrase)
	if assert.NoError(t, err) {
		assert.Equal(t, expected, actual)
	}

	// sadpath: empty passphrase
	_, err = HashTransaction(&txe.Tx, "")
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "empty network passphrase")
	}
}
