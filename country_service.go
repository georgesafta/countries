package countries

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	baseURL        = "https://restcountries.eu/rest/v2/%s"
	queryDelimiter = "?"
	and            = "&"
	fieldsFilter   = "fields"
	codesFilter    = "codes"
)

// ByName calls the country API filtered by country partial name or native name.
// Optionally, we can filter the fields by name.
// Returns the list of countries matching the filters.
func ByName(name string, fields ...string) ([]Country, error) {
	data, err := get(fmt.Sprintf("name/%s%s", name, filter(queryDelimiter, fieldsFilter, fields...)))
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

// ByFullName calls the country API filtered by country full name.
// Optionally, we can filter the fields by name.
// Returns the list of countries matching the filters.
func ByFullName(name string, fields ...string) ([]Country, error) {
	data, err := get(fmt.Sprintf("name/%s?fullText=true%s", name, filter(and, fieldsFilter, fields...)))
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

// ByCode calls the country API filtered by country ISO 3166 code.
// Optionally, we can filter the fields by name.
// Returns the list of countries matching the filters.
func ByCode(code string, fields ...string) ([]Country, error) {
	data, err := get(fmt.Sprintf("alpha/%s%s", code, filter(queryDelimiter, fieldsFilter, fields...)))
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

// ByCodes calls the country API filtered by country ISO 3166 codes.
// Optionally, we can filter the fields by name.
// Returns the list of countries matching the filters.
func ByCodes(codes []string, fields ...string) ([]Country, error) {
	if len(codes) == 0 {
		e := fmt.Errorf("Empty list of codes")
		log.Println(e)
		return nil, e
	}
	data, err := get(fmt.Sprintf("alpha%s%s", filter(queryDelimiter, codesFilter, codes...), filter(and, fieldsFilter, fields...)))
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
	resData, err := get(fmt.Sprintf("capital/%s%s", name, filter(queryDelimiter, fieldsFilter, fields...)))
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
	resData, err := get(fmt.Sprintf("all%s", filter(queryDelimiter, fieldsFilter, fields...)))
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

// ByCurrency calls the country API filtered by ISO 4217 currency code.
// Optionally, we can filter the fields by name.
// Returns the list of countries matching the filters.
func ByCurrency(currency string, fields ...string) ([]Country, error) {
	resData, err := get(fmt.Sprintf("currency/%s%s", currency, filter(queryDelimiter, fieldsFilter, fields...)))
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

// ByLanguage calls the country API filtered by ISO 639-1 language code.
// Optionally, we can filter the fields by name.
// Returns the list of countries matching the filters.
func ByLanguage(language string, fields ...string) ([]Country, error) {
	resData, err := get(fmt.Sprintf("lang/%s%s", language, filter(queryDelimiter, fieldsFilter, fields...)))
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

// ByCallingCode calls the country API filtered by calling code.
// Optionally, we can filter the fields by name.
// Returns the list of countries matching the filters.
func ByCallingCode(callingCode string, fields ...string) ([]Country, error) {
	resData, err := get(fmt.Sprintf("callingcode/%s%s", callingCode, filter(queryDelimiter, fieldsFilter, fields...)))
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

// ByRegion calls the country API filtered by region.
// Optionally, we can filter the fields by name.
// Returns the list of countries matching the filters.
func ByRegion(region string, fields ...string) ([]Country, error) {
	resData, err := get(fmt.Sprintf("region/%s%s", region, filter(queryDelimiter, fieldsFilter, fields...)))
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

// ByRegionalBloc calls the country API filtered by regional bloc.
// Optionally, we can filter the fields by name.
// Returns the list of countries matching the filters.
func ByRegionalBloc(regionalBloc string, fields ...string) ([]Country, error) {
	resData, err := get(fmt.Sprintf("regionalbloc/%s%s", regionalBloc, filter(queryDelimiter, fieldsFilter, fields...)))
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

func filter(prefix, fieldName string, fields ...string) string {
	if fields == nil {
		return ""
	}

	sb := strings.Builder{}
	sb.WriteString(prefix)
	sb.WriteString(fieldName)
	sb.WriteString("=")
	for i := 0; i < len(fields); i++ {
		sb.WriteString(fields[i])
		if i != len(fields)-1 {
			sb.WriteString(";")
		}
	}

	return sb.String()
}
