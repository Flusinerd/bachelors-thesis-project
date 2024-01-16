import { Outlet } from 'react-router-dom'
import './App.css'
import NavBar from './components/NavBar'

function App() {
  return (
    <>
      <NavBar></NavBar>
      <main>
        <Outlet></Outlet>
      </main>
    </>
  )
}

export default App
