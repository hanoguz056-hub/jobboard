import { useQuery } from "@tanstack/react-query";
import { Link } from "react-router-dom";
import { getJobs } from "../api/jobs";
import type { Job } from "../types";

function EmployerDashboard() {
    const { data: jobs, isLoading } = useQuery({
        queryKey: ['my-jobs'],
        queryFn: getJobs
    })


    if (isLoading) return <div className="text-center mt-10">Загрузка...</div>

    return (
        <div className="max-w-5xl mx-auto px-4 py-8">
            <div className="flex justify-between items-center mb-8">
                <h1 className="text-3xl font-bold">Мои вакансии</h1>
                <Link to="/jobs/new" className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700">+ Создать вакансию</Link>
            </div>
            <div className="grid grid-cols-1 gap-4">
                {jobs?.map((job: Job) => (
                    <div key={job.id} className="bg-white rounded-lg shadow p-6 flex justify-between items-center">
                        <div>
                            <h2 className="text-xl font-bold">{job.title}</h2>
                            <p className="text-gray-500">{job.city}</p>
                            <span className={`text-sm px-2 py-1 rounded ${job.status === "open" ? "bg-green-100 text-green-700" : "bg-gray-100 text-gray-700"}`}>{job.status}</span>
                        </div>
                        <div>
                            <Link to={`/jobs/${job.id}/applications`} className="bg-gray-100 px-3 py-2 rounded hover:bg-gray-200">
                                Отклики
                            </Link>
                            <Link to={`/jobs/${job.id}/edit`} className="bg-blue-100 text-blue-700 px-3 py-2 rounded hover:bg-blue-200">Редактировать</Link>
                        </div>
                    </div>
                ))}
            </div>
        </div>
    )

}


export default EmployerDashboard