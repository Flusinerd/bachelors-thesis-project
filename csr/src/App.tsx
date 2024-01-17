import { useEffect, useState } from 'react'
import { Outlet } from 'react-router-dom'
import './App.css'
import { State, StateContext } from './StateContext'
import NavBar from './components/NavBar'

function App() {
  const [shoppingCart, setShoppingCart] = useState<State['shoppingCart']>([])
  const value = { shoppingCart, setShoppingCart }

  useEffect(() => {
    const cart = localStorage.getItem('shoppingCart')
    if (cart) {
      setShoppingCart(JSON.parse(cart))
    }
  }, [])

  useEffect(() => {
    localStorage.setItem('shoppingCart', JSON.stringify(shoppingCart))
  }, [shoppingCart])

  return (
    <>
      <StateContext.Provider value={value}>
        <NavBar></NavBar>
        <main>
          <Outlet></Outlet>
        </main>
      </StateContext.Provider>
    </>
  )
}

export default App
