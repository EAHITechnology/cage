package utils

type ConfigInfo struct {
	Test1 string `toml:"test1"`
}

var Config = ConfigInfo{}
