package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)



func TestKeypair_Sign_Verify_Success(t *testing.T) {
	privKey := GeneratePrivateKey()
	PublicKey := privKey.PublicKey()
	msg := []byte("merhaba dunya")

	sig, err := privKey.Sign(msg)
	assert.Nil(t, err)
	assert.True(t, sig.Verify(PublicKey, msg))
}


func TestKeypair_Sign_Verify_Fail(t *testing.T) {
	privKey := GeneratePrivateKey()
	PublicKey := privKey.PublicKey()
	msg := []byte("merhaba dunya")

	sig, err := privKey.Sign(msg)
	assert.Nil(t, err)

	otherPrivKey := GeneratePrivateKey()
	otherPublicKey := otherPrivKey.PublicKey()

	assert.False(t, sig.Verify(otherPublicKey, msg))
	assert.False(t, sig.Verify(PublicKey, []byte("xxxxxx")))
}
