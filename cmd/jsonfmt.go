package cmd

import (
	"encoding/json"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	jsonPretty "github.com/ansidev/json-pretty"
	"github.com/ansidev/jsonfmt/config"
	"github.com/urfave/cli/v2"
)

const (
	SPACE         byte   = ' '
	TAB           byte   = '\t'
	INDENT_SPACE  string = "w"
	INDENT_TAB    string = "t"
	INDENT_STRING string = "s"
)

func init() {
	log.SetFlags(0)
}

func commandFlags() []cli.Flag {
	return []cli.Flag{
		&cli.IntFlag{
			Name:    "width",
			Aliases: []string{"w"},
			Usage:   "Width is a max column width for single line arrays. Default is 80.",
			Value:   80,
		},
		&cli.StringFlag{
			Name:    "prefix",
			Aliases: []string{"p"},
			Usage:   "Prefix is a prefix for all lines. Default is an empty string.",
			Value:   "",
		},
		&cli.StringFlag{
			Name:    "indent",
			Aliases: []string{"i"},
			Usage: `Indent config. Format: '{indent_style}:{indent_value}'.
	Indent style: w (white space), t (tab), s (string).
	If indent_style is w or t, the indent value will be a number (indent size).
	If indent_style is s, the indent value will be indent string.
	Default indentation is 2 spaces.`,
			Value: "w:2",
		},
		&cli.BoolFlag{
			Name:    "sort-keys",
			Aliases: []string{"s"},
			Usage:   "Sort keys will sort the keys alphabetically. Default is false.",
			Value:   false,
		},
		&cli.BoolFlag{
			Name:    "color",
			Aliases: []string{"c"},
			Usage:   "Enable color, only applied if the output is terminal. Default is false.",
			Value:   false,
		},
		&cli.BoolFlag{
			Name:    "minify",
			Aliases: []string{"m"},
			Usage:   "Minify formatted json. These options will be ignored: -i, -w. Default is false.",
			Value:   false,
		},
		&cli.StringFlag{
			Name:    "output",
			Aliases: []string{"o"},
			Usage:   "Output path. Use this option if you want to save to file, by default the result will be printed to the stdout.",
		},
		&cli.BoolFlag{
			Name:    "debug",
			Aliases: []string{"d"},
			Usage:   "Enable debug mode. Default is false.",
			Value:   false,
		},
	}
}

func commandHandler() func(c *cli.Context) error {
	return func(c *cli.Context) error {
		indent := getIndent(c.String("indent"))
		cfg := config.JsonFmtConfig{
			Input:       c.Args().First(),
			Width:       c.Int("width"),
			Prefix:      c.String("prefix"),
			Indent:      indent,
			SortKeys:    c.Bool("sort-keys"),
			EnableColor: c.Bool("color"),
			Minify:      c.Bool("minify"),
			Output:      c.String("output"),
			Debug:       c.Bool("debug"),
		}

		if c.IsSet("output") {
			cfg.PrintToConsole = false
		} else {
			cfg.PrintToConsole = true
		}

		if c.IsSet("debug") {
			b, err := json.Marshal(cfg)
			if err != nil {
				return err
			} else {
				log.Println("jsonfmt config")
				log.Println("==============")
				log.Print(string(jsonPretty.Pretty(b)))
				log.Println("==============")
			}
		}

		jsonByte, err := os.ReadFile(cfg.Input)

		if err != nil {
			return err
		}

		out := formatJson(jsonByte, cfg)

		if cfg.PrintToConsole {
			log.Print(string(out))
			return nil
		}

		werr := os.WriteFile(cfg.Output, out, 0644)

		if werr != nil {
			log.Println("Failed to save formatted JSON to file", cfg.Output)
		} else {
			log.Println("Saved formatted JSON to file", cfg.Output, "successfully.")
		}

		return nil
	}
}

func getIndent(indentConfig string) string {
	re := regexp.MustCompile(`^(?P<indent_style>[wts]):(?P<indent_value>.+)$`)
	matches := re.FindStringSubmatch(indentConfig)

	if len(matches) != 3 {
		log.Fatalln("Invalid value for --indent:", indentConfig)
	}

	indentStyleIndex := re.SubexpIndex("indent_style")
	indentStyle := matches[indentStyleIndex]

	indentValueIndex := re.SubexpIndex("indent_value")
	indentValue := matches[indentValueIndex]

	switch indentStyle {
	case INDENT_SPACE, INDENT_TAB:
		indentSize, err := strconv.Atoi(indentValue)
		if err != nil {
			log.Fatalln("Invalid indent size value:", indentValue, err)
		}

		indentCharMap := map[string]byte{
			INDENT_SPACE: SPACE,
			INDENT_TAB:   TAB,
		}

		indentChar := indentCharMap[indentStyle]
		return generateIndent(indentChar, indentSize)
	case string(INDENT_STRING):
		return indentValue
	default:
		log.Fatalln("Invalid value for --indent:", indentConfig)
	}

	return ""
}

func generateIndent(char byte, size int) string {
	var sb strings.Builder
	for i := 0; i < size; i++ {
		sb.WriteByte(char)
	}

	return sb.String()
}

func formatJson(input []byte, cfg config.JsonFmtConfig) []byte {
	out := jsonPretty.PrettyOptions(input, &jsonPretty.Options{
		Width:    cfg.Width,
		Prefix:   cfg.Prefix,
		Indent:   cfg.Indent,
		SortKeys: cfg.SortKeys,
	})

	if cfg.Minify {
		out = jsonPretty.UglyInPlace(out)
	}

	if cfg.PrintToConsole && cfg.EnableColor {
		return jsonPretty.Color(out, nil)
	}

	return out
}

func initCliApp() *cli.App {
	return &cli.App{
		Name:                 "jsonfmt",
		Usage:                "JSON Formatter CLI",
		Flags:                commandFlags(),
		Action:               commandHandler(),
		EnableBashCompletion: true,
	}
}

func Run(args []string) {
	app := initCliApp()

	err := app.Run(args)
	if err != nil {
		log.Fatalln(err)
	}
}
