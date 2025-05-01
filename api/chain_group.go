package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	stander_req "github.com/Mr-LvGJ/stander/pkg/service/req"
	stander_resp "github.com/Mr-LvGJ/stander/pkg/service/resp"
	"naive-admin-go/pkg/client"
)

var ChainGroup = &chainGroup[any]{}

type chainGroup[T any] struct {
}

func (ch *chainGroup[T]) Handle(c *gin.Context) {
	action := c.Query("Action")
	switch action {
	case "ListChainGroups":
		var req stander_req.ListChainGroupReq
		if err := c.Bind(&req); err != nil {
			Resp.Err(c, 20001, err.Error())
			return
		}
		chains, err := (&chainGroup[stander_resp.ListChainGroupsResp]{}).Do(c, action, req)
		if err != nil {
			Resp.Err(c, 20001, err.Error())
			return
		}
		Resp.Succ(c, map[string]interface{}{
			"pageData": chains.ChainGroups,
			"total":    0,
		})
	case "AddChainGroup":
		var req stander_req.AddChainGroupReq
		if err := c.Bind(&req); err != nil {
			Resp.Err(c, 20001, err.Error())
			return
		}
		chains, err := (&chainGroup[stander_resp.AddChainGroupResp]{}).Do(c, action, req)
		if err != nil {
			Resp.Err(c, 20001, err.Error())
			return
		}
		Resp.Succ(c, chains)
	case "DeleteChainGroup":
		var req stander_req.DelChainGroupReq
		if err := c.Bind(&req); err != nil {
			Resp.Err(c, 20001, err.Error())
			return
		}
		chains, err := (&chainGroup[stander_resp.EmptyResp]{}).Do(c, action, req)
		if err != nil {
			Resp.Err(c, 20001, err.Error())
			return
		}
		Resp.Succ(c, chains)
	case "EditChain":
		var req stander_req.EditChainReq
		if err := c.Bind(&req); err != nil {
			Resp.Err(c, 20001, err.Error())
			return
		}
		_, err := (&chain[stander_resp.EditChainResp]{}).Do(c, action, req)
		if err != nil {
			Resp.Err(c, 20001, err.Error())
			return
		}
		Resp.Succ(c, nil)
	}
}
func (ch *chainGroup[T]) Do(c *gin.Context, action string, req any) (*T, error) {
	resp, err := client.DoRequest[T](c, http.MethodPost, "chain", action, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
