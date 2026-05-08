package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/dansimco/go-design-tokens/color"
	"github.com/dansimco/go-design-tokens/figma"
	"github.com/dansimco/go-design-tokens/theme"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: go-ds <command>")
		fmt.Fprintln(os.Stderr, "  figma  export tokens to a Figma file as variables")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "figma":
		runFigma(os.Args[2:])
	default:
		fmt.Fprintf(os.Stderr, "unknown command: %s\n", os.Args[1])
		os.Exit(1)
	}
}

func runFigma(args []string) {
	fs := flag.NewFlagSet("figma", flag.ExitOnError)
	endpoint := fs.String("mcp-endpoint", figma.DefaultMCPEndpoint, "Figma local MCP server endpoint")
	dryRun := fs.Bool("dry-run", false, "print Plugin API script to stdout without calling the MCP server")
	listTools := fs.Bool("list-tools", false, "print available MCP tools and their schemas then exit")
	fs.Parse(args)

	client := figma.NewMCPClient(*endpoint)

	if *listTools {
		tools, err := client.ListTools()
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(tools)
		return
	}

	t := buildTheme()
	script := figma.ThemeToPluginScript(&t)

	if *dryRun {
		fmt.Println(script)
		return
	}

	cmd := exec.Command("pbcopy")
	cmd.Stdin = strings.NewReader(script)
	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "error copying to clipboard: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Script copied to clipboard.")
	fmt.Println("In Figma: Plugins → Development → Open console → paste → Enter.")
}

func buildTheme() theme.Theme {
	t := theme.New()

	ink := color.NewRamp()
	ink.AddKey("#FEFCFF", 0)
	ink.AddKey("#64617A", 0.562)
	ink.AddKey("#020103", 1)
	t.AddRamp("ink", ink)

	intlOrange := color.NewRamp()
	intlOrange.AddKey("#FEFCFF", 0)
	intlOrange.AddKey("#FF4F01", 0.406)
	intlOrange.AddKey("#010000", 1)
	t.AddRamp("intlOrange", intlOrange)

	azimuth := color.NewRamp()
	azimuth.AddKey("#F9F9FB", 0)
	azimuth.AddKey("#6857DD", 0.4)
	azimuth.AddKey("#05000f", 1)
	t.AddRamp("azimuth", azimuth)

	patrick := color.NewRamp()
	patrick.AddKey("#FFFFFF", 0)
	patrick.AddKey("#DD79D8", 0.4)
	patrick.AddKey("#05000f", 1)
	t.AddRamp("patrick", patrick)

	kelp := color.NewRamp()
	kelp.AddKey("#EBFFE0", 0)
	kelp.AddKey("#6E9B55", 0.53)
	kelp.AddKey("#000000", 1)
	t.AddRamp("kelp", kelp)

	seaFoam := color.NewRamp()
	seaFoam.AddKey("#E2FFFF", 0)
	seaFoam.AddKey("#7AD0D3", 0.375)
	seaFoam.AddKey("#010511", 1)
	t.AddRamp("seaFoam", seaFoam)

	electric := color.NewRamp()
	electric.AddKey("#FFFFD8", 0)
	electric.AddKey("#F8FF6C", 0.375)
	electric.AddKey("#080a00", 1)
	t.AddRamp("electric", electric)

	neutral := t.AddColorMode("neutral")
	content := neutral.AddRole("content")
	content.AddState("default", color.UIColor{Light: ink.At(1), Dark: ink.At(1)})
	content.AddState("hover", color.UIColor{Light: ink.At(0.28), Dark: ink.At(0.78)})

	return t
}
