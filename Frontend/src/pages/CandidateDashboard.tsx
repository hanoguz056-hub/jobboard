import { useQuery } from "@tanstack/react-query";
import { Link } from "react-router-dom";
import { getMyApplications } from "../api/applications";
import type { Application } from "../types";


function CandidateDashboard() {
    const { data: applications, isLoading } = useQuery({
        queryKey: ['my-applications'],
        queryFn: getMyApplications
    })

    if (isLoading) return <div className="text-center mt-10">Загрузка...</div>

    return (
        <div className="max-w-5xl mx-auto px-4 py-8">
            <div className="flex justify-between items-center mb-8">
                <h1 className="text-3xl font-bold">Мои Отклики</h1>
                <Link to="/jobs" className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700">Найти вакансии</Link>
            </div>
            <div className="grid grid-cols-1 gap-4">
                {applications?.map((application: Application) => (
                    <div key={application.id} className="bg-white rounded-lg shadow p-6 flex justify-between items-center">
                        <div>
                            <p className="text-gray-500">Вакансия: {application.job_id}</p>
                            <p className="text-gray-500">Сопроводительное: {application.cover_letter}</p>
                        </div>
                        <span className={`text-sm px-3 py-1 rounded-full font-medium ${application.status === "accepted" ? "bg-green-100 text-green-700" : ""} ${application.status === "rejected" ? "bg-red-100 text-red-700" : ""} ${application.status === "pending" ? "bg-yellow-100 text-yellow-700" : ""} ${application.status === "interview" ? "bg-blue-100 text-blue-700" : ""}`}>
                            {application.status}
                        </span>
                    </div>
                ))}
            </div>
        </div>
    )
}

export default CandidateDashboard