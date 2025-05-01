package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	client = http.DefaultClient
	prefix = "/api/v1/"
)

func getUrl(group, action string) string {
	url := os.Getenv("STANDER_URL")
	return url + prefix + group + "?Action=" + action
}

func DoRequest[T any](ctx *gin.Context, method, group, action string, request any) (*T, error) {
	resp := struct {
		Result *T
	}{}
	zap.S().Infow("begin do request",
		"method", method,
		"group", group,
		"action", action,
		"X-User-Id", strconv.Itoa(ctx.GetInt("uid")),
		"X-Role-Id", ctx.GetString("roleId"),
		"request", request)
	bs, _ := json.Marshal(request)
	req, err := http.NewRequestWithContext(ctx, method, getUrl(group, action), bytes.NewBuffer(bs))
	if err != nil {
		return nil, err
	}
	req.Header.Set("X-User-Id", strconv.Itoa(ctx.GetInt("uid")))
	req.Header.Set("X-Role-Id", ctx.GetString("roleId"))
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		zap.S().Errorw("client do err", "err", err)
		return nil, err
	}
	if res.Body != nil {
		defer res.Body.Close()
	}
	respBytes, err := io.ReadAll(res.Body)
	if err != nil {
		zap.S().Errorw("io read err", "err", err)
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		r := make(map[string]interface{})
		if err := json.Unmarshal(respBytes, &r); err != nil {
			zap.S().Errorw("unmarshal err", "err", err)
			return nil, err
		}
		return nil, errors.New(r["Error"].(string))
	}
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		zap.S().Errorw("unmarshal err", "err", err)
		return nil, err
	}
	zap.S().Infow("after do request", "result", resp)
	return resp.Result, nil
}
