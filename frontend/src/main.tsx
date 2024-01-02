import React from 'react';
import ReactDOM from 'react-dom/client';
import { BrowserRouter } from 'react-router-dom';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

import App from './App.tsx';
import './index.css';

const queryClient = new QueryClient({});

ReactDOM.createRoot(document.getElementById('root') as Element).render(
  <React.StrictMode>
    <BrowserRouter>
      <QueryClientProvider client={ queryClient }>
        <App />
      </QueryClientProvider>
    </BrowserRouter>
  </React.StrictMode>,
);
