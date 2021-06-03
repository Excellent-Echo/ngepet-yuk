package entity

type DetailItem struct {
	ID            string  `json:"id"`
	CoinId        int     `json:"coin_id"`
	Name          string  `json:"name"`
	Symbol        string  `json:"symbol"`
	MarketCapRank int     `json:"market_cap_rank"`
	Thumb         string  `json:"thumb"`
	Small         string  `json:"small"`
	Large         string  `json:"large"`
	Slug          string  `json:"slug"`
	PriceBtc      float32 `json:"price_btc"`
	Score         int     `json:"score"`
}

type Item struct {
	Item DetailItem `json:"item"`
}

type Coins struct {
	Crypto []Item `json:"coins"`
}
