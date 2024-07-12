package request

type SignUpRequest struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
	FullName string `form:"full_name" json:"full_name"`
}

type SignInRequest struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}
