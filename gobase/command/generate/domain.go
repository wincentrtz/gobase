package generate

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/urfave/cli"
)

func GenerateRepository(domain string) {
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

func GenerateRepositoryImpl(domain string) {
	input, err := ioutil.ReadFile("gobase/template/repository_impl.go")
	if err != nil {
		fmt.Printf("File is Not Exist: %v", err)
	}
	file := string(input)
	file = strings.Replace(file, "NewTemplateRepository", "New"+strings.Title(domain)+"Repository", -1)
	file = strings.Replace(file, "TemplateRepository", domain+"."+strings.Title(domain)+"Repository", -1)
	file = strings.Replace(file, "Template", strings.Title(domain), -1)
	file = strings.Replace(file, "templateRepository", domain+"Repository", -1)
	file = strings.Replace(file, "template", "repository", -1)

	os.Mkdir("."+string(filepath.Separator)+"domains/"+domain+"/repository", 0775)
	err = ioutil.WriteFile("domains/"+domain+"/repository/"+domain+"_repository.go", []byte(file), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
	err = ioutil.WriteFile("domains/"+domain+"/repository/"+domain+"_repository_test.go", []byte("package repository_test"), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
}

func GenerateUsecase(domain string) {
	input, err := ioutil.ReadFile("gobase/template/usecase.go")
	if err != nil {
		fmt.Printf("File is Not Exist: %v", err)
	}
	file := string(input)
	file = strings.Replace(file, "Template", strings.Title(domain), -1)
	file = strings.Replace(file, "template", domain, -1)
	os.Mkdir("."+string(filepath.Separator)+"domains/"+domain+"/usecase", 0775)
	err = ioutil.WriteFile("domains/"+domain+"/usecase.go", []byte(file), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
}

func GenerateUsecaseImpl(domain string) {
	input, err := ioutil.ReadFile("gobase/template/usecase_impl.go")
	if err != nil {
		fmt.Printf("File is Not Exist: %v", err)
	}
	file := string(input)
	file = strings.Replace(file, "NewTemplateUsecase", "New"+strings.Title(domain)+"Usecase", -1)
	file = strings.Replace(file, "TemplateUsecase", domain+"."+strings.Title(domain)+"Usecase", -1)
	file = strings.Replace(file, "TemplateRepository", domain+"."+strings.Title(domain)+"Repository", -1)
	file = strings.Replace(file, "Template", strings.Title(domain), -1)
	file = strings.Replace(file, "templateRepo", domain+"Repo", -1)
	file = strings.Replace(file, "templateUsecase", domain+"Usecase", -1)
	file = strings.Replace(file, "template", "usecase", -1)
	err = ioutil.WriteFile("domains/"+domain+"/usecase/"+domain+"_usecase.go", []byte(file), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
	err = ioutil.WriteFile("domains/"+domain+"/usecase/"+domain+"_usecase_test.go", []byte("package usecase_test"), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
}

func GenerateHandler(domain string) {
	input, err := ioutil.ReadFile("gobase/template/handler.go")
	if err != nil {
		fmt.Printf("File is Not Exist: %v", err)
	}
	file := string(input)
	file = strings.Replace(file, "NewTemplateRepository", "repository.New"+strings.Title(domain)+"Repository", -1)
	file = strings.Replace(file, "NewTemplateUsecase", "usecase.New"+strings.Title(domain)+"Usecase", -1)
	file = strings.Replace(file, "NewTemplateHandler", "New"+strings.Title(domain)+"Handler", -1)
	file = strings.Replace(file, "TemplateUsecase", domain+"."+strings.Title(domain)+"Usecase", -1)
	file = strings.Replace(file, "Template", strings.Title(domain), -1)
	file = strings.Replace(file, "templateUsecase", domain+"Usecase", -1)
	file = strings.Replace(file, "templateHandler", domain+"Handler", -1)
	file = strings.Replace(file, "/api/template/{id}", "/api/"+domain+"/{id}", -1)
	file = strings.Replace(file, "template", "handler", -1)
	os.Mkdir("."+string(filepath.Separator)+"domains/"+domain+"/handler", 0775)
	err = ioutil.WriteFile("domains/"+domain+"/handler/"+domain+"_handler.go", []byte(file), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
	err = ioutil.WriteFile("domains/"+domain+"/handler/"+domain+"_handler_test.go", []byte("package handler_test"), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
}

func GenerateMocksRepository(domain string) {
	input, err := ioutil.ReadFile("gobase/template/repository_mocks.go")
	if err != nil {
		fmt.Printf("File is Not Exist: %v", err)
	}
	os.Mkdir("."+string(filepath.Separator)+"domains/"+domain+"/mocks", 0775)
	file := string(input)
	file = strings.Replace(file, "template", "mocks", -1)

	err = ioutil.WriteFile("domains/"+domain+"/mocks/repository.go", []byte(file), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
}

func GenerateMocksUsecase(domain string) {
	input, err := ioutil.ReadFile("gobase/template/usecase_mocks.go")
	if err != nil {
		fmt.Printf("File is Not Exist: %v", err)
	}
	file := string(input)
	file = strings.Replace(file, "template", "mocks", -1)

	err = ioutil.WriteFile("domains/"+domain+"/mocks/usecase.go", []byte(file), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
}

func Domain(c *cli.Context) {
	domain := c.Args().Get(2)
	os.Mkdir("."+string(filepath.Separator)+"domains/"+domain, 0775)

	GenerateRepository(domain)
	GenerateRepositoryImpl(domain)
	GenerateUsecase(domain)
	GenerateUsecaseImpl(domain)
	GenerateHandler(domain)
	GenerateMocksRepository(domain)
	GenerateMocksUsecase(domain)
	Migration(c)

	fmt.Printf("Successfully generate " + domain + " domain")
}
