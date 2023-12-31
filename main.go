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

// Server — интерфейс веб-сервера.
type Server interface {
	handleRequest(string, string) (int, string)
}

// WebServer — сервер по умолчанию.
type WebServer struct{}

// Proxy — прокси-сервер.
type Proxy struct {
	WebServer *WebServer
}

func NewProxyServer(webServer *WebServer) *Proxy {
	return &Proxy{
		WebServer: webServer,
	}
}

func (p *Proxy) handleRequest(url, method string) (int, string) {
	// запрещает доступ к /api/admin
	if strings.HasPrefix(url, "/api/admin") {
		return 403, "Forbidden"
	}
	return p.WebServer.handleRequest(url, method)
}

func (webServer *WebServer) handleRequest(url, method string) (int, string) {
	if !strings.HasPrefix(url, "/api/") {
		return 400, "Bad Request"
	}
	return 200, fmt.Sprintf("%s Request: %s", method, url)
}

func DemoProxy() {
	proxyServer := NewProxyServer(&WebServer{})
	for _, v := range []string{"/api/info", "/api/admin", "/api/version", "/admin"} {
		httpCode, body := proxyServer.handleRequest(v, "GET")
		log.Println(v, httpCode, body)
	}
}

// Computer — абстракция компьютера.
type Computer interface {
	Print()
	SetPrinter(Printer)
}

// Mac — компьютер Mac.
type Mac struct {
	printer Printer
}

func (m *Mac) Print() {
	fmt.Println("Печать для Mac.")
	m.printer.PrintFile()
}

func (m *Mac) SetPrinter(p Printer) {
	m.printer = p
}

// Windows — компьютер Windows.
type Windows struct {
	printer Printer
}

func (w *Windows) Print() {
	fmt.Println("Печать для Windows.")
	w.printer.PrintFile()
}

func (w *Windows) SetPrinter(p Printer) {
	w.printer = p
}

// Printer — интерфейс для принтера.
type Printer interface {
	PrintFile()
}

type Epson struct{}

func (p *Epson) PrintFile() {
	fmt.Println("Печать на принтере Epson.")
}

type HP struct{}

func (p *HP) PrintFile() {
	fmt.Println("Печать на принтере HP.")
}

func DemoBridge() {
	// создаём два принтера
	hp := &HP{}
	epson := &Epson{}

	// печать на Mac
	mac := &Mac{}
	mac.SetPrinter(hp)
	mac.Print()
	mac.SetPrinter(epson)
	mac.Print()

	// печать на Windows
	win := &Windows{}
	win.SetPrinter(hp)
	win.Print()
	win.SetPrinter(epson)
	win.Print()
}

func main() {
	border("Decorator")
	DemoDecorator()
	border("Adapter")
	DemoAdapter()
	border("Composite")
	DemoComposite()
	border("Proxy")
	DemoProxy()
	border("Bridge")
	DemoBridge()
}

func border(name string) {
	line := strings.Repeat("=", 80)
	out := fmt.Sprintf("%s\n\t\t\t\t%s\n%s", line, name, line)
	fmt.Println(out)
}
