package cryptoHelper

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
)

//https://ipfans.github.io/2014/03/golang-rsa/

// loadPrivateKey loads an parses a PEM encoded private key file.
func loadPublicKey(key string) (Unsigner, error) {

	return parsePublicKey([]byte(key))
}

// parsePublicKey parses a PEM encoded private key.
func parsePublicKey(pemBytes []byte) (Unsigner, error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, errors.New("ssh: no key found")
	}

	var rawkey interface{}
	switch block.Type {
	case "PUBLIC KEY":
		rsa, err := x509.ParsePKIXPublicKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		rawkey = rsa
	default:
		return nil, fmt.Errorf("ssh: unsupported key type %q", block.Type)
	}

	return newUnsignerFromKey(rawkey)
}

// loadPrivateKey loads an parses a PEM encoded private key file.
func loadPrivateKey(key string) (Signer, error) {
	return parsePrivateKey([]byte(key))
}

// parsePublicKey parses a PEM encoded private key.
func parsePrivateKey(pemBytes []byte) (Signer, error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, errors.New("ssh: no key found")
	}

	var rawkey interface{}
	switch block.Type {
	case "RSA PRIVATE KEY":
		rsa, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			return nil, err
		}
		rawkey = rsa
	default:
		return nil, fmt.Errorf("ssh: unsupported key type %q", block.Type)
	}
	return newSignerFromKey(rawkey)
}

// A Signer is can create signatures that verify against a public key.
type Signer interface {
	// Sign returns raw signature for the given data. This method
	// will apply the hash specified for the keytype to the data.
	Sign(data []byte) ([]byte, error)
}

// A Signer is can create signatures that verify against a public key.
type Unsigner interface {
	// Sign returns raw signature for the given data. This method
	// will apply the hash specified for the keytype to the data.
	Unsign(data []byte, sig []byte) error
}

func newSignerFromKey(k interface{}) (Signer, error) {
	var sshKey Signer
	switch t := k.(type) {
	case *rsa.PrivateKey:
		sshKey = &rsaPrivateKey{t}
	default:
		return nil, fmt.Errorf("ssh: unsupported key type %T", k)
	}
	return sshKey, nil
}

func newUnsignerFromKey(k interface{}) (Unsigner, error) {
	var sshKey Unsigner
	switch t := k.(type) {
	case *rsa.PublicKey:
		sshKey = &rsaPublicKey{t}
	default:
		return nil, fmt.Errorf("ssh: unsupported key type %T", k)
	}
	return sshKey, nil
}

type rsaPublicKey struct {
	*rsa.PublicKey
}

type rsaPrivateKey struct {
	*rsa.PrivateKey
}

// Sign signs data with rsa-sha1
func (r *rsaPrivateKey) Sign(data []byte) ([]byte, error) {
	h := sha1.New()
	h.Write(data)
	d := h.Sum(nil)
	return rsa.SignPKCS1v15(rand.Reader, r.PrivateKey, crypto.SHA1, d)
}

// Unsign verifies the message using a rsa-sha1 signature
func (r *rsaPublicKey) Unsign(message []byte, sig []byte) error {
	h := sha1.New()
	h.Write(message)
	d := h.Sum(nil)
	return rsa.VerifyPKCS1v15(r.PublicKey, crypto.SHA1, d, sig)
}

// func CheckPubKey(text string, signed string, pubKey string) bool {

// 	notifyData, berr := json.Marshal(text)
// 	if berr != nil {
// 		return false
// 	}
// 	fmt.Println(text, signed, pubKey)

// 	prePubKey := "-----BEGIN PUBLIC KEY-----\n"
// 	sufPubKey := "\n-----END PUBLIC KEY-----"

// 	parser, perr := loadPublicKey(prePubKey + pubKey + sufPubKey)
// 	if perr != nil {
// 		return false
// 		//fmt.Errorf("could not sign request: %v", err)
// 	}

// 	//strings.Replace(signed, "/", "\\/", -1)
// 	signData, nerr := base64.StdEncoding.DecodeString(signed)
// 	if nerr != nil {
// 		return false
// 	}

// 	err := parser.Unsign(notifyData, signData)
// 	if err != nil {
// 		return false
// 		//fmt.Errorf("could not sign request: %v", err)
// 	}
// 	return true
// 	//fmt.Printf("Unsign error: %v\n", err)
// }
