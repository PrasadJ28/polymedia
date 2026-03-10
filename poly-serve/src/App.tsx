import { useRef } from "react";
import "./index.css";


export function App() {
  const inputFile = useRef(null);
  const uploadFile = () => {

  }
  return (
    <div className="app">
      <input type="file" id="file" ref={inputFile} />
      <button onClick={uploadFile}>Upload</button>
    </div>
  );
}

export default App;
