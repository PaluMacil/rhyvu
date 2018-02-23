package ry

type Permission struct {
	ID          int
	Name        string `storm:"unique"`
	Description string
	Audit       `storm:"inline"`
}
