package easyhmac

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEncode(t *testing.T) {
	easyhmac.Secret = "test"

	// Initialize SignedMessage struct with secret
	key := easyhmac.SignedMessage{}

	// Add payload data
	key.Payload = "test message"

	// Create HMAC signature
	key.Sign()

	// Marshal message to JSON and encode in url-safe base64
	signedkey, err := key.Encode()

	assert.NoError(t, err, "should be no error")

	assert.Equal(t, "eyJwIjoiZEdWemRDQnRaWE56WVdkbCIsInMiOiJZbTVpYVUxUmRrNURkVEUzY0U4clFVRkVaV1Z5Wnk5bmFYZFZjV3RPTUVweVFYTlRkRWRpUVUxcVVUMD0ifQ==", signedkey, "they should be equal")

}
