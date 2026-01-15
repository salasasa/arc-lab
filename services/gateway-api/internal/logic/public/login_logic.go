// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package public

import (
	"context"
	"errors"
	"time"

	"gateway-api/internal/svc"
	"gateway-api/internal/types"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户登录
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line
	// TODO: 使用真正逻辑
	// 1. Mock 验证账号密码
	if req.Username == "admin" && req.Password == "123456" {
		now := time.Now().Unix()
		accessExpire := l.svcCtx.Config.Auth.AccessExpire

		// 2. 签发 JWT Token
		token, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, req.Username)
		if err != nil {
			return nil, err
		}

		return &types.LoginResp{
			AccessToken:  token,
			AccessExpire: now + accessExpire,
			Role:         "admin",
		}, nil
	}

	return nil, errors.New("用户名或密码错误")

}

// 辅助函数：生成 JWT 字符串
func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds int64, username string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["username"] = username
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}
