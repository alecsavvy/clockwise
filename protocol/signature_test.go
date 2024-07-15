package protocol_test

import (
	"testing"
)

func TestSignature(t *testing.T) {
	// inputMsg := &gen.CreateUserV2{
	// 	Handle: "LemonadeJetpack",
	// 	Bio:    "to the moon",
	// }

	// signer, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	// require.Nil(t, err)

	// envelope, err := protocol.Sign(signer, inputMsg)
	// require.Nil(t, err)

	// var outputMsg gen.CreateUserV2
	// recoveredSigner, err := protocol.Recover(envelope, &outputMsg)
	// require.Nil(t, err)

	// expectedSigner := crypto.PubkeyToAddress(signer.PublicKey)
	// require.EqualValues(t, expectedSigner, *recoveredSigner)

	// require.True(t, proto.Equal(inputMsg, &outputMsg), "expected and actual message not equal")
}
