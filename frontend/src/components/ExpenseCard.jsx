export default function ExpenseCard({ expense, currentUserId, onSettled }) {


  const round2 = (value) => Number(value.toFixed(2))

  const isPayer = expense.paidById === currentUserId
  const myShare = round2(expense.myShare) || 0

  const balanceAmount =
    isPayer && expense.totalOwedToMe > 0
      ? round2(Math.max(expense.totalOwedToMe - myShare, 0))
      : 0

  let label = ""
  let color = ""
  let accent = "from-gray-300 to-gray-200"
  let showPayButton = false

  if (isPayer) {
    label = `You paid ₹${expense.amount}`
    color = "text-green-600"
    accent = "from-green-400 to-green-200"
  } else if (expense.isSettled) {
    label = "Settled"
    color = "text-gray-500"
  } else if (myShare > 0) {
    label = `You owe ₹${myShare}`
    color = "text-red-600"
    accent = "from-red-400 to-red-200"
    showPayButton = true
  }

  return (
    <div className="relative">
      <div
        className={`absolute left-0 top-0 h-full w-1.5 rounded-l-xl bg-linear-to-b ${accent}`}
      />

      <div
        className="
          flex justify-between items-center
          bg-white/80 backdrop-blur
          border border-gray-100
          shadow-sm hover:shadow-lg
          transition-all duration-300
          rounded-xl p-6 pl-8
        "
      >
        <div>
          <h3 className="font-semibold text-lg text-gray-900">
            {expense.title}
          </h3>

          <p className="text-sm text-gray-500 mt-1">
            Paid by {expense.paidByName}
          </p>

          <p className="text-sm text-gray-700 mt-2">
            Total{" "}
            <span className="font-semibold text-gray-900">
              ₹{expense.amount}
            </span>
          </p>

          <div className="mt-4">
            <span className="
              inline-flex items-center gap-1
              px-3 py-1 text-xs font-medium
              bg-gray-100 text-gray-700
              rounded-full
            ">
              👥 {expense.splitCount} members
            </span>
          </div>
        </div>

        <div className="text-right">
          <p className={`font-semibold text-base ${color}`}>
            {label}
          </p>

          {isPayer && balanceAmount > 0 && (
            <p className="text-sm text-green-700 mt-1">
              You’ll receive{" "}
              <span className="font-semibold">
                ₹{balanceAmount}
              </span>
            </p>
          )}

          {showPayButton && (
            <button
              onClick={() => onSettled(expense)}
              className="
                mt-4 text-sm font-medium
                bg-blue-600 text-white
                px-5 py-2
                rounded-full
                hover:bg-blue-700
                shadow-sm hover:shadow-md
                transition
              "
            >
              Mark as paid
            </button>
          )}
        </div>
      </div>
    </div>
  )
}
