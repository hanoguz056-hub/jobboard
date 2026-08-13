import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.tsx'

const queryclient = new QueryClient()

createRoot(document.getElementById('root')!).render(

  < QueryClientProvider client={queryclient} >
    <App />
  </QueryClientProvider >

)
