import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import { createJob } from "../api/jobs";
import type { Job } from "../types";



function CreateJobPage() {
    const [form, setForm] = useState({
        title: '',
        description: '',
        city: '',
        type: 'full',
        salary_min: 0,
        salary_max: 0,
        company_id: '',
    })

    const navigate = useNavigate()
    const [error, setError] = useState('')

    const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement | HTMLTextAreaElement>) => {
        const { name, value } = e.target

        setForm({ ...form, [name]: name === 'salary_min' || name === 'salary_max' ? Number(value) : value })
    }

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault()
        try {
            await createJob(form as Job)
            navigate('/dashboard')
        } catch {
            setError('Ошибка создания вакансии')
        }
    }

    return (
        <div className="max-w-2xl mx-auto px-4 py-8">
            <h1>Создать вакансию</h1>
            <form onSubmit={handleSubmit} className="bg-white rounded-lg shadow p-6 flex flex-col gap-4">
                <input name="title" onChange={handleChange} placeholder="Название" className="border rounded px-3 py-2" />
                <textarea name="description" onChange={handleChange} placeholder="Описание" className="border rounded px-3 py-2 h-32" />
                <input name="city" onChange={handleChange} placeholder="Город" className="border rounded px-3 py-2" />
                <input name="company_id" onChange={handleChange} placeholder="ID компании" className="border rounded px-3 py-2" />
                <select name="type" onChange={handleChange} className="border rounded px-3 py-2">
                    <option value="full">Полная занятость</option>
                    <option value="part">Частичная</option>
                    <option value="remote">Удаленно</option>
                </select>
                <input name="salary_min" type="number" onChange={handleChange} placeholder="Зарплата от" className="border rounded px-3 py-2" />
                <input name="salary_max" type="number" onChange={handleChange} placeholder="Зарплата до" className="border rounded px-3 py-2" />
                {error && <p className="text-red-500">{error}</p>}
                <button type="submit" className="bg-blue-600 text-white py-2 rounded hover:bg-blue-700" >Создать</button>
            </form>
        </div>
    )
}


export default CreateJobPage