package blockchain

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

func (setup *FabricSetup) QueryCandidate() ([]byte, error) {

	// Prepare arguments
	var args []string

	args = append(args, "queryOne")
	args = append(args, "wangyujiang")

	response, err := setup.client.Query(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{[]byte(args[1])}})
	if err != nil {
		return []byte(""), fmt.Errorf("failed to query: %v", err)
	}

	return response.Payload, nil
}
