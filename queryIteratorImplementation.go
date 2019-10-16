package mockstubimpl

import (
	"fmt"

	"github.com/hyperledger/fabric/protos/ledger/queryresult"
)

type commonIteratorImplementation struct {
	responses  []queryresult.KV
	currentLoc int
}

func (iterator *commonIteratorImplementation) Next() (*queryresult.KV, error) {
	if iterator.currentLoc == len(iterator.responses) {
		return nil, fmt.Errorf("No more entities left")
	}
	iterator.currentLoc += 1
	return &iterator.responses[iterator.currentLoc], nil
}

func (iterator *commonIteratorImplementation) HasNext() bool {
	if iterator.currentLoc < len(iterator.responses)-1 {
		return true
	}
	return false
}

func (iterator *commonIteratorImplementation) Close() error {
	return nil
}

func newCommonIterator(responses []queryresult.KV) *commonIteratorImplementation {
	return &commonIteratorImplementation{
		responses:  responses,
		currentLoc: -1,
	}
}
