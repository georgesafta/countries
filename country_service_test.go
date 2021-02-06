package countries_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/georgesafta/countries"
)

func TestByName(t *testing.T) {
	expectedResponse := []countries.Country{
		{
			Name:              "Colombia",
			TopLevelDomain:    []string{".co"},
			Alpha2Code:        "CO",
			Alpha3Code:        "COL",
			CallingCodes:      []string{"57"},
			Capital:           "Bogotá",
			AltSpellings:      []string{"CO", "Republic of Colombia", "República de Colombia"},
			Region:            "Americas",
			Subregion:         "South America",
			Population:        48759958,
			LatitudeLongitude: []float32{4.0, -72.0},
			Demonym:           "Colombian",
			Area:              1141748.0,
			Gini:              55.9,
			Timezones:         []string{"UTC-05:00"},
			Borders:           []string{"BRA", "ECU", "PAN", "PER", "VEN"},
			NativeName:        "Colombia",
			NumericCode:       "170",
			Currencies: []countries.Currency{
				{
					Code:   "COP",
					Name:   "Colombian peso",
					Symbol: "$",
				},
			},
			Languages: []countries.Language{
				{
					Iso6391:    "es",
					Iso6392:    "spa",
					Name:       "Spanish",
					NativeName: "Español",
				},
			},
			Translations: map[string]string{
				"de": "Kolumbien",
				"es": "Colombia",
				"fr": "Colombie",
				"ja": "コロンビア",
				"it": "Colombia",
				"br": "Colômbia",
				"pt": "Colômbia",
			},
			FlagURL: "https://restcountries.eu/data/col.svg",
			RegionalBlocs: []countries.RegionalBloc{
				{
					Acronym:       "PA",
					Name:          "Pacific Alliance",
					OtherAcronyms: []string{},
					OtherNames:    []string{"Alianza del Pacífico"},
				},
				{
					Acronym:       "USAN",
					Name:          "Union of South American Nations",
					OtherAcronyms: []string{"UNASUR", "UNASUL", "UZAN"},
					OtherNames:    []string{"Unión de Naciones Suramericanas", "União de Nações Sul-Americanas", "Unie van Zuid-Amerikaanse Naties", "South American Union"},
				},
			},
			Cioc: "COL",
		},
	}

	ts := httptest.NewServer(http.HandlerFunc(handleGetCountries))
	defer ts.Close()

	client := countries.NewHTTPClient(ts.URL)
	resp, err := client.ByName("test")
	if err != nil {
		t.Fatal("Call unsuccessful")
	}
	if len(resp) != 1 {
		t.Fatal("Response should have size 1")
	}
	if !reflect.DeepEqual(expectedResponse, resp) {
		t.Fatalf("Response not matching, expected : %v, got : %v", expectedResponse, resp)
	}
}

func handleGetCountries(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadFile("mock/data.json")
	if err == nil {
		w.Write(data)
	}
}
