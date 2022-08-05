package config

type JsonFmtConfig struct {
	Input          string `json:"input"`
	Width          int    `json:"width"`
	Prefix         string `json:"prefix"`
	Indent         string `json:"indent"`
	SortKeys       bool   `json:"sort-keys"`
	PrintToConsole bool   `json:"print-to-console"`
	EnableColor    bool   `json:"enable_color"`
	Minify         bool   `json:"minify"`
	Output         string `json:"output"`
	Debug          bool   `json:"debug"`
}
