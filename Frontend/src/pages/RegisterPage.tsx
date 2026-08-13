import React, { useState } from "react"
import { useNavigate } from "react-router-dom"
import { register } from "../api/auth"

function RegisterPage() {

    const [email, setEmail] = useState('')
    const [password, setPassword] = useState('')
    const [error, setError] = useState('')
    const navigate = useNavigate()
    const [role, setRole] = useState<"employer" | "candidate">('candidate')

    const handlerSumbit = async (e: React.FormEvent) => {
        e.preventDefault()
        try {
            const response = await register({ email, password, role })
            localStorage.setItem("token", response.access_token)
            localStorage.setItem("user", JSON.stringify(response.user))
            navigate("/")
        } catch {
            setError("Ошибочка")
        }
    }


    return (
        <div className="min-h-screen flex items-center justify-center bg-gray-50">
            <div className="bg-white p-8 rounded-lg shadow w-full max-w-md">
                <form onSubmit={handlerSumbit}>
                    <input type="email" onChange={(e) => setEmail(e.target.value)} placeholder="Email" className="w-full border rounded px-3 py-2 mb-4" />
                    <input type="password" onChange={(e) => setPassword(e.target.value)} placeholder="Password" className="w-full border rounded px-3 py-2 mb-4" />
                    <select value={role} onChange={(e) => { setRole(e.target.value as "employer" | "candidate") }} className="w-full border rounded px-3 py-2 mb-4">
                        <option value="candidate">Соискатель</option>
                        <option value="employer">Работадатель</option>
                    </select>
                    {error && <p className="text-red-500 mb-4">{error}</p>}
                    <button type="submit" className="w-full bg-blue-600 text-white py-2 rounded">Регистрация</button>
                </form>
            </div>
        </div>
    )
}

export default RegisterPage