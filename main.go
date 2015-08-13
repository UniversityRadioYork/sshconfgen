package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/BurntSushi/toml"
)

type server struct {
	Fqdn    string
	Aliases []string
}

type Servers []server

type Conf struct {
	Username string
	Proxy    string
	Servers  Servers
}

type TemplateData struct {
	Conf
	DateTime string
}

func gethostname(fqdn string) string {
	return strings.Split(fqdn, ".")[0]
}

func main() {
	rawconf, err := ioutil.ReadFile("config.toml")
	if err != nil {
		log.Fatal(err)
	}
	var conf Conf
	if err = toml.Unmarshal(rawconf, &conf); err != nil {
		log.Fatal(err)
	}
	data := TemplateData{
		conf,
		time.Now().Format("2006-01-02"),
	}
	funcmap := template.FuncMap{
		"gethostname": gethostname,
	}
	tmpl := template.Must(template.New("template.txt").Funcs(funcmap).ParseFiles("template.txt"))
	if err := tmpl.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}
