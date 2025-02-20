package shortener

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"

	"github.com/itchyny/base58-go"
)

func sha2560f(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

func base58encoding(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}

func GeneratorShortLink(initialLink string, userId string) string {
	urlHashBytes := sha2560f(initialLink + userId)
	generatedNumber := new(big.Int).SetBytes(urlHashBytes).Uint64()
	finalString := base58encoding([]byte(fmt.Sprintf("%d", generatedNumber)))
	return finalString[:8]

}
