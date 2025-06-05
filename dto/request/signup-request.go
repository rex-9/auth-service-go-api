package request

type SignUpRequest struct {
    Password string `json:"password" binding:"required,min=8"`
    Email    string `json:"email" binding:"required,email"`
    UserName string `json:"userName" binding:"required"`
}