package controllers

import (
	"github.com/beego/beego/v2/server/web"
	session2 "github.com/beego/beego/v2/server/web/session"
)

type SessionController struct {
	web.Controller
}

func (c *SessionController) SessionInfo() {
	ctx := c.Ctx.Request.Context()
	session := c.StartSession()
	defer session.SessionRelease(ctx, c.Ctx.ResponseWriter)
	sessionID := session.SessionID(ctx)

	provider, err := session2.GetProvider("redis")
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

	exist, err := provider.SessionExist(ctx, sessionID)
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

	if !exist {
		_ = c.DestroySession()
		c.Ctx.Output.SetStatus(500)
		_ = c.Ctx.Output.JSON(
			ErrorResponse{
				Message: "non-exist sessionID",
			},
			false,
			false,
		)
	}
	_ = c.Ctx.JSONResp(
		&FormatResponse{
			Data: map[string]any{
				"session_id": session.SessionID(ctx),
				"session_data": map[string]any{
					"account":  c.GetSession("account"),
					"password": c.GetSession("password"),
				},
			},
		},
	)
}
