import { useContext, useState } from 'react'
import { StateContext } from './StateContext'

export default function Cart() {
  const { shoppingCart, setShoppingCart } = useContext(StateContext)
  const [cartOpen, setCartOpen] = useState(false)
  const [leave, setLeave] = useState(false)

  const toggleCart = () => {
    if (cartOpen) {
      setLeave(true)
      setTimeout(() => {
        setCartOpen(false)
        setLeave(false)
      }, 500)
    } else {
      setCartOpen(true)
    }
  }

  const removeProduct = (id: string) => {
    const newCart = structuredClone(shoppingCart).filter(
      product => product.id !== id,
    )
    setShoppingCart(newCart)
  }

  return (
    <div className='cart'>
      <button
        className='cart-button'
        id='cart-button'
        onClick={() => toggleCart()}
      >
        <img src='/icons/shopping-bag.svg' alt='Cart' />
        {shoppingCart.length > 0 && (
          <span className='cart-count'>{`${shoppingCart.length} Item(s) in Cart`}</span>
        )}
        {shoppingCart.length === 0 && (
          <span className='cart-count'>0 Items in Cart</span>
        )}
      </button>
      {shoppingCart.length > 0 && (
        <div
          className={`cart-container ${cartOpen ? '' : 'hidden'} ${
            leave ? 'leave' : ''
          }`}
          id='cart-container'
        >
          <ul>
            {shoppingCart.map(product => (
              <li key={product.id}>
                <div className='cart-item'>
                  <span>{product.name}</span>
                  <button
                    className='icon-button'
                    type='button'
                    onClick={() => removeProduct(product.id)}
                  >
                    <img src='/icons/trash.svg' alt='Remove' />
                  </button>
                </div>
              </li>
            ))}
            <button
              className='icon-button'
              style={{
                gap: '1rem',
                padding: '1rem',
                background: '#3880ff',
                color: '#fff',
              }}
            >
              <img
                src='/icons/shopping-cart.svg'
                alt=''
                style={{ filter: 'invert(100%)' }}
              />
              <span>Checkout</span>
            </button>
          </ul>
        </div>
      )}
    </div>
  )
}
