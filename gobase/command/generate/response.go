package generate

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/urfave/cli"
)

func Response(c *cli.Context) {
	response := c.Args().Get(2)
	responseFileName := response + "_response" + ".go"
	input, err := ioutil.ReadFile("gobase/template/response.go")
	if err != nil {
		fmt.Printf("File is Not Exist: %v", err)
	}
	file := string(input)
	file = strings.Replace(file, "Template", strings.Title(response), -1)
	file = strings.Replace(file, "template", "responses", -1)
	err = ioutil.WriteFile("models/dto/responses/" + responseFileName, []byte(file), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}

	fmt.Printf("Successfully generate " + response + " response")
}
