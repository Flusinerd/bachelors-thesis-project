export type ProductCardProps = {
  product: {
    id: string
    name: string
    description: string
    price: number
  }
}

export default function ProductCard({ product }: ProductCardProps) {
  const formatedPrice = new Intl.NumberFormat('en-US', {
    style: 'currency',
    currency: 'USD',
  }).format(product.price)

  return (
    <li>
      <a className='product-card' href={`/products/${product.id}`}>
        <h3>{product.name}</h3>
        <img className='product-preview' src='/images/preview.svg'></img>
        <p className='product-description'>{product.description}</p>
        <p className='product-price'>{formatedPrice}</p>
      </a>
    </li>
  )
}
