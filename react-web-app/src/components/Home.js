import Request from "./Request";
import Menu from "./Menu";
import { useState } from "react";
import { Row, Col } from "reactstrap";
const Home = () => {
	const[command, setCommand] = useState("Convert Private Key");
	return (
		<div className="home">
				<Row xs="2">
					<Col style={{width: "20%"}}>
						<Menu setHomeCommand={setCommand}/>
					</Col>
					<Col style={{width: "75%"}}>
						<Request command={command} />
					</Col>
				</Row>
		</div>
	);
}
 
export default Home;