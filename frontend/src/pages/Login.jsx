import { useState } from "react";
import api from "../api";

export default function Login({ setUser }) {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [error, setError] = useState("");
    const [loading, setLoading] = useState(false);

    const handleLogin = async (e) => {
        e.preventDefault();
        setError("");

        if (!email || !password) {
            setError("Email and password are required");
            return;
        }

        try {
            setLoading(true);
            await api.post("/auth/login", { email, password });
            const res = await api.get("/api/me");
            setUser(res.data);
        } catch {
            setError("Invalid email or password");
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
                        Split expenses. Stay fair.
                    </p>
                </div>

                {/* LOGIN CARD */}
                <div className="bg-white rounded-xl shadow-sm p-8">
                    <h2 className="text-xl font-semibold text-gray-900 mb-6">
                        Sign in to your account
                    </h2>

                    <form onSubmit={handleLogin} className="space-y-4">
                        <div>
                            <label className="block text-sm font-medium text-gray-700 mb-1">
                                Email
                            </label>
                            <input
                                type="email"
                                className="w-full border border-gray-300 rounded-lg px-4 py-2
                           focus:outline-none focus:ring-2 focus:ring-blue-500
                           focus:border-blue-500 transition"
                                placeholder="you@example.com"
                                value={email}
                                onChange={(e) => setEmail(e.target.value)}
                            />
                        </div>

                        <div>
                            <label className="block text-sm font-medium text-gray-700 mb-1">
                                Password
                            </label>
                            <input
                                type="password"
                                className="w-full border border-gray-300 rounded-lg px-4 py-2
                           focus:outline-none focus:ring-2 focus:ring-blue-500
                           focus:border-blue-500 transition"
                                placeholder="••••••••"
                                value={password}
                                onChange={(e) => setPassword(e.target.value)}
                            />
                        </div>

                        {error && (
                            <p className="text-red-500 text-sm">{error}</p>
                        )}

                        <button
                            disabled={loading}
                            className="w-full bg-blue-600 text-white py-2 rounded-lg
                         hover:bg-blue-700 transition
                         disabled:opacity-50 disabled:cursor-not-allowed"
                        >
                            {loading ? "Signing in..." : "Sign In"}
                        </button>
                        <p className="text-sm text-center text-gray-600 mt-4">
                            New here?{" "}
                            <a href="/register" className="text-blue-600 hover:underline">
                                Create an account
                            </a>
                        </p>
                    </form>
                </div>
            </div>
        </div>
    );
}
