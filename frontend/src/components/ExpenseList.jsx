import ExpenseCard from "./ExpenseCard"

export default function ExpenseList({ expenses, currentUserId, onSettled }) {
  if (expenses == null) {
    return <p className="text-gray-500">No expenses yet</p>
  }

  return (
    <div className="space-y-4">
      {expenses.map(exp => (
        <ExpenseCard
          key={exp.id}
          expense={exp}
          currentUserId={currentUserId}
          onSettled={onSettled}
        />
      ))}
    </div>
  )
}
