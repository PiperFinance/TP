package configs

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"

	"TP/schema"
)

var (
	CMCPlatformChainMap = make(map[int64]int64)
	CMCPlatforms        schema.Platforms
)

func LoadCMC() {
	var byteValue []byte
	if _, err := os.Stat(Config.CMCPlatformsDir); errors.Is(err, os.ErrNotExist) {
		resp, err := http.Get(Config.CMCPlatformsURL.String())
		if err != nil {
			Logger.Panicln(err)
		}
		byteValue, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			Logger.Panicf("HTTPCMCPlatformLoader: %s", err)
		}
	} else {
		jsonFile, err := os.Open(Config.CMCPlatformsDir)
		defer func(jsonFile *os.File) {
			err := jsonFile.Close()
			if err != nil {
				Logger.Error(err)
			}
		}(jsonFile)
		if err != nil {
			Logger.Panicf("JSON:CMCPlatformLoader %s", err)
		}
		byteValue, err = ioutil.ReadAll(jsonFile)
		if err != nil {
			Logger.Panicf("JSONTokenLoader: %s", err)
		}
	}
	err := json.Unmarshal(byteValue, &CGPlatforms)
	if err != nil {
		Logger.Panicf("CMCPlatformLoader: %s", err)
	}
	for _, platform := range CGPlatforms {
		CGPlatformChainMap[platform.ID] = platform.ChainID
	}
}
