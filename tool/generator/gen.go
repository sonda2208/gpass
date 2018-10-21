package main

import (
	"html/template"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	ClientName         string `envconfig:"OBJECT_NAME" required:"true"`
	TemplateFilePath   string
	RequiredFields     string `envconfig:"REQUIRED_FIELDS" required:"true"`
	SampleID           string
	SampleDataPath     string
	SampleListDataPath string
}

type GeneratorConfig struct {
	ClientName           string
	CapitalizeClientName string
	IsClass              bool
	RequiredFields       []string
	SampleID             string
	SampleData           string
	SampleListData       string
}

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

func generateCode(conf *GeneratorConfig, tmlpPath, outputFile string) error {
	t, err := template.ParseFiles(tmlpPath)
	if err != nil {
		return err
	}

	output, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer output.Close()

	err = t.Execute(output, conf)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	conf := loadConfig("GEN")

	genConf := &GeneratorConfig{}
	genConf.ClientName = conf.ClientName
	genConf.CapitalizeClientName = strings.Title(conf.ClientName)
	genConf.IsClass = strings.HasSuffix(strings.ToLower(conf.ClientName), "class")
	genConf.RequiredFields = strings.Split(conf.RequiredFields, ",")

	err := generateCode(genConf, conf.TemplateFilePath, "./"+strings.ToLower(conf.ClientName)+".go")
	if err != nil {
		log.Fatal(err)
	}
}
