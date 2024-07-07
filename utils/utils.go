package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sort"

	"github.com/a-h/templ"
)

type CountryList struct {
	Name struct {
		Common string `json:"common"`
	} `json:"name"`
}

var Test *[]CountryList

func Render(w http.ResponseWriter, r *http.Request, component templ.Component) error {
	return component.Render(r.Context(), w)
}

func GetCountryData() {
	var (
		countryList []CountryList
	)
	url := "https://restcountries.com/v3.1/all"
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)

	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		fmt.Printf("Unexpected API status code %s\n", res.Status)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}

	err = json.Unmarshal(body, &countryList)

	if err != nil {
		fmt.Println(err)
	}

	sort.Slice(countryList, func(i, j int) bool {
		return countryList[i].Name.Common < countryList[j].Name.Common
	})

	Test = &countryList
}
