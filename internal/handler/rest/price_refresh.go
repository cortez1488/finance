package rest

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

const APIIURI = "https://www.binance.com/api/v3/ticker/price"

func (h *Handler) RefreshPrices() {
	resp, err := http.Get(APIIURI)
	if err != nil {
		log.Fatalln("Request to API doesn't work " + err.Error())
	}

	responseBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln("Can't get bytes from response, read bytes: " + err.Error())
	}
	resp.Body.Close()
	if err != nil {
		log.Fatalln("Can't get bytes from response, read bytes: " + err.Error())
	}

	var structResponse []RefreshDBDTO
	err = json.Unmarshal(responseBytes, &structResponse)
	if err != nil {
		log.Fatalln("Unmarshalling doesnt work" + err.Error())
	}

	err = h.priceRefreshService.RefreshPrices(&structResponse)
	if err != nil {
		log.Fatalln("h.priceRefreshService.RefreshPrices(): " + err.Error())
	}
}
