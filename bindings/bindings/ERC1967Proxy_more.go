// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/specularL2/specular/bindings/solc"
)

const ERC1967ProxyStorageLayoutJSON = "{\"storage\":null,\"types\":{}}"

var ERC1967ProxyStorageLayout = new(solc.StorageLayout)

var ERC1967ProxyDeployedBin = "0x60806040523661001357610011610017565b005b6100115b610027610022610067565b61009f565b565b606061004e838360405180606001604052806027815260200161025f602791396100c3565b9392505050565b6001600160a01b03163b151590565b90565b600061009a7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b905090565b3660008037600080366000845af43d6000803e8080156100be573d6000f35b3d6000fd5b6060600080856001600160a01b0316856040516100e0919061020f565b600060405180830381855af49150503d806000811461011b576040519150601f19603f3d011682016040523d82523d6000602084013e610120565b606091505b50915091506101318683838761013b565b9695505050505050565b606083156101af5782516000036101a8576001600160a01b0385163b6101a85760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064015b60405180910390fd5b50816101b9565b6101b983836101c1565b949350505050565b8151156101d15781518083602001fd5b8060405162461bcd60e51b815260040161019f919061022b565b60005b838110156102065781810151838201526020016101ee565b50506000910152565b600082516102218184602087016101eb565b9190910192915050565b602081526000825180602084015261024a8160408501602087016101eb565b601f01601f1916919091016040019291505056fe416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220df5e088a16650cdc3010db20ecd1e2b26729605480c947e2f8d77aac24e8ff8464736f6c63430008110033"
func init() {
	if err := json.Unmarshal([]byte(ERC1967ProxyStorageLayoutJSON), ERC1967ProxyStorageLayout); err != nil {
		panic(err)
	}

	layouts["ERC1967Proxy"] = ERC1967ProxyStorageLayout
	deployedBytecodes["ERC1967Proxy"] = ERC1967ProxyDeployedBin
}
