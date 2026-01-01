import { useState } from "react"; // Added useState
import { LogOut, UserCircle, Wallet, History, Menu, X } from "lucide-react"; // Added Menu, X
import api from "../api";
import { useNavigate } from "react-router-dom";

export default function Navbar({ user, onLogout }) {
  const navigate = useNavigate();
  const [isMobileMenuOpen, setIsMobileMenuOpen] = useState(false); // New State

  const handleLogout = async () => {
    try {
      await api.post("/auth/logout");
      onLogout();
    } catch (err) {
      console.error("Logout failed", err);
    }
  };

  return (
    <nav className="sticky top-0 z-50 bg-white/80 backdrop-blur-md border-b border-slate-100">
      <div className="max-w-7xl mx-auto px-6 py-6 flex justify-between items-center">
        <div
          className="flex items-center gap-3 cursor-pointer group"
          onClick={() => navigate("/groups")}
        >
          <div className="p-2.5 bg-purple-50 rounded-xl group-hover:bg-[#7c3aed] transition-colors duration-300">
            <Wallet className="w-7 h-7 text-[#7c3aed] group-hover:text-white transition-colors" />
          </div>
          <h1 className="text-2xl font-black text-slate-900 tracking-tight">
            ExpenseSplitter
          </h1>
        </div>
        <div className="hidden md:flex items-center gap-4">
          <div className="flex items-center gap-3 bg-slate-50 border border-slate-100 px-5 py-3 rounded-2xl">
            <UserCircle className="w-6 h-6 text-slate-400" />
            <span className="text-sm font-bold text-slate-600">
              {user?.email}
            </span>
          </div>

          <div className="h-8 w-px bg-slate-200 mx-2"></div>

          <button
            onClick={() => navigate("/paymenthistory")}
            className="flex items-center gap-2.5 bg-white px-6 py-3 rounded-2xl text-sm font-black text-slate-600 border border-slate-200 hover:border-purple-200 hover:text-purple-600 hover:shadow-lg hover:shadow-purple-50 transition-all active:scale-95"
          >
            <History className="w-5 h-5" />
            <span className="hidden sm:inline">HISTORY</span>
          </button>

          <button
            onClick={handleLogout}
            className="flex items-center gap-2.5 bg-slate-900 text-white px-6 py-3 rounded-2xl text-sm font-black hover:bg-rose-600 shadow-lg shadow-slate-200 hover:shadow-rose-100 transition-all active:scale-95"
          >
            <LogOut className="w-5 h-5" />
            <span className="hidden sm:inline">LOGOUT</span>
          </button>
        </div>
        <button 
          className="md:hidden p-2 text-slate-600 hover:bg-slate-100 rounded-xl transition-colors"
          onClick={() => setIsMobileMenuOpen(!isMobileMenuOpen)}
        >
          {isMobileMenuOpen ? <X className="w-7 h-7" /> : <Menu className="w-7 h-7" />}
        </button>
      </div>
      {isMobileMenuOpen && (
        <div className="md:hidden border-t border-slate-100 bg-white/95 backdrop-blur-xl absolute w-full left-0 px-6 py-6 shadow-xl space-y-4 animate-in slide-in-from-top-5 duration-200">
                    <div className="flex items-center gap-3 bg-slate-50 p-4 rounded-2xl">
            <UserCircle className="w-6 h-6 text-slate-400" />
            <span className="text-sm font-bold text-slate-600 break-all">
              {user?.email}
            </span>
          </div>

          <button
            onClick={() => { navigate("/paymenthistory"); setIsMobileMenuOpen(false); }}
            className="w-full flex items-center justify-center gap-2.5 bg-white p-4 rounded-2xl text-sm font-black text-slate-600 border border-slate-200 active:bg-slate-50"
          >
            <History className="w-5 h-5" />
            PAYMENT HISTORY
          </button>

          <button
            onClick={handleLogout}
            className="w-full flex items-center justify-center gap-2.5 bg-slate-900 text-white p-4 rounded-2xl text-sm font-black hover:bg-rose-600 active:scale-95 transition-all"
          >
            <LogOut className="w-5 h-5" />
            LOGOUT
          </button>
        </div>
      )}
    </nav>
  );
}