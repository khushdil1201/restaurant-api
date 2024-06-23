package controllers

import (
    "github.com/gin-gonic/gin"
    "github.com/dgrijalva/jwt-go"
    "restaurant-api/models"
    "golang.org/x/crypto/bcrypt"
    "time"
    "os"
)

func GenerateJWT(user models.User) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": user.ID,
        "username": user.Username,
        "role": user.Role,
        "exp": time.Now().Add(time.Hour * 72).Unix(),
    })

    return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

// @Summary Login a user
// @Description Login a user and return a JWT token
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.User true "User credentials"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /login [post]
func Login(c *gin.Context) {
    var input models.User
    var user models.User

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    models.DB.Where("username = ?", input.Username).First(&user)

    if user.ID == 0 {
        c.JSON(401, gin.H{"error": "Invalid username or password"})
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
        c.JSON(401, gin.H{"error": "Invalid username or password"})
        return
    }

    token, err := GenerateJWT(user)
    if err != nil {
        c.JSON(500, gin.H{"error": "Failed to generate token"})
        return
    }

    c.JSON(200, gin.H{"token": token})
}

// @Summary Register a new user
// @Description Register a new user with the provided credentials
// @Tags auth
// @Accept json
// @Produce json
// @Param user body models.User true "User credentials"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Router /register [post]
func RegisterUser(c *gin.Context) {
    var input models.User

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }

    passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(500, gin.H{"error": "Failed to hash password"})
        return
    }
    input.Password = string(passwordHash)

    models.DB.Create(&input)
    c.JSON(200, gin.H{"message": "User registered successfully"})
}