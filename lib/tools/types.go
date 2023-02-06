package tools

type GetAddressFromPrivkeyResult struct {
	PrivateKeyHex string `json:"privateKeyHex"`
	PrivateKeyWIF string `json:"privateKeyWIF"`
	PublicKeyHex  string `json:"publicKeyHex"`
	AddressHex    string `json:"addressHex"`
	AddressBase58 string `json:"addressBase58"`
}

func (g *GetAddressFromPrivkeyResult) String() string {
	return g.AddressBase58
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
	WIF        string `json:"wif"`
	PrivateKey string `json:"privateKey"`
}

func (c *ConvertPrivateKeyResult) String() string {
	return c.WIF
}

type ConvertAddressResult struct {
	Address string `json:"address"`
}

func (c *ConvertAddressResult) String() string {
	return c.Address
}
