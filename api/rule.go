package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	stander_resp "github.com/Mr-LvGJ/stander/pkg/service/resp"
	"naive-admin-go/pkg/client"
)

var Rule = &rule[any]{}

type rule[T any] struct {
}

func (r *rule[T]) Handle(c *gin.Context) {
	action := c.Query("Action")

	switch action {
	case "ListRules":
		nodes, err := (&rule[stander_resp.ListRuleResp]{}).Do(c, action)
		if err != nil {
			Resp.Err(c, 20001, err.Error())
			return
		}
		Resp.Succ(c, nodes)

	}
}

func (r *rule[T]) Do(c *gin.Context, action string) (*T, error) {
	resp, err := client.DoRequest[T](c.Request.Context(), http.MethodPost, "rule", action, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
