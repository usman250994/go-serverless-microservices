package user

type userType string

const (
	userTypeAdmin   userType = "admin"
	userTypeRegular userType = "regular"
)

type profile struct {
	Id       string   `json:"id"`
	Name     string   `json:"name"`
	Email    string   `json:"email"`
	UserType userType `json:"type"`
	Rating   float64  `json:"rating"`
}
