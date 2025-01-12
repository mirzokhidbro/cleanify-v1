package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	// Generate private key
	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		log.Fatal(err)
	}

	// Get public key
	publicKey := privateKey.PublicKey

	// Convert to base64
	privateKeyBytes := privateKey.D.Bytes()
	publicKeyBytes := elliptic.Marshal(elliptic.P256(), publicKey.X, publicKey.Y)

	privateKeyBase64 := base64.URLEncoding.EncodeToString(privateKeyBytes)
	publicKeyBase64 := base64.URLEncoding.EncodeToString(publicKeyBytes)

	fmt.Printf("VAPID Private Key: %s\n", privateKeyBase64)
	fmt.Printf("VAPID Public Key: %s\n", publicKeyBase64)
}
