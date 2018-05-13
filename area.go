package main

type area struct {
	columnStart int
	columnEnd   int
	rowStart    int
	rowEnd      int
	view        view
}

func (a *area) SetView(v view) {
	a.view = v
}
