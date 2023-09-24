package person

type Person struct {
	ID            int
	Count         int
	Name          string
	Surname       string
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
