#! /usr/bin/env bash
PRIVKEY_HEX=00821d8c8a3627adc68aa4034fea953b2f5da553fab312db3fa274240bd49f35
PRIVKEY_B58=cMbgxCJrTYUqgcmiC1berh5DFrtY1KeU4PXZ6NZxgenniF1mXCRk
ADDRESS_HEX=7926223070547d2d15b2ef5e7383e541c338ffe9
ADDRESS_B58=qUbxboqjBRp96j3La8D1RYkyqx5uQbJPoW
SCRIPT_PUBKEY_33=210299d391f528b9edd07284c7e23df8415232a8ce41531cf460a390ce32b4efd112ac
SCRIPT_PUBKEY_65=4104ae1a62fe09c5f51b13905f07f06b99a2f7159b2225f374cd378d71302fa28414e7aab37397f554a7df5f142c21c1b7303b8a0626f1baded5c72a704f7e6cd84cac

echo ""
echo ">>>>>>>>>>>>>>>> CMD: convertaddress"
echo ""
echo \> qtool convertaddress  qUbxboqjBRp96j3La8D1RYkyqx5uQbJPoW -f b58
./bin/qtool convertaddress  qUbxboqjBRp96j3La8D1RYkyqx5uQbJPoW -f b58
echo ""
echo \> qtool convertaddress  7926223070547d2d15b2ef5e7383e541c338ffe9 -f hex
./bin/qtool convertaddress  7926223070547d2d15b2ef5e7383e541c338ffe9 -f hex

echo ""
echo ">>>>>>>>>>>>>>>> CMD: convertprivkey"
echo ""
echo \> qtool convertprivkey $PRIVKEY_B58 -f b58
./bin/qtool convertprivkey $PRIVKEY_B58 -f b58
echo ""
echo \> qtool convertprivkey $PRIVKEY_HEX -f hex -n testnet
./bin/qtool convertprivkey $PRIVKEY_HEX -f hex -n testnet

echo ""
echo ">>>>>>>>>>>>>>>> CMD: keytoaddress"
echo ""
echo \> qtool getaddrfromprivkey $PRIVKEY_B58 -n testnet -b qtum -f b58 -v
./bin/qtool getaddrfromprivkey $PRIVKEY_B58 -n testnet -b qtum -f b58 -v
echo ""
echo \> qtool getaddrfromprivkey $PRIVKEY_HEX -f hex -n testnet -v
./bin/qtool getaddrfromprivkey $PRIVKEY_HEX -f hex -n testnet -v

echo ""
echo ">>>>>>>>>>>>>>>> CMD: p2pktoaddr"
echo ""
echo \> qtool p2pktoaddr $SCRIPT_PUBKEY_33 -n testnet -v
./bin/qtool p2pktoaddr $SCRIPT_PUBKEY_33 -n testnet  -v
echo ""
echo \> qtool p2pktoaddr $SCRIPT_PUBKEY_65 -n testnet -v
./bin/qtool p2pktoaddr $SCRIPT_PUBKEY_65 -n testnet  -v
echo ""