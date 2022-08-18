package tools

const (
	// mainnet prefixes
	qtumMainPubKeyHashAddrID byte = 58
	qtumMainScriptHashAddrID byte = 50
	btcMainPubKeyHashAddrID  byte = 0x00
	mainnetVersion           byte = 0x80

	// testnet prefixes
	qtumTestNetPubKeyHashAddrID byte = 120 // hex: 0x78 dec: 120
	qtumTestNetScriptHashAddrID byte = 110 // hex: 0x6e dec: 110
	btcTestNetPubKeyHashAddrID  byte = 0x6F
	testnetVersion              byte = 0xEF

	OP_CHECKSIG = "ac"
	NETWORK     = "testnet"
)
