## qtool completion bash

Generate the autocompletion script for bash

### Synopsis

Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(qtool completion bash)

To load completions for every new session, execute once:

#### Linux:

	qtool completion bash > /etc/bash_completion.d/qtool

#### macOS:

	qtool completion bash > $(brew --prefix)/etc/bash_completion.d/qtool

You will need to start a new shell for this setup to take effect.


```
qtool completion bash
```

### Options

```
  -h, --help              help for bash
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

###### Auto generated by spf13/cobra on 21-Aug-2022
