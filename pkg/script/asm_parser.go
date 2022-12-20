package script

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/qtumproject/btcd/txscript"
	"github.com/qtumproject/qtool/pkg/common"
)

// Returns the opcodes found in the scriptPubKey
func GetScriptPubKeyOpcodes(scriptPubKey string) ([]byte, error) {
	scriptPubKeyBytes, err := hex.DecodeString(scriptPubKey)
	var opcodes []byte
	if err != nil {
		return opcodes, err
	}
	// Create a tokenizer to iterate the script and count the number of opcodes.
	const scriptVersion = 0
	var numOpcodes int
	tokenizer := txscript.MakeScriptTokenizer(scriptVersion, scriptPubKeyBytes)
	for tokenizer.Next() {
		opcodes = append(opcodes, tokenizer.Opcode())
		numOpcodes++
	}
	if tokenizer.Err() != nil {
		return opcodes, tokenizer.Err()
	}
	return opcodes, nil

}

// Returns a string with the disassembled code of the scriptPubKey passed as 1 line hex
func DisasmScript(scriptPubKey string) (string, error) {
	scriptBytes, err := hex.DecodeString(scriptPubKey)
	if err != nil {
		return "", err
	}
	disasm, err := txscript.DisasmString(scriptBytes)
	if err != nil {
		return "", err
	}
	return disasm, nil
}

// Returns the contract params found in the Vout of scriptPubKey
func ParseContractParams(disasm string, finalOp byte) (*ContractInvokeInfo, error) {
	contractInvokeInfo := ContractInvokeInfo{}
	script := strings.Split(disasm, " ")
	if len(script) < 5 {
		return nil, nil
	}
	switch finalOp {
	case txscript.OP_CALL:
		if script[3] == "OP_SENDER" {
			// 1 <sender's pubkeyhash address> < {signature, pubkey}> OP_SENDER <evm version>  <gas limit> <gas price> [byte code] <contract address> OP_CALL
			// https://qtum.info/tx/0425fa39feed4cd6c93998159901095c147f8b0043823067dc1d25dabf950ac9
			if len(script) != 10 {
				return nil, errors.New(fmt.Sprintf("invalid OP_SENDER script for parts 10: %v", len(script)))
			}
			contractInvokeInfo.From = script[1]
			contractInvokeInfo.GasLimit = script[5]
			contractInvokeInfo.GasPrice = script[6]
			contractInvokeInfo.CallData = script[7]
			contractInvokeInfo.To = script[8]
		} else {
			// <evm version> <gas limit> <gas price> [byte code] <contract address> OP_CALL
			// https://qtum.info/tx/ea860543077aca9fd3c43c9d414a0c21de364db9b2db25a22c1c1b6083c62ec2
			if len(script) != 6 {
				return nil, errors.New(fmt.Sprintf("invalid OP_CALL script for parts 6: %v", len(script)))
			}
			contractInvokeInfo.GasLimit = script[1]
			contractInvokeInfo.GasPrice = script[2]
			contractInvokeInfo.CallData = script[3]
			contractInvokeInfo.To = script[4]
		}
	case txscript.OP_CREATE:
		if script[3] == "OP_SENDER" {
			if len(script) != 9 {
				return nil, errors.New(fmt.Sprintf("invalid OP_SENDER script for parts 9: %v", len(script)))
			}
			contractInvokeInfo.From = script[1]
			contractInvokeInfo.GasLimit = script[5]
			contractInvokeInfo.GasPrice = script[6]
			contractInvokeInfo.CallData = script[7]
			contractInvokeInfo.To = ""
		} else {
			// <evm version> <gas limit> <gas price> [byte code] OP_CREATE
			// https://qtum.info/tx/9b11819d254cc1877353cf02b587308ccf2e79c61d9f5ca726d744fcea2fd37c
			if len(script) != 5 {
				return nil, errors.New(fmt.Sprintf("invalid OP_CREATE script for parts 5: %v", len(script)))
			}
			contractInvokeInfo.GasLimit = script[1]
			contractInvokeInfo.GasPrice = script[2]
			contractInvokeInfo.CallData = script[3]
			contractInvokeInfo.To = ""
		}
	case txscript.OP_SPEND:
		return nil, fmt.Errorf("OP_SPEND contract parsing partially implemented")
	default:
		return nil, fmt.Errorf("unknown final opcode: %v", finalOp)
	}
	if contractInvokeInfo != (ContractInvokeInfo{}) {
		gasLimit, err := common.ConvertToBigEndian(contractInvokeInfo.GasLimit)
		if err != nil {
			return nil, err
		}
		contractInvokeInfo.GasLimit = gasLimit
	}

	return &contractInvokeInfo, nil
}
