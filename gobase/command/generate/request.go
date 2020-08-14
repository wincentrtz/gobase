package generate

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/urfave/cli"
)

func Request(c *cli.Context) {
	request := c.Args().Get(2)
	requestFileName := request + "_request" + ".go"
	input, err := ioutil.ReadFile("gobase/template/request.go")
	if err != nil {
		fmt.Printf("File is Not Exist: %v", err)
	}
	file := string(input)
	file = strings.Replace(file, "Template", strings.Title(request), -1)
	file = strings.Replace(file, "template", "requests", -1)
	err = ioutil.WriteFile("models/dto/requests/" + requestFileName, []byte(file), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}

	fmt.Printf("Successfully generate " + request + " request")
}
