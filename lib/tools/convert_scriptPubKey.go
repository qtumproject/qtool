package tools

import (
	"encoding/hex"
	"fmt"
	"strconv"

	"github.com/qtumproject/qtool/lib/common"
)

// AddressFromP2PKScript extracts the pubkey from a pay-to-pubkey script and returns a 'P2pkToAddressResult' struct.
func AddressFromP2PKScript(scriptPubKeyHexStr string, blockchain, network string) (*P2pkToAddressResult, error) {
	dataLength, _ := strconv.ParseInt(scriptPubKeyHexStr[:2], 16, 8)
	if dataLength != 33 && dataLength != 65 {
		return nil, fmt.Errorf("data length found %v, expected 33 or 65", dataLength)
	}

	opcode := scriptPubKeyHexStr[len(scriptPubKeyHexStr)-2:]
	if opcode != common.OP_CHECKSIG {
		return nil, fmt.Errorf("expected opcode '0xac' (OP_CHECKSIG) but found opcode '0x%s'", opcode)
	}

	scriptPubKeyBytes, err := hex.DecodeString(scriptPubKeyHexStr[2:])
	if err != nil {
		return nil, err
	}
	if dataLength == 33 && len(scriptPubKeyBytes) != 34 || dataLength == 65 && len(scriptPubKeyBytes) != 66 {
		return nil, fmt.Errorf("script length %v, expected length 34 or 65", len(scriptPubKeyBytes))
	}
	pubKey := scriptPubKeyBytes[:dataLength]
	pubKeyStr := hex.EncodeToString(pubKey)

	pubKeyToHexAddrResult, err := ConvertPubkeyToAddrHash160(pubKeyStr)
	if err != nil {
		return nil, err
	}

	addrHexToB58Result, err := ConvertAddressHexToBase58(pubKeyToHexAddrResult.Address, blockchain, network)
	if err != nil {
		return nil, err
	}

	script := ScriptPubKey{
		Hex: scriptPubKeyHexStr,
		Asm: fmt.Sprintf("%x OP_CHECKSIG", pubKey),
	}

	return &P2pkToAddressResult{
		ScriptPubKey:  script,
		PubKey:        hex.EncodeToString(pubKey),
		AddressHex:    pubKeyToHexAddrResult.Address,
		AddressBase58: addrHexToB58Result.Address,
	}, nil

}
