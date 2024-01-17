import { useContext } from 'react'
import { LoaderFunctionArgs, useLoaderData } from 'react-router-dom'
import { StateContext } from '../StateContext'
import { Product } from '../product'

export async function loader({ params }: LoaderFunctionArgs<unknown>) {
  const productDataReq = await fetch(
    `http://localhost:8080/api/products/${params.productId}`,
  )
  const productData = await productDataReq.json()

  return {
    product: productData,
  }
}

export default function ProductDetailPage() {
  const { product } = useLoaderData() as { product: Product }
  const { shoppingCart, setShoppingCart } = useContext(StateContext)

  const addToCart = () => {
    const cart = structuredClone(shoppingCart)
    if (cart.some(item => item.id === product.id)) {
      return
    }
    cart.push(product)
    setShoppingCart(cart)
  }

  return (
    <section id='product-detail-section'>
      <div className='product-slideshow'>
        <div className='slide'>
          <img
            className='product-slideshow-image'
            src='/images/macbook1.jpg'
            alt='Product preview'
          />
        </div>
        <div className='slide'>
          <img
            className='product-slideshow-image'
            src='/images/macbook2.jpg'
            alt='Product preview'
            loading='lazy'
          />
        </div>
      </div>
      <div className='product-details-wrapper'>
        <h1>{product.name}</h1>
        <p>{product.description}</p>
        <p>${Math.round(product.price * 100) / 100}</p>
        <button
          type='submit'
          className='add-to-cart-btn'
          onClick={() => addToCart()}
        >
          Add to cart
        </button>
      </div>
    </section>
  )
}
