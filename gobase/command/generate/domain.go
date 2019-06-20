package generate

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/urfave/cli"
)

func generateRepository(domain string) {
	os.Mkdir("."+string(filepath.Separator)+"domains/"+domain+"/repository", 0775)
	err := ioutil.WriteFile("domains/"+domain+"/repository.go", []byte("package "+domain), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
	err = ioutil.WriteFile("domains/"+domain+"/repository/"+domain+"_repository.go", []byte("package repository"), 0755)
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
}

func generateHandler(domain string) {
	os.Mkdir("."+string(filepath.Separator)+"domains/"+domain+"/handler", 0775)
	err := ioutil.WriteFile("domains/"+domain+"/handler/"+domain+"_handler.go", []byte("package handler"), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
}

func Domain(c *cli.Context) {
	domain := c.Args().Get(2)
	os.Mkdir("."+string(filepath.Separator)+"domains/"+domain, 0775)

	generateRepository(domain)
	generateUsecase(domain)
	generateHandler(domain)

	fmt.Printf("Successfully " + domain + " domain")
}
