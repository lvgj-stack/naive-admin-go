package api

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"naive-admin-go/db"
	"naive-admin-go/inout"
	"naive-admin-go/model"
	"naive-admin-go/utils"
)

var Auth = &auth{}

type auth struct {
}

func (auth) Captcha(c *gin.Context) {
	id, b64s, _ := utils.GetCaptcha()
	session := sessions.Default(c)
	session.Set("captch", id)
	session.Save()

	parts := strings.SplitN(b64s, ",", 2)
	imgData, _ := base64.StdEncoding.DecodeString(parts[1])

	c.Header("Content-Type", parts[0]) // 使用提取的 MIME 类型

	c.Data(http.StatusOK, parts[0], imgData) // 使用提取的 MIME 类型
}

func (auth) Login(c *gin.Context) {
	var params inout.LoginReq
	err := c.Bind(&params)
	if err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	}
	session := sessions.Default(c)
	if !utils.VerifyCaptcha(session.Get("captch").(string), params.Captcha) {
		Resp.Err(c, 20001, "验证码不正确")
		return
	}
	var info *model.User
	db.Dao.Model(model.User{}).
		Where("username =? ", params.Username).
		Where("password=?", fmt.Sprintf("%x", md5.Sum([]byte(params.Password)))).
		Find(&info)
	if info.ID == 0 {
		Resp.Err(c, 20001, "账号或密码不正确")
		return
	}

	var roleIds, roleNames []string
	db.Dao.Model(model.UserRolesRole{}).
		Where("userId = ?", info.ID).
		Select("roleId").Find(&roleIds)
	db.Dao.Model(model.Role{}).
		Where("id in (?)", roleIds).Order("id asc").
		Select("code").Find(&roleNames)

	currentRole := ""
	if len(roleNames) > 1 {
		currentRole = roleNames[0]
	}
	Resp.Succ(c, inout.LoginRes{
		AccessToken: utils.GenerateToken(info.ID, info.ID, info.Username, currentRole, roleNames),
	})
}

func (auth) password(c *gin.Context) {
	var params inout.AuthPwReq
	err := c.Bind(&params)
	if err != nil {
		Resp.Err(c, 20001, err.Error())
		return
	}
	uid, _ := c.Get("uid")
	var oldCun int64
	db.Dao.Model(model.User{}).Where("id=? and password=?", uid, fmt.Sprintf("%x", md5.Sum([]byte(params.OldPassword))))
	if oldCun > 0 {
		db.Dao.Model(model.User{}).
			Where("id=? ", uid).
			Update("password", fmt.Sprintf("%x", md5.Sum([]byte(params.NewPassword))))
	}
	Resp.Succ(c, true)
}
func (auth) Logout(c *gin.Context) {
	Resp.Succ(c, true)
}

func (auth) SwitchRole(c *gin.Context) {
	roleName := c.Param("role")
	jwtToken, _ := c.Get("jwt_token")
	claim := jwtToken.(*utils.CustomClaims)
	claim.CurrentRoleCode = roleName
	Resp.Succ(c, inout.LoginRes{
		AccessToken: utils.GenerateToken(claim.UID, claim.UID, claim.Username, roleName, claim.RoleCodes),
	})
}
