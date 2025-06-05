package response

import (
	"time"
)



type UserDto struct {
    ID        uint      `json:"id"`
    Username  string    `json:"username"`
    Email     string    `json:"email"`
    CreatedAt time.Time `json:"createdAt"`
    IsVerified bool     `json:"isVerified"`
    Token string        `json:"token,omitempty"`
    Password string      `json:"password,omitempty"`
    Role     string   `json:"role,omitempty"`
}