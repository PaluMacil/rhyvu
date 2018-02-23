package ry

var ItemTypes = []string{"search", "link", "login", "logout", "header", "profile"}

type MenuItem struct {
	ID                 int
	Text               string
	ItemType           string
	Link               string
	Icon               string
	RequiredPermission Permission
	Order              int
	SubMenu            Menu
}

type Menu []MenuItem

func (m Menu) Len() int           { return len(m) }
func (m Menu) Swap(i, j int)      { m[i], m[j] = m[j], m[i] }
func (m Menu) Less(i, j int) bool { return m[i].Order < m[j].Order }
