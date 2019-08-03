package migrations

import (
	"github.com/wincentrtz/gobase/gobase/utils"
)

func UserSchema() string {
	return utils.Migration().Table("users").Column(initializeUserColumns()).Build()
}

func initializeUserColumns() []utils.Column {
	return []utils.Column{
		utils.NewColumn().Name("id").Primary().Build(),
		utils.NewColumn().Name("name").TypeData("VARCHAR").NotNull().Build(),
		utils.NewColumn().Name("email").TypeData("VARCHAR").NotNull().Build(),
		utils.NewColumn().Name("created_on").TypeData("TIMESTAMP").NotNull().Build(),
	}
}
