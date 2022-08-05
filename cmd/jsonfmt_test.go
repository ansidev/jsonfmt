package cmd

import (
	"bytes"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	TestFile = "../test/sample.json"
)

func runTest(args []string) error {
	app := initCliApp()

	err := app.Run(args)
	if err != nil {
		return err
	}

	return nil
}

func getOutputPath() string {
	return "output.json"
}

func Test_Command_WithValidArguments(t *testing.T) {
	type args struct {
		args   []string
		testFn func(t *testing.T, out bytes.Buffer)
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "format with -c",
			args: args{
				args: []string{"-c", TestFile},
				testFn: func(t *testing.T, out bytes.Buffer) {
					expected := "{\n  \x1b[94m\"id\"\x1b[0m: \x1b[92m\"56a4f58f-745b-402d-90d5-61cc863bc588\"\x1b[0m,\n  \x1b[94m\"first_name\"\x1b[0m: \x1b[92m\"John\"\x1b[0m,\n  \x1b[94m\"last_name\"\x1b[0m: \x1b[92m\"Doe\"\x1b[0m,\n  \x1b[94m\"email\"\x1b[0m: \x1b[92m\"john.doe@example.com\"\x1b[0m\n}\n"
					assert.Equal(t, expected, out.String())
				},
			},
		},
		{
			name: "format with -d",
			args: args{
				args: []string{"-d", TestFile},
				testFn: func(t *testing.T, out bytes.Buffer) {
					expected := `jsonfmt config
==============
{
  "input": "../test/sample.json",
  "width": 80,
  "prefix": "",
  "indent": "  ",
  "sort-keys": false,
  "print-to-console": true,
  "enable_color": false,
  "minify": false,
  "output": "",
  "debug": true
}
==============
{
  "id": "56a4f58f-745b-402d-90d5-61cc863bc588",
  "first_name": "John",
  "last_name": "Doe",
  "email": "john.doe@example.com"
}
`
					assert.Equal(t, expected, out.String())
				},
			},
		},
		{
			name: "format with -i='w:2'",
			args: args{
				args: []string{"-i", "w:2", TestFile},
				testFn: func(t *testing.T, out bytes.Buffer) {
					expected := `{
  "id": "56a4f58f-745b-402d-90d5-61cc863bc588",
  "first_name": "John",
  "last_name": "Doe",
  "email": "john.doe@example.com"
}
`
					assert.Equal(t, expected, out.String())
				},
			},
		},
		{
			name: "format with -i='t:1'",
			args: args{
				args: []string{"-i", "t:1", TestFile},
				testFn: func(t *testing.T, out bytes.Buffer) {
					expected := `{
	"id": "56a4f58f-745b-402d-90d5-61cc863bc588",
	"first_name": "John",
	"last_name": "Doe",
	"email": "john.doe@example.com"
}
`
					assert.Equal(t, expected, out.String())
				},
			},
		},
		{
			name: "format with -i='s:foo'",
			args: args{
				args: []string{"-i", "s:foo", TestFile},
				testFn: func(t *testing.T, out bytes.Buffer) {
					expected := `{
foo"id": "56a4f58f-745b-402d-90d5-61cc863bc588",
foo"first_name": "John",
foo"last_name": "Doe",
foo"email": "john.doe@example.com"
}
`
					assert.Equal(t, expected, out.String())
				},
			},
		},
		{
			name: "format with -m",
			args: args{
				args: []string{"-m", TestFile},
				testFn: func(t *testing.T, out bytes.Buffer) {
					expected := `{"id":"56a4f58f-745b-402d-90d5-61cc863bc588","first_name":"John","last_name":"Doe","email":"john.doe@example.com"}
`
					assert.Equal(t, expected, out.String())
				},
			},
		},
		{
			name: "format with -o",
			args: args{
				args: []string{"-o", getOutputPath(), TestFile},
				testFn: func(t *testing.T, out bytes.Buffer) {
					b, err := os.ReadFile(getOutputPath())
					assert.NoError(t, err)
					expected := `{
  "id": "56a4f58f-745b-402d-90d5-61cc863bc588",
  "first_name": "John",
  "last_name": "Doe",
  "email": "john.doe@example.com"
}
`
					assert.Equal(t, expected, string(b))
					os.Remove(getOutputPath())
				},
			},
		},
		{
			name: "format with -p",
			args: args{
				args: []string{"-p", "foo", TestFile},
				testFn: func(t *testing.T, out bytes.Buffer) {
					expected := `foo{
foo  "id": "56a4f58f-745b-402d-90d5-61cc863bc588",
foo  "first_name": "John",
foo  "last_name": "Doe",
foo  "email": "john.doe@example.com"
foo}
`
					assert.Equal(t, expected, out.String())
				},
			},
		},
		{
			name: "format with -s",
			args: args{
				args: []string{"-s", TestFile},
				testFn: func(t *testing.T, out bytes.Buffer) {
					expected := `{
  "email": "john.doe@example.com",
  "first_name": "John",
  "id": "56a4f58f-745b-402d-90d5-61cc863bc588",
  "last_name": "Doe"
}
`
					assert.Equal(t, expected, out.String())
				},
			},
		},
		{
			name: "format with -w",
			args: args{
				args: []string{"-w", "30", TestFile},
				testFn: func(t *testing.T, out bytes.Buffer) {
					expected := `{
  "id": "56a4f58f-745b-402d-90d5-61cc863bc588",
  "first_name": "John",
  "last_name": "Doe",
  "email": "john.doe@example.com"
}
`
					assert.Equal(t, expected, out.String())
				},
			},
		},
	}

	defer func() {
		log.SetOutput(os.Stderr)
	}()
	for _, tt := range tests {
		var buf bytes.Buffer
		log.SetOutput(&buf)
		t.Run(tt.name, func(t *testing.T) {
			testArgs := os.Args[0:1]
			testArgs = append(testArgs, tt.args.args...)

			err := runTest(testArgs)
			assert.NoError(t, err)

			tt.args.testFn(t, buf)
		})
	}
}
