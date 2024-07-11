package domain

import "time"

type Comments struct {
    ID        uint      `json:"id"`
    CourseID  uint      `json:"course_id"`
    UserID    uint      `json:"user_id"`
    Content   string    `json:"content"`
    CreatedAt time.Time `json:"created_at"`
}
