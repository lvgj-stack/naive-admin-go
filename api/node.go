package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

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
		nodes, err := (&node[stander_resp.ListNodeResp]{}).Do(c, action)
		if err != nil {
			Resp.Err(c, 20001, err.Error())
			return
		}
		Resp.Succ(c, nodes)

	}
}

func (n *node[T]) Do(c *gin.Context, action string) (*T, error) {
	resp, err := client.DoRequest[T](c.Request.Context(), http.MethodPost, "node", action, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
