import NavBar from './components/NavBar';
import Home from './components/Home'
import {
  Routes,
  Route,
} from "react-router-dom";

function App() {
  return (
    <div className="App">
          <NavBar/>
        <Routes>
            <Route exact path="/" element={<Home/>}/>
        </Routes>
    </div>
  );
}

export default App;
