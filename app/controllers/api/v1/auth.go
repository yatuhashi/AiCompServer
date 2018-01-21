package controllers

import (
	"Base/app/db"
	"Base/app/models"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
	"github.com/revel/revel/cache"
	"golang.org/x/crypto/scrypt"
	"gopkg.in/validator.v2"
	"log"
	"strconv"
	"time"
)

type Auth struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func TokenGenerator(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

type ApiAuth struct {
	ApiV1Controller
}

func (c ApiAuth) GetSessionID() revel.Result {
	session := TokenGenerator(32)
	cache.Set(session, session, 2*time.Minute)
	c.Response.Out.Header().Add("token", session)
	r := Response{"Get Session ID"}
	log.Print("&&", session, "&&")
	return c.RenderJSON(r)
}

func (c ApiAuth) SignIn() revel.Result {
	session := c.Request.Header.Get("token")
	if session == "" {
		return c.HandleNotFoundError("Retry Session")
	}
	log.Print("&&", session, "&&")
	var res string
	if err := cache.Get(session, &res); err != nil {
		r := Response{"Session Timeout"}
		return c.RenderJSON(r)
	}
	go cache.Delete(session)

	jsonData := &Auth{}
	if err := c.BindParams(jsonData); err != nil {
		return c.HandleBadRequestError(err.Error())
	}
	log.Print(jsonData.Password)
	salt := []byte("yatuhashi")
	converted, _ := scrypt.Key([]byte(jsonData.Password), salt, 16384, 8, 1, 32)
	password := hex.EncodeToString(converted[:])

	userOld := &models.User{}
	if err := db.DB.Find(&userOld, models.User{Username: jsonData.Username}).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}

	if userOld.Password != password {
		return c.HandleNotFoundError("password is incorrect")
	}

	userNew := &models.User{}
	userNew = userOld
	userNew.Token = TokenGenerator(32)
	if err := validator.Validate(userNew); err != nil {
		return c.HandleBadRequestError(err.Error())
	}
	if err := db.DB.Model(&userOld).Update(&userNew).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}
	r := Response{userNew.Token}
	go cache.Set(userNew.Token, userNew.Username, 30*time.Minute)
	c.Response.Out.Header().Add("token", userNew.Token)
	return c.RenderJSON(r)
}

func (c ApiAuth) SignOut() revel.Result {
	token := c.Request.Header.Get("token")
	if token == "" {
		return c.HandleNotFoundError("Not SignIn")
	}
	user := &models.User{}
	if err := db.DB.Find(&user, models.User{Token: token}).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}
	if err := db.DB.Model(&user).Update("Token", gorm.Expr("NULL")).Error; err != nil {
		return c.HandleBadRequestError(err.Error())
	}
	go cache.Delete(token)
	r := Response{"Sign Out"}
	return c.RenderJSON(r)
}

func CheckRole(c ApiV1Controller, AllowRole []string) revel.Result {
	log.Print("CheckRole")
	token := c.Request.Header.Get("token")
	if token == "" {
		return c.HandleNotFoundError("Not SignIn")
	}
	user := &models.User{}
	if err := db.DB.Find(&user, models.User{Token: token}).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}
	for _, role := range AllowRole {
		if role == user.Role || role == strconv.FormatUint(user.ID, 10) {
			return nil
		}
	}
	r := Response{"Permission Denied"}
	return c.RenderJSON(r)
}

func CheckToken(c ApiV1Controller) revel.Result {
	log.Print("CheckToken")
	token := c.Request.Header.Get("token")
	if token == "" {
		return c.HandleNotFoundError("Not SignIn")
	}
	user := &models.User{}
	if err := db.DB.Find(&user, models.User{Token: token}).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}
	// Check Token Timeout
	var res string
	if err := cache.Get(token, &res); err != nil {
		r := Response{"Session Timeout"}
		return c.RenderJSON(r)
	}
	go cache.Set(user.Token, user.Username, 30*time.Minute)
	return nil
}
