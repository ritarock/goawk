package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/urfave/cli/v2"
)

func main() {
	var command string
	var file string
	var field string

	app := cli.App{
		Name:  "goawk",
		Usage: "scanning word",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "command",
				Usage:       "set command",
				Aliases:     []string{"c"},
				Required:    true,
				Destination: &command,
			},
			&cli.StringFlag{
				Name:        "file",
				Usage:       "set file",
				Aliases:     []string{"f"},
				Required:    true,
				Destination: &file,
			},
			&cli.StringFlag{
				Name:        "field",
				Usage:       "set file",
				Aliases:     []string{"F"},
				Value:       " ",
				Destination: &field,
			},
		},
		Action: func(c *cli.Context) error {
			text := read(file)
			row, err := parseCommand(command)
			if err != nil {
				fmt.Println(err)
			}
			field, err := parseField(field)
			if err != nil {
				fmt.Println(err)
			}
			action(text, row, field)
			return nil
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func read(file string) (result []string) {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		panic(err)
	}
	return
}

func parseCommand(command string) (int, error) {
	if !strings.Contains(command, "{") || !strings.Contains(command, "}") {
		return 0, errors.New(command + "is not defined")
	}
	command = strings.TrimLeft(command, "{")
	command = strings.TrimRight(command, "}")

	commandParse := strings.Split(command, " ")
	if commandParse[0] != "print" {
		return 0, errors.New(commandParse[0] + "is not defined")
	}
	if !strings.Contains(commandParse[1], "$") {
		return 0, errors.New(commandParse[1] + "is not defined")
	}

	tmp := strings.TrimLeft(commandParse[1], "$")
	i, err := strconv.Atoi(tmp)
	if err != nil {
		return 0, errors.New(commandParse[1] + "is not defined")
	}
	return i - 1, nil
}

func parseField(field string) (string, error) {
	if !strings.Contains(field, "[") || !strings.Contains(field, "]") {
		return "", errors.New(field + "is not defined")
	}
	field = strings.TrimLeft(field, "[")
	field = strings.TrimRight(field, "]")

	return field, nil
}

func action(text []string, row int, field string) {
	switch row {
	case -1:
		for _, v := range text {
			fmt.Println(v)
		}
	default:
		for _, v := range text {
			arr := strings.Split(v, field)
			fmt.Println(arr[row])
		}
	}
}
