import { BrowserRouter, Route, Routes } from 'react-router-dom'
import './App.css'
import LoginPage from './pages/LoginPage'
import RegisterPage from './pages/RegisterPage'
import Navbar from './components/Navbar'
import JobsPage from './pages/JobsPage'
import JobPage from './pages/JobPage'
import DashboardPage from './pages/DashboardPage'
import CreateJobPage from './pages/CreateJobPage'
import ApplyJob from './pages/ApplyJob'
import ProtectedRoute from './components/ProtectedRoute'

function App() {


  return (
    <BrowserRouter>
      <Navbar />
      <Routes>
        <Route path='/login' element={<LoginPage />} />
        <Route path='/register' element={<RegisterPage />} />
        <Route path='/' element={<JobsPage />} />
        <Route path='/jobs' element={<JobsPage />} />
        <Route path='/dashboard' element={<ProtectedRoute><DashboardPage /></ProtectedRoute>} />
        <Route path='/jobs/new' element={<ProtectedRoute>< CreateJobPage /></ProtectedRoute>} />
        <Route path='/jobs/:id' element={<JobPage />} />
        <Route path='/jobs/:id/apply' element={<ApplyJob />} />
      </Routes>
    </BrowserRouter>
  )
}

export default App
