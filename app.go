package main

type app struct {
	currentPage *page
}

func NewApp() *app {
	return &app{
		currentPage: &page{},
	}
}

func (a *app) GetCurrentPage() *page {
	return a.currentPage
}

func (a *app) Run() {
}
