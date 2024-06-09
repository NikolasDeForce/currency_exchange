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
	Dates []exchange.CSVData
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
	tmpl, _ := template.ParseFiles("./templates/index.html")

	dates := exchange.ReadCSV("exchange.csv")

	data := ViewData{
		Dates: dates,
	}
	exchange.WriteCSV("exchange.csv")

	err := tmpl.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	}
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	dates := exchange.ReadCSV("exchange.csv")

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
