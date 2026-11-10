package data

import (
	"gorm.io/gen"
)

func gopsTable(g *gen.Generator) {

	user := g.GenerateModel("user")
	g.ApplyBasic(user)

}
