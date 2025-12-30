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
    } catch (err) {
      alert(err.response?.data?.error || "Failed to add member")
    } finally {
      setLoading(false)
    }
  }

  return (
    <div className="
      flex items-center gap-3
      bg-white/80 backdrop-blur
      border border-gray-100
      rounded-xl p-3
      shadow-sm
    ">
      <input
        type="email"
        placeholder="Add member by email"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
        className="
          flex-1
          bg-transparent
          px-4 py-2
          text-sm
          rounded-full
          border border-gray-200
          focus:outline-none
          focus:ring-2 focus:ring-blue-500
          focus:border-blue-500
          transition
        "
      />

      <button
        onClick={addMember}
        disabled={loading}
        className="
          px-5 py-2
          text-sm font-medium
          text-white
          rounded-full
          bg-blue-600
          hover:bg-blue-700
          shadow-sm hover:shadow-md
          disabled:opacity-50
          disabled:cursor-not-allowed
          transition
        "
      >
        {loading ? "Adding…" : "Add"}
      </button>
    </div>
  )
}
