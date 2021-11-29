package domain

type CitiesCollection struct {
	Cities []City
}

type City struct {
	Name string `json:"name"`
}
