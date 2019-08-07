package tokens

import (
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
	"log"
	"os"
)

const (
	privKeyPath = "./pkg/tokens/keys/app.rsa"
	pubKeyPath  = "./pkg/tokens/keys/app.rsa.pub"
)

func init() {

	file, err := os.OpenFile("petbook.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open log file", err)
	}
	Logger = log.New(file, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)

	signBytes, err := ioutil.ReadFile(privKeyPath)
	if err != nil {
		Error(err, "Error occurred while trying to read private key from file.\n")
	}

	SignKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		Error(err, "Error occurred while trying to parse private key from file.\n")
	}

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	if err != nil {
		Error(err, "Error occurred while trying to read public key from file.\n")
	}

	VerifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		Error(err, "Error occurred while trying to read private key from file.\n")
	}
}
