package shell

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	"golang.org/x/crypto/ssh"
)

// GenerateKeys generate SSH keys.
func GenerateKeys() (*string, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2014)
	if err != nil {
		return nil, err
	}

	privateKeyDer := x509.MarshalPKCS1PrivateKey(privateKey)
	privateKeyBlock := pem.Block{
		Type:    "RSA PRIVATE KEY",
		Headers: nil,
		Bytes:   privateKeyDer,
	}
	privateKeyPem := string(pem.EncodeToMemory(&privateKeyBlock))

	publicKey := privateKey.PublicKey
	publicKeyDer, err := x509.MarshalPKIXPublicKey(&publicKey)
	if err != nil {
		return nil, err
	}

	publicKeyBlock := pem.Block{
		Type:    "PUBLIC KEY",
		Headers: nil,
		Bytes:   publicKeyDer,
	}
	publicKeyPem := string(pem.EncodeToMemory(&publicKeyBlock))

	fmt.Println(privateKeyPem)
	fmt.Println(publicKeyPem)

	return &privateKeyPem, nil
}

// MakeSSHKeyPair makes keys...
func MakeSSHKeyPair(pubKeyPath, privateKeyPath string) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, 1024)
	if err != nil {
		return err
	}

	// generate and write private key as PEM
	privateKeyFile, err := os.Create(privateKeyPath)
	defer privateKeyFile.Close()
	if err != nil {
		return err
	}
	privateKeyPEM := &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(privateKey)}
	if err := pem.Encode(privateKeyFile, privateKeyPEM); err != nil {
		return err
	}

	// generate and write public key
	_, err = ssh.NewPublicKey(&privateKey.PublicKey)
	if err != nil {
		return err
	}
	//return ioutil.WriteFile(pubKeyPath, ssh.MarshalAuthorizedKey(pub), 0655)

	fmt.Print(privateKey)

	return nil
}
