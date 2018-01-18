package controllers

import (
	"Base/app/controllers"
	"base/app/models"
	"github.com/revel/revel"
	"gopkg.in/validator.v2"
	"log"
)

type ApiUser struct {
	ApiV1Controller
}

func (c ApiUser) Index() revel.Result {
	users := []models.User{}
	if err := controllers.DB.Order("id desc").Find(&users).Error; err != nil {
		return c.HandleInternalServerError("Record Find Failure")
	}
	r := Response{users}
	return c.RenderJSON(r)
}

func (c ApiUser) Show(id int) revel.Result {
	users := []models.User{}
	if err := controllers.DB.First(&users, id).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}
	r := Response{users}
	return c.RenderJSON(r)
}

func (c ApiUser) Create() revel.Result {
	user := &models.User{}
	if err := c.BindParams(user); err != nil {
		log.Printf("test@@@@@@@@@@")
		return c.HandleBadRequestError(err.Error())
	}
	log.Print(user.Username)
	log.Print(user.Password)
	log.Print(user.Role)
	if err := validator.Validate(user); err != nil {
		log.Printf("test@@@@@@@@@@")
		log.Print(err)
		return c.HandleBadRequestError(err.Error())
	}
	if err := controllers.DB.Create(user).Error; err != nil {
		return c.HandleInternalServerError("Record Create Failure")
	}
	r := Response{user}
	return c.RenderJSON(r)
}

func (c ApiUser) Update() revel.Result {
	// user := &models.User{}
	r := Response{"update"}
	return c.RenderJSON(r)
}

func (c ApiUser) Delete(id int) revel.Result {
	user := &models.User{}
	if err := controllers.DB.First(&user, id).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}
	if err := controllers.DB.Delete(&user).Error; err != nil {
		return c.HandleInternalServerError("Record Delete Failure")
	}
	r := Response{"success delete"}
	return c.RenderJSON(r)
}
