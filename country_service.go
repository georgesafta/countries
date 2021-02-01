package countries

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const baseURL = "https://restcountries.eu/rest/v2/%s"

// ByName calls the country API filtered by country name.
// Optionally, we can filter the fields by name.
// Returns the list of countries matching the filters.
func ByName(name string, fields ...string) ([]Country, error) {
	data, err := get(fmt.Sprintf("name/%s%s", name, filter(fields...)))
	if err != nil {
		return nil, err
	}

	var c []Country
	err = json.Unmarshal(data, &c)
	if err != nil {
		log.Println("Error deserializing data", err)
		return nil, err
	}

	return c, nil
}

// ByCapital calls the country API filtered by capital city name.
// Optionally, we can filter the fields by name.
// Returns the list of countries matching the filters.
func ByCapital(name string, fields ...string) ([]Country, error) {
	resData, err := get(fmt.Sprintf("capital/%s%s", name, filter(fields...)))
	if err != nil {
		return nil, err
	}

	var c []Country
	err = json.Unmarshal(resData, &c)
	if err != nil {
		log.Println("Error deserializing data", err)
		return nil, err
	}

	return c, nil
}

// All retrieves all the countries by calling the country API.
// Optionally, we can filter the fields by name.
// Returns the list of countries matching the filters.
func All(fields ...string) ([]Country, error) {
	resData, err := get(fmt.Sprintf("all%s", filter(fields...)))
	if err != nil {
		return nil, err
	}

	var c []Country
	err = json.Unmarshal(resData, &c)
	if err != nil {
		log.Println("Error deserializing data", err)
		return nil, err
	}

	return c, nil
}

func get(endpoint string) ([]byte, error) {
	url := fmt.Sprintf(baseURL, endpoint)
	res, err := http.Get(url)
	if err != nil {
		log.Println("Error calling the API", err)
		return []byte{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		e := fmt.Errorf("Unexpected API status code %s", res.Status)
		log.Println("Unsucessfull call", e)
		return []byte{}, e
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Error reading the response", err)
		return []byte{}, err
	}

	return body, nil
}

func filter(fields ...string) string {
	if fields == nil {
		return ""
	}

	sb := strings.Builder{}
	sb.WriteString("?fields=")
	for i := 0; i < len(fields); i++ {
		sb.WriteString(fields[i])
		if i != len(fields)-1 {
			sb.WriteString(";")
		}
	}

	return sb.String()
}
