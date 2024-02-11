import { State, StateContext } from '@/components/StateContext'
import '@/styles/globals.css'
import type { AppProps } from 'next/app'
import { useEffect, useRef, useState } from 'react'

export default function App({ Component, pageProps }: AppProps) {
  const [shoppingCart, setShoppingCart] = useState<State['shoppingCart']>([])
  const value = { shoppingCart, setShoppingCart }
  const isInitialLoad = useRef(true)

  useEffect(() => {
    const cart = localStorage.getItem('shoppingCart')
    if (cart) {
      setShoppingCart(JSON.parse(cart))
    }
  }, [])

  useEffect(() => {
    if (isInitialLoad.current) {
      isInitialLoad.current = false
      return
    }
    localStorage.setItem('shoppingCart', JSON.stringify(shoppingCart))
  }, [shoppingCart])

  return (
    <>
      <StateContext.Provider value={value}>
        <Component {...pageProps} />
      </StateContext.Provider>
    </>
  )
}
