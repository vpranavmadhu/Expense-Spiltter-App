import api from "../api";

export default function Navbar({ user, onLogout }) {
  const handleLogout = async () => {
    await api.post("/auth/logout"); // optional if backend supports
    onLogout();
  };

  return (
    <div className="bg-white border-b shadow-sm">
      <div className="max-w-6xl mx-auto px-6 py-3 flex justify-between items-center">
        {/* APP NAME */}
        <h1 className="text-xl font-semibold text-gray-900">
          ExpenseSplitter
        </h1>

        {/* USER INFO */}
        <div className="flex items-center gap-4">
          <span className="text-sm text-gray-600">
            {user?.email}
          </span>

          <button
            onClick={handleLogout}
            className="text-sm bg-red-500 text-white px-3 py-1 rounded
                       hover:bg-red-600 transition"
          >
            Logout
          </button>
        </div>
      </div>
    </div>
  );
}
