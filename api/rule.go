package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	stander_req "github.com/Mr-LvGJ/stander/pkg/service/req"
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
		var req stander_req.ListRuleReq
		if err := c.Bind(&req); err != nil {
			Resp.Err(c, 20001, err.Error())
			return
		}
		rules, err := (&rule[stander_resp.ListRuleResp]{}).Do(c, action, req)
		if err != nil {
			Resp.Err(c, 20001, err.Error())
			return
		}
		Resp.Succ(c, map[string]interface{}{
			"pageData": rules.Rules,
			"total":    rules.TotalCount,
		})
	case "AddRule":
		var req stander_req.AddRuleReq
		if err := c.Bind(&req); err != nil {
			Resp.Err(c, 20001, err.Error())
			return
		}
		rules, err := (&rule[stander_resp.AddRuleResp]{}).Do(c, action, req)
		if err != nil {
			Resp.Err(c, 20001, err.Error())
			return
		}
		Resp.Succ(c, rules)
	case "DeleteRule":
		var req stander_req.DelRuleReq
		if err := c.Bind(&req); err != nil {
			Resp.Err(c, 20001, err.Error())
			return
		}
		rules, err := (&rule[stander_resp.DelRuleResp]{}).Do(c, action, req)
		if err != nil {
			Resp.Err(c, 20001, err.Error())
			return
		}
		Resp.Succ(c, rules)
	case "ModifyRule":
		var req stander_req.ModifyRuleReq
		if err := c.Bind(&req); err != nil {
			Resp.Err(c, 20001, err.Error())
			return
		}
		rules, err := (&rule[stander_resp.ModifyRuleResp]{}).Do(c, action, req)
		if err != nil {
			Resp.Err(c, 20001, err.Error())
			return
		}
		Resp.Succ(c, rules)
	case "TestRule":
		var req stander_req.TestRuleReq
		if err := c.Bind(&req); err != nil {
			Resp.Err(c, 20001, err.Error())
			return
		}
		rules, err := (&rule[stander_resp.TestRuleResp]{}).Do(c, action, req)
		if err != nil {
			Resp.Err(c, 20001, err.Error())
			return
		}
		Resp.Succ(c, rules)
	}
}

func (r *rule[T]) Do(c *gin.Context, action string, req any) (*T, error) {
	resp, err := client.DoRequest[T](c, http.MethodPost, "rule", action, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
