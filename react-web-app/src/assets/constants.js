
const methods = [
	"convertprivkey",
	"convertaddress",
	"getaddrfromprivkey",
	"getaddressfromscriptpubKey",
];

const commands = [
	"Convert Private Key",
	"Convert Address",
	"Get Address from Private Key",
	"Get Address from ScriptPubKey"
]

const sample = {
	"qtum_addressHex" : "7926223070547d2d15b2ef5e7383e541c338ffe9",
	"qtum_addressBase58Testnet" : "qUbxboqjBRp96j3La8D1RYkyqx5uQbJPoW",
	"qtum_addressBase58Mainnet" : "QXeZZ5MsAF5pPrPy47ZFMmtCpg7RExT4mi",
	"qtum_privateKeyWIF" : "cMbgxCJrTYUqgcmiC1berh5DFrtY1KeU4PXZ6NZxgenniF1mXCRk",
	"qtum_privateKeyHex" : "00821d8c8a3627adc68aa4034fea953b2f5da553fab312db3fa274240bd49f35",
	"qtum_pubKey" : "0299d391f528b9edd07284c7e23df8415232a8ce41531cf460a390ce32b4efd112", 
	"qtum_scriptpubkey_33" : "210299d391f528b9edd07284c7e23df8415232a8ce41531cf460a390ce32b4efd112ac",
	"qtum_scriptpubkey_65" : "4104ae1a62fe09c5f51b13905f07f06b99a2f7159b2225f374cd378d71302fa28414e7aab37397f554a7df5f142c21c1b7303b8a0626f1baded5c72a704f7e6cd84cac",
	"btc_addressHex" : "977AE6E32349B99B72196CB62B5EF37329ED81B4",
	"btc_addressBase58Mainnet" : "1EoxGLjv4ZADtRBjTVeXY35czVyDdp7rU4",
	"btc_addressBase58Testnet" : "muKuZPptsabUfXfMB4cuMxHwrVZvaRBqJp",
	"btc_privateKeyWIF" : "L2WmFR8WMr5GSprjt7UTA7WQ23WDEZPVRimrZv1dmz7e4JzxqSNq",
	"btc_privateKeyHex" : "9DF5A907FF17ED6A4E02C00C2C119049A045F52A4E817B06B2EC54EB68F70079",
	"btc_pubKey" : "0299d391f528b9edd07284c7e23df8415232a8ce41531cf460a390ce32b4efd112", 
}
export {methods, commands, sample };