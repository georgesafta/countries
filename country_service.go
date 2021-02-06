package countries

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// BaseURL is the base url of the countries API, use it when initialising the client.
const (
	BaseURL        = "https://restcountries.eu/rest/v2"
	queryDelimiter = "?"
	and            = "&"
	fieldsFilter   = "fields"
	codesFilter    = "codes"
)

// HTTPClient is a wrapper over http.Client.
type HTTPClient struct {
	Client  *http.Client
	baseURL string
}

// NewHTTPClient returns a new HTTPClient.
func NewHTTPClient(baseURL string) *HTTPClient {
	return &HTTPClient{
		Client:  &http.Client{},
		baseURL: baseURL,
	}
}

// ByName calls the country API filtered by country partial name or native name.
// Optionally, we can filter the fields by name.
// Returns the list of countries matching the filters.
func (c *HTTPClient) ByName(name string, fields ...string) ([]Country, error) {
	data, err := c.get(fmt.Sprintf("/name/%s%s", name, filter(queryDelimiter, fieldsFilter, fields...)))
	if err != nil {
		return nil, err
	}

	return unmarshal(data)
}

// ByFullName calls the country API filtered by country full name.
// Optionally, we can filter the fields by name.
// Returns the list of countries matching the filters.
func (c *HTTPClient) ByFullName(name string, fields ...string) ([]Country, error) {
	data, err := c.get(fmt.Sprintf("/name/%s?fullText=true%s", name, filter(and, fieldsFilter, fields...)))
	if err != nil {
		return nil, err
	}

	return unmarshal(data)
}

// ByCode calls the country API filtered by country ISO 3166 code.
// Optionally, we can filter the fields by name.
// Returns the list of countries matching the filters.
func (c *HTTPClient) ByCode(code string, fields ...string) ([]Country, error) {
	data, err := c.get(fmt.Sprintf("/alpha/%s%s", code, filter(queryDelimiter, fieldsFilter, fields...)))
	if err != nil {
		return nil, err
	}

	return unmarshal(data)
}

// ByCodes calls the country API filtered by country ISO 3166 codes.
// Optionally, we can filter the fields by name.
// Returns the list of countries matching the filters.
func (c *HTTPClient) ByCodes(codes []string, fields ...string) ([]Country, error) {
	if len(codes) == 0 {
		e := fmt.Errorf("Empty list of codes")
		log.Println(e)
		return nil, e
	}
	data, err := c.get(fmt.Sprintf("/alpha%s%s", filter(queryDelimiter, codesFilter, codes...), filter(and, fieldsFilter, fields...)))
	if err != nil {
		return nil, err
	}

	return unmarshal(data)
}

// ByCapital calls the country API filtered by capital city name.
// Optionally, we can filter the fields by name.
// Returns the list of countries matching the filters.
func (c *HTTPClient) ByCapital(name string, fields ...string) ([]Country, error) {
	data, err := c.get(fmt.Sprintf("/capital/%s%s", name, filter(queryDelimiter, fieldsFilter, fields...)))
	if err != nil {
		return nil, err
	}

	return unmarshal(data)
}

// All retrieves all the countries by calling the country API.
// Optionally, we can filter the fields by name.
// Returns the list of countries matching the filters.
func (c *HTTPClient) All(fields ...string) ([]Country, error) {
	data, err := c.get(fmt.Sprintf("/all%s", filter(queryDelimiter, fieldsFilter, fields...)))
	if err != nil {
		return nil, err
	}

	return unmarshal(data)
}

// ByCurrency calls the country API filtered by ISO 4217 currency code.
// Optionally, we can filter the fields by name.
// Returns the list of countries matching the filters.
func (c *HTTPClient) ByCurrency(currency string, fields ...string) ([]Country, error) {
	data, err := c.get(fmt.Sprintf("/currency/%s%s", currency, filter(queryDelimiter, fieldsFilter, fields...)))
	if err != nil {
		return nil, err
	}

	return unmarshal(data)
}

// ByLanguage calls the country API filtered by ISO 639-1 language code.
// Optionally, we can filter the fields by name.
// Returns the list of countries matching the filters.
func (c *HTTPClient) ByLanguage(language string, fields ...string) ([]Country, error) {
	data, err := c.get(fmt.Sprintf("/lang/%s%s", language, filter(queryDelimiter, fieldsFilter, fields...)))
	if err != nil {
		return nil, err
	}

	return unmarshal(data)
}

// ByCallingCode calls the country API filtered by calling code.
// Optionally, we can filter the fields by name.
// Returns the list of countries matching the filters.
func (c *HTTPClient) ByCallingCode(callingCode string, fields ...string) ([]Country, error) {
	data, err := c.get(fmt.Sprintf("/callingcode/%s%s", callingCode, filter(queryDelimiter, fieldsFilter, fields...)))
	if err != nil {
		return nil, err
	}

	return unmarshal(data)
}

// ByRegion calls the country API filtered by region.
// Optionally, we can filter the fields by name.
// Returns the list of countries matching the filters.
func (c *HTTPClient) ByRegion(region string, fields ...string) ([]Country, error) {
	data, err := c.get(fmt.Sprintf("/region/%s%s", region, filter(queryDelimiter, fieldsFilter, fields...)))
	if err != nil {
		return nil, err
	}

	return unmarshal(data)
}

// ByRegionalBloc calls the country API filtered by regional bloc.
// Optionally, we can filter the fields by name.
// Returns the list of countries matching the filters.
func (c *HTTPClient) ByRegionalBloc(regionalBloc string, fields ...string) ([]Country, error) {
	data, err := c.get(fmt.Sprintf("/regionalbloc/%s%s", regionalBloc, filter(queryDelimiter, fieldsFilter, fields...)))
	if err != nil {
		return nil, err
	}

	return unmarshal(data)
}

func (c *HTTPClient) get(endpoint string) ([]byte, error) {
	url := fmt.Sprintf(c.baseURL+"%s", endpoint)
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

func unmarshal(data []byte) ([]Country, error) {
	var c []Country
	err := json.Unmarshal(data, &c)
	if err != nil {
		log.Println("Error deserializing data", err)
		return nil, err
	}

	return c, nil
}
