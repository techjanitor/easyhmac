package easyhmac

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
)

var MaxSize = 1024
var Secret = "CHANGEME"

// SignedMessage contains a payload and a signature with the hmac secret
type SignedMessage struct {
	Payload   []byte `json:"p"`
	Signature []byte `json:"s"`
}

// Encode marshals the data to JSON and url-safe base64 encodes it
func (sm *SignedMessage) Encode() (message string, err error) {

	msg, err := json.Marshal(sm)
	if err != nil {
		return
	}

	message = base64.URLEncoding.EncodeToString(msg)

	if len(message) > MaxSize {
		err = fmt.Errorf("Message exceeds %d bytes", MaxSize)
		return
	}

	return

}

// Decode will unencode the url-safe base64 message and unmarshal to JSON
func (sm *SignedMessage) Decode(message string) (err error) {

	// Decode message
	msg, err := base64.URLEncoding.DecodeString(message)
	if err != nil {
		return
	}

	// Unmarshal JSON into struct
	err = json.Unmarshal(msg, sm)
	if err != nil {
		return
	}

	return

}

// Sign creates a HMAC signature for the message
func (sm *SignedMessage) Sign() {

	mac := hmac.New(sha256.New, []byte(Secret))
	mac.Write(sm.Payload)

	sm.Signature = []byte(base64.StdEncoding.EncodeToString(mac.Sum(nil)))

}

// CheckSignature takes the base64 encoded message and signature
func (sm *SignedMessage) Verify() bool {

	mac := hmac.New(sha256.New, []byte(Secret))
	mac.Write(sm.Payload)

	expected := []byte(base64.StdEncoding.EncodeToString(mac.Sum(nil)))

	return hmac.Equal(sm.Signature, expected)

}
