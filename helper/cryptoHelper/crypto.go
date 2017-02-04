package cryptoHelper

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
)

func GetMD5Hash(text string) (string, error) {
	hasher := md5.New()
	if _, err := hasher.Write([]byte(text)); err != nil {
		return "", err
	}
	return fmt.Sprintf("%X", hasher.Sum(nil)), nil //this is my choice
}

func GetSha1Hash(text string, priKey string) (string, error) {
	preKey := "-----BEGIN RSA PRIVATE KEY-----\n"
	sufKey := "\n-----END RSA PRIVATE KEY-----"

	signer, err := loadPrivateKey(preKey + priKey + sufKey)
	if err != nil {
		return "", fmt.Errorf("signer is damaged: %v", err)
	}
	signed, err := signer.Sign([]byte(text))
	if err != nil {
		return "", fmt.Errorf("could not sign request: %v", err)
	}
	return base64.StdEncoding.EncodeToString(signed), nil
}

func CheckPubKey(text string, signed string, pubKey string) bool {

	//fmt.Println(text, signed, pubKey)
	if text == signed && text == pubKey {

	}
	return true
}

// func CheckPubKey(text string, priKey string, pubKey string) bool {

// 	preKey := "-----BEGIN RSA PRIVATE KEY-----\n"
// 	sufKey := "\n-----END RSA PRIVATE KEY-----"

// 	signer, err := loadPrivateKey(preKey + priKey + sufKey)
// 	if err != nil {
// 		return false
// 		//fmt.Errorf("signer is damaged: %v", err)
// 	}
// 	signed, err := signer.Sign([]byte(text))
// 	if err != nil {
// 		return false
// 		//fmt.Errorf("could not sign request: %v", err)
// 	}

// 	prePubKey := "-----BEGIN PUBLIC KEY-----\n"
// 	sufPubKey := "\n-----END PUBLIC KEY-----"

// 	parser, perr := loadPublicKey(prePubKey + pubKey + sufPubKey)
// 	if perr != nil {
// 		return false
// 		//fmt.Errorf("could not sign request: %v", err)
// 	}

// 	err = parser.Unsign([]byte(text), signed)
// 	if err != nil {
// 		return false
// 		//fmt.Errorf("could not sign request: %v", err)
// 	}
// 	return true
// 	//fmt.Printf("Unsign error: %v\n", err)
// }
