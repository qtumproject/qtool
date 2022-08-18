import { commands, methods, sample } from "./constants";

export const createJSONRequest = (command, data, format, blockchain, network) => {
	let method, params;

	params = {
		"data": data,
		"format": format,
		"blockchain": blockchain,
		"network": network
	};

	switch (command) {
		case commands[0]:
			method = methods[0];
			// params[0].key = data;
			break;
		case commands[1]:
			method = methods[1];
			// params[0].key = data;
			break;
		case commands[2]:
			method = methods[2];
			// params[0].key = data;
			break;
		case commands[3]:
			method = methods[3];
			// params[0].key = data;
			break;
		default:
	}
	const jsonrpc = "2.0";
	const id = 1;
	let request = {
		method,
		params,
		jsonrpc,
		id
	};
	return request;
}

export const getAPIEndpoint = (command) => {
	let endpoint;
	switch (command) {
		case commands[0]:
			endpoint = "http://localhost:8080/privatekey";
			break;
		case commands[1]:
			endpoint = "http://localhost:8080/address";
			break;
		case commands[2]:
			endpoint = "http://localhost:8080/privatekey";
			break;
		case commands[3]:
			endpoint = "http://localhost:8080/script";
			break;
		default:
	}
	return endpoint;
}

export const displayDefaultValues = (setData, command, format, blockchain, network) => {
	switch (command) {
		case commands[0]:
			if (format==="b58") {
				setData(sample.qtum_privateKeyWIF);
			} else {
				setData(sample.qtum_privateKeyHex);
			}
			break;
		case commands[1]:
			if (format==="b58") {
				switch (network) {
					case "testnet":
						blockchain === "qtum" ? setData(sample.qtum_addressBase58Testnet) 
						: setData(sample.btc_addressBase58Testnet);
						break;
					case "mainnet":
						blockchain === "qtum" ? setData(sample.qtum_addressBase58Mainnet)
						: setData(sample.btc_addressBase58Mainnet);
						break;
					default:
						break;
				}
			} else {
				blockchain === "qtum" ? setData(sample.qtum_addressHex) : setData(sample.btc_addressHex);
			}
			break;
		case commands[2]:
			if (format==="b58") {
				blockchain === "qtum" ? setData(sample.qtum_privateKeyWIF) : setData(sample.btc_privateKeyWIF);
			} else {
				blockchain === "qtum" ? setData(sample.qtum_privateKeyHex) : setData(sample.btc_privateKeyHex);
			}
			break;
		case commands[3]:
			setData(sample.qtum_scriptpubkey_33);
			break;
		default:
			break;
	}
}
