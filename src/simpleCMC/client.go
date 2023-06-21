package simplecmc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"TP/configs"
	"TP/schema"
)

type Client struct {
	apiKeyIndex   int
	ApiKeys       []string
	selectorMutex sync.Mutex
	Currency      string
}

func (c *Client) nextApiKey() string {
	c.selectorMutex.Lock()
	k := c.ApiKeys[c.apiKeyIndex]
	c.apiKeyIndex++
	if c.apiKeyIndex >= len(c.ApiKeys) {
		c.apiKeyIndex = 0
	}
	c.selectorMutex.Unlock()
	return k
}

func (c *Client) AllPrices() ([]Coin, schema.PriceResponse, error) {
	result := make([]Coin, 0)
	priceRes := make(schema.PriceResponse)
	start, limit := 1, 5000
	for {
		url := fmt.Sprintf("https://pro-api.coinmarketcap.com/v1/cryptocurrency/listings/latest?convert=%s&start=%d&limit=%d&", c.Currency, start, limit) + "aux=platform%2Ccmc_rank%2Cmax_supply%2Ctotal_supply%2Ctags%2Cdate_added"
		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Add("X-CMC_PRO_API_KEY", c.nextApiKey())
		fmt.Println(url)
		res, err := http.DefaultClient.Do(req)
		if err != nil {
			return nil, nil, err
		}
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		resp := CMCListResp{}
		if err := json.Unmarshal(body, &resp); err != nil {
			return nil, nil, err
		}
		if len(resp.Data) < limit {
			break
		} else {
			for _, coin := range resp.Data {
				if coin.Platform == nil {
					// NOTE - platform not found !
					continue
				}
				var ok bool
				chainId, ok := configs.CGPlatformChainMap[coin.Platform.ID]
				if !ok {
					// NOTE - not found !
					continue
				}
				key := configs.GenTokenIdExtra(coin.Platform.TokenAddress, schema.ChainId(chainId))
				quote, ok := coin.Quote[Currency(c.Currency)]
				if ok {
					result = append(result, coin)
					priceRes[key] = quote.Price
				}
			}
			start += limit

		}
	}
	return result, priceRes, nil
}
