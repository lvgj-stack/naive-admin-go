package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

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
		nodes, err := (&chain[stander_resp.ListChainResp]{}).Do(c, action)
		if err != nil {
			Resp.Err(c, 20001, err.Error())
			return
		}
		Resp.Succ(c, nodes)
	}
}
func (ch *chain[T]) Do(c *gin.Context, action string) (*T, error) {
	resp, err := client.DoRequest[T](c.Request.Context(), http.MethodPost, "chain", action, nil)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
