import ExpenseCard from "./ExpenseCard"

export default function ExpenseList({
  expenses,
  balances,
  currentUserId,
  onSettled
}) {
  if (!expenses.length) {
    return (
      <p className="text-gray-500 text-sm">
        No expenses yet
      </p>
    )
  }

  return (
    <div className="space-y-2">
      {expenses.map((expense) => (
        <ExpenseCard
          key={expense.id}
          expense={expense}
          balances={balances}
          currentUserId={currentUserId}
          onSettled={onSettled}
        />
      ))}
    </div>
  )
}

