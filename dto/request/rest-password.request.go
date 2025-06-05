package request

type ResetPasswordRequest struct {
    NewPassword string `json:"newPassword" binding:"required,min=6"`
    ResetToken string `json:"resetToken" binding:"required"`
}