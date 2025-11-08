import { Suspense } from 'react';
import { Outlet } from 'react-router-dom';
import { Provider } from 'react-redux';
import { store } from './Reducers/store';
import { CssVarsProvider } from '@mui/joy/styles';
import { theme } from './Style/theme';
import './App.css';

function App() {
  return (
    <Provider store={store}>
      <CssVarsProvider theme={theme} defaultMode="system">
        <main>
          <Suspense fallback={<div>Loading...</div>}>
            <Outlet />
          </Suspense>
        </main>
      </CssVarsProvider>
    </Provider>
  );
}

export default App;
