package models

import (
	"time"

	"github.com/uadmin/uadmin"
)

type Справочники struct {
	uadmin.Model
	Name        string
	Description string `uadmin:"html"`
	TargetDate  time.Time
}
