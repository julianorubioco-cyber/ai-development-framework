package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/julianorubioco-cyber/ai-development-framework/internal/adf"
)

func usage() {
	fmt.Fprintf(os.Stderr, `ADF %s

Uso:
  adf install [--dry-run]
  adf init [caminho] [--dry-run]
  adf detect [caminho]
  adf doctor
  adf uninstall [--restore-backup] [--dry-run]
  adf version

`, adf.Version)
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}

	var result any
	var err error

	switch os.Args[1] {
	case "install":
		flags := flag.NewFlagSet("install", flag.ExitOnError)
		dryRun := flags.Bool("dry-run", false, "simula sem alterar arquivos")
		_ = flags.Parse(os.Args[2:])
		result, err = adf.Install(*dryRun)

	case "init":
		flags := flag.NewFlagSet("init", flag.ExitOnError)
		dryRun := flags.Bool("dry-run", false, "simula sem alterar arquivos")
		_ = flags.Parse(os.Args[2:])
		path := "."
		if flags.NArg() > 0 {
			path = flags.Arg(0)
		}
		result, err = adf.InitWorkspace(path, *dryRun)

	case "detect":
		path := "."
		if len(os.Args) > 2 {
			path = os.Args[2]
		}
		result, err = adf.DetectCompatibility(path)

	case "doctor":
		result, err = adf.Doctor()

	case "uninstall":
		flags := flag.NewFlagSet("uninstall", flag.ExitOnError)
		restore := flags.Bool("restore-backup", false, "restaura Skills anteriores")
		dryRun := flags.Bool("dry-run", false, "simula sem alterar arquivos")
		_ = flags.Parse(os.Args[2:])
		result, err = adf.Uninstall(*restore, *dryRun)

	case "version", "--version", "-v":
		fmt.Printf("ADF %s\n", adf.Version)
		return

	case "help", "--help", "-h":
		usage()
		return

	default:
		err = errors.New("comando desconhecido: " + os.Args[1])
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, "Erro:", err)
		os.Exit(1)
	}
	if err := adf.PrintJSON(os.Stdout, result); err != nil {
		fmt.Fprintln(os.Stderr, "Erro:", err)
		os.Exit(1)
	}
}
