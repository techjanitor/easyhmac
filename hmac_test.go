package easyhmac

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	testmessage = "eyJwIjoiZEdWemRDQnRaWE56WVdkbCIsInMiOiJZbTVpYVUxUmRrNURkVEUzY0U4clFVRkVaV1Z5Wnk5bmFYZFZjV3RPTUVweVFYTlRkRWRpUVUxcVVUMD0ifQ=="
	Secret      = "test"
)

func TestEncode(t *testing.T) {

	// Initialize SignedMessage struct with secret
	key := SignedMessage{}

	// Add payload data
	key.Payload = []byte("test message")

	// Create HMAC signature
	key.Sign()

	// Marshal message to JSON and encode in url-safe base64
	signedkey, err := key.Encode()

	assert.NoError(t, err, "should be no error")

	assert.Equal(t, testmessage, signedkey, "they should be equal")

}

func TestDecode(t *testing.T) {

	// Initialize SignedMessage struct with secret
	key := SignedMessage{}

	// Decode message
	err := message.Decode(testmessage)

	assert.NoError(t, err, "should be no error")

	check := message.Verify()

	assert.True(t, check, "should be true")

	assert.Equal(t, "test message", key.Payload, "they should be equal")

}
