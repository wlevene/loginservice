package model

import (
	"github.com/wlevene/mgm/v3"
)

type {{.Type}} struct {
	mgm.DefaultModel `bson:",inline"`

	// TODO: Fill your own fields
}
