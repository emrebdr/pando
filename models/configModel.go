package models

type ConfigModel struct {
	Version     string       `yaml:"version"`
	Language    string       `yaml:"language"`
	Name        string       `yaml:"name"`
	Description string       `yaml:"description"`
	Destination string       `yaml:"destination"`
	Maintainers []Maintainer `yaml:"maintainers,omitempty"`
	Struct      Struct       `yaml:"struct"`
}

type Maintainer struct {
	Name  string `yaml:"name"`
	Email string `yaml:"email,omitempty"`
	Url   string `yaml:"url,omitempty"`
}

type Struct struct {
	Folders []Folder `yaml:"folders,omitempty"`
	Files   []File   `yaml:"files,omitempty"`
}

type Folder struct {
	Name        string   `yaml:"name"`
	Description string   `yaml:"description,omitempty"`
	Files       []File   `yaml:"files"`
	Folder      []Folder `yaml:"folder,omitempty"`
}

type File struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description,omitempty"`
	ContentType string `yaml:"content_type"`
	Content     string `yaml:"content"`
}
