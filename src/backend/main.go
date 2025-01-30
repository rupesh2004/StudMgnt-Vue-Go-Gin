package main

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Student struct {
	StudentID    int    `json:"studentID"`
	FullName     string `json:"fullName" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	Phone        string `json:"phone" binding:"required"`
	Dob          string `json:"dob"`
	ProfilePhoto string `json:"profilePhoto"`
}

var students []Student

func setStudentID() int {
	return rand.Intn(9000) + 1000
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	router.Static("/profile_photos", "./profile_photos")

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": students})
	})

	router.POST("/studData", func(c *gin.Context) {
		var newStudent Student
		if err := c.ShouldBindJSON(&newStudent); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}


		if newStudent.ProfilePhoto != "" {
			decodedImage, err := base64.StdEncoding.DecodeString(newStudent.ProfilePhoto)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to decode base64 image"})
				return
			}

			fileName := fmt.Sprintf("%s.jpg", newStudent.FullName)
			filePath := filepath.Join("profile_photos", fileName)

			err = os.WriteFile(filePath, decodedImage, 0644)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
				return
			}

			newStudent.ProfilePhoto = fmt.Sprintf("http://localhost:3000/profile_photos/%s", fileName)
		}

		newStudent.StudentID = setStudentID()
		students = append(students, newStudent)

		c.JSON(http.StatusOK, gin.H{
			"message":  "Student data received successfully",
			"students": students,
		})
	})

	router.Run(":3000")
}
