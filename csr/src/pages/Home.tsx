import { useEffect, useState } from 'react'
import ProductCard from '../components/ProductCard'
import { Product } from '../product'

export default function Home() {
  const [products, setProducts] = useState<Product[]>([])

  useEffect(() => {
    async function fetchProducts() {
      const response = await fetch('http://localhost:8080/api/products')
      const products = await response.json()
      setProducts(products)
    }

    fetchProducts()
  }, [])

  return (
    <ul className='products'>
      {products.map(product => (
        <ProductCard key={product.id} product={product}></ProductCard>
      ))}
    </ul>
  )
}
