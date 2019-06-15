package main

import (
	"github.com/iotaledger/goshimmer/packages/node"
	"github.com/iotaledger/goshimmer/plugins/analysis"
	"github.com/iotaledger/goshimmer/plugins/autopeering"
	"github.com/iotaledger/goshimmer/plugins/cli"
	"github.com/iotaledger/goshimmer/plugins/fpc"
	"github.com/iotaledger/goshimmer/plugins/gossip"
	"github.com/iotaledger/goshimmer/plugins/gracefulshutdown"
	"github.com/iotaledger/goshimmer/plugins/statusscreen"
	statusscreen_tps "github.com/iotaledger/goshimmer/plugins/statusscreen-tps"
	"github.com/iotaledger/goshimmer/plugins/tangle"
	"github.com/iotaledger/goshimmer/plugins/tipselection"
	"github.com/iotaledger/goshimmer/plugins/webapi"
	webapi_gtta "github.com/iotaledger/goshimmer/plugins/webapi-gtta"
	webapi_spammer "github.com/iotaledger/goshimmer/plugins/webapi-spammer"
)

func main() {
	node.Run(
		cli.PLUGIN,
		autopeering.PLUGIN,
		fpc.PLUGIN,
		gossip.PLUGIN,
		tangle.PLUGIN,
		analysis.PLUGIN,
		gracefulshutdown.PLUGIN,
		tipselection.PLUGIN,

		statusscreen.PLUGIN,
		statusscreen_tps.PLUGIN,

		webapi.PLUGIN,
		webapi_gtta.PLUGIN,
		webapi_spammer.PLUGIN,
	)
}
