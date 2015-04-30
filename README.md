# easyhmac
[![GoDoc](http://img.shields.io/badge/go-documentation-brightgreen.svg?style=flat-square)](https://godoc.org/github.com/techjanitor/easyhmac)
[![Build Status](https://travis-ci.org/techjanitor/easyhmac.svg)](https://travis-ci.org/techjanitor/easyhmac)

Sign arbitrary payloads with HMAC in url safe base64

```
import "github.com/techjanitor/easyhmac"
```

 Encode and sign message:
```
	// Initialize SignedMessage struct with secret
	key := easyhmac.SignedMessage{
		Secret: "secret",
	}

	// Add payload data
	key.Payload = "a cool message"
	
	// Create HMAC signature
	key.Sign()
	
	// Marshal message to JSON and encode in url-safe base64
	signedkey, err := key.Encode()
	if err != nil {
		return
	}

```

Decode and verify message:

```
	// Initialize SignedMessage struct with secret
	message := easyhmac.SignedMessage{
		Secret: "secret",
	}

	// Decode message
	err := message.Decode(ourmessage)
	if err != nil {
		return
	}

	// Verify signature, returns a bool (true if verified)
	check := message.Verify()
	if !check {
		return
	}
```
