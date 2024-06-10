package exchange

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/gocolly/colly/v2"
)

type CSVCryptoData struct {
	Name              string `csv:"Name"`
	Symbol            string `csv:"Symbol"`
	MarketCap         string `csv:"Market Cap (USD)"`
	Price             string `csv:"Price (USD)"`
	CirculatingSupply string `csv:"Circulating Supply (USD)"`
	Volume            string `csv:"Volume (24h)"`
	Changeh           string `csv:"Change (1h)"`
	Change24h         string `csv:"Change (24h)"`
	Changed           string `csv:"Change (7d)"`
}

func WriteCryptoExchange(fName string) {
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	writer.Write([]string{"Name", "Symbol", "Market Cap (USD)", "Price (USD)", "Circulating Supply (USD)", "Volume (24h)", "Change (1h)", "Change (24h)", "Change (7d)"})

	// Instantiate default collector
	c := colly.NewCollector()

	c.OnHTML("tbody tr", func(e *colly.HTMLElement) {
		writer.Write([]string{
			e.ChildText(".cmc-table__column-name"),
			e.ChildText(".cmc-table__cell--sort-by__symbol"),
			e.ChildText(".cmc-table__cell--sort-by__market-cap"),
			e.ChildText(".cmc-table__cell--sort-by__price"),
			e.ChildText(".cmc-table__cell--sort-by__circulating-supply"),
			e.ChildText(".cmc-table__cell--sort-by__volume-24-h"),
			e.ChildText(".cmc-table__cell--sort-by__percent-change-1-h"),
			e.ChildText(".cmc-table__cell--sort-by__percent-change-24-h"),
			e.ChildText(".cmc-table__cell--sort-by__percent-change-7-d"),
		})
	})

	c.Visit("https://coinmarketcap.com/all/views/all/")
}

func ReadCryptoCSV(fName string) []CSVCryptoData {
	file, err := os.Open(fName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	dates := []CSVCryptoData{}
	if err := gocsv.UnmarshalFile(file, &dates); err != nil {
		fmt.Println(err)
	}
	return dates
}
