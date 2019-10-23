# Hyperledger Fabric Go MockStub implementation

This implementation of ChaincodeStubInterface almost entirely repeats shim.MockStub except for the function GetQueryResult that was changed and now supports simple queries like:
```
{"selector": {
    "parameter1": "value",
    "parameter2": true,
    "parameter3": 112
    }
}
```

### Installation
To include the module in your project just add it to your dependencies:
```Go
import "github.com/intellecteu/fabric-go-mockstub-impl"
```

