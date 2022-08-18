import { useEffect } from "react";
import { useState } from "react";
import { Card, CardBody, Row, Col, CardHeader } from "reactstrap";
import Result from "./Response";
import { getAPIEndpoint, createJSONRequest, displayDefaultValues } from "../assets/utils";
import { commands} from "../assets/constants";

const Request = ({command}) => {
	const [data, setData] = useState("");
	const [format, setFormat] = useState("b58");
	const [blockchain, setBlockchain] = useState("qtum");
	const [network, setNetwork] = useState("testnet");
	const [isLoading, setIsLoading] = useState(false);
	const [result, setResult] = useState(null);

	useEffect(() => {
		setResult(null);
		console.log("Component => Params : command = ", command);
		displayDefaultValues(setData, command, format, blockchain, network);
	}, [command, format, blockchain, network]);

	const handleSubmit = (e) => {
		e.preventDefault();
		setIsLoading(true);

		const url = getAPIEndpoint(command);
		const jsonRequest = createJSONRequest(command, data, format, blockchain, network);
		console.log("request: ", JSON.stringify(jsonRequest));
	    
		fetch(url, {
			method: 'POST',
			headers: { "Content-Type": "application/json" },
			body: JSON.stringify(jsonRequest)
			})
			.then((result) => {
				return result.json();
			})
			.then((data) => {
				console.log("data: ", data);
				setResult(data);
			})
			.catch(err => {
				console.log("Error: ", err);
				setResult({error: err});
			});

		setIsLoading(false);
	}


	return (
		<div>
			<Row xs="2">
				<Col style={{
					width: "50%",
					}}>
					<Card className="card">
						<CardHeader style={{
							color: "#ffffff",
							background: "#0074d9",
							width: "100%",
							fontSize: "15px"
						}}>
							{command}
						</CardHeader>
						<CardBody>
							<form onSubmit={handleSubmit} className="form">
							<label>
								Input {command === commands[0] ? "Private key" 
								: command === commands[1] ? "Address" 
								: command === commands[2] ? "Private Key" 
								: "scriptPubKey"}
							</label>
								<input 
								type="text" 
								value={data}
								onChange={(e) => setData(e.target.value)} 
								style={{
									width: "530px",
									fontSize: "12px"
								}}
							/>
							<div className="options">
								{command !== commands[0] ?
									<div>
										<label>Select blockchain</label>
										<select 
										value={blockchain} 
										onChange={(e) => setBlockchain(e.target.value)}>
												<option value="qtum">Qtum</option>
												<option value="btc">Bitcoin</option>
										</select>
									</div>
								: null}

								{command === commands[0] || commands[2] ?
									<div>
										<label>Select network</label>
										<select 
										value={network} 
										onChange={(e) => setNetwork(e.target.value)}>
												<option value="testnet">Testnet</option>
												<option value="mainnet">Mainnet</option>
										</select>
									</div>
								: null}

								{command !== commands[3] ?
									<div>
										<label>Select input format</label>
										<select 
										value={format} 
										onChange={(e) => setFormat(e.target.value)}>
												<option value="b58">Base58</option>
												<option value="hex">Hexadecimal</option>
										</select>
									</div>
								: null}
							</div>

								{!isLoading ? <button>Submit</button> : <button disabled>waiting...</button>}
							</form>
						</CardBody>
					</Card>
				</Col>
				<Col style={{
					width: "50%"
				}}>
							<Result output={result} command={command}/>

				</Col>
			</Row>
		</div>
	);
}
 
export default Request;