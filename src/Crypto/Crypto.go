package Crypto

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	//	"os"
)

type Crypto struct {
	PKey   *rsa.PrivateKey
	PubKey *rsa.PublicKey
}go run

func (C *Crypto) Decrypt(src string) string {
	hash := sha256.New()
	plainText, err := rsa.DecryptOAEP(hash, rand.Reader, C.PKey, []byte(src), []byte(""))

	if err != nil {
		fmt.Println(err)
	}
	return string(plainText)
}
func (C *Crypto) SignVerify(src string, signature string) bool {
	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto // for simple example
	message := []byte(src)
	PSSmessage := message
	newhash := crypto.SHA256
	pssh := newhash.New()
	pssh.Write(PSSmessage)
	hashed := pssh.Sum(nil)
	err := rsa.VerifyPSS(C.PubKey, newhash, hashed, []byte(signature), &opts)
	if err != nil {
		fmt.Println("Who are U? Verify Signature failed")
		return false
	} else {
		return true
	}
}
func (C *Crypto) Sign(src string) []byte {
	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto // for simple example
	message := []byte(src)
	PSSmessage := message
	newhash := crypto.SHA256
	pssh := newhash.New()
	pssh.Write(PSSmessage)
	hashed := pssh.Sum(nil)

	signature, err := rsa.SignPSS(rand.Reader, C.PKey, newhash, hashed, &opts)

	if err != nil {
		fmt.Println(err)
	}

	return signature
}
func (C *Crypto) Encrypt(src string) string {
	message := []byte(src)
	lab := []byte("")
	hash := sha256.New()
	ciphertext, err := rsa.EncryptOAEP(hash, rand.Reader, C.PubKey, message, lab)
	if err != nil {
		fmt.Println(err)
	}
	return string(ciphertext)

}
func (C *Crypto) GenerateKey() {
	C.PKey, _ = rsa.GenerateKey(rand.Reader, 2048)
}
