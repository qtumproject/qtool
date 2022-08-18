import { Link } from "react-router-dom";
import logo from "../assets/qtum.png";
const NavVar = () => {
	return (
		<nav className="navBar">
			<h1> Qtool Web</h1>
			<div className="links">
				<Link to="/">Home</Link>
				<a href="http://qtum.org"><img src={logo} alt="logo" width="50px" /></a>

			</div>
		</nav>
	  );
}
 
export default NavVar;
