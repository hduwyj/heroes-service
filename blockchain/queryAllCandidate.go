package blockchain

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

func (setup *FabricSetup) QueryAllCandidate() ([]byte, error) {

	// Prepare arguments
	var args []string
	args = append(args, "queryAllCandidates")

	response, err := setup.client.Query(channel.Request{ChaincodeID: setup.ChainCodeID, Fcn: args[0], Args: [][]byte{}})
	if err != nil {
		return []byte(""), fmt.Errorf("failed to query: %v", err)
	}

	return response.Payload, nil
}
