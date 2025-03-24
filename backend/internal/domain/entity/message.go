package entity

import "time"

type Message struct {
	ID         int64      `json:"id"`
	SenderID   int64      `json:"sender_id"`
	ReceiverID int64      `json:"receiver_id"`
	Content    string     `json:"content"`
	CreatedAt  time.Time  `json:"created_at"`
	ReadAt     *time.Time `json:"read_at,omitempty"`
	Sender     *User      `json:"sender,omitempty"`   // 送信者情報への参照（必要に応じて）
	Receiver   *User      `json:"receiver,omitempty"` // 受信者情報への参照（必要に応じて）
}

func NewMessage(senderID, receiverID int64, content string) *Message {
	return &Message{
		SenderID:   senderID,
		ReceiverID: receiverID,
		Content:    content,
		CreatedAt:  time.Now(),
	}
}

func (m *Message) MarkAsRead() {
	now := time.Now()
	m.ReadAt = &now
}

func (m *Message) IsRead() bool {
	return m.ReadAt != nil
}
