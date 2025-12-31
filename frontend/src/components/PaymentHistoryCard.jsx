import { ArrowUpRight, ArrowDownLeft } from "lucide-react"

export default function PaymentHistoryCard({ item }) {
  const isPaid = item.direction === "paid"

  return (
    <div
      className={`
        flex justify-between items-center
        p-6 rounded-3xl border
        shadow-sm hover:shadow-lg transition-all duration-300
        ${isPaid ? "bg-rose-50 border-rose-100" : "bg-emerald-50 border-emerald-100"}
      `}
    >
      <div className="flex items-center gap-5">
        <div
          className={`
            w-12 h-12 flex items-center justify-center rounded-2xl
            ${isPaid ? "bg-white text-rose-500 shadow-sm" : "bg-white text-emerald-500 shadow-sm"}
          `}
        >
          {isPaid ? <ArrowUpRight className="w-6 h-6" /> : <ArrowDownLeft className="w-6 h-6" />}
        </div>

        <div>
          <p className="font-black text-slate-900 text-lg">
            {isPaid ? "You Paid" : "You Received"}
          </p>

          <p className="text-sm font-bold text-slate-600">
            {item.fromUser} <span className="text-slate-400 font-medium text-xs">({item.fromEmail})</span>
          </p>

          <p className="text-[10px] font-black text-slate-400 uppercase tracking-widest mt-1">
            Group: {item.groupName}
          </p>
        </div>
      </div>

      <div className="text-right">
        <p
          className={`text-2xl font-black ${
            isPaid ? "text-rose-600" : "text-emerald-600"
          }`}
        >
          ₹{item.amount}
        </p>

        <p className="text-[10px] font-bold text-slate-400 uppercase tracking-wider mt-1">
          {new Date(item.createdAt).toLocaleDateString()} •{" "}
          {new Date(item.createdAt).toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })}
        </p>
      </div>
    </div>
  )
}