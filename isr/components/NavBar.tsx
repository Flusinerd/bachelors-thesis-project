import Link from 'next/link'
import Cart from './Cart'

export default function NavBar() {
  return (
    <div className='page-header'>
      <nav>
        <ul>
          <h1>Product Store</h1>
          <li>
            <Link href='/'>Home</Link>
          </li>
          <li>
            <Link href='/about'>About</Link>
          </li>
        </ul>
      </nav>
      <Cart />
    </div>
  )
}
