# Hyperledger Fabric Go MockStub implementation

This implementation of ChaincodeStubInterface almost entirely repeats shim.MockStub except for several function like GetQueryResult and GetQueryResultWithPagination that were changed and now support simple queries like:
```
{"selector": {
    "parameter1": "value",
    "parameter2": true,
    "parameter3": 112
    }
}
```
GetQueryResultWithPagination expects an empty bookmark to start search from the beginning of the record list. If the end of the list was reached, the bookmark returned will contain the "end symbol" `%%end$$`

### Installation
To include the module in your project just add it to your dependencies:
```Go
import "github.com/intellecteu/fabric-go-mockstub-impl"
```

If Fabric packages don't install correctly, run the following commands
```
go get github.com/hyperledger/fabric/common/util@v1.4
go get github.com/hyperledger/fabric/core/chaincode/shim@v1.4
go get github.com/hyperledger/fabric/protos/peer@v1.4
```
