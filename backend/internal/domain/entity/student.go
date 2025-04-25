package entity

import "time"

type Student struct {
	ID             int64     `json:"id"`
	UserID         int64     `json:"user_id"`
	University     string    `json:"university"`
	GraduationYear int       `json:"graduation_year"`
	Major          string    `json:"major"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func NewStudent(userID int64, university string, graduationYear int, major string) *Student {
	return &Student{
		UserID:         userID,
		University:     university,
		GraduationYear: graduationYear,
		Major:          major,
	}
}

// UpdateProfile は学生プロフィールを更新する
func (s *Student) UpdateProfile(university string, graduationYear int, major string) {
	s.University = university
	s.GraduationYear = graduationYear
	s.Major = major
	s.UpdatedAt = time.Now()
}
