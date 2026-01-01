import { useEffect, useState } from "react"
import { useParams } from "react-router-dom"

import AddMember from "../components/AddMember"
import ExpenseList from "../components/ExpenseList"
import AddExpenseModal from "../components/AddExpenseModel"
import SettlementSuggestions from "../components/SettlementSuggestions"
import api from "../api"

export default function GroupDetail() {
  const { groupId } = useParams()

  const [user, setUser] = useState(null)
  const [group, setGroup] = useState(null)
  const [members, setMembers] = useState([])
  const [expenses, setExpenses] = useState([])
  const [balances, setBalances] = useState({})
  const [showAddExpense, setShowAddExpense] = useState(false)

  const fetchMe = async () => {
    const res = await api.get("/api/me")
    setUser(res.data)
  }

  const fetchGroup = async () => {
    const res = await api.get(`/api/groups/${groupId}`)
    setGroup(res.data)
  }

  const fetchMembers = async () => {
    const res = await api.get(`/api/groups/${groupId}/members`)
    setMembers(res.data || [])
  }

  const fetchExpenses = async () => {
    const res = await api.get(`/api/groups/${groupId}/expenses`)
    setExpenses(res.data || [])
  }

  const fetchBalances = async () => {
    const res = await api.get(`/api/groups/${groupId}/balances`)
    setBalances(res.data || {})
  }

  const handleSettle = async (expense) => {
    await api.post("/api/settlements",
      {
        group_id: Number(groupId),
        expense_id: expense.id,
        to_user_id: expense.paidById,
        amount: expense.myShare
      })
    await fetchExpenses()
    await fetchBalances()
  }

  useEffect(() => {
    if (!groupId) return
    fetchMe()
    fetchGroup()
    fetchMembers()
    fetchExpenses()
    fetchBalances()
  }, [groupId])

  if (!group || !user) return <div className="flex justify-center items-center h-screen font-black text-slate-300 uppercase tracking-widest animate-pulse">Initializing Dashboard...</div>

  return (
    <div className="min-h-screen bg-slate-200 p-6 lg:p-12">
      <div className="max-w-8xl mx-auto">

        <div className="mb-14">
          <span className="text-[15px] font-black text-purple-500 uppercase tracking-[0.4em] mb-3 block">Group Workspace</span>
          <h1 className="text-6xl font-black text-slate-900 tracking-tighter leading-none">{group.name}</h1>
        </div>

        <div className="grid grid-cols-1 lg:grid-cols-12 gap-14 items-start">

          <div className="lg:col-span-3">
            <div className="bg-white border border-gray-50 rounded-[3rem] p-10 shadow-sm">
              <h2 className="text-[20px] font-black text-slate-500 uppercase tracking-[0.2em] mb-8">Group Members</h2>
              <div className="space-y-6 mb-10">
                {members.map(m => (
                  <div key={m.id} className="flex items-center gap-4 group cursor-default">
                    <div className="w-12 h-12 rounded-2xl bg-slate-50 text-slate-400 flex items-center justify-center text-sm font-black group-hover:bg-[#7c3aed] group-hover:text-white transition-all duration-300">
                      {m.username[0].toUpperCase()}
                    </div>
                    <div className="flex flex-col">
                      <span className="text-slate-800 font-bold text-base tracking-tight">{m.username}</span>
                      {m.id === user.id && <span className="text-[9px] font-black text-purple-500 uppercase tracking-wider">You</span>}
                    </div>
                  </div>
                ))}
              </div>
              <AddMember groupId={groupId} onAdded={fetchMembers} />
            </div>
          </div>

          <div className="lg:col-span-6 space-y-8">

            <div className="bg-slate-700 rounded-[2.5rem] p-8 flex items-center justify-between shadow-2xl shadow-slate-200">
              <div className="flex flex-col">
                <h2 className="text-2xl font-black text-white tracking-tight italic">Activity Feed</h2>
                <span className="text-[10px] font-bold text-slate-400 uppercase tracking-widest">{expenses.length} Total Transactions</span>
              </div>

              <button
                onClick={() => setShowAddExpense(true)}
                className="bg-[#7c3aed] hover:bg-[#8b5cf6] text-white px-8 py-3.5 rounded-2xl font-black text-xs shadow-lg transition-all active:scale-95 flex items-center gap-2 group uppercase tracking-widest"
              >
                <span className="text-xl group-hover:rotate-90 transition-transform duration-300">+</span>
                New Expense
              </button>
            </div>

            <div className="space-y-6">
              <ExpenseList expenses={expenses} currentUserId={user.id} onSettled={handleSettle} />
            </div>
          </div>

          <div className="lg:col-span-3">
            <SettlementSuggestions
              balances={balances}
              members={members}
              currentUserId={user.id}
            />
          </div>

        </div>
      </div>

      {showAddExpense && (
        <AddExpenseModal
          groupId={groupId}
          members={members}
          onClose={() => setShowAddExpense(false)}
          onAdded={() => { fetchExpenses(); fetchBalances(); }}
        />
      )}
    </div>
  )
}