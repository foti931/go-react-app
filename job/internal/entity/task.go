package entity

import (
	"github.com/uptrace/bun"
)

type Task struct {
	bun.BaseModel `bun:"table:tasks"`
	ID            int `bun:",pk,autoincrement"`
	Title         string
	User          *User `bun:"rel:belongs-to, join:user_id=id"`
	UserId        uint
}
