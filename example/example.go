package example

type Address struct {
	// Used in html
	City    string  `json:"city"`
	Number  float64 `json:"number"`
	Country string  `json:"country,omitempty"`
}

type PersonalInfo struct {
	Hobbies []string `json:"hobby"`
	PetName string   `json:"pet_name"`
}

type Person struct {
	Name         string       `json:"name"`
	PersonalInfo PersonalInfo `json:"personal_info"`
	Nicknames    []string     `json:"nicknames"`
	Addresses    []Address    `json:"addresses"`
}

