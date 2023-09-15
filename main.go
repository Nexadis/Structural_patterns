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

type Operation int

const (
	Add Operation = iota
	Sub
	Mul
	Div
)

type Calculator interface {
	Calculate() int
}

type Oper struct {
	Type  Operation
	Left  Calculator
	Right Calculator
}

func (o Oper) Calculate() int {
	switch {
	case o.Type == Add:
		log.Printf("%d+%d\n", o.Left.Calculate(), o.Right.Calculate())
		return o.Left.Calculate() + o.Right.Calculate()
	case o.Type == Sub:
		log.Printf("%d-%d\n", o.Left.Calculate(), o.Right.Calculate())
		return o.Left.Calculate() - o.Right.Calculate()
	case o.Type == Mul:
		log.Printf("%d*%d\n", o.Left.Calculate(), o.Right.Calculate())
		return o.Left.Calculate() * o.Right.Calculate()
	case o.Type == Div:
		log.Printf("%d/%d\n", o.Left.Calculate(), o.Right.Calculate())
		return o.Left.Calculate() / o.Right.Calculate()
	}
	return 0
}

type Number struct {
	Value int
}

func (n *Number) Calculate() int {
	return n.Value
}

func DemoComposite() {
	root := &Oper{
		Type: Div,
		Left: &Oper{
			Type: Mul,
			Left: &Oper{
				Type:  Add,
				Left:  &Number{Value: 2},
				Right: &Number{Value: 3},
			},
			Right: &Oper{
				Type:  Sub,
				Left:  &Number{Value: 77},
				Right: &Number{Value: 55},
			},
		},
		Right: &Number{Value: 2},
	}
	log.Println("Result:", root.Calculate())
}

func main() {
	border("Decorator")
	DemoDecorator()
	border("Adapter")
	DemoAdapter()
	border("Composite")
	DemoComposite()
}

func border(name string) {
	line := strings.Repeat("=", 80)
	out := fmt.Sprintf("%s\n\t\t\t\t%s\n%s", line, name, line)
	fmt.Println(out)
}
