package main

import "github.com/sparkymat/grip2/size"

type page struct {
	areas []area
}

func (p *page) ResetLayout(columnSizes []size.Size, rowSize []size.Size) {
}

func (p *page) Area(columnStart, columnEnd, rowStart, rowEnd int) *area {
	return &area{columnStart, columnEnd, rowStart, rowEnd, nil}
}

func (p *page) NewTextView(defaultText string) *textView {
	return &textView{
		text: defaultText,
	}
}
