package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	// the api isn't available and one might face panic
	resp, _ := http.Get("http://dev.markitondemand.com/Api/v2/Quote?symbol=googl")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	quote := new(QuoteResponse)
	xml.Unmarshal(body, &quote)

	fmt.Printf("%s: %.2f", quote.Name, quote.LastPrice)

}

type QuoteResponse struct {
	Status           string
	Name             string
	LastPrice        float32
	Change           float32
	ChangePercent    float32
	TimeStamp        string
	MSDate           float32
	MarketCap        int
	Volume           int
	ChangeYTD        float32
	ChangePercentYTD float32
	High             float32
	Low              float32
	Open             float32
}

/*
Step-1: Simple execution of single web service call
-----
package main

import (
	"net/http"
	"io/ioutil"
	"encoding/xml"
	"fmt"
)

func main() {

	resp, _ := http.Get("http://dev.markitondemand.com/Api/v2/Quote?symbol=googl")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	quote := new(QuoteResponse)
	xml.Unmarshal(body, &quote)

	fmt.Printf("%s: %.2f", quote.Name, quote.LastPrice)

}


type QuoteResponse struct {
	Status string
	Name string
	LastPrice float32
	Change float32
	ChangePercent float32
	TimeStamp string
	MSDate float32
	MarketCap int
	Volume int
	ChangeYTD float32
	ChangePercentYTD float32
	High float32
	Low float32
	Open float32
}
*/

/*
Step-2 Execution of multiple execution calls using sequential programming
-----
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	start := time.Now()

	stockSymbols := []string{
		"googl",
		"msft",
		"aapl",
		"bbry",
		"hpq",
		"vz",
		"t",
		"tmus",
		"s",
	}

	for _, symbol := range stockSymbols {
		resp, _ := http.Get("http://dev.markitondemand.com/Api/v2/Quote?symbol=" + symbol)
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)

		quote := new(QuoteResponse)
		xml.Unmarshal(body, &quote)

		fmt.Printf("%s: $%.2f\n", quote.Name, quote.LastPrice)
	}

	elapsed := time.Since(start)

	fmt.Printf("Execution Time: %s", elapsed)
}

type QuoteResponse struct {
	Status           string
	Name             string
	LastPrice        float32
	Change           float32
	ChangePercent    float32
	TimeStamp        string
	MSDate           float32
	MarketCap        int
	Volume           int
	ChangeYTD        float32
	ChangePercentYTD float32
	High             float32
	Low              float32
	Open             float32
}
*/

/*
Step-3 Execution of multiple asychronous web service calls using concurrent programming
-----
package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	start := time.Now()

	stockSymbols := []string{
		"googl",
		"msft",
		"aapl",
		"bbry",
		"hpq",
		"vz",
		"t",
		"tmus",
		"s",
	}

	numComplete := 0

	for _, symbol := range stockSymbols {
		go func(symbol string){
			resp, _ := http.Get("http://dev.markitondemand.com/Api/v2/Quote?symbol=" + symbol)
			defer resp.Body.Close()
			body, _ := ioutil.ReadAll(resp.Body)

			quote := new(QuoteResponse)
			xml.Unmarshal(body, &quote)

			fmt.Printf("%s: $%.2f\n", quote.Name, quote.LastPrice)
			numComplete++
		}(symbol)
	}

	for numComplete < len(stockSymbols) {
		time.Sleep(10 * time.Millisecond)
	}
	elapsed := time.Since(start)

	fmt.Printf("Execution Time: %s", elapsed)
}

type QuoteResponse struct {
	Status           string
	Name             string
	LastPrice        float32
	Change           float32
	ChangePercent    float32
	TimeStamp        string
	MSDate           float32
	MarketCap        int
	Volume           int
	ChangeYTD        float32
	ChangePercentYTD float32
	High             float32
	Low              float32
	Open             float32
}
*/
