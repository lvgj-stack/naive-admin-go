package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	stander_req "github.com/Mr-LvGJ/stander/pkg/service/req"
	stander_resp "github.com/Mr-LvGJ/stander/pkg/service/resp"
	"naive-admin-go/pkg/client"
)

var Node = &node[any]{}

type node[T any] struct {
}

func (n *node[T]) Handle(c *gin.Context) {
	action := c.Query("Action")
	switch action {
	case "ListNodes":
		var req stander_req.ListNodeReq
		if err := c.Bind(&req); err != nil {
			Resp.Err(c, 20001, err.Error())
			return
		}
		nodes, err := (&node[stander_resp.ListNodeResp]{}).Do(c, action, req)
		if err != nil {
			Resp.Err(c, 20001, err.Error())
			return
		}
		Resp.Succ(c, map[string]interface{}{
			"pageData": nodes.Nodes,
			"total":    nodes.TotalCount,
		})
	case "AddNode":
		var req stander_req.AddNodeReq
		if err := c.Bind(&req); err != nil {
			Resp.Err(c, 20001, err.Error())
			return
		}
		nodes, err := (&node[stander_resp.AddNodeResp]{}).Do(c, action, req)
		if err != nil {
			Resp.Err(c, 20001, err.Error())
			return
		}
		Resp.Succ(c, nodes.Key)
	case "DeleteNode":
		var req stander_req.DelNodeReq
		if err := c.Bind(&req); err != nil {
			Resp.Err(c, 20001, err.Error())
			return
		}
		nodes, err := (&node[stander_resp.DelNodeResp]{}).Do(c, action, req)
		if err != nil {
			Resp.Err(c, 20001, err.Error())
			return
		}
		Resp.Succ(c, nodes.ID)
	}
}

func (n *node[T]) Do(c *gin.Context, action string, req any) (*T, error) {
	resp, err := client.DoRequest[T](c.Request.Context(), http.MethodPost, "node", action, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
