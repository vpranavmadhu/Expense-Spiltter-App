import { useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { User, Mail, Phone, Lock, ArrowRight, Wallet } from "lucide-react";
import api from "../api";
import { toast } from "react-toastify";

export default function Register() {
  const [form, setForm] = useState({ name: "", email: "", phone: "", password: "" });
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);

  const navigate = useNavigate();

  const handleChange = (e) => {
    const { name, value } = e.target;

    if (name === "phone") {
      if (!/^\d*$/.test(value)) return;
      if (value.length > 10) return;
    }

    setForm({ ...form, [name]: value });
  };

  const handleRegister = async (e) => {
    e.preventDefault();
    setError("");

    if (!form.name || !form.email || !form.password || !form.phone) {
      setError("Please fill in all required fields");
      return;
    }

    if (form.phone.length !== 10) {
      setError("Phone number must be exactly 10 digits");
      return;
    }

    try {
      setLoading(true);
      
      await api.post("/auth/register", {
        name: form.name,
        email: form.email,
        phone: form.phone, 
        password: form.password
      });

      toast("Registration Successful! Please login.");
      navigate("/login");

    } catch (err) {
      console.error(err); 
      setError(err.response?.data?.error || "Registration failed. Please check your details.");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen bg-slate-200 flex items-center justify-center p-6 relative overflow-hidden">
      <div className="absolute top-[-10%] right-[-10%] w-160 h-160 bg-purple-100/50 rounded-full blur-3xl -z-10"></div>
      <div className="absolute bottom-[-10%] left-[-10%] w-120 h-120 bg-indigo-100/50 rounded-full blur-3xl -z-10"></div>

      <div className="w-full max-w-md">
        
        <div className="text-center mb-10">
          <div className="inline-flex items-center justify-center w-16 h-16 bg-white rounded-3xl shadow-lg shadow-purple-100 mb-6 text-[#7c3aed]">
            <Wallet className="w-8 h-8" />
          </div>
          <h1 className="text-4xl font-black text-slate-900 tracking-tighter">ExpenseSplitter</h1>
          <p className="text-slate-400 font-bold uppercase tracking-widest text-xs mt-3">Join the Workspace</p>
        </div>

        <div className="bg-white/80 backdrop-blur-xl rounded-[2.5rem] shadow-2xl shadow-slate-200/50 border border-white p-10">
          <h2 className="text-xl font-black text-slate-800 mb-8">Create Account</h2>

          <form onSubmit={handleRegister} className="space-y-4">
            
            <div className="relative group">
              <User className="absolute left-5 top-1/2 -translate-y-1/2 w-5 h-5 text-slate-400 group-focus-within:text-[#7c3aed] transition-colors" />
              <input 
                name="name"
                type="text"
                className="w-full bg-slate-50 border border-transparent rounded-2xl pl-14 pr-6 py-4 font-bold text-slate-700 placeholder:text-slate-300 focus:bg-white focus:border-purple-100 focus:outline-none focus:ring-4 focus:ring-purple-50 transition-all"
                placeholder="Full Name" 
                value={form.name} 
                onChange={handleChange} 
              />
            </div>

            <div className="relative group">
              <Mail className="absolute left-5 top-1/2 -translate-y-1/2 w-5 h-5 text-slate-400 group-focus-within:text-[#7c3aed] transition-colors" />
              <input 
                name="email"
                type="email"
                className="w-full bg-slate-50 border border-transparent rounded-2xl pl-14 pr-6 py-4 font-bold text-slate-700 placeholder:text-slate-300 focus:bg-white focus:border-purple-100 focus:outline-none focus:ring-4 focus:ring-purple-50 transition-all"
                placeholder="Email Address" 
                value={form.email} 
                onChange={handleChange} 
              />
            </div>

            <div className="relative group">
              <Phone className="absolute left-5 top-1/2 -translate-y-1/2 w-5 h-5 text-slate-400 group-focus-within:text-[#7c3aed] transition-colors" />
              <input 
                name="phone"
                type="tel" 
                maxLength={10}
                className="w-full bg-slate-50 border border-transparent rounded-2xl pl-14 pr-6 py-4 font-bold text-slate-700 placeholder:text-slate-300 focus:bg-white focus:border-purple-100 focus:outline-none focus:ring-4 focus:ring-purple-50 transition-all"
                placeholder="Phone Number (10 digits)" 
                value={form.phone} 
                onChange={handleChange} 
              />
            </div>

            <div className="relative group">
              <Lock className="absolute left-5 top-1/2 -translate-y-1/2 w-5 h-5 text-slate-400 group-focus-within:text-[#7c3aed] transition-colors" />
              <input 
                name="password"
                type="password"
                className="w-full bg-slate-50 border border-transparent rounded-2xl pl-14 pr-6 py-4 font-bold text-slate-700 placeholder:text-slate-300 focus:bg-white focus:border-purple-100 focus:outline-none focus:ring-4 focus:ring-purple-50 transition-all"
                placeholder="Password" 
                value={form.password} 
                onChange={handleChange} 
              />
            </div>

            {error && (
               <div className="bg-rose-50 text-rose-500 text-xs font-bold px-4 py-3 rounded-xl border border-rose-100 text-center">
                 {error}
               </div>
            )}

            <button 
                disabled={loading}
                className="w-full bg-[#7c3aed] hover:bg-[#6d28d9] text-white py-4 rounded-2xl font-black text-xs uppercase tracking-[0.2em] shadow-xl shadow-purple-200 transition-all active:scale-95 disabled:opacity-50 disabled:cursor-not-allowed flex items-center justify-center gap-2 group mt-2"
            >
                {loading ? "CREATING..." : "CREATE ACCOUNT"}
                {!loading && <ArrowRight className="w-4 h-4 group-hover:translate-x-1 transition-transform" />}
            </button>
          </form>

          <div className="mt-8 pt-6 border-t border-slate-100 text-center">
            <p className="text-xs font-bold text-slate-400">
                Already have an account?
            </p>
            <Link 
                to="/login" 
                className="inline-block mt-2 text-sm font-black text-[#7c3aed] hover:text-[#6d28d9] hover:underline"
            >
                Sign In
            </Link>
          </div>
        </div>
      </div>
    </div>
  );
}