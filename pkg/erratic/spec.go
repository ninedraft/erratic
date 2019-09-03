package erratic

type Package struct {
	Name   string  `json:"package" yaml:"package" toml:"package"`
	Errors []Error `json:"errors" toml:"error"`
}

type Error struct {
	Name   string  `json:"name" toml:"name"`
	Code   int     `json:"code" toml:"code"`
	Fields []Field `json:"fields" toml:"field"`
}

type Field struct {
	Name    string            `json:"name"`
	Type    string            `json:"type"`
	Tags    map[string]string `json:"tags"`
	Default interface{}       `json:"default"`
}
