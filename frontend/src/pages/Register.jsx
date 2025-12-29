import { useState } from "react";
import api from "../api";

export default function Register({ setUser }) {
  const [form, setForm] = useState({
    name: "",
    email: "",
    phone: "",
    password: "",
  });

  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);

  const handleChange = (e) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleRegister = async (e) => {
    e.preventDefault();
    setError("");

    if (!form.name || !form.email || !form.phone || !form.password) {
      setError("All fields are required");
      return;
    }

    try {
      setLoading(true);

      await api.post("/auth/register", {
        name: form.name,
        email: form.email,
        phone: form.phone,
        password: form.password,
      });

      // auto login
      const res = await api.get("/api/me");
      setUser(res.data);
    } catch {
      setError("Registration failed");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="min-h-screen flex items-center justify-center bg-gray-100">
      <div className="w-full max-w-md">
        {/* APP NAME */}
        <div className="text-center mb-6">
          <h1 className="text-3xl font-bold text-gray-900">
            ExpenseSplitter
          </h1>
          <p className="text-gray-600 mt-1">
            Create your account
          </p>
        </div>

        {/* REGISTER CARD */}
        <div className="bg-white rounded-xl shadow-sm p-8">
          <h2 className="text-xl font-semibold text-gray-900 mb-6">
            Sign up
          </h2>

          <form onSubmit={handleRegister} className="space-y-4">
            <Input label="Name" name="name" value={form.name} onChange={handleChange} />
            <Input label="Email" name="email" type="email" value={form.email} onChange={handleChange} />
            <Input label="Phone" name="phone" value={form.phone} onChange={handleChange} />
            <Input label="Password" name="password" type="password" value={form.password} onChange={handleChange} />

            {error && (
              <p className="text-red-500 text-sm">{error}</p>
            )}

            <button
              disabled={loading}
              className="w-full bg-blue-600 text-white py-2 rounded-lg
                         hover:bg-blue-700 transition
                         disabled:opacity-50"
            >
              {loading ? "Creating account..." : "Create Account"}
            </button>
            <p className="text-sm text-center text-gray-600 mt-4">
                            Already a User?{" "}
                            <a href="/login" className="text-blue-600 hover:underline">
                                Log In here
                            </a>
                        </p>
          </form>
        </div>
      </div>
    </div>
  );
}

function Input({ label, ...props }) {
  return (
    <div>
      <label className="block text-sm font-medium text-gray-700 mb-1">
        {label}
      </label>
      <input
        {...props}
        className="w-full border border-gray-300 rounded-lg px-4 py-2
                   focus:outline-none focus:ring-2 focus:ring-blue-500
                   focus:border-blue-500 transition"
      />
    </div>
  );
}
