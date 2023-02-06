package common

const (
	// mainnet prefixes
	QtumMainPubKeyHashAddrID byte = 0x3a // hex: 0x3a dec: 58
	QtumMainScriptHashAddrID byte = 0x32 // hex: 0x32 dec: 50
	BtcMainPubKeyHashAddrID  byte = 0x00
	MainnetVersion           byte = 0x80

	// testnet prefixes
	QtumTestNetPubKeyHashAddrID byte = 0x78 // hex: 0x78 dec: 120
	QtumTestNetScriptHashAddrID byte = 0x6e // hex: 0x6e dec: 110
	BtcTestNetPubKeyHashAddrID  byte = 0x6F
	TestnetVersion              byte = 0xEF

	OP_CHECKSIG = "ac"
)
