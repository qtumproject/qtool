package rpc

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/qtumproject/btcd/btcjson"
	"github.com/qtumproject/btcd/txscript"
	"github.com/qtumproject/qtool/pkg/common"
	"github.com/qtumproject/qtool/pkg/script"
	qtool "github.com/qtumproject/qtool/pkg/tools"
)

type QtumTxRawResult struct {
	*btcjson.TxRawResult
	receiverAddr string
	senderAddr   string
}

// Returns the sender address found in the scriptSig in Hex format
func (r *QtumTxRawResult) FindSenderAddress() (string, error) {
	if len(r.Vin) == 0 {
		return "", fmt.Errorf("no vins found in tx")
	}
	sender := r.Vin[0].Address
	hexFromAddrResult, err := qtool.ConvertAddressBase58ToHex(sender)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Error converting sender address to hex. Error: %s", err))
	}
	r.senderAddr = common.AddHexPrefix(hexFromAddrResult.Address)
	return r.senderAddr, nil
}

// Returns the destination address(es) found in the scriptPubKey in Hex format
// The addresses returned could be a contract address and/or a regular address
func (r *QtumTxRawResult) GetReceiverAddressList() ([]string, error) {
	if len(r.Vout) == 0 {
		return nil, errors.New("No vouts found in tx")
	}

	recvrAddrList := make([]string, 0)

	for _, vout := range r.Vout {
		if vout.ScriptPubKey.Type == "pubkeyhash" {
			if len(vout.ScriptPubKey.Addresses) == 0 {
				return nil, errors.New("No addresses found in vout")
			}

			receiver, err := qtool.ConvertAddressBase58ToHex(vout.ScriptPubKey.Addresses[0])
			if err != nil {
				return nil, err
			}
			recvrAddrList = append(recvrAddrList, common.AddHexPrefix(receiver.Address))
			continue
		}
		opcodes, err := script.GetScriptPubKeyOpcodes(vout.ScriptPubKey.Hex)
		if err != nil {
			return nil, err
		}
		finalOpcode := opcodes[len(opcodes)-1]
		if finalOpcode == txscript.OP_CREATE || finalOpcode == txscript.OP_CALL || finalOpcode == txscript.OP_SENDER {
			scriptAsm, err := script.DisasmScript(vout.ScriptPubKey.Hex)
			if err != nil {
				return nil, err
			}
			contractInvokeInfo, err := script.ParseContractParams(scriptAsm, finalOpcode)
			if err != nil {
				return nil, errors.New(fmt.Sprintf("Error parsing contract params. Error: %s", err))
			}
			receiver := contractInvokeInfo.To
			r.receiverAddr = common.AddHexPrefix(receiver)
			recvrAddrList = append(recvrAddrList, r.receiverAddr)
		}
	}

	if len(recvrAddrList) == 0 {
		return nil, errors.New("No receiver addresses found in tx")
	}
	return recvrAddrList, nil
}

func (r *QtumTxRawResult) FindReceiverAddress() (string, error) {
	if len(r.Vout) == 0 {
		return "", errors.New("No vouts found in tx")
	}
	if r.senderAddr == "" {
		_, err := r.FindSenderAddress()
		if err != nil {
			return "", err
		}
	}

	receiver := ""
	for _, vout := range r.Vout {
		if len(vout.ScriptPubKey.Addresses) > 0 {
			address, err := qtool.ConvertAddressBase58ToHex(vout.ScriptPubKey.Addresses[0])
			if err != nil {
				return "", err
			}
			addressHex := common.AddHexPrefix(address.Address)
			if receiver == "" || addressHex != r.senderAddr {
				receiver = addressHex
			}
		}
		opcodes, err := script.GetScriptPubKeyOpcodes(vout.ScriptPubKey.Hex)
		if err != nil {
			return "", err
		}
		finalOpcode := opcodes[len(opcodes)-1]
		if finalOpcode == txscript.OP_CREATE || finalOpcode == txscript.OP_CALL {
			scriptAsm, err := script.DisasmScript(vout.ScriptPubKey.Hex)
			if err != nil {
				return "", err
			}
			contractInvokeInfo, err := script.ParseContractParams(scriptAsm, finalOpcode)
			if err != nil {
				return "", errors.New(fmt.Sprintf("Error parsing contract params. Error: %s", err))
			}
			contractAddr := contractInvokeInfo.To
			receiver = common.AddHexPrefix(contractAddr)
			break
		}

	}
	if receiver == "" {
		return "", errors.New("No receiver addresses found in tx")
	}
	r.receiverAddr = receiver
	return receiver, nil
}

// ExtractContractInfo checks Vouts of given pubKey scripts and returns
// ContractInvoke params (i.e., gasLimit, gasPrice, callData)
func (r *QtumTxRawResult) ExtractContractInfo() (*script.ContractInvokeInfo, error) {
	for _, vout := range r.Vout {
		opcodes, err := script.GetScriptPubKeyOpcodes(vout.ScriptPubKey.Hex)
		if err != nil {
			return nil, err
		}
		if len(opcodes) == 0 {
			continue
		}
		finalOpcode := opcodes[len(opcodes)-1]
		if finalOpcode == txscript.OP_CREATE || finalOpcode == txscript.OP_CALL || finalOpcode == txscript.OP_SENDER {
			scriptAsm, err := script.DisasmScript(vout.ScriptPubKey.Hex)
			if err != nil {
				return nil, err
			}
			contractInvokeInfo, err := script.ParseContractParams(scriptAsm, finalOpcode)
			if err != nil {
				return nil, err
			}
			if contractInvokeInfo.GasPrice != "" {
				gasPriceWei, err := common.ConvertHexFromSatoshiToWei(contractInvokeInfo.GasPrice)
				if err != nil {
					return nil, err
				}
				contractInvokeInfo.GasPrice = common.AddHexPrefix(gasPriceWei)
			}
			contractInvokeInfo.CallData = common.AddHexPrefix(contractInvokeInfo.CallData)
			contractInvokeInfo.To = common.AddHexPrefix(contractInvokeInfo.To)
			contractInvokeInfo.GasLimit = common.AddHexPrefix(contractInvokeInfo.GasLimit)

			return contractInvokeInfo, nil
		}
	}
	return nil, nil

}

// Returns the initial gas fee paid by the sender for the TX in units of satoshis
func (r *QtumTxRawResult) CalculateInitialGasFee() int64 {
	vinTotal := vinTotalValue(r.Vin)
	voutTotal := voutTotalValue(r.Vout)
	if vinTotal-voutTotal < 0 {
		// This is a coinbase/coinstake transaction
		return 0
	}
	return vinTotal - voutTotal
}

func vinTotalValue(vins []btcjson.Vin) int64 {
	var vinValue int64
	for _, vin := range vins {
		if vin.IsCoinBase() {
			return 0
		}
		vinValue += vin.AmountSatoshi
	}
	return vinValue
}

func voutTotalValue(vouts []btcjson.Vout) int64 {
	var voutValue int64
	for _, vout := range vouts {
		voutValue += vout.AmountSatoshi
	}
	return voutValue
}
