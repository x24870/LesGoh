package main

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	// Generate a private key
	privateKey, err := ecdsa.GenerateKey(crypto.S256(), rand.Reader)
	if err != nil {
		log.Fatalf("Failed to generate private key: %v", err)
	}

	// Convert the private key to bytes
	privateKeyBytes := crypto.FromECDSA(privateKey)

	// Convert bytes to hex string for readability
	privateKeyHex := hex.EncodeToString(privateKeyBytes)

	// Derive the public key from the private key
	publicKey := privateKey.Public()

	// Convert the public key to bytes
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Failed to assert public key type")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	// Derive the Ethereum address from the public key
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	// Output
	fmt.Println("Private Key:", privateKeyHex)
	fmt.Println("Public Key:", hex.EncodeToString(publicKeyBytes))
	fmt.Println("Ethereum Address:", address.Hex())
}
