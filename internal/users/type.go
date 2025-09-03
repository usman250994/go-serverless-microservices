package user

type updateProfileReq struct {
	Email    string   `json:"email"`
	Name     string   `json:"name"`
	UserType userType `json:"userType"`
	Ratings  float64  `json:"ratings"`
}
