import { useEffect, useState } from "react"
import { useParams } from "react-router-dom"
import axios from "axios"

import AddMember from "../components/AddMember"
import ExpenseList from "../components/ExpenseList"
import AddExpenseModal from "../components/AddExpenseModel"

export default function GroupDetail() {
  const { groupId } = useParams()

  const [user, setUser] = useState(null)
  const [group, setGroup] = useState(null)
  const [members, setMembers] = useState([])
  const [expenses, setExpenses] = useState([])
  const [balances, setBalances] = useState({})

  const [showAddExpense, setShowAddExpense] = useState(false)
  const fetchMe = async () => {
    const res = await axios.get(
      "http://localhost:8080/api/me",
      { withCredentials: true }
    )
    setUser(res.data)
  }

  const fetchGroup = async () => {
    const res = await axios.get(
      `http://localhost:8080/api/groups/${groupId}`,
      { withCredentials: true }
    )
    setGroup(res.data)
  }

  const fetchMembers = async () => {
    const res = await axios.get(
      `http://localhost:8080/api/groups/${groupId}/members`,
      { withCredentials: true }
    )
    setMembers(res.data)
  }

  const fetchExpenses = async () => {
    const res = await axios.get(
      `http://localhost:8080/api/groups/${groupId}/expenses`,
      { withCredentials: true }
    )
    setExpenses(res.data)
  }

  const fetchBalances = async () => {
    const res = await axios.get(
      `http://localhost:8080/api/groups/${groupId}/balances`,
      { withCredentials: true }
    )
    setBalances(res.data)
  }

  useEffect(() => {
    if (!groupId) return

    fetchMe()
    fetchGroup()
    fetchMembers()
    fetchExpenses()
    fetchBalances()
  }, [groupId])

  if (!group || !user) {
    return (
      <div className="flex justify-center items-center h-screen text-gray-500">
        Loading...
      </div>
    )
  }

  console.log(expenses);
  

  return (
    <div className="max-w-4xl mx-auto p-6 space-y-6">

      <h1 className="text-3xl font-bold">{group.name}</h1>

      <div className="bg-white shadow rounded p-4">
        <h2 className="text-lg font-semibold mb-4">Members</h2>

        <div className="flex flex-wrap gap-2 mb-4">
          {members.map((m) => (
            <span
              key={m.id}
              className="px-4 py-1 bg-gray-100 rounded-full text-sm font-medium"
            >
              {m.username}
            </span>
          ))}
        </div>

        <AddMember groupId={groupId} onAdded={fetchMembers} />
      </div>

      <button
        onClick={() => setShowAddExpense(true)}
        className="bg-green-600 text-white px-4 py-2 rounded hover:bg-green-700"
      >
        + Add Expense
      </button>

      <div className="bg-white shadow rounded p-4">
        <h2 className="text-lg font-semibold mb-3">Expenses</h2>
        <ExpenseList
          expenses={expenses}
          balances={balances}
          currentUserId={user.id}
          onSettled={() => {
            fetchExpenses()
            fetchBalances()
          }}
        />
      </div>

      {showAddExpense && (
        <AddExpenseModal
          groupId={groupId}
          members={members}
          onClose={() => setShowAddExpense(false)}
          onAdded={() => {
            fetchExpenses()
            fetchBalances()
          }}
        />
      )}
    </div>
  )
}
