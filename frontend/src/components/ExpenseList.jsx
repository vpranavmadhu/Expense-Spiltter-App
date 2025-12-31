import ExpenseCard from "./ExpenseCard"

export default function ExpenseList({ expenses, currentUserId, onSettled }) {
  if (!expenses || expenses.length === 0) {
    return (
      <div className="bg-white rounded-4xl p-16 text-center text-gray-400 border border-dashed border-gray-200">
        <p className="font-bold uppercase text-xs tracking-widest">No expenses recorded yet.</p>
        <p className="text-[10px] mt-2 italic">Click "+ Add Expense" to get started</p>
      </div>
    )
  }

  return (
    <div className="space-y-4">
      {expenses.map(exp => (
        <ExpenseCard key={exp.id} expense={exp} currentUserId={currentUserId} onSettled={onSettled} />
      ))}
    </div>
  )
}