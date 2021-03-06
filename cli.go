package main

import (
	"fmt"

	"github.com/amirography/rose/internal"
	"github.com/atotto/clipboard"
	"github.com/charmbracelet/lipgloss"
	"github.com/urfave/cli/v2"
)

var printFlags []cli.Flag = []cli.Flag{
	&cli.StringFlag{
		Name:    "colorname",
		Aliases: []string{"c"},
		Value:   "",
		Usage:   "Color name.",
	},
	&cli.StringFlag{
		Name:    "swatch",
		Aliases: []string{"s"},
		Value:   "rp",
		Usage:   "Swatch to use.",
	},
	&cli.StringFlag{
		Name:    "palette",
		Aliases: []string{"p"},
		Value:   "rosePine",
		Usage:   "Palette to use.(not developed for now)",
	},
	&cli.BoolFlag{
		Name:    "all",
		Aliases: []string{"a"},
		Usage:   "Palette to use.",
	},
}

var cpFlags []cli.Flag = []cli.Flag{
	&cli.StringFlag{
		Name:     "colorname",
		Aliases:  []string{"c"},
		Value:    "",
		Usage:    "Color name.",
		Required: true,
	},
	&cli.StringFlag{
		Name:    "swatch",
		Aliases: []string{"s"},
		Value:   "rp",
		Usage:   "Swatch to use.",
		EnvVars: []string{"DEFAULTSWATCH"},
	},
	&cli.StringFlag{
		Name:    "palette",
		Aliases: []string{"p"},
		Value:   "rosePine",
		Usage:   "Palette to use. (not developed for now)",
		EnvVars: []string{"DEFAULTPALLETE"},
	},
}

var App = &cli.App{
	Name:  "rose",
	Usage: "Your friendly Rosé Pine helper.",
	Commands: []*cli.Command{
		{
			Name:    "print",
			Aliases: []string{"p"},
			Usage:   "Print out the color code of the given string",
			Flags:   printFlags,
			Action:  printer,
		},
		{
			Name:    "copy",
			Aliases: []string{"cp"},
			Usage:   "Store code inside clipboard",
			Flags:   cpFlags,
			Action:  cp,
		},
	},
	Action: printall,

}

func printer(c *cli.Context) error {
	if c.Bool("all") {
		swatch := internal.List(c.String("swatch"))
		for _, ing := range swatch {
			col := ing.Hex
			block := fmt.Sprint(lipgloss.NewStyle().Background(lipgloss.Color(col)).PaddingRight(2).Render(""))
			code := fmt.Sprint(lipgloss.NewStyle().Bold(true).Italic(true).Render(block + "\t" + col + "\t\t" + ing.Name))
			fmt.Println(code)
		}
		return nil

	}

	col := internal.Get(c.String("colorname"), c.String("swatch"))
	block := fmt.Sprint(lipgloss.NewStyle().Background(lipgloss.Color(col)).PaddingRight(2).Render(""))
	code := fmt.Sprint(lipgloss.NewStyle().Bold(true).Italic(true).Render(block + " " + col))
	fmt.Println(code)
	return nil
}
func printall(c *cli.Context) error {
	swatch := internal.List("default")
	for _, ing := range swatch {
		col := ing.Hex
		block := fmt.Sprint(lipgloss.NewStyle().Background(lipgloss.Color(col)).PaddingRight(2).Render(""))
		code := fmt.Sprint(lipgloss.NewStyle().Bold(true).Italic(true).Render(block + "\t" + col + "\t\t" + ing.Name))
		fmt.Println(code)
	}
	return nil
}

func cp(c *cli.Context) error {
	col := internal.Get(c.String("colorname"), c.String("swatch"))
	err := clipboard.WriteAll(col)
	if err != nil {
		return err
	}
	return nil
}
