package parser

import (
	"ena/models"
	"ena/utils"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Parser struct {
	config  models.ConfigModel
	folders []string
}

func (p *Parser) readConfigFile(filename string) {
	f, err := os.ReadFile(filename)
	if err != nil {
		utils.LogFatal("Error while reading config file", err)
	}

	var config models.ConfigModel

	err = yaml.Unmarshal(f, &config)
	if err != nil {
		utils.LogFatal("Error while unmarshalling config file", err)
	}

	p.config = config
}

func InitParser(filename string) *Parser {
	p := &Parser{}
	p.readConfigFile(filename)
	return p
}

func (p *Parser) Parse() {

}

func (p *Parser) ParseFolders() []string {
	for _, folder := range p.config.Struct.Folders {
		for _, zf := range folder.Folders {
			fmt.Println(zf.Name)
		}
	}

	return []string{}
}
