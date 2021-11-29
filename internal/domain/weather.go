package domain

type Weather struct {
	Forecast struct {
		Forecastday []struct {
			Date      string   `json:"date"`
			Day       struct {
				Condition struct {
					Text string `json:"text"`
				} `json:"condition"`
			} `json:"day"`			
		} `json:"forecastday"`
	} `json:"forecast"`
}