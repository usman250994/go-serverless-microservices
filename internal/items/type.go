package item

type addProductReq struct {
	Name    string  `json:"name"`
	Details string  `json:"details"`
	Lat     float64 `json:"lat"`
	Lng     float64 `json:"lng"`
}

type ProductQuery struct {
	Name    string  `json:"name"`
	Details string  `json:"details"`
	Lat     float64 `json:"lat"`
	Lng     float64 `json:"lng"`
}
