package crypto

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"ngepet-yuk/entity"
)

type Service interface {
	GetTopCoins() (entity.Coins, error)
}

type coinService struct {
}

func NewService() *coinService {
	return &coinService{}
}

func (s *coinService) GetTopCoins() (entity.Coins, error) {
	response, err := http.Get("https://api.coingecko.com/api/v3/search/trending")

	if err != nil {
		fmt.Print(err.Error())

	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	jsonData := responseData

	var coins entity.Coins
	err = json.Unmarshal(jsonData, &coins)

	if err != nil {
		fmt.Println(err.Error())
	}

	return coins, nil
}
