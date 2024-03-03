package main

import (
    "encoding/json"
    "fmt"
    "net/http"
)

type CoinGeckoResponse struct {
    Bitcoin struct {
        USD float64 `json:"usd"`
    } `json:"bitcoin"`
    Ethereum struct {
        USD float64 `json:"usd"`
    } `json:"ethereum"`
    // Add more cryptocurrency fields as needed
}

func main() {
    // Create a new HTTP client
    client := http.Client{}

    // Specify the URL of the CoinGecko API
    url := "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin,ethereum&vs_currencies=usd"

    // Make a GET request to the API
    response, err := client.Get(url)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }

    defer response.Body.Close()

    // Decode the JSON response
    var data CoinGeckoResponse
    if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
        fmt.Println("Error decoding JSON:", err)
        return
    }

    // Print the cryptocurrency prices
    fmt.Println("Bitcoin Price:", data.Bitcoin.USD)
    fmt.Println("Ethereum Price:", data.Ethereum.USD)
}
