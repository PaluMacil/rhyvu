package ry

type Group struct {
	ID       int
	HomePage string
	Primary  bool
}

type GroupPermission struct {
	ID           int
	GroupID      int
	PermissionID int
}
