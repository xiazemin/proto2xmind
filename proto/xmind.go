package gen

import (
	"encoding/json"

	xmindgo "github.com/xiazemin/xmind-go"
	xmlXmind "github.com/xiazemin/xmind-go/xxml"
)

type Xmind struct {
	sheet []*xmlXmind.XmindNode
}

func NewXmind() *Xmind {
	return &Xmind{
		sheet: make([]*xmlXmind.XmindNode, 0),
	}
}

func (x *Xmind) Marshal(dst string) {
	sheetData, _ := json.Marshal(x.sheet)
	err := xmindgo.SaveSheets(dst, string(sheetData))
	if err != nil {
		panic(err)
	}
}
