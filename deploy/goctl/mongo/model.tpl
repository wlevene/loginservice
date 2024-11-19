// Code generated by goctl. DO NOT EDIT.
package {{.lowerType}}

import (
    "context"
	"errors"

    "github.com/wlevene/mgm/v3"
    "go.mongodb.org/mongo-driver/bson/primitive"
    {{if .Cache}}"github.com/zeromicro/go-zero/core/stores/monc"{{else}}"github.com/zeromicro/go-zero/core/stores/mon"{{end}}
   
)

{{if .Cache}}var prefix{{.Type}}CacheKey = "cache:{{.lowerType}}:"{{end}}

type {{.lowerType}}Model interface{
    Insert(ctx context.Context,data *{{.Type}}) error
    FindOne(ctx context.Context,id string) (*{{.Type}}, error)
    Update(ctx context.Context,data *{{.Type}}) error
    Delete(ctx context.Context,id string) error
	DeleteForce(ctx context.Context, id string) error
}

type default{{.Type}}Model struct {
    conn {{if .Cache}}*monc.Model{{else}}*mon.Model{{end}}
}

func newDefault{{.Type}}Model(conn {{if .Cache}}*monc.Model{{else}}*mon.Model{{end}}) *default{{.Type}}Model {
    return &default{{.Type}}Model{conn: conn}
}


func (m *default{{.Type}}Model) Insert(ctx context.Context, data *{{.Type}}) error {
   return mgm.Coll(data).Create(data)
}

func (m *default{{.Type}}Model) FindOne(ctx context.Context, id string) (*{{.Type}}, error) {
    _, err := primitive.ObjectIDFromHex(id)
    if err != nil {
        return nil, ErrInvalidObjectId
    }

    var data {{.Type}}
 
	coll := mgm.Coll(&data)
	err = coll.FindByID(id, &data)

	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (m *default{{.Type}}Model) Update(ctx context.Context, data *{{.Type}}) error {
   return mgm.Coll(data).Update(data)
}

func (m *default{{.Type}}Model) Delete(ctx context.Context, id string) error {
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


func (m *default{{.Type}}Model) DeleteForce(ctx context.Context, id string) error {
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