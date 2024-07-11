package controllers

import (
	"backend/clients"
	"backend/dao"
	"backend/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func AddComment(c *gin.Context) {
    var comment domain.CommentDomain
    if err := c.ShouldBindJSON(&comment); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    // Convertir domain.CommentDomain a dao.Comment
    daoComment := dao.Comment{
        CourseID:  comment.CourseID,
        UserID:    comment.UserID,
        Content:   comment.Content,
        CreatedAt: comment.CreatedAt,
    }

    if err := clients.DB.Create(&daoComment).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"status": "comment added"})
}

func GetComments(c *gin.Context) {
    courseID, err := strconv.Atoi(c.Param("course_id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid course id"})
        return
    }

    var comments []dao.Comment
    if err := clients.DB.Where("course_id = ?", courseID).Find(&comments).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, comments)
}
