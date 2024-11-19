package user

import (
	"time"

	"github.com/wlevene/mgm/v3"
)

type User struct {
	mgm.DefaultModel `bson:",inline"`

	Password      string    `bson:"password"`
	AvatarURL     string    `bson:"avatar_url"`
	Gender        string    `bson:"gender"`
	EMail         string    `bson:"email"`
	UserName      string    `bson:"user_name"`
	Bio           string    `bson:"bio"`
	Nick          string    `bson:"nick"`
	Level         int64     `bson:"level"`
	LastLoginIP   string    `bson:"last_login_ip"`
	LastLoginTime time.Time `bson:"last_login_at"`
}
