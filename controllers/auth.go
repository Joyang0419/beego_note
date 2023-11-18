package controllers

import (
	"time"

	"github.com/Joyang0419/beego_note/models"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web"
)

type AuthController struct {
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
		_ = c.Ctx.Output.JSON(
			ErrorResponse{
				Message: err.Error(),
			},
			false,
			false,
		)
	}

	o := orm.NewOrm()
	id, err := o.Insert(
		&models.User{
			Account:  requestBody.Account,
			Password: requestBody.Password,
		})

	if err != nil {
		c.Ctx.Output.SetStatus(500)
		_ = c.Ctx.Output.JSON(
			ErrorResponse{
				Message: err.Error(),
			},
			false,
			false,
		)
	}

	_ = c.Ctx.JSONResp(
		&FormatResponse{
			Data: map[string]any{
				"userID": id,
			}},
	)
}

type SignInRequestBody struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func (c *AuthController) SignIn() {
	var (
		userContainer = new(models.User)
		requestBody   SignInRequestBody
	)
	if err := c.BindJSON(&requestBody); err != nil {
		c.Ctx.Output.SetStatus(400)
		_ = c.Ctx.Output.JSON(
			ErrorResponse{
				Message: err.Error(),
			},
			false,
			false,
		)
	}
	o := orm.NewOrm()
	query := o.QueryTable(new(models.User)).
		Filter("account", requestBody.Account).
		Filter("password", requestBody.Password)

	if err := query.One(userContainer); err != nil {
		if !query.Exist() {
			c.Ctx.Output.SetStatus(400)
			_ = c.Ctx.Output.JSON(
				ErrorResponse{
					Message: "user not found",
				},
				false,
				false,
			)
		}
		_ = c.Ctx.Output.JSON(
			ErrorResponse{
				Message: err.Error(),
			},
			false,
			false,
		)
	}

	userContainer.LoginTime = time.Now()
	if c.CruSession != nil {
		_ = c.DestroySession()
	}
	_ = c.SessionRegenerateID()
	_ = c.Ctx.Input.CruSession.Flush(c.Ctx.Request.Context())
	_ = c.SetSession("account", userContainer.Account)
	if _, err := o.Update(userContainer, "login_time"); err != nil {
		c.Ctx.Output.SetStatus(500)
		_ = c.Ctx.Output.JSON(
			ErrorResponse{
				Message: err.Error(),
			},
			false,
			false,
		)
	}

	_ = c.Ctx.JSONResp(
		&FormatResponse{
			Data: map[string]any{
				"userID": userContainer.ID,
			},
		},
	)
}
