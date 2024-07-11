package dao

import "time"

type Comment struct {
    ID        uint      `gorm:"primaryKey"`
    CourseID  uint      `gorm:"not null"`
    UserID    uint      `gorm:"not null"`
    Content   string    `gorm:"type:text;not null"`
    CreatedAt time.Time `gorm:"autoCreateTime"`
}
