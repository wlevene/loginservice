// Code generated by goctl. DO NOT EDIT.
package accountlocks

import (
	"context"
	"errors"

	"github.com/wlevene/mgm/v3"
	"github.com/zeromicro/go-zero/core/stores/monc"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var prefixAccountLocksCacheKey = "cache:accountLocks:"

type accountLocksModel interface {
	Insert(ctx context.Context, data *AccountLocks) error
	FindOne(ctx context.Context, id string) (*AccountLocks, error)
	Update(ctx context.Context, data *AccountLocks) error
	Delete(ctx context.Context, id string) error
	DeleteForce(ctx context.Context, id string) error
}

type defaultAccountLocksModel struct {
	conn *monc.Model
}

func newDefaultAccountLocksModel(conn *monc.Model) *defaultAccountLocksModel {
	return &defaultAccountLocksModel{conn: conn}
}

func (m *defaultAccountLocksModel) Insert(ctx context.Context, data *AccountLocks) error {
	return mgm.Coll(data).Create(data)
}

func (m *defaultAccountLocksModel) FindOne(ctx context.Context, id string) (*AccountLocks, error) {
	_, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, ErrInvalidObjectId
	}

	var data AccountLocks

	coll := mgm.Coll(&data)
	err = coll.FindByID(id, &data)

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (m *defaultAccountLocksModel) Update(ctx context.Context, data *AccountLocks) error {
	return mgm.Coll(data).Update(data)
}

func (m *defaultAccountLocksModel) Delete(ctx context.Context, id string) error {
	_, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrInvalidObjectId
	}

	save_obj, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	if save_obj == nil {
		return errors.New("data is nil")
	}

	save_obj.SetDeleted(true)
	return nil
}

func (m *defaultAccountLocksModel) DeleteForce(ctx context.Context, id string) error {
	_, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return ErrInvalidObjectId
	}

	save_obj, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	if save_obj == nil {
		return errors.New("data is nil")
	}

	mgm.Coll(save_obj).Delete(save_obj)
	return nil
}
