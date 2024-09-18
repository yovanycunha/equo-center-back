package controllers

import (
	"equocenterback/pkg/models"
	"equocenterback/pkg/repositories"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type UserController struct {
	UserRepository repositories.UserRepository
}

func NewUserController(userService repositories.UserRepository) UserController {
	return UserController{
		UserRepository: userService,
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user models.User

	if err := ctx.ShouldBindJSON(&user)
	err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error message": err.Error()})
		return
	}

	passwordHash, errHash := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if errHash != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error message": errHash.Error()})
		return
	}

	newUser := models.User{
		Email: user.Email,
		Password: string(passwordHash),
	}

	err := uc.UserRepository.CreateUser(&newUser)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error message": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "user created successfully"})
}

func (uc *UserController) Login(ctx *gin.Context) {
	var inputUser models.User

	if err := ctx.ShouldBindJSON(&inputUser)
	err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error message": err.Error()})
		return
	}

	userFound, err := uc.UserRepository.GetUser(&inputUser.Email)

	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error message": err.Error() + " - No user found"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(inputUser.Password))
	err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error message": "invalid password"})
		return
	}

	generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": userFound.Email,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	token, err := generateToken.SignedString([]byte("secret"))
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"error message": err.Error() + " - Error generating token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (uc *UserController) VerifyToken(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

	authToken := strings.Split(authHeader, " ")
	if len(authToken) != 2 || authToken[0] != "Bearer" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token format"})
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	tokenString := authToken[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})
	if err != nil || !token.Valid {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		ctx.Abort()
		return
	}

	if float64(time.Now().Unix()) > claims["exp"].(float64) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token expired"})
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	var email = claims["email"].(string)
	found,err := uc.UserRepository.GetUser(&email)

	if found.ID == [12]byte{} {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Set("currentUser", found)

	ctx.Next()
}

func (uc *UserController) RegisterUserRoutes(routes *gin.RouterGroup) {
	userRoutes := routes.Group("/user")

	userRoutes.POST("/create", uc.CreateUser)
	userRoutes.POST("/login", uc.Login)
	userRoutes.POST("/verify", uc.VerifyToken)
}