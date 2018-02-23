package ry

var ItemTypes = []string{""}

type MenuItem struct {
	ID                  int
	Text                string
	ItemType            string
	Link                string
	Icon                string
	RequiredPermissions []Permission
	Order               int
	SubMenu             []MenuItem
}

type Menu []MenuItem
