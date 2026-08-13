import { useQuery } from "@tanstack/react-query";
import { Link, useParams } from "react-router-dom";
import { getJobByID } from "../api/jobs";


function JobPage() {
    const { id } = useParams<{ id: string }>()

    const { data: job, isLoading } = useQuery({
        queryKey: ['job', id],
        queryFn: () => getJobByID(id!)
    })


    const userStr = localStorage.getItem('user')
    let user: { role?: string } = {}
    try {
        user = userStr ? JSON.parse(userStr) : {}
    } catch {
        user = {}
    }

    const isCandidate = user.role === 'candidate'

    if (isLoading) return <div className="text-center mt-10">Загрузка...</div>
    if (!job) return <div className="text-center mt-10">Вакансия не найдена</div>

    return (
        <div className="max-w-3xl mx-auto px-4 py-8">
            <h1 className="text-3xl font-bold mb-4">
                {job.title}
            </h1>
            <p className="text-gray-500 mb-2">{job.city}</p>
            <p className="text-green-600 mb-4">{job.salary_min} - {job.salary_max} $</p>
            <p className="text-gray-700 mb-6">{job.description}</p>

            {isCandidate && (
                <Link to={`/jobs/${job.id}/apply`} className="bg-blue-600 text-white px-6 py-3 rounded hover:bg-blue-700">Откликнуться</Link>
            )}
        </div>
    )

}

export default JobPage