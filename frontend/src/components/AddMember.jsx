import { useState } from "react"
import axios from "axios"

export default function AddMember({ groupId, onAdded }) {
  const [email, setEmail] = useState("")
  const [loading, setLoading] = useState(false)

  const addMember = async () => {
    if (!email) return alert("Email required")
    try {
      setLoading(true)
      await axios.post(
        `http://localhost:8080/api/groups/${groupId}/addmember`,
        { email },
        { withCredentials: true }
      )
      setEmail("")
      onAdded()
    }
    catch (err) { alert(err.response?.data?.error || "Failed") } finally { setLoading(false) }
  }

  return (
    <div className="flex flex-col gap-2 pt-6 border-t border-gray-50">
      <input
        type="email"
        placeholder="Add member by email"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
        className="w-full bg-slate-50 border-none px-4 py-3 text-sm rounded-2xl focus:ring-2 focus:ring-purple-500 transition-all outline-none"
      />
      <button
        onClick={addMember}
        disabled={loading}
        className="w-full py-3 text-xs font-black text-white rounded-2xl bg-slate-900 hover:bg-black transition-all shadow-lg active:scale-95"
      >
        {loading ? "ADDING..." : "ADD MEMBER"}
      </button>
    </div>
  )
}