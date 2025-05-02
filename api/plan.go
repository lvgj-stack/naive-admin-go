package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	stander_req "github.com/Mr-LvGJ/stander/pkg/service/req"
	stander_resp "github.com/Mr-LvGJ/stander/pkg/service/resp"
	"naive-admin-go/pkg/client"
)

var Plan = &plan[any]{}

type plan[T any] struct {
}

func (r *plan[T]) Handle(c *gin.Context) {
	action := c.Query("Action")

	switch action {
	case "ListPlans":
		var req stander_req.ListPlansReq
		if err := c.Bind(&req); err != nil {
			Resp.Err(c, 20001, err.Error())
			return
		}
		res, err := (&plan[stander_resp.ListPlansResp]{}).Do(c, action, req)
		if err != nil {
			Resp.Err(c, 20001, err.Error())
			return
		}
		Resp.Succ(c, res)
	}
}

func (r *plan[T]) Do(c *gin.Context, action string, req any) (*T, error) {
	resp, err := client.DoRequest[T](c, http.MethodPost, "plan", action, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
