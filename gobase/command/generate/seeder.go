package generate

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"github.com/wincentrtz/gobase/gobase/utils"
	"io/ioutil"
	"strings"
)

func Seeder(c *cli.Context) {
	domain := c.Args().Get(2)
	domainFileName := domain + ".go"
	if !utils.IsFileExist("models/entity", domainFileName) {
		log.Error(fmt.Sprintf("File %v is Not Exist", domainFileName))
	}

	input, err := ioutil.ReadFile("gobase/template/seeder.go")
	if err != nil {
		log.Fatal("File is Not Exist: %v", err)
	}
	file := string(input)
	file = strings.Replace(file, "template", "seeders", 1)
	file = file[:strings.Index(file, "(")+1]+
		"\n\t\"github.com/wincentrtz/gobase/models/entity\""+ file[strings.Index(file, "(")+1:]
	file = strings.Replace(file, "template", domain, -1)
	file = strings.Replace(file, "struct{}", "entity."+strings.Title(domain), -1)
	err = ioutil.WriteFile("seeders/"+domain+"_seeder.go", []byte(file), 0755)
	if err != nil {
		fmt.Printf("Unable to write file: %v", err)
	}
}
