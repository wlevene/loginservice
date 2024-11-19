package login

import (
	"context"
	"errors"
	"time"

	"github.com/wlevene/loginservice/internal/dao/accountlocks"
	"github.com/wlevene/loginservice/internal/dao/loginattempts"
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

	resp = &types.LoginReply{}

	// check account is locked
	lock, _ := l.svcCtx.Dao.AccountLocksModel.FindByUserId(user.ID.Hex())
	if lock != nil &&
		lock.UnlockTime < time.Now().Unix() {
		err = errors.New(lock.ToString())
		return
	}

	attempts, _ := l.svcCtx.Dao.LoginAttemptsModel.FindByDefaultTimeWindowSize(user.ID.Hex())

	if len(attempts) >= 5 {
		err = errors.New("Too many login attempts. Please try again later")
		return
	}

	failed_attempts, _ := l.svcCtx.Dao.LoginAttemptsModel.FindFailAttemptsByTimeWindowSize(
		5*time.Minute,
		user.ID.Hex())

	if len(failed_attempts) >= 5 {
		lock = &accountlocks.AccountLocks{
			LockReason: "Too many login attempts",
			LockTime:   time.Now().Unix(),
			UnlockTime: time.Now().Add(5 * time.Minute).Unix(),
			UseId:      user.ID.Hex(),
		}
		l.svcCtx.Dao.AccountLocksModel.Insert(context.Background(), lock)
	}

	login_attempts := &loginattempts.LoginAttempts{
		UseId:       user.ID.Hex(),
		AttemptTime: time.Now().Unix(),
		IPAddress:   "",
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Pwd))
	if err != nil {
		login_attempts.Success = false
		l.svcCtx.Dao.LoginAttemptsModel.Insert(context.Background(), login_attempts)

		// find recently  login attempts
		return nil, errors.New("invalid email or password")
	}

	login_attempts.Success = true
	l.svcCtx.Dao.LoginAttemptsModel.Insert(context.Background(), login_attempts)

	// delete account lock if exists
	lock, _ = l.svcCtx.Dao.AccountLocksModel.FindByUserId(user.ID.Hex())
	if lock != nil {
		l.svcCtx.Dao.AccountLocksModel.Delete(context.Background(), lock.ID.Hex())
	}

	token, err := util.GenerateJWT(user.EMail, l.svcCtx.Config.JwtAuth.AccessSecret)
	if err != nil {
		logx.Errorf("Failed to generate JWT for user: %v, error: %v", req.Email, err)
		return nil, err
	}

	resp.Token = token

	return
}
