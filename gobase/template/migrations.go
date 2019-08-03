package template

import "github.com/wincentrtz/gobase/gobase/utils"

func _TEMPLATE_Schema() string {
	return utils.Migration().Table("_TABLE_").Column(initialize_TEMPLATE_Columns()).Build()
}

func initialize_TEMPLATE_Columns() []utils.Column {
	return []utils.Column{
		utils.NewColumn().Name("id").Primary().Build(),
	}
}
