package models

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	SeqDate string
	SendSeq int64
	RecvSeq int64
}
