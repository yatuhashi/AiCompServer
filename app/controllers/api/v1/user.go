package controllers

import (
	"Base/app/db"
	"Base/app/models"
	"github.com/revel/revel"
	"gopkg.in/validator.v2"
)

type ApiUser struct {
	ApiV1Controller
}

func (c ApiUser) Index() revel.Result {
	users := []models.User{}
	if err := db.DB.Order("id desc").Find(&users).Error; err != nil {
		return c.HandleInternalServerError("Record Find Failure")
	}
	r := Response{users}
	return c.RenderJSON(r)
}

func (c ApiUser) Show(id int) revel.Result {
	user := &models.User{}
	if err := db.DB.First(&user, id).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}
	r := Response{user}
	return c.RenderJSON(r)
}

func (c ApiUser) Create() revel.Result {
	user := &models.User{}
	if err := c.BindParams(user); err != nil {
		return c.HandleBadRequestError(err.Error())
	}
	if err := validator.Validate(user); err != nil {
		return c.HandleBadRequestError(err.Error())
	}
	if err := db.DB.Create(user).Error; err != nil {
		return c.HandleBadRequestError(err.Error())
	}
	r := Response{user}
	return c.RenderJSON(r)
}

func (c ApiUser) Update(id int) revel.Result {
	userOld := &models.User{}
	if err := db.DB.First(&userOld, id).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}
	userNew := &models.User{}
	if err := c.BindParams(userNew); err != nil {
		return c.HandleBadRequestError(err.Error())
	}
	if err := validator.Validate(userNew); err != nil {
		return c.HandleBadRequestError(err.Error())
	}
	if err := db.DB.Model(&userOld).Update(&userNew).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}
	r := Response{userNew}
	return c.RenderJSON(r)
}

func (c ApiUser) Delete(id int) revel.Result {
	user := &models.User{}
	if err := db.DB.First(&user, id).Error; err != nil {
		return c.HandleNotFoundError(err.Error())
	}
	if err := db.DB.Delete(&user).Error; err != nil {
		return c.HandleInternalServerError("Record Delete Failure")
	}
	r := Response{"success delete"}
	return c.RenderJSON(r)
}
