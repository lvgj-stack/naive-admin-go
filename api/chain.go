package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	stander_req "github.com/Mr-LvGJ/stander/pkg/service/req"
	stander_resp "github.com/Mr-LvGJ/stander/pkg/service/resp"
	"naive-admin-go/pkg/client"
)

var Chain = &chain[any]{}

type chain[T any] struct {
}

func (ch *chain[T]) Handle(c *gin.Context) {
	action := c.Query("Action")
	switch action {
	case "ListChains":
		var req stander_req.ListChainReq
		if err := c.Bind(&req); err != nil {
			Resp.Err(c, 20001, err.Error())
			return
		}
		chains, err := (&chain[stander_resp.ListChainResp]{}).Do(c, action, req)
		if err != nil {
			Resp.Err(c, 20001, err.Error())
			return
		}
		Resp.Succ(c, map[string]interface{}{
			"pageData": chains.Chains,
			"total":    chains.TotalCount,
		})
	case "AddChain":
		var req stander_req.AddChainReq
		if err := c.Bind(&req); err != nil {
			Resp.Err(c, 20001, err.Error())
			return
		}
		chains, err := (&chain[stander_resp.AddChainResp]{}).Do(c, action, req)
		if err != nil {
			Resp.Err(c, 20001, err.Error())
			return
		}
		Resp.Succ(c, chains)
	case "DeleteChain":
		var req stander_req.DelChainReq
		if err := c.Bind(&req); err != nil {
			Resp.Err(c, 20001, err.Error())
			return
		}
		chains, err := (&chain[stander_resp.DelChainResp]{}).Do(c, action, req)
		if err != nil {
			Resp.Err(c, 20001, err.Error())
			return
		}
		Resp.Succ(c, chains)
	}
}
func (ch *chain[T]) Do(c *gin.Context, action string, req any) (*T, error) {
	resp, err := client.DoRequest[T](c.Request.Context(), http.MethodPost, "chain", action, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
