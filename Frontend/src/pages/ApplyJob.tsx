import React, { useState } from "react"
import { useNavigate, useParams } from "react-router-dom"
import { applyToJob } from "../api/applications"


function ApplyJob() {
    const [file, setFile] = useState<File | null>(null)
    const [coverLetter, setCoverLetter] = useState('')

    const { id } = useParams<{ id: string }>()
    const navigate = useNavigate()
    const [error, setError] = useState('')

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault()
        if (!file) {
            setError("Выберите файл резюме")
            return
        }
        try {
            const formData = new FormData()
            formData.append('resume', file)
            formData.append('cover_letter', coverLetter)
            await applyToJob(id!, formData)
            navigate("/dashboard")
        } catch {
            setError('Ошибка при отклике')
        }
    }


    return (
        <div className="max-w-2xl mx-auto px-4 py-8">
            <h1 className="text-3xl font-bold mb-8">Откликнуться</h1>
            <form onSubmit={handleSubmit} className="bg-white rounded-lg shadow p-6 flex flex-col gap-4">
                <textarea value={coverLetter} onChange={(e) => setCoverLetter(e.target.value)} placeholder="Сопроводительное письмо" className="border rounded px-3 py-2 h-32" />
                <input type="file" accept=".pdf" onChange={(e) => setFile(e.target.files?.[0] || null)} className="border rounded px-3 py-2" />
                {error && <p className="text-red-500">{error}</p>}
                <button type="submit" className="bg-blue-600 text-white py-2 rounded hover:bg-blue-700">Отправить отклик</button>
            </form>
        </div>
    )

}

export default ApplyJob