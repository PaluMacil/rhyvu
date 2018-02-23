package ry

import (
	"fmt"
	"html/template"
	"sort"
	"strings"
)

type Application struct {
	Settings Settings
	Session  *Session
	Menu     Menu
}

func (app Application) Init(token string) (Application, error) {
	// Load Settings
	if err := Db.All(&app.Settings); err != nil {
		return app, err
	}

	// Load Session (if token provided) and only add if found
	if token != "" {
		var s Session
		if err := Db.One("Token", token, &s); err == nil {
			*app.Session = s
		}
	}

	// Load Menu
	if err := Db.Select().OrderBy("Order").Find(&app.Menu); err != nil {
		return app, err
	}
	// Sort submenus
	for _, m := range app.Menu {
		if len(m.SubMenu) > 1 {
			sort.Sort(m.SubMenu)
		}
	}

	return app, nil
}

func (app Application) RenderMenu() (template.HTML, error) {
	var b strings.Builder
	var menu Menu
	err := Db.Select().OrderBy("Order").Find(&menu)
	if err != nil {
		return "", fmt.Errorf("could not load menu: %s", err)
	}
	for _, m := range menu {
		switch m.ItemType {
		case "search":
		case "link":

		case "login":
		case "logout":
		case "header":
		case "profile":
		}
	}
	return template.HTML(b.String()), nil
}
