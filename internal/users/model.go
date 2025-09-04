package user

type userType string

const (
	userTypeAdmin   userType = "rentee"
	userTypeRegular userType = "renter"
)

type profile struct {
	Id       string   `json:"id"`
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	UserType userType `json:"type"`
	Rating   float64  `json:"rating"`
}
