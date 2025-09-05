package item

type product struct {
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	Details string  `json:"details"`
	Lat     float64 `json:"lat"`
	Lng     float64 `json:"lng"`
}
