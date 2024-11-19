package login

import (
	"context"
	"errors"

	"github.com/wlevene/loginservice/internal/svc"
	"github.com/wlevene/loginservice/internal/types"
	"github.com/wlevene/loginservice/internal/util"
	"golang.org/x/crypto/bcrypt"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginReply, err error) {
	logx.Debug("LoginLogic.Login", req)
	user, err := l.svcCtx.Dao.UserModel.FindByEmail(req.Email)
	if err != nil {
		logx.Errorf("Failed to find user by email: %v, error: %v", req.Email, err)
		return nil, errors.New("invalid email or password")
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Pwd))
	if err != nil {
		logx.Errorf("Password does not match for user: %v, error: %v", req.Email, err)
		return nil, errors.New("invalid email or password")
	}

	// 生成JWT
	token, err := util.GenerateJWT(user.EMail, l.svcCtx.Config.JwtAuth.AccessSecret)
	if err != nil {
		logx.Errorf("Failed to generate JWT for user: %v, error: %v", req.Email, err)
		return nil, err
	}

	// 构建响应
	resp = &types.LoginReply{
		Token: token,
	}
	return
}
