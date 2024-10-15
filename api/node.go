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
		Resp.Succ(c, nodes.Nodes)

	}
}

func (n *node[T]) Do(c *gin.Context, action string, req any) (*T, error) {
	resp, err := client.DoRequest[T](c.Request.Context(), http.MethodPost, "node", action, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
