import { useQuery } from "@tanstack/react-query"
import { Link } from "react-router-dom"
import { getJobs } from "../api/jobs"
import type { Job } from "../types"

function JobsPage() {
    const { data: jobs, isLoading, isError } = useQuery({
        queryKey: ['jobs'],
        queryFn: getJobs
    })

    if (isLoading) return <div className="text-center mt-10">Загрузка...</div>
    if (isError) return <div className="text-center mt-10 text-red-500">Ошибка загрузки!</div>

    return (
        <div className="max-w-5xl mx-auto px-4 py-8">
            <h1 className="text-3xl font-bold mb-8">Вакансии</h1>
            <div className="grid grid-cols-1 gap-6"> {jobs?.map((job: Job) => (
                <div key={job.id} className="bg-white rounded-lg shadow p-6">
                    <h2 className="text-xl font-bold">{job.title}</h2>
                    <p className="text-gray-500 mt-1">{job.city}</p>
                    <p className="text-gray-500 mt-1">{job.salary_min} - {job.salary_max} $</p>
                    <Link to={`/jobs/${job.id}`} className="mt-4 inline-block bg-blue-600 text-white px-4 py-2 rounded">Подробнее</Link>
                </div>
            ))}</div>
        </div>
    )

}


export default JobsPage