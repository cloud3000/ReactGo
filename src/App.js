import logo from "./logo.svg";
import gopher from "./go.png";
import "./App.css";

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>This react app is served by Golang.</p>
        <img className="gopher" src={gopher} alt="gopher" />
      </header>
    </div>
  );
}

export default App;
