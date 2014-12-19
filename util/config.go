package util

import (
	"bytes"
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
)

type Config struct {
	Ref    string `toml:"ref"`
	Path   string `toml:"path"`
	Banner string `toml:"banner"`
}

const defaultConfig = `
banner = "This is the test world"
ref = "QmZXxbfUdRYi578pectWLFNFv5USQrsXdYAGeCsMJ6X8Zt"
path = "world.json"
`

func GetConfig(path string) (c *Config) {
	if _, err := toml.DecodeFile(path, &c); err != nil {
		fmt.Println("Config file does not exists,create config")

		fmt.Println(err)
		return
	}
	return
}

func (c *Config) SaveConfig() (err error) {
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(c); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf.String)
	return err
}
