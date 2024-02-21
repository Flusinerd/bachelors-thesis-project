import { createContext } from 'react'
import { Product } from './product'

export type State = {
  shoppingCart: Product[]
  setShoppingCart: (products: Product[]) => void
}

export const StateContext = createContext<State>({
  shoppingCart: [],
  setShoppingCart: () => {},
})
