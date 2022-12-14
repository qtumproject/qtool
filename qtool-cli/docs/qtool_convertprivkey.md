## qtool convertprivkey

Converts the encoding of a ECDSA private key

### Synopsis

Converts a ECDSA private key from base58 ('b58') encoding to hexadecimal ('hex') encoding and vice versa.	
When converting from 'hex' is important to explicity set the flags for 'network'.
Example:
qtool convertprivkey   cMbgxCJrTYUqgcmiC1berh5DFrtY1KeU4PXZ6NZxgenniF1mXCRk -f b58
00821d8c8a3627adc68aa4034fea953b2f5da553fab312db3fa274240bd49f35


```
qtool convertprivkey [flags]
```

### Examples

```
qtool convertprivkey cMbgxCJrTYUqgcmiC1berh5DFrtY1KeU4PXZ6NZxgenniF1mXCRk -f b58
```

### Options

```
  -f, --from string   The current encoding format of the address (b58 or hex)
  -h, --help          help for convertprivkey
```

### Options inherited from parent commands

```
  -b, --blockchain string   blockchain: "qtum" or "btc" (default "qtum")
  -n, --network string      network type: "testnet" or "mainnet" (default "mainnet")
  -v, --verbose             verbose output
```

### SEE ALSO

* [qtool](qtool.md)	 - 

###### Auto generated by spf13/cobra on 21-Aug-2022
