package main

import (
	"encoding/base64"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Student structure
type Student struct {
	StudentID    int    `json:"studentID"`
	FullName     string `json:"fullName" binding:"required"`
	Email        string `json:"email" binding:"required,email"`
	Phone        string `json:"phone" binding:"required"`
	Dob          string `json:"dob"`
	ProfilePhoto string `json:"profilePhoto"`
}

var students []Student

// Generate a random Student ID
func setStudentID() int {
	return rand.Intn(9000) + 1000
}

func main() {
	router := gin.Default()
	router.Use(cors.Default())

	// Serve profile photos statically
	router.Static("/profile_photos", "./profile_photos")

	// Fetch all students
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": students})
	})

	// Add a new student
	router.POST("/studData", func(c *gin.Context) {
		var newStudent Student
		if err := c.ShouldBindJSON(&newStudent); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Handle Profile Photo
		if newStudent.ProfilePhoto != "" {
			decodedImage, err := base64.StdEncoding.DecodeString(newStudent.ProfilePhoto)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to decode base64 image"})
				return
			}

			fileName := fmt.Sprintf("%d.jpg", setStudentID())
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

	// Delete a student
	router.DELETE("/deleteStudent/:studentID", func(c *gin.Context) {
		id := c.Param("studentID")
		studentID, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
			return
		}

		for i, student := range students {
			if student.StudentID == studentID {
				students = append(students[:i], students[i+1:]...)
				c.JSON(http.StatusOK, gin.H{"message": "Student deleted successfully"})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Student with ID %d not found", studentID)})
	})

	// Update a student
	router.PUT("/updateStudent/:studentID", func(c *gin.Context) {
		id := c.Param("studentID")
		studentID, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid student ID"})
			return
		}
	
		var updatedData Student
		if err := c.ShouldBindJSON(&updatedData); err != nil {
			fmt.Println("❌ Request Body Error:", err) // Debugging Log
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request data"})
			return
		}
	
		fmt.Printf("🔹 Received update request for StudentID %d: %+v\n", studentID, updatedData) // Debugging Log
		fmt.Printf("🔹 Base64 Image Length: %d\n", len(updatedData.ProfilePhoto))
	
		// If profile photo is updated
		if updatedData.ProfilePhoto != "" {
			decodedImage, err := base64.StdEncoding.DecodeString(updatedData.ProfilePhoto)
			if err != nil {
				fmt.Println("❌ Base64 Decode Error:", err) // Debugging Log
				c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to decode base64 image"})
				return
			}
	
			fileName := fmt.Sprintf("%d_updated.jpg", studentID)
			filePath := filepath.Join("profile_photos", fileName)
	
			err = os.WriteFile(filePath, decodedImage, 0644)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save image"})
				return
			}
	
			updatedData.ProfilePhoto = fmt.Sprintf("http://localhost:3000/profile_photos/%s", fileName)
		}
	
		// Update student data in list
		for i, student := range students {
			if student.StudentID == studentID {
				students[i] = updatedData
				c.JSON(http.StatusOK, gin.H{"message": "Student updated successfully", "student": students[i]})
				return
			}
		}
		c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Student with ID %d not found", studentID)})
	})
	
	
	

	router.Run(":3000")
}
