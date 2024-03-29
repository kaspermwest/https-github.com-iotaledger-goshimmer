package tipselection

import (
	"github.com/iotaledger/goshimmer/packages/events"
	"github.com/iotaledger/goshimmer/packages/model/value_transaction"
	"github.com/iotaledger/goshimmer/packages/node"
	"github.com/iotaledger/goshimmer/plugins/tangle"
)

var PLUGIN = node.NewPlugin("Tipselection", configure, run)

func configure(node *node.Plugin) {
	tangle.Events.TransactionSolid.Attach(events.NewClosure(func(transaction *value_transaction.ValueTransaction) {
		go func() {
			tips.Delete(transaction.GetBranchTransactionHash())
			tips.Delete(transaction.GetTrunkTransactionHash())
			tips.Set(transaction.GetHash(), transaction.GetHash())
		}()
	}))
}

func run(run *node.Plugin) {
}
