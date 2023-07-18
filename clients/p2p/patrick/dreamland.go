package p2p

import (
	dreamlandRegistry "bitbucket.org/taubyte/dreamland/registry"
	"github.com/taubyte/go-interfaces/common"
	p2p "github.com/taubyte/go-interfaces/p2p/peer"
)

func init() {
	dreamlandRegistry.Registry.Patrick.Client = createPatrickClient
}

func createPatrickClient(node p2p.Node, config *common.ClientConfig) (common.Client, error) {
	return New(node.Context(), node)
}