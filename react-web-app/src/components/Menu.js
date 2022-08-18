import { Button, Nav, Card, CardTitle, CardBody} from "reactstrap";
import { commands } from "../assets/constants";
import { useState } from "react";

const Menu = ({setHomeCommand}) => {
	let activeButtons = [false, false, false, false] 
	const [active, setActive] = useState(activeButtons);
	const handleClick = (i) => {
		for (let j = 0; j < 4 ; j++) {
			if (j === i) {
				activeButtons[j] = true;
			} else {
				activeButtons[j] = false;
			}
		}
		setHomeCommand(commands[i]);
		setActive(activeButtons);
	}

	return (
		<div className="Menu">
			<Card>
				<CardBody>
					<CardTitle>Menu</CardTitle>
			<Nav>
				<Button color="primary" active={active[0] ? true : false} onClick={()=>handleClick(0)}>{commands[0]}</Button>
				<Button color="primary" active={active[1] ? true : false} onClick={()=>handleClick(1)}>{commands[1]}</Button>
				<Button color="primary" active={active[2] ? true : false} onClick={()=>handleClick(2)}>{commands[2]}</Button>
				<Button color="primary" active={active[3] ? true : false} onClick={()=>handleClick(3)}>{commands[3]}</Button>
			</Nav>
			</CardBody>
			</Card>


		</div>
		
	);
}
 
export default Menu;