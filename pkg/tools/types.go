package tools

type GetAddressFromPrivkeyResult struct {
	PrivateKeyHex string `json:"privateKeyHex"`
	PrivateKeyWIF string `json:"privateKeyWIF"`
	PublicKeyHex  string `json:"publicKeyHex"`
	AddressHex    string `json:"addressHex"`
	AddressBase58 string `json:"addressBase58"`
}

type ScriptPubKey struct {
	Hex string `json:"hex"`
	Asm string `json:"asm"`
}

type P2pkToAddressResult struct {
	ScriptPubKey  ScriptPubKey `json:"scriptPubKey"`
	PubKey        string       `json:"pubKey"`
	AddressHex    string       `json:"addressHex"`
	AddressBase58 string       `json:"addressBase58"`
}

type ConvertPrivateKeyResult struct {
	PrivateKey string `json:"privKey"`
}

// func NewConvertPrivateKeyResult(privateKey string) *ConvertPrivateKeyResult {
// 	return &ConvertPrivateKeyResult{
// 		PrivateKey: privateKey,
// 	}
// }

type ConvertAddressResult struct {
	Address string `json:"address"`
}

// func NewConvertAddressResult(address string) *ConvertAddressResult {
// 	return &ConvertAddressResult{
// 		Address: address,
// 	}
// }
