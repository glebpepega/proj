package person

type Person struct {
	Count      int
	Name       string `validate:"required"`
	Surname    string `validate:"required"`
	Patronymic string
	Age        int
	Gender     string
	Country    []countryResp
	Error      string
}

type countryResp struct {
	Country_id  string
	Probability float64
}

//{"name":"Gleb","surname":"Lol"}
