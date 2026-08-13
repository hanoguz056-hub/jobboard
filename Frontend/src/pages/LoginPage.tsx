import React, { useState } from "react"
import { useNavigate } from "react-router-dom"
import { login } from "../api/auth"



function LoginPage() {


    const [email, setEmail] = useState('')
    const [password, setPassword] = useState('')
    const [error, setError] = useState('')
    const navigate = useNavigate()


    const handlerSumbit = async (e: React.FormEvent) => {
        e.preventDefault()
        try {
            const response = await login({ email, password })
            console.log('response', response)
            console.log('token', response.access_token)
            localStorage.setItem('token', response.access_token)
            localStorage.setItem('user', JSON.stringify(response.user))
            navigate('/')
        } catch {
            setError('Неверный email или пароль')
            console.log('error:', error)
        }
    }

    return (
        <div className="min-h-screen flex items-center justify-center bg-gray-50" >
            <div className="bg-white p-8 rounded-lg shadow w-full max-w-md">
                <form onSubmit={handlerSumbit}>
                    <input type="email" value={email} onChange={(e) => setEmail(e.target.value)} placeholder="Email" className="w-full border rounded px-3 py-2 mb-4" />
                    <input type="password" value={password} onChange={(e) => setPassword(e.target.value)} placeholder="password" className="w-full border rounded px-3 py-2 mb-4" />
                    {error && <p className="text-red-500 mb-4">{error}</p>}
                    <button type="submit" className="w-full bg-blue-600 text-white py-2 rounded">Войти</button>
                </form>
            </div>
        </div>
    )
}

export default LoginPage