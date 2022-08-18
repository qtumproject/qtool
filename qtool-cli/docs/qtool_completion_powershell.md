## qtool completion powershell

Generate the autocompletion script for powershell

### Synopsis

Generate the autocompletion script for powershell.

To load completions in your current shell session:

	qtool completion powershell | Out-String | Invoke-Expression

To load completions for every new session, add the output of the above command
to your powershell profile.


```
qtool completion powershell [flags]
```

### Options

```
  -h, --help              help for powershell
      --no-descriptions   disable completion descriptions
```

### Options inherited from parent commands

```
  -b, --blockchain string   blockchain: "qtum" or "btc" (default "qtum")
  -n, --network string      network type: "testnet" or "mainnet" (default "mainnet")
  -v, --verbose             verbose output
```

### SEE ALSO

* [qtool completion](qtool_completion.md)	 - Generate the autocompletion script for the specified shell

###### Auto generated by spf13/cobra on 4-Aug-2022