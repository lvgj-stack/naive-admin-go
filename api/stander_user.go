package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	stander_req "github.com/Mr-LvGJ/stander/pkg/service/req"
	stander_resp "github.com/Mr-LvGJ/stander/pkg/service/resp"
	"naive-admin-go/pkg/client"
)

var StanderUser = &standerUser[any]{}

type standerUser[T any] struct {
}

func (r *standerUser[T]) Handle(c *gin.Context) {
	action := c.Query("Action")

	switch action {
	case "GetUserPlanInfo":
		var req stander_req.GetUserPlanInfoReq
		if err := c.Bind(&req); err != nil {
			Resp.Err(c, 20001, err.Error())
			return
		}
		rules, err := (&standerUser[stander_resp.GetUserPlanInfoResp]{}).Do(c, action, req)
		if err != nil {
			Resp.Err(c, 20001, err.Error())
			return
		}
		Resp.Succ(c, rules)
	case "ListUsers":
		var req stander_req.ListUsersReq
		if err := c.Bind(&req); err != nil {
			Resp.Err(c, 20001, err.Error())
			return
		}
		rules, err := (&standerUser[stander_resp.ListUsersResp]{}).Do(c, action, req)
		if err != nil {
			Resp.Err(c, 20001, err.Error())
			return
		}
		Resp.Succ(c, map[string]any{
			"pageData": rules.Users,
			"total":    rules.TotalCount,
		})
	case "EditUser":
		var req stander_req.EditUserReq
		if err := c.Bind(&req); err != nil {
			Resp.Err(c, 20001, err.Error())
			return
		}
		_, err := (&standerUser[stander_resp.EmptyResp]{}).Do(c, action, req)
		if err != nil {
			Resp.Err(c, 20001, err.Error())
			return
		}
		Resp.Succ(c, nil)
	}
}

func (r *standerUser[T]) Do(c *gin.Context, action string, req any) (*T, error) {
	resp, err := client.DoRequest[T](c, http.MethodPost, "user", action, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
