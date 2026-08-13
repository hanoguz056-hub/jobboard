import EmployerDashboard from "./EmployerDashboard"
import CandidateDashboard from "./CandidateDashboard"

function DashboardPage() {

    const userStr = localStorage.getItem('user')

    let user: { role?: string } = {}

    try {
        user = userStr ? JSON.parse(userStr) : {}
    } catch {
        user = {}
    }

    if (user.role === 'employer') return <EmployerDashboard />
    if (user.role === 'candidate') return <CandidateDashboard />

    return (
        <div>Unauthorized</div>
    )
}


export default DashboardPage