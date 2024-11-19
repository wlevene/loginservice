package user

import (
	"github.com/wlevene/mgm/v3"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		FindByUserID(user_id string) (*User, error)
		FindByEmail(email string) (*User, error)

		All() ([]*User, error)
	}

	customUserModel struct {
		*defaultUserModel
	}
)

func (model *customUserModel) FindByUserID(user_id string) (*User, error) {

	if user_id == "" {
		return nil, nil
	}

	filter := bson.D{
		{Key: "user_id", Value: user_id},
		// {"creator", creator},
	}

	user := &User{}

	err := mgm.Coll(user).FindOne(
		mgm.Ctx(),
		filter,
		nil).Decode(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (model *customUserModel) FindByEmail(email string) (*User, error) {

	if email == "" {
		return nil, nil
	}

	filter := bson.D{
		{Key: "email", Value: email},
		// {"creator", creator},
	}

	user := &User{}

	err := mgm.Coll(user).FindOne(
		mgm.Ctx(),
		filter,
		nil).Decode(user)

	if err != nil {
		return nil, err
	}
	return user, nil
}

// NewUserModel returns a model for the mongo.
func NewUserModel(url, db, collection string, c cache.CacheConf) UserModel {
	conn := monc.MustNewModel(url, db, collection, c)
	return &customUserModel{
		defaultUserModel: newDefaultUserModel(conn),
	}
}

func NewUserModelV2() UserModel {
	return &customUserModel{
		defaultUserModel: newDefaultUserModel(nil),
	}
}

func (model *customUserModel) All() ([]*User, error) {
	var allUsers []*User

	coll := mgm.Coll(&User{})

	cursor, err := coll.Find(mgm.Ctx(), bson.M{}, nil)

	if err != nil {
		return nil, err
	}
	defer cursor.Close(mgm.Ctx())

	for cursor.Next(mgm.Ctx()) {
		user := &User{}
		if err := cursor.Decode(user); err != nil {
			return nil, err
		}
		allUsers = append(allUsers, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return allUsers, nil
}
