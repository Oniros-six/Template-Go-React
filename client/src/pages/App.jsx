import { useState } from 'react'
import axios from 'axios'

function App() {
  const [count, setCount] = useState(0)

  const consulta = async() => {
    const resp = await axios.get("/prueba")
    console.log(resp)
  }

  return (
    <div className="App">
      <div>
        {count}
      </div>
     <button onClick={()=>setCount(count+1)}>+1</button>
     <button onClick={()=>consulta()}>Consultamos?</button>

    </div>
  )
}

export default App
