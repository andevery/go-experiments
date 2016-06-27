package main

import (
	"bytes"
	"encoding/json"
	"github.com/andevery/go-experiments/marshal-vs-template/qtmp"
	"log"
	"sync"
	"text/template"
)

type App struct {
	data    []*qtmp.Data
	tmpl    *template.Template
	buffers sync.Pool
}

func NewApp(dataCount int) *App {
	data := make([]*qtmp.Data, dataCount)
	for i, _ := range data {
		data[i] = &qtmp.Data{ID: int64(i), Name: "Name"}
	}
	tmpl, err := template.New("Data").
		Parse(`[{{range $ind, $data := .}}{{if gt $ind 0}},{{end}}{"ID":{{.ID}},"Name":"{{.Name}}"}{{end}}]`)
	if err != nil {
		log.Fatal(err)
	}
	return &App{data, tmpl, sync.Pool{}}
}

func (a *App) newBuffer() *bytes.Buffer {
	if b := a.buffers.Get(); b != nil {
		buf := b.(*bytes.Buffer)
		buf.Reset()
		return buf
	}
	return new(bytes.Buffer)
}

func (a *App) MarshalRender() []byte {
	buf := a.newBuffer()
	defer a.buffers.Put(buf)
	json.NewEncoder(buf).Encode(&a.data)
	return buf.Bytes()
}

func (a *App) TemplateRender() []byte {
	buf := a.newBuffer()
	defer a.buffers.Put(buf)
	err := a.tmpl.Execute(buf, a.data)
	if err != nil {
		log.Fatal(err)
	}
	buf.WriteByte('\n')
	return buf.Bytes()
}

func (a *App) QuickTemplateRender() []byte {
	buf := a.newBuffer()
	defer a.buffers.Put(buf)
	qtmp.WriteRender(buf, a.data)
	buf.WriteByte('\n')
	return buf.Bytes()
}

func main() {
	a := NewApp(2)
	log.Println(string(a.TemplateRender()))
	log.Println(string(a.QuickTemplateRender()))
	log.Println(string(a.MarshalRender()))
}
