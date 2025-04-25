package entity

import "time"

// AppointmentStatus は予約のステータスを表す
type AppointmentStatus string

const (
	AppointmentStatusPending   AppointmentStatus = "pending"
	AppointmentStatusConfirmed AppointmentStatus = "confirmed"
	AppointmentStatusCancelled AppointmentStatus = "cancelled"
)

// Appointment は学生と企業の面談予約を表すエンティティ
type Appointment struct {
	ID        int64             `json:"id"`
	StudentID int64             `json:"student_id"`
	CompanyID int64             `json:"company_id"`
	DateTime  time.Time         `json:"date_time"`
	Status    AppointmentStatus `json:"status"`
	Notes     string            `json:"notes,omitempty"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}

func NewAppointment(studentID, companyID int64, dateTime time.Time, notes string) *Appointment {
	now := time.Now()
	return &Appointment{
		StudentID: studentID,
		CompanyID: companyID,
		DateTime:  dateTime,
		Status:    AppointmentStatusPending,
		Notes:     notes,
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (a *Appointment) IsPending() bool {
	return a.Status == AppointmentStatusPending
}

func (a *Appointment) IsConfirmed() bool {
	return a.Status == AppointmentStatusConfirmed
}

func (a *Appointment) IsCancelled() bool {
	return a.Status == AppointmentStatusCancelled
}
