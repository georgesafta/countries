package countries_test

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/georgesafta/countries"
)

var expectedFullResponse = []countries.Country{
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

var expectedFilteredResponse = []countries.Country{
	{
		Name:       "Colombia",
		Alpha3Code: "COL",
		Capital:    "Bogotá",
		Region:     "Americas",
		Subregion:  "South America",
		Population: 48759958,
		Demonym:    "Colombian",
		Area:       1141748.0,
		NativeName: "Colombia",
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
		FlagURL: "https://restcountries.eu/data/col.svg",
	},
}

var fullMockPath = "mock/full_data.json"
var partialMockPath = "mock/partial_data.json"

func TestByName(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(checkedHandler(t, fullMockPath, "/name/test")))
	defer ts.Close()

	client := countries.NewHTTPClient(ts.URL)
	resp, err := client.ByName("test")
	if err != nil {
		t.Fatal("Call unsuccessful")
	}
	if len(resp) != 1 {
		t.Fatal("Response should have size 1")
	}
	if !reflect.DeepEqual(expectedFullResponse, resp) {
		t.Fatalf("Response not matching, expected : %v, got : %v", expectedFullResponse, resp)
	}
}

func TestByNameFiltered(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(checkedHandler(t, partialMockPath, "/name/test?fields=name;capital;currencies")))
	defer ts.Close()

	client := countries.NewHTTPClient(ts.URL)
	resp, err := client.ByName("test", []string{"name", "capital", "currencies"}...)
	if err != nil {
		t.Fatal("Call unsuccessful")
	}
	if len(resp) != 1 {
		t.Fatal("Response should have size 1")
	}
	if !reflect.DeepEqual(expectedFilteredResponse, resp) {
		t.Fatalf("Response not matching, expected : %v, got : %v", expectedFilteredResponse, resp)
	}
}

func TestByFullName(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(checkedHandler(t, fullMockPath, "/name/test?fullText=true")))
	defer ts.Close()

	client := countries.NewHTTPClient(ts.URL)
	resp, err := client.ByFullName("test")
	if err != nil {
		t.Fatal("Call unsuccessful")
	}
	if len(resp) != 1 {
		t.Fatal("Response should have size 1")
	}
	if !reflect.DeepEqual(expectedFullResponse, resp) {
		t.Fatalf("Response not matching, expected : %v, got : %v", expectedFullResponse, resp)
	}
}

func TestByFullNameFiltered(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(checkedHandler(t, partialMockPath, "/name/test?fullText=true&fields=name;capital;currencies")))
	defer ts.Close()

	client := countries.NewHTTPClient(ts.URL)
	resp, err := client.ByFullName("test", []string{"name", "capital", "currencies"}...)
	if err != nil {
		t.Fatal("Call unsuccessful")
	}
	if len(resp) != 1 {
		t.Fatal("Response should have size 1")
	}
	if !reflect.DeepEqual(expectedFilteredResponse, resp) {
		t.Fatalf("Response not matching, expected : %v, got : %v", expectedFilteredResponse, resp)
	}
}

func TestByCode(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(checkedHandler(t, fullMockPath, "/alpha/test")))
	defer ts.Close()

	client := countries.NewHTTPClient(ts.URL)
	resp, err := client.ByCode("test")
	if err != nil {
		t.Fatal("Call unsuccessful")
	}
	if len(resp) != 1 {
		t.Fatal("Response should have size 1")
	}
	if !reflect.DeepEqual(expectedFullResponse, resp) {
		t.Fatalf("Response not matching, expected : %v, got : %v", expectedFullResponse, resp)
	}
}

func TestByCodeFiltered(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(checkedHandler(t, partialMockPath, "/alpha/test?fields=name;capital;currencies")))
	defer ts.Close()

	client := countries.NewHTTPClient(ts.URL)
	resp, err := client.ByCode("test", []string{"name", "capital", "currencies"}...)
	if err != nil {
		t.Fatal("Call unsuccessful")
	}
	if len(resp) != 1 {
		t.Fatal("Response should have size 1")
	}
	if !reflect.DeepEqual(expectedFilteredResponse, resp) {
		t.Fatalf("Response not matching, expected : %v, got : %v", expectedFilteredResponse, resp)
	}
}

func TestByCapital(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(checkedHandler(t, fullMockPath, "/capital/test")))
	defer ts.Close()

	client := countries.NewHTTPClient(ts.URL)
	resp, err := client.ByCapital("test")
	if err != nil {
		t.Fatal("Call unsuccessful")
	}
	if len(resp) != 1 {
		t.Fatal("Response should have size 1")
	}
	if !reflect.DeepEqual(expectedFullResponse, resp) {
		t.Fatalf("Response not matching, expected : %v, got : %v", expectedFullResponse, resp)
	}
}

func TestByCapitalFiltered(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(checkedHandler(t, partialMockPath, "/capital/test?fields=name;capital;currencies")))
	defer ts.Close()

	client := countries.NewHTTPClient(ts.URL)
	resp, err := client.ByCapital("test", []string{"name", "capital", "currencies"}...)
	if err != nil {
		t.Fatal("Call unsuccessful")
	}
	if len(resp) != 1 {
		t.Fatal("Response should have size 1")
	}
	if !reflect.DeepEqual(expectedFilteredResponse, resp) {
		t.Fatalf("Response not matching, expected : %v, got : %v", expectedFilteredResponse, resp)
	}
}

func checkedHandler(t *testing.T, filePath, expectedURL string) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			t.Fatalf("Expected GET call, but got: %s", r.Method)
		}

		query := ""
		if r.URL.RawQuery != "" {
			query += "?" + r.URL.RawQuery
		}
		if r.URL.Path+query != expectedURL {
			t.Fatalf("Expected call to url: %s, but got: %s", expectedURL, r.URL.Path)
		}

		data, err := ioutil.ReadFile(filePath)
		if err == nil {
			w.Write(data)
		}
	}
}
