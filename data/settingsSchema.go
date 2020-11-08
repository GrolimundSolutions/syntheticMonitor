package data

// SyntheticSettings obj. with the Settings and a list of URS they you want to call
type SyntheticSettings struct {
	Location      string          `json:"Location"`
	SyntheticUrls []SyntheticUrls `json:"SyntheticUrls"`
}

// SyntheticUrls is an Object with som test information about the test
type SyntheticUrls struct {
	URL    string `json:"URL"`
	Name   string `json:"Name"`
	Expect string `json:"Expect"`
}
