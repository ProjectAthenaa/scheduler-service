package modules

import (
	"github.com/ProjectAthenaa/sonic-core/protos/module"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent/product"
)

var modules = map[product.Site]module.ModuleClient{}

func GetConnection(site product.Site) module.ModuleClient {
	return modules[site]
}
