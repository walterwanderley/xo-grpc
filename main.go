package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"

	"golang.org/x/mod/modfile"

	"github.com/walterwanderley/xo-grpc/metadata"
)

var (
	databaseDriverModule string
	databaseDriverName   string
	module               string
	workers              int
	showVersion          bool
	help                 bool
	verbose              bool
)

func main() {
	flag.BoolVar(&help, "h", false, "Help for this program")
	flag.BoolVar(&showVersion, "v", false, "Show version")
	flag.BoolVar(&verbose, "verbose", false, "Verbose")
	flag.IntVar(&workers, "w", 8, "Number of go routines to call protoc")
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

	if workers <= 0 {
		fmt.Println("\n-w must be a positive number")
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
	fmt.Println("Running compile.sh...")
	protos := make([]string, 0)
	protos = append(protos, "typespb")
	for _, pkg := range def.Packages {
		pkgName := metadata.ToSnakeCase(pkg.Package)
		protos = append(protos, pkgName)
	}

	newDir := filepath.Join(workingDirectory, "proto")
	if _, err := os.Stat(newDir); os.IsNotExist(err) {
		err := os.MkdirAll(newDir, 0750)
		if err != nil {
			panic(err)
		}
	}

	if err := os.Chdir(newDir); err != nil {
		panic(err)
	}

	pkgs := make(chan string, workers)
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go compileProto(pkgs, workingDirectory, &wg)
	}

	for _, p := range protos {
		newDir := filepath.Join(workingDirectory, "proto", p)
		if _, err := os.Stat(newDir); os.IsNotExist(err) {
			err := os.MkdirAll(newDir, 0750)
			if err != nil {
				panic(err)
			}
		}
		pkgs <- p
	}

	close(pkgs)
	wg.Wait()

	fmt.Println("Generating OpenAPIv2 specs...")
	execCommand("protoc -I. -I3rd-party --openapiv2_out . --openapiv2_opt logtostderr=true,allow_repeated_fields_in_body=true,generate_unbound_methods=true,allow_merge=true " + strings.Join(protos, ".proto ") + ".proto")

	if err := os.Chdir(workingDirectory); err != nil {
		panic(err)
	}

	fmt.Printf("Configuring project %s...\n", def.GoModule)
	execCommand("go mod init " + def.GoModule)
	execCommand("go mod tidy")

	fmt.Println("Finished!")
}

func compileProto(pkgs <-chan string, wd string, wg *sync.WaitGroup) {
	defer wg.Done()
	for pkg := range pkgs {
		fmt.Printf("Compiling %s.proto...\n", pkg)
		err := execCommand(fmt.Sprintf("protoc -I. -I3rd-party --go_out %s --go_opt paths=source_relative --go-grpc_out %s --go-grpc_opt paths=source_relative %s.proto", pkg, pkg, pkg))
		if err != nil {
			return
		}
		fmt.Printf("Generating reverse proxy (grpc-gateway) %s.proto...\n", pkg)
		execCommand(fmt.Sprintf("protoc -I. -I3rd-party --grpc-gateway_out %s --grpc-gateway_opt logtostderr=true,paths=source_relative,allow_repeated_fields_in_body=true,generate_unbound_methods=true %s.proto", pkg, pkg))
	}
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
