import { LogOut, UserCircle, Wallet, History } from "lucide-react";
import api from "../api";
import { useNavigate } from "react-router-dom";

export default function Navbar({ user, onLogout }) {
  const navigate = useNavigate();

  const handleLogout = async () => {
    await api.post("/auth/logout");
    onLogout();
  };

  return (
    <nav className="bg-blue-50 border-b shadow-sm">
      <div className="max-w-6xl mx-auto px-6 py-5 flex justify-between items-center">
        
        <div
          className="flex items-center gap-3 cursor-pointer"
          onClick={() => navigate("/groups")}
        >
          <Wallet className="w-7 h-7 text-blue-600" />
          <h1 className="text-2xl font-bold text-gray-900 tracking-tight">
            ExpenseSplitter
          </h1>
        </div>

        <div className="flex items-center gap-4">
          <div className="flex items-center gap-2 bg-gray-100 px-4 py-2 rounded-full">
            <UserCircle className="w-6 h-6 text-gray-600" />
            <span className="text-sm text-gray-800 font-medium">
              {user?.email}
            </span>
          </div>

          <button
            onClick={() => navigate("/paymenthistory")}
            className="flex items-center gap bg-white  px-4 py-2 rounded-full text-sm font-medium text-gray-700 border border-gray-200 hover:bg-gray-100
              hover:text-gray-900
              transition">
            <History className="w-4 h-4" />
            History
          </button>

          <button
            onClick={handleLogout}
            className="flex items-center gap-2 text-sm bg-red-500 text-white px-4 py-2 rounded-full hover:bg-red-600 transition">
            <LogOut className="w-4 h-4" />
            Logout
          </button>
        </div>
      </div>
    </nav>
  );
}
