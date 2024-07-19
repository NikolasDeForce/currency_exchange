package handler

import (
	core "exchange/core"
	exchange "exchange/parse"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

type ViewData struct {
	DatesCurrency []exchange.CSVCurrencyData
	DatesCrypto   []exchange.CSVCryptoData
}

var c = core.CalcCode{
	Name:    "",
	Code:    "",
	Value:   "",
	Price:   0,
	Change:  "",
	Percent: "",
}

type CalcVal struct {
	Val float64
}

var v = CalcVal{
	Val: 0.0,
}

func MainHandler(w http.ResponseWriter, r *http.Request) {
	exchange.WriteCurrencyCSV("./data/exchange.csv")

	dates := exchange.ReadCurrencyCSV("./data/exchange.csv")

	tmpl, _ := template.ParseFiles("./templates/index.html")

	data := ViewData{
		DatesCurrency: dates,
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	}
}

func CryptoHandler(w http.ResponseWriter, r *http.Request) {
	exchange.WriteCryptoExchange("./data/crypto.csv")

	dates := exchange.ReadCryptoCSV("./data/crypto.csv")

	tmpl, _ := template.ParseFiles("./templates/crypto.html")

	data := ViewData{
		DatesCrypto: dates,
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	}
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	dates := exchange.ReadCurrencyCSV("./data/exchange.csv")

	v.Val, _ = strconv.ParseFloat(r.FormValue("Value"), 64)

	c.Code = r.FormValue("Code")
	for _, date := range dates {
		price, err := strconv.ParseFloat(date.Price, 64)
		if err != nil {
			fmt.Println(err)
		}
		price = price * v.Val
		if date.Code == c.Code {
			c = core.CalcCode{
				Name:    date.Name,
				Code:    date.Code,
				Value:   date.Value,
				Price:   price,
				Change:  date.Change,
				Percent: date.Percent,
			}
		}
	}

	tmpl, _ := template.ParseFiles("./templates/calc.html")

	err := tmpl.Execute(w, nil)
	if err != nil {
		fmt.Println(err)
	}
}

func OkHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	tmpl, _ := template.ParseFiles("./templates/ok.html")

	data := core.CalcCode{
		Name:    c.Name,
		Code:    c.Code,
		Value:   c.Value,
		Price:   c.Price,
		Change:  c.Change,
		Percent: c.Percent,
	}

	err := tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	}
}
