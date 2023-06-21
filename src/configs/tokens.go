package configs

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/common"
	log "github.com/sirupsen/logrus"

	"TP/schema"
)

var (
	onceForChainTokens sync.Once
	// CD chain Tokens URL
	allTokensArray       = make([]schema.Token, 0)
	allTokens            = make(schema.TokenMapping)
	geckoTokens          = make(map[string]schema.Token)
	geckoTokenIds        = make(map[string]schema.TokenId)
	chainTokens          = make(map[schema.ChainId]schema.TokenMapping)
	NULL_TOKEN_ADDRESS   = common.HexToAddress("0x0000000000000000000000000000000000000000")
	NATIVE_TOKEN_ADDRESS = common.HexToAddress("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE")
	tokensUrl            = "https://raw.githubusercontent.com/PiperFinance/CD/main/tokens/outVerified/all_tokens.json"
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
			geckoTokens[token.Detail.CoingeckoId] = token
			geckoTokenIds[token.Detail.CoingeckoId] = tokenId
		}
	})
}

func GenTokenId(address common.Address, chainId schema.ChainId) schema.TokenId {
	return schema.TokenId(fmt.Sprintf("%s-%d", strings.ToLower(address.String()), chainId))
}

func GenTokenIdExtra(address string, chainId schema.ChainId) schema.TokenId {
	return schema.TokenId(fmt.Sprintf("%s-%d", strings.ToLower(address), chainId))
}

func GetToken(id schema.TokenId) schema.Token {
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

func GeckoIdToToken(geckoId string) schema.Token {
	return geckoTokens[geckoId]
}

func GeckoIdToTokenId(geckoId string) schema.TokenId {
	return geckoTokenIds[geckoId]
}

//func ChainTokensArray(id schema.ChainId) []schema.Token {
//	return chainTokens[id]
//}
