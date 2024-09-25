package main

type HomeworldDto struct {
	Name    string `json:"name"`
	Climate string `json:"climat"`
	Terrain string `json:"terrain"`
}

type CharacterDto struct {
	Name      string       `json:"name"`
	Height    *int         `json:"height"`
	Mass      *float64     `json:"mass"`
	HairColor string       `json:"hair_color"`
	SkinColor string       `json:"skin_color"`
	EyeColor  string       `json:"eye_color"`
	BirthYear string       `json:"birth_year"`
	Gender    string       `json:"gender"`
	Homeworld HomeworldDto `json:"homeworld"`
}
