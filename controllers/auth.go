package controllers

import (
	"github.com/Joyang0419/beego_note/models"
	"github.com/astaxie/beego/logs"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
)

type AuthController struct {
	Ormer orm.Ormer
	web.Controller
}

type RegisterRequestBody struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func (c *AuthController) Register() {
	var requestBody RegisterRequestBody
	if err := c.BindJSON(&requestBody); err != nil {
		c.Ctx.Output.SetStatus(400)
		if jsonErr := c.Ctx.Output.JSON(
			ErrorResponse{
				Message: err.Error(),
			},
			false,
			false,
		); jsonErr != nil {
			logs.Error(jsonErr)
		}
		return
	}

	id, err := c.Ormer.Insert(
		&models.User{
			Account:  requestBody.Account,
			Password: requestBody.Password,
		})

	if err != nil {
		c.Ctx.Output.SetStatus(500)
		if jsonErr := c.Ctx.Output.JSON(
			ErrorResponse{
				Message: err.Error(),
			},
			false,
			false,
		); jsonErr != nil {
			logs.Error(jsonErr)
		}
		return
	}

	if jsonErr := c.Ctx.JSONResp(&FormatResponse{
		Data: map[string]any{
			"userID": id,
		}}); jsonErr != nil {
		logs.Error(jsonErr)
	}
}
