package person

type Person struct {
	ID            int
	Count         int
	Name          string `validate:"required"`
	Surname       string `validate:"required"`
	Patronymic    string
	Age           int
	Gender        string
	OriginCountry string
	Country       []countryResp
}

type People []Person

type countryResp struct {
	Country_id  string
	Probability float64
}
