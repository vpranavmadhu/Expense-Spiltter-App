export default function ExpenseCard({
  expense,
  currentUserId,
  onSettled
}) {
  const isPayer = expense.paidById === currentUserId
  const myShare = expense.myShare || 0

  let label = ""
  let color = ""
  let bg = "bg-gray-50"
  let showPayButton = false

  if (isPayer) {
    label = `You paid ₹${expense.amount}`
    color = "text-green-600"
    bg = "bg-green-50"
  } else if (myShare > 0) {
    label = `You owe ₹${myShare}`
    color = "text-red-600"
    bg = "bg-red-50"
    showPayButton = true
  } else {
    label = "Settled"
    color = "text-gray-500"
    bg = "bg-gray-50"
  }

  return (
    <div
      className={`
        flex justify-between items-center
        ${bg} border border-gray-100
        shadow-sm hover:shadow-md transition
        rounded-lg p-6
      `}
    >
      <div>
        <h3 className="font-semibold text-lg text-gray-900">
          {expense.title}
        </h3>

        <p className="text-sm text-gray-500 mt-1">
          Paid by {expense.paidByName}
        </p>

        <p className="text-sm text-gray-700 mt-2">
          Total: <span className="font-medium">₹{expense.amount}</span>
        </p>
      </div>

      <div className="text-right">
        <p className={`font-semibold ${color}`}>
          {label}
        </p>

        {showPayButton && (
          <button
            onClick={() => onSettled(expense)}
            className="
              mt-3 text-sm font-medium
              bg-blue-600 text-white px-4 py-1.5
              rounded hover:bg-blue-700 transition
            "
          >
            Mark as paid
          </button>
        )}
      </div>
    </div>
  )
}
