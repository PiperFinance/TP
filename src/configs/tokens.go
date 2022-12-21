package configs

import (
	"TP/schema"
	"encoding/json"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

var (
	onceForChainTokens sync.Once
	// CD chain Tokens URL
	allTokensArray       = make([]schema.Token, 0)
	allTokens            = make(schema.TokenMapping)
	chainTokens          = make(map[schema.ChainId]schema.TokenMapping)
	NULL_TOKEN_ADDRESS   = common.HexToAddress("0x0000000000000000000000000000000000000000")
	NATIVE_TOKEN_ADDRESS = common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE")
	tokensUrl            = "https://github.com/PiperFinance/CD/blob/main/tokens/outVerified/all_tokens.json?raw=true"
	tokensDir            = "data/all_tokens.json"
)

func init() {

	onceForChainTokens.Do(func() {
		// Load Tokens ...
		var byteValue []byte
		if _, err := os.Stat(tokensDir); errors.Is(err, os.ErrNotExist) {
			resp, err := http.Get(tokensUrl)
			if err != nil {
				log.Fatalln(err)
			}
			byteValue, err = ioutil.ReadAll(resp.Body)
			if err != nil {
				log.Fatalf("HTTPTokenLoader: %s", err)
			}
		} else {
			jsonFile, err := os.Open(tokensDir)
			defer func(jsonFile *os.File) {
				err := jsonFile.Close()
				if err != nil {
					log.Error(err)
				}
			}(jsonFile)
			if err != nil {
				log.Fatalf("JSONTokenLoader: %s", err)
			}
			byteValue, err = ioutil.ReadAll(jsonFile)
			if err != nil {
				log.Fatalf("JSONTokenLoader: %s", err)
			}
		}
		err := json.Unmarshal(byteValue, &allTokens)
		if err != nil {
			log.Fatalf("TokenLoader: %s", err)
		}
		for tokenId, token := range allTokens {
			chainId := token.Detail.ChainId
			if chainTokens[chainId] == nil {
				chainTokens[chainId] = make(schema.TokenMapping)
			}
			chainTokens[chainId][tokenId] = token
			allTokensArray = append(allTokensArray, token)
		}
	})

}

func GetToken(id schema.TokenId) schema.Token{
	return allTokens[id]
}

func AllChainsTokens() schema.TokenMapping {
	return allTokens
}
func AllChainsTokensArray() []schema.Token {
	return allTokensArray
}
func ChainTokens(id schema.ChainId) schema.TokenMapping {
	return chainTokens[id]
}

//func ChainTokensArray(id schema.ChainId) []schema.Token {
//	return chainTokens[id]
//}
