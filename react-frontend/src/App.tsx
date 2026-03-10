import { Suspense } from 'react';
import { Outlet } from 'react-router-dom';
import { Provider } from 'react-redux';
import { store } from './Reducers/store';
import './App.css';

function App() {
  return (
    <Provider store={store}>
        <main>
          <Suspense fallback={<div>Loading...</div>}>
            <Outlet />
          </Suspense>
        </main>
    </Provider>
  );
}

export default App;
