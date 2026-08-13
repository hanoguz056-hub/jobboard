import { Link, useNavigate } from "react-router-dom";

function Navbar() {
    const token = localStorage.getItem('token')
    const navigate = useNavigate()

    const handleLogout = () => {
        localStorage.removeItem("token")
        localStorage.removeItem("user")
        navigate('/login')
    }

    return (
        <nav className="bg-white shadow px-6 py-4 flex justify-between items-center">
            <Link to="/" className="text-xl font-bold text-blue-600">
                <h1>JobBoard</h1>
            </Link>
            <div className="flex gap-4 items-center">
                {token ? (
                    <>
                        <Link to="/dashboard" className="text-gray-600 hover:text-blue-600">
                            Кабинет
                        </Link>
                        <button onClick={handleLogout} className="bg-red-500 text-white px-4 py-2 rounded hover:bg-red-600">Выйти</button>
                    </>
                ) : (
                    <><Link to="/login" className="text-gray-600 hover:text-blue-600">
                        Войти
                    </Link>
                        <Link to="/register" className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700">
                            Регистрация
                        </Link>
                    </>
                )}
            </div>
        </nav>
    )
}


export default Navbar