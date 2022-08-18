import { useEffect } from "react";
import { CardBody, Card, CardHeader } from "reactstrap";
import Output from "./Output";

const Response = ({output, command}) => {

useEffect(() => {
	console.log("Component Result ->output: ", output);
}, [output, command]);


	return (
		<div className="Output">
			<form>
				<Card className="card">
					<CardHeader style={{
						backgroundColor: "#f1356d",
						color: "#ffffff",
						width: "100%",
						fontSize: "15px"
					
					}}>
						Output
					</CardHeader>
					<CardBody>
							<Output output={output} />
					</CardBody>
				</Card>
			</form>
		</div>
	);
}
 
export default Response;