import ReactDOM from 'react-dom/client'
import { RouterProvider, createBrowserRouter } from 'react-router-dom'
import App from './App.tsx'
import './index.css'
import Home from './pages/Home.tsx'
import ProductDetailPage, {
  loader as ProductDetailsLoader,
} from './pages/ProductDetail.tsx'

const router = createBrowserRouter([
  {
    path: '/',
    element: <App />,
    children: [
      {
        path: '/',
        element: <Home />,
      },
      {
        path: 'products/:productId',
        element: <ProductDetailPage />,
        loader: ProductDetailsLoader,
      },
    ],
  },
])

ReactDOM.createRoot(document.getElementById('root')!).render(
  <RouterProvider router={router}></RouterProvider>,
)
