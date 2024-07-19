package exchange

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/gocolly/colly/v2"
)

type CSVCurrencyData struct {
	Name    string `csv:"Name"`
	Code    string `csv:"Code"`
	Value   string `csv:"Value"`
	Price   string `csv:"Price"`
	Change  string `csv:"Change"`
	Percent string `csv:"Percent"`
}

func WriteCurrencyCSV(fName string) {
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write CSV header
	writer.Write([]string{"Name", "Code", "Value", "Price", "Change", "Percent"})

	// Instantiate default collector
	c := colly.NewCollector()

	c.OnHTML("tr.A1NefxsU", func(e *colly.HTMLElement) {
		writer.Write([]string{
			e.ChildText("span.OYTwr2Ke"),
			e.ChildText("span.Vf1AWW7q"),
			e.ChildText("span.PjocFlvi"),
			e.ChildText("span._ZXx92_y"),
			e.ChildTexts("span.M6nt2YLN")[0],
			e.ChildTexts("span.M6nt2YLN")[1],
		})
	})

	c.Visit("https://finance.rambler.ru/currencies/")
}

func ReadCurrencyCSV(fName string) []CSVCurrencyData {
	file, err := os.Open(fName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	dates := []CSVCurrencyData{}
	if err := gocsv.UnmarshalFile(file, &dates); err != nil {
		fmt.Println(err)
	}
	return dates
}
