package generate

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/urfave/cli"
)

func generateRepository(domain string) {
	input, err := ioutil.ReadFile("gobase/template/repository.go")
	if err != nil {
		fmt.Printf("File is Not Exist: %v", err)
	}
	file := string(input)
	file = strings.Replace(file, "Template", strings.Title(domain), -1)
	file = strings.Replace(file, "template", domain, -1)
	err = ioutil.WriteFile("domains/"+domain+"/repository.go", []byte(file), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
}

func generateRepositoryImpl(domain string) {
	input, err := ioutil.ReadFile("gobase/template/repository_impl.go")
	if err != nil {
		fmt.Printf("File is Not Exist: %v", err)
	}
	file := string(input)
	file = strings.Replace(file, "Template", strings.Title(domain), -1)
	file = strings.Replace(file, "templateRepository", domain+"repository", -1)
	file = strings.Replace(file, "template", "repository", -1)

	os.Mkdir("."+string(filepath.Separator)+"domains/"+domain+"/repository", 0775)
	err = ioutil.WriteFile("domains/"+domain+"/repository/"+domain+"_repository.go", []byte(file), 0755)
	err = ioutil.WriteFile("domains/"+domain+"/repository/"+domain+"_repository_test.go", []byte("package repository_test"), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
}

func generateUsecase(domain string) {
	os.Mkdir("."+string(filepath.Separator)+"domains/"+domain+"/usecase", 0775)
	err := ioutil.WriteFile("domains/"+domain+"/usecase.go", []byte("package "+domain), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}

	err = ioutil.WriteFile("domains/"+domain+"/usecase/"+domain+"_usecase.go", []byte("package usecase"), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
	err = ioutil.WriteFile("domains/"+domain+"/usecase/"+domain+"_usecase_test.go", []byte("package usecase_test"), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
}

func generateHandler(domain string) {
	os.Mkdir("."+string(filepath.Separator)+"domains/"+domain+"/handler", 0775)
	err := ioutil.WriteFile("domains/"+domain+"/handler/"+domain+"_handler.go", []byte("package handler"), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
	err = ioutil.WriteFile("domains/"+domain+"/handler/"+domain+"_handler_test.go", []byte("package handler_test"), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
}

func Domain(c *cli.Context) {
	domain := c.Args().Get(2)
	os.Mkdir("."+string(filepath.Separator)+"domains/"+domain, 0775)

	generateRepository(domain)
	generateRepositoryImpl(domain)
	generateUsecase(domain)
	generateHandler(domain)

	fmt.Printf("Successfully " + domain + " domain")
}
