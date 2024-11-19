package util

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/zeromicro/go-zero/core/logx"
)

var CtxKeyJwtUserId = "jwtUserID"

type Claims struct {
	UserID string `json:"jwtUserID"`
	jwt.StandardClaims
}

// GenerateJWT 生成一个JWT字符串
func GenerateJWT(userID, jwtKey string) (string, error) {
	expirationTime := time.Now().Add(15 * 24 * time.Hour) // 令牌将在15天后过期
	claims := &Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 注意：这里将jwtKey从字符串转换为字节切片
	tokenString, err := token.SignedString([]byte(jwtKey))

	return tokenString, err
}
func VerifyJWT(tokenString, jwtKey string) (*Claims, error) {
	claims := &Claims{}

	// 注意：这里的回调函数提供了jwtKey的字节切片形式
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		// 确保token的签名方法符合预期
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(jwtKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}

func GetUidFromCtx(ctx context.Context) string {
	var uid string
	if jsonUid, ok := ctx.Value(CtxKeyJwtUserId).(string); ok {
		uid = jsonUid
	}
	return uid
}

func GetIntUidFromCtx(ctx context.Context) int64 {
	var uid int64
	if jsonUid, ok := ctx.Value(CtxKeyJwtUserId).(json.Number); ok {
		if int64Uid, err := jsonUid.Int64(); err == nil {
			uid = int64Uid
		} else {
			logx.WithContext(ctx).Errorf("GetUidFromCtx err : %+v", err)
		}
	}
	return uid
}
