package controllers

import (
	"gudang-obat/config"
	"gudang-obat/models"
	"gudang-obat/token"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type RegisterInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {

	var input RegisterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)

	duplicate := config.DB.Where("username = ?", input.Username).First(&u)

	u.Username = input.Username
	u.Password = string(hashedPassword)

	if duplicate.Error != nil {
		config.DB.Save(&u)
		c.JSON(http.StatusOK, gin.H{"message": "registration success"})
	} else {
		c.JSON(http.StatusConflict, gin.H{"message": "username already exist"})
	}
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Login(c *gin.Context) {

	var input LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}

	u.Username = input.Username
	u.Password = input.Password

	user, token, err := LoginCheck(u.Username, u.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}

	s := models.Session{}
	s.CreatedAt = time.Now()
	s.UpdatedAt = time.Now()
	s.UserID = user.ID
	s.Token = token
	s.ExpiredAt = time.Now().Add(24 * time.Hour)
	result := config.DB.Create(&s).Error

	if result != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed set session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": s})

}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(username string, password string) (models.User, string, error) {

	var err error

	u := models.User{}

	// err = config.DB.Model(models.User{}).Where("username = ?", username).Take(&u).Error
	user, err := FindUser(username)

	if err != nil {
		return user, "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return user, "", err
	}

	token, err := token.GenerateToken(u.ID)

	if err != nil {
		return user, "", err
	}

	return user, token, nil

}

func FindUser(username string) (models.User, error) {
	var users models.User
	err := config.DB.Where("username = ?", username).Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil

}
