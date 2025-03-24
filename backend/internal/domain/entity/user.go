package entity

import "time"

type User struct {
	ID           int64     `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	UserType     string    `json:"user_type"`
	Name         string    `json:"name"`
	ProfileImage string    `json:"profile_image,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func NewUser(email, passwordHash, userTipe, name string) *User {
	now := time.Now()
	return &User{
		Email:        email,
		PasswordHash: passwordHash,
		UserType:     userTipe,
		Name:         name,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
}

func (u *User) IsStudent() bool {
	return u.UserType == "student"
}

func (u *User) IsCompany() bool {
	return u.UserType == "company"
}

func (u *User) UpdateProfileImage(name, image string) {
	u.Name = name
	u.ProfileImage = image
	u.UpdatedAt = time.Now()
}
