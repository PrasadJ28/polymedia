import { Outlet, RouterProvider } from 'react-router-dom'
import './App.css'
import router from './UI/Routes/routes'

function App() {
  return (
    <div className="App">
      <RouterProvider router={router} />
      <Outlet />
    </div>
  )
}

export default App
