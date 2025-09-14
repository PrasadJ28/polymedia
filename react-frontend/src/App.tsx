import { Suspense } from 'react'
import './App.css'
import { Outlet } from 'react-router-dom'
import { Provider } from 'react-redux'
import { store } from './Reducers/store'
import { theme } from './Style/theme'
import { CssVarsProvider } from '@mui/joy/styles'

function App() {

  return (
    <>
      <Provider store={store}>
        <CssVarsProvider theme={theme} defaultMode="system">

          <div className="text-3xl font-bold underline">
            <main>
              <Suspense fallback={<div>Loading...</div>}>
                <Outlet/>
              </Suspense>
            </main>
            Hello
          </div>
        </CssVarsProvider>
      </Provider>
    </>
  )
}

export default App
