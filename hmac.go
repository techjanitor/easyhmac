package easyhmac

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
)

// SignedMessage contains a payload and a signature with the hmac secret
type SignedMessage struct {
	Payload   []byte `json:"p"`
	signature []byte `json:"s"`
	Secret    string
}

// Encode marshals the data to JSON and url-safe base64 encodes it
func (sm *SignedMessage) Encode() (message string, err error) {

	msg, err := json.Marshal(sm)
	if err != nil {
		return
	}

	message = base64.URLEncoding.EncodeToString(msg)

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

	mac := hmac.New(sha256.New, []byte(sm.Secret))
	mac.Write(sm.Payload)

	sm.signature = []byte(base64.StdEncoding.EncodeToString(mac.Sum(nil)))

}

// CheckSignature takes the base64 encoded message and signature
func (sm *SignedMessage) Verify() bool {

	mac := hmac.New(sha256.New, []byte(sm.Secret))
	mac.Write(sm.Payload)

	expected := []byte(base64.StdEncoding.EncodeToString(mac.Sum(nil)))

	return hmac.Equal(sm.signature, expected)

}
