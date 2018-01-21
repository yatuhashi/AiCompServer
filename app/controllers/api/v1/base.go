package controllers

import (
	"encoding/json"
	"github.com/revel/revel"
	"net/http"
)

type ApiV1Controller struct {
	*revel.Controller
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Response struct {
	Results interface{} `json:"results"`
}

func (c *ApiV1Controller) BindParams(s interface{}) error {
	return json.Unmarshal(c.Params.JSON, &s)
}

// Bad Request Error
func (c *ApiV1Controller) HandleBadRequestError(s string) revel.Result {
	c.Response.Status = http.StatusBadRequest
	r := ErrorResponse{c.Response.Status, s}
	return c.RenderJSON(r)
}

// Not Found Error
func (c *ApiV1Controller) HandleNotFoundError(s string) revel.Result {
	c.Response.Status = http.StatusNotFound
	r := ErrorResponse{c.Response.Status, s}
	return c.RenderJSON(r)
}

// Internal Server Error
func (c *ApiV1Controller) HandleInternalServerError(s string) revel.Result {
	c.Response.Status = http.StatusInternalServerError
	r := ErrorResponse{c.Response.Status, s}
	return c.RenderJSON(r)
}
