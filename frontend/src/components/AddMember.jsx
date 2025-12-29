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
    <div className="flex gap-2">
      <input
        type="email"
        placeholder="Enter email to add"
        value={email}
        onChange={(e) => setEmail(e.target.value)}
        className="border rounded px-3 py-2 flex-1 focus:outline-none focus:ring-2 focus:ring-blue-500"
      />
      <button
        onClick={addMember}
        disabled={loading}
        className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700 disabled:opacity-50"
      >
        Add
      </button>
    </div>
  )
}
