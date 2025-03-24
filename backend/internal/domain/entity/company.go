package entity

import "time"

type Company struct {
	ID          int64     `json:"id"`
	UserID      int64     `json:"user_id"`
	Description string    `json:"description"`
	Industry    string    `json:"industry"`
	Location    string    `json:"location"`
	Website     string    `json:"website"`
	Logo        string    `json:"logo,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	User        *User     `json:"user,omitempty"` // ユーザー情報への参照（必要に応じて）
}

func NewCompany(userID int64, description, industry, location, website string) *Company {
	now := time.Now()
	return &Company{
		UserID:      userID,
		Description: description,
		Industry:    industry,
		Location:    location,
		Website:     website,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

// UpdateProfile は企業プロフィールを更新する
func (c *Company) UpdateProfile(description, industry, location, website, logo string) {
	c.Description = description
	c.Industry = industry
	c.Location = location
	c.Website = website
	if logo != "" {
		c.Logo = logo
	}
	c.UpdatedAt = time.Now()
}
