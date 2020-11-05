package data

type SyntheticSettings struct {
	Location      string          `json:"Location"`
	SyntheticUrls []SyntheticUrls `json:"SyntheticUrls"`
}
type SyntheticUrls struct {
	URL    string `json:"URL"`
	Name   string `json:"Name"`
	Expect string `json:"Expect"`
}
