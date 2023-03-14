package authcontroller

import (
	"encoding/json"
	"net/http"
	"time"
	middleware "viska/middleware/jwt"
	"viska/storage/models"

	"github.com/golang-jwt/jwt/v4"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/thedevsaddam/renderer"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	logger *logrus.Logger
	render *renderer.Render
	db     *gorm.DB
}

func NewAuth(logger *logrus.Logger, render *renderer.Render, db *gorm.DB) *Auth {
	return &Auth{
		logger: logger,
		render: render,
		db:     db,
	}
}

func (a *Auth) Register(w http.ResponseWriter, r *http.Request) {
	logger := a.logger.WithFields(logrus.Fields{
		"Layer":     "Auth",
		"Func Name": "Register",
	})

	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		logger.Errorf(`Error: %v`, err)
		a.render.JSON(w, http.StatusBadRequest, &models.Template1{
			Message: "Error",
			Data:    nil,
		})
		return
	}
	defer r.Body.Close()

	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(userInput.Password), bcrypt.DefaultCost)
	userInput.Password = string(hashPassword)

	if err := a.db.Create(&userInput).Error; err != nil {
		logger.Errorf(`Error something: %v`, err)
		a.render.JSON(w, http.StatusInternalServerError, models.Template1{
			Message: "Error",
			Data:    nil,
		})
		return
	}

	a.render.JSON(w, http.StatusOK, &models.Template1{
		Message: "Success",
	})
}

func (a *Auth) Login(w http.ResponseWriter, r *http.Request) {
	logger := a.logger.WithFields(logrus.Fields{
		"Layer":     "Auth",
		"Func Name": "Login",
	})

	var userInput models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&userInput); err != nil {
		logger.Errorf(`Error: %v`, err)
		a.render.JSON(w, http.StatusBadRequest, &models.Template1{
			Message: "Error",
			Data:    nil,
		})
		return
	}

	var user models.User
	if err := a.db.Where("username = ?", userInput.Username).First(&user).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			logger.Errorf(`Error record not found`)
			a.render.JSON(w, http.StatusUnauthorized, &models.Template1{
				Message: "Error unauthorized.",
				Data:    nil,
			})
			return
		default:
			logger.Errorf(`Error server error`)
			a.render.JSON(w, http.StatusInternalServerError, &models.Template1{
				Message: "Error server error",
				Data:    nil,
			})
			return
		}

	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		logger.Errorf(`Username atau Password salah`)
		a.render.JSON(w, http.StatusUnauthorized, &models.Template1{
			Message: "Username atau Password salah.",
			Data:    nil,
		})
		return
	}

	expTime := time.Now().Add(time.Minute * 10)
	claims := &middleware.JWTClaims{
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "viska-balita",
			ExpiresAt: jwt.NewNumericDate(expTime),
		},
	}

	tokenAlgo := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenAlgo.SignedString(middleware.JWT_KEY)
	if err != nil {
		logger.Errorf(`Error signed token`)
		a.render.JSON(w, http.StatusInternalServerError, &models.Template1{
			Message: "Error signed token",
			Data:    nil,
		})
		return
	}

	a.render.JSON(w, http.StatusOK, &models.Template1{
		Message: "Login berhasil",
		Data:    token,
	})
}
