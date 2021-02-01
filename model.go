package countries

// Country contains all informations related to a country.
type Country struct {
	Name              string            `json:"name,omitempty"`
	Capital           string            `json:"capital,omitempty"`
	TopLevelDomain    []string          `json:"topLevelDomain,omitempty"`
	Alpha2Code        string            `json:"alpha2Code,omitEmpty"`
	Alpha3Code        string            `json:"alpha3Code,omitempty"`
	AltSpellings      []string          `json:"altSpelling,omitempty"`
	Region            string            `json:"region,omitempty"`
	Subregion         string            `json:"subregion,omitempty"`
	Population        int32             `json:"population"`
	LatitudeLongitude []float32         `json:"latlng,omitempty"`
	Demonym           string            `json:"denonym,omitempty"`
	Area              float32           `json:"area,omitempty"`
	Gini              float32           `json:"gini,omitempty"`
	Timezones         []string          `json:"timeszones,omitempty"`
	Borders           []string          `json:"borders,omitempty"`
	NativeName        string            `json:"nativeName,omitempty"`
	NumericCode       string            `json:"numericCode,omitempty"`
	Currencies        []Currency        `json:"currencies,omitempty"`
	Languages         []Language        `json:"languages,omitempty"`
	Translations      map[string]string `json:"translations,omitempty"`
	FlagURL           string            `json:"flag,omitempty"`
	RegionalBlocs     []RegionalBloc    `json:"regionalBlocs,omitempty"`
	Cioc              string            `json:"cioc,omitempty"`
}

// Currency contains all information related to currency.
type Currency struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

// Language contains data related to a language.
type Language struct {
	Iso6391    string `json:"iso639_1"`
	Iso6392    string `json:"iso639_2"`
	Name       string `json:"name"`
	NativeName string `json:"nativeName"`
}

// RegionalBloc contains the data of a country's regional bloc.
type RegionalBloc struct {
	Acronyn       string   `json:"acronym"`
	Name          string   `json:"name"`
	OtherAcronyms []string `json:"otherAcronyms"`
	OtherNames    []string `json:"otherNames"`
}
