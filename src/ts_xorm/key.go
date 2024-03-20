package ts_xorm

import (
	"database/sql"
	"time"
)

type Key struct {
	Id         sql.NullInt64  `xorm:"pk autoincr comment('ID') bigint 'id'"`
	Name       string         `xorm:"not null varchar(255) 'name'"`
	NbId       sql.NullInt64  `xorm:"bigint 'nb_id'"`
	PubKey     string         `xorm:"not null varchar(255) 'pub_key'"`
	PvtKey     sql.NullString `xorm:"varchar(255) 'pvt_key'"`
	CreateTime time.Time      `xorm:"not null datetime 'create_time' created"`
	UserId     string         `xorm:"not null varchar(255) 'user_id'"`
}

func (k *Key) TableName() string {
	return "keys"
}
