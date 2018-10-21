package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ClientName           string `envconfig:"OBJECT_NAME" required:"true"`
	CapitalizeClientName string
	IsClass              bool
}

const (
	templateFilePath = "./client.tmpl"
)

func loadConfig(prefixEnv string) *Config {
	conf := &Config{}

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	err = envconfig.Process(prefixEnv, conf)
	if err != nil {
		log.Fatal(err)
	}

	return conf
}

func main() {
	conf := loadConfig("GEN")
	conf.CapitalizeClientName = strings.Title(conf.ClientName)
	conf.IsClass = strings.HasSuffix(strings.ToLower(conf.ClientName), "class")

	fmt.Println(conf)

	t, err := template.ParseFiles(templateFilePath)
	if err != nil {
		log.Fatal(err)
	}

	output, err := os.Create("./" + strings.ToLower(conf.ClientName) + ".go")
	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()

	err = t.Execute(output, conf)
	if err != nil {
		log.Fatal(err)
	}
}
