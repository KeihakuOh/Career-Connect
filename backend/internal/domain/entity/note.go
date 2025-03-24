package entity

import "time"

type Note struct {
	ID        int64     `json:"id"`
	StudentID int64     `json:"student_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewNote(studentID int64, title, content string) *Note {
	now := time.Now()
	return &Note{
		StudentID: studentID,
		Title:     title,
		Content:   content,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (n *Note) Update(title, content string) {
	n.Title = title
	n.Content = content
	n.UpdatedAt = time.Now()
}
