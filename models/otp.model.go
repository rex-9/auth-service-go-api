package models

import (
    "time"
    "gorm.io/gorm"
)

type OTP struct {
    gorm.Model
    OTP       uint   `gorm:"not null"`
    Email     string `gorm:"not null"`
    ExpiresAt time.Time `gorm:"not null"`
}

// IsExpired checks if the OTP has expired (after 5 minutes)
func (otp *OTP) IsExpired() bool {
    return time.Now().After(otp.ExpiresAt)
}

// BeforeCreate will set the ExpiresAt to 5 minutes from creation
func (otp *OTP) BeforeCreate(tx *gorm.DB) error {
    otp.ExpiresAt = time.Now().Add(5 * time.Minute)
    return nil
}