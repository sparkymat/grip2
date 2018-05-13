package main

import "github.com/sparkymat/grip2/size"

func main() {
	app := NewApp()
	p := app.GetCurrentPage()

	p.ResetLayout(
		[]size.Size{size.Auto, size.WithCells(40)},
		[]size.Size{size.Auto, size.WithCells(3)},
	)

	mainArea := p.Area(0, 0, 0, 0)
	commandArea := p.Area(0, 0, 1, 1)
	sidebarArea := p.Area(1, 1, 0, 1)

	mainArea.SetView(p.NewTextView("Main area"))
	commandArea.SetView(p.NewTextView("Command area"))
	sidebarArea.SetView(p.NewTextView("Sidebar area"))

	app.Run()
}
