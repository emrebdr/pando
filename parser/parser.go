package parser

import (
	"ena/models"
	"ena/utils"
	"io/ioutil"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

type Parser struct {
	config models.ConfigModel
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

func Parse(filename string, destination string) {
	p := &Parser{}
	p.readConfigFile(filename)
	p.setDestinationAddress(destination)
	p.parse()
}

func (p *Parser) parse() {
	p.createFolderTemplateStruct()
	p.createFilesTemplateSturct()
}

func (p *Parser) createStruct(folders []models.Folder, path string) {
	path = path + "/"
	for _, folder := range folders {
		p.createFolder(path + folder.Name)
		p.createFiles(folder.Files, path+folder.Name)
		if folder.Folder != nil {
			p.createStruct(folder.Folder, path+folder.Name)
		}
	}
}

func (p *Parser) createFolder(name string) {
	err := os.Mkdir(name, 0755)
	if err != nil {
		utils.LogFatal("Occurs error while creating a folder", err)
	}
}

func (p *Parser) createFiles(files []models.File, path string) {
	for _, file := range files {
		f, err := os.Create(path + "/" + file.Name)
		defer f.Close()
		if err != nil {
			utils.LogFatal("Invalid file name. Occurs error while creating a file.", err)
			return
		}

		p.writeToFileContent(f, file.ContentType, file.Content)
	}
}

func (p *Parser) writeToFileContent(file *os.File, content_type string, content string) {
	data := p.getConent(content_type, content)

	_, err := file.Write([]byte(data))
	if err != nil {
		utils.LogFatal("Occurs error while writing content to file.", err)
	}
}

func (p *Parser) getConent(content_type string, content string) string {
	switch content_type {
	case "text":
		return content
	case "file":
		data, err := os.ReadFile(p.config.Destination + "/" + content)
		if err != nil {
			utils.LogFatal("File not found. Does not exist file path.", err)
			return ""
		}
		return string(data)
	case "url":
		resp, err := http.Get(content)
		if err != nil {
			utils.LogFatal("Something went wrong...", err) // TBC
			return ""
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			utils.LogFatal("Something went wrong...", err) // TBC
			return ""
		}

		return string(body)
	}

	return ""
}

func (p *Parser) checkDestinationAddressAlreadyExist() {
	if p.config.Destination == "" {
		p.config.Destination = "./"
	} else {
		if _, err := os.Stat(p.config.Destination); os.IsNotExist(err) {
			os.Mkdir(p.config.Destination, 0755)
		}
	}
}

func (p *Parser) createFolderTemplateStruct() {
	p.checkDestinationAddressAlreadyExist()
	p.createStruct(p.config.Struct.Folders, p.config.Destination)
}

func (p *Parser) createFilesTemplateSturct() {
	p.createFiles(p.config.Struct.Files, p.config.Destination)
}

func (p *Parser) setDestinationAddress(address string) {
	if p.config.Destination == "" {
		p.config.Destination = address
		return
	}
}
