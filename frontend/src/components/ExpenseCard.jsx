export default function ExpenseCard({ expense, currentUserId, onSettled }) {
  const round2 = (value) => Number(value.toFixed(2))
  const isPayer = expense.paidById === currentUserId
  const myShare = round2(expense.myShare) || 0

  const balanceAmount = isPayer && expense.totalOwedToMe > 0
      ? round2(Math.max(expense.totalOwedToMe - myShare, 0))
      : 0

  let label = "", color = "", accent = "from-gray-300 to-gray-200", showPayButton = false

  if (isPayer) {
    label = `You paid ₹${expense.amount}`; color = "text-emerald-500"; accent = "from-emerald-400 to-emerald-200"
  } else if (expense.isSettled) {
    label = "Settled"; color = "text-gray-400"
  } else if (myShare > 0) {
    label = `You owe ₹${myShare}`; color = "text-rose-500"; accent = "from-rose-400 to-rose-200"; showPayButton = true
  }

  return (
    <div className="relative group">
      <div className={`absolute left-0 top-0 h-full w-1.5 rounded-l-3xl bg-linear-to-b ${accent}`} />
      <div className="flex justify-between items-center bg-white border border-gray-100 shadow-sm hover:shadow-md transition-all rounded-3xl p-6 pl-8">
        <div>
          <h3 className="font-extrabold text-xl text-slate-800 group-hover:text-purple-600 transition-colors">{expense.title}</h3>
          <p className="text-sm text-gray-400 mt-0.5">Paid by {expense.paidByName}</p>
          <div className="mt-4 flex items-center gap-3">
            <span className="text-sm font-bold text-slate-700">Total ₹{expense.amount}</span>
            <span className="bg-slate-100 text-[10px] font-black uppercase text-gray-500 px-2 py-0.5 rounded">👥 {expense.splitCount} members</span>
          </div>
        </div>

        <div className="text-right flex flex-col items-end">
          <p className={`font-black text-lg ${color}`}>{label}</p>
          
          {/* REQUESTED SUB-INFO PRESERVED */}
          {isPayer && balanceAmount > 0 && (
            <p className="text-[11px] font-bold text-emerald-600 mt-1 bg-emerald-50 px-2 py-0.5 rounded-lg animate-pulse">
              You’ll receive <span className="underline decoration-2 underline-offset-2">₹{balanceAmount}</span>
            </p>
          )}

          {showPayButton && (
            <button
              onClick={() => onSettled(expense)}
              className="mt-4 text-[10px] font-black uppercase tracking-wider bg-blue-600 text-white px-5 py-2.5 rounded-xl hover:bg-blue-700 shadow-lg shadow-blue-100 transition-all active:scale-95"
            >
              Mark as paid
            </button>
          )}
        </div>
      </div>
    </div>
  )
}