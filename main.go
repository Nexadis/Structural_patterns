package main

import (
	"fmt"
	"log"
	"strings"
)

type Modifier interface {
	Modify() string
}

type Original struct {
	Value string
}

func (o *Original) Modify() string {
	return o.Value
}

// Upper возвращает строку в верхнем регистре.
type Upper struct {
	modifier Modifier
}

func (u *Upper) Modify() string {
	return strings.ToUpper(u.modifier.Modify())
}

// Replace заменяет строки old на new.
type Replace struct {
	modifier Modifier
	old      string
	new      string
}

func (r *Replace) Modify() string {
	return strings.Replace(r.modifier.Modify(), r.old, r.new, -1)
}

func DemoDecorator() {
	original := &Original{Value: "Привет, гофер!"}
	log.Println(original.Modify())
	replace := &Replace{
		modifier: original,
		old:      "гофер",
		new:      "мир",
	}
	upper := &Upper{
		modifier: replace,
	}
	log.Println(upper.Modify())
}

// JSONData — интерфейс для декодирования JSON.
type JSONData interface {
	DecodeJSON() interface{}
}

// YAMLData — интерфейс для декодирования YAML.
type YAMLData interface {
	DecodeYAML() interface{}
}

type Client struct {
	Data interface{}
}

func (client *Client) Decode(input JSONData) {
	client.Data = input.DecodeJSON()
}

type Adapter struct {
	yaml YAMLData
}

func (a *Adapter) DecodeJSON() interface{} {
	return a.yaml.DecodeYAML()
}

func Load(client *Client, input YAMLData) {
	adapter := &Adapter{
		yaml: input,
	}
	client.Decode(adapter)
}

type someYaml struct{}

func (s someYaml) DecodeYAML() interface{} {
	return "Some YAML"
}

func DemoAdapter() {
	client := &Client{}
	input := someYaml{}
	Load(client, input)
	fmt.Printf("json client: %v, yaml: %v\n", client.Data, input.DecodeYAML())
}

func main() {
	border("Decorator")
	DemoDecorator()
	border("Adapter")
	DemoAdapter()
}

func border(name string) {
	line := strings.Repeat("=", 80)
	out := fmt.Sprintf("%s\n\t\t\t\t%s\n%s", line, name, line)
	fmt.Println(out)
}
