package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"

	"golang.org/x/mod/modfile"

	"github.com/walterwanderley/xo-grpc/metadata"
)

var (
	databaseDriverModule string
	databaseDriverName   string
	module               string
	showVersion          bool
	help                 bool
	verbose              bool
)

func main() {
	flag.BoolVar(&help, "h", false, "Help for this program")
	flag.BoolVar(&showVersion, "v", false, "Show version")
	flag.BoolVar(&verbose, "verbose", false, "Verbose")
	flag.StringVar(&module, "m", "my-project", "Go module name if there are no go.mod")
	flag.StringVar(&databaseDriverModule, "db-module", "github.com/jackc/pgx/v4/stdlib", "Database driver module")
	flag.StringVar(&databaseDriverName, "db-driver", "pgx", "Database driver name")
	flag.Parse()

	if help {
		printHelp()
		return
	}

	if showVersion {
		fmt.Println(version)
		return
	}

	if len(flag.Args()) < 1 {
		fmt.Printf("\nmodelsPath is required\n\n")
		printHelp()
		os.Exit(1)
	}

	modelsPath := strings.TrimPrefix(strings.TrimPrefix(strings.TrimSuffix(flag.Arg(0), "/"), "."), "/")

	if m := moduleFromGoMod(); m != "" {
		fmt.Println("Using module path from go.mod:", m)
		module = m
	}

	def := metadata.Definition{
		GoModule:             module,
		DatabaseDriverModule: databaseDriverModule,
		DatabaseDriverName:   databaseDriverName,
		ModelsPath:           modelsPath,
		Packages:             make([]*metadata.Package, 0),
	}

	pkgs, err := metadata.ParsePackages(modelsPath, module)
	if err != nil {
		log.Fatal("parser error:", err.Error())
	}
	def.Packages = pkgs

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("unable to get working directory:", err.Error())
	}

	err = process(&def, wd)
	if err != nil {
		log.Fatal("unable to process templates:", err.Error())
	}

	postProcess(&def, wd)
}

func moduleFromGoMod() string {
	f, err := os.Open("go.mod")
	if err != nil {
		return ""
	}
	defer f.Close()

	b, err := io.ReadAll(f)
	if err != nil {
		return ""
	}

	return modfile.ModulePath(b)
}

func postProcess(def *metadata.Definition, workingDirectory string) {
	fmt.Printf("Configuring project %s...\n", def.GoModule)
	execCommand("go mod init " + def.GoModule)
	execCommand("go mod tidy")
	execCommand("go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway " +
		"github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 " +
		"google.golang.org/protobuf/cmd/protoc-gen-go " +
		"google.golang.org/grpc/cmd/protoc-gen-go-grpc " +
		"github.com/bufbuild/buf/cmd/buf")
	fmt.Println("Compiling protocol buffers...")
	if err := os.Chdir("proto"); err != nil {
		panic(err)
	}
	execCommand("buf mod update")
	if err := os.Chdir(workingDirectory); err != nil {
		panic(err)
	}
	execCommand("buf generate")
	execCommand("go mod tidy")
	fmt.Println("Finished!")
}

func execCommand(command string) error {
	line := strings.Split(command, " ")
	cmd := exec.Command(line[0], line[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("[error] %q: %w", command, err)
	}
	return nil
}

func printHelp() {
	fmt.Println("xo-grpc [flags] modelsPath")
	fmt.Println("\nflags:")
	flag.PrintDefaults()
	fmt.Println("\nFor more information, please visit https://github.com/walterwanderley/xo-grpc")
}
