package ry

import (
	"time"
)

type Audit struct {
	ModifiedBy int `storm:"index"`
	Modified   time.Time
	CreatedBy  int `storm:"index"`
	Created    time.Time
}
