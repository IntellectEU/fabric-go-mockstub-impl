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
To install the package simply run in your project
```
go get bitbucket.org/yurii_uhlanov_intellecteu/fabric-go-mockstub-impl
```
