package modules

import (
	"github.com/ProjectAthenaa/sonic-core/protos/module"
	"github.com/ProjectAthenaa/sonic-core/sonic/database/ent/product"
	"google.golang.org/grpc"
)

//moduleMap holds the DNS names for all the modules
var siteMap = map[product.Site]string{
	product.SiteFinishLine:     "finishline.modules.svc.cluster.local:3000",
	product.SiteJD_Sports:      "jdsports.modules.svc.cluster.local:3000",
	product.SiteYeezySupply:    "yeezysupply.modules.svc.cluster.local:3000",
	product.SiteSupreme:        "supreme.modules.svc.cluster.local:3000",
	product.SiteEastbay_US:     "eastbayUS.modules.svc.cluster.local:3000",
	product.SiteChamps_US:      "champsUS.modules.svc.cluster.local:3000",
	product.SiteFootaction_US:  "footactionUS.modules.svc.cluster.local:3000",
	product.SiteFootlocker_US:  "footlockerUS.modules.svc.cluster.local:3000",
	product.SiteBestbuy:        "bestbuy.modules.svc.cluster.local:3000",
	product.SitePokemon_Center: "pokemon-center.modules.svc.cluster.local:3000",
	product.SitePanini_US:      "paniniUS.modules.svc.cluster.local:3000",
	product.SiteTopss:          "topps.modules.svc.cluster.local:3000",
	product.SiteNordstorm:      "nordstorm.modules.svc.cluster.local:3000",
	product.SiteEnd:            "end.modules.svc.cluster.local:3000",
	product.SiteTarget:         "target.modules.svc.cluster.local:3000",
	product.SiteAmazon:         "amazon.modules.svc.cluster.local:3000",
	product.SiteSolebox:        "solebox.modules.svc.cluster.local:3000",
	product.SiteOnygo:          "onygo.modules.svc.cluster.local:3000",
	product.SiteSnipes:         "snipes.modules.svc.cluster.local:3000",
	product.SiteSsense:         "ssense.modules.svc.cluster.local:3000",
	product.SiteWalmart:        "walmart.modules.svc.cluster.local:3000",
	product.SiteHibbet:         "hibbet.modules.svc.cluster.local:3000",
	product.SiteNewBalance:     "newbalance.modules.svc.cluster.local:3000",
}

//populateMap connects to all the services shown in siteMap and populates the Modules map accordingly
func populateMap() error {
	for k, v := range siteMap {
		conn, err := grpc.Dial(v, grpc.WithInsecure())
		if err != nil {
			panic(err)
		}

		modules[k] = module.NewModuleClient(conn)
	}
	return nil
}
