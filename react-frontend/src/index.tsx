import { StrictMode, Suspense } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import { RouterProvider } from 'react-router-dom'
import router from './Router/router'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <Suspense fallback={<div>Booting…</div>}>
      <RouterProvider router={router} />
    </Suspense>
  </StrictMode>,
)
