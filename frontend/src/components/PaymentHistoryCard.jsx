import { ArrowUpRight, ArrowDownLeft } from "lucide-react"

export default function PaymentHistoryCard({ item }) {
  const isPaid = item.direction === "paid"

  return (
    <div
      className={`
        flex justify-between items-center
        p-5 rounded-xl border
        shadow-sm hover:shadow-md transition
        ${isPaid ? "bg-red-50 border-red-100" : "bg-green-50 border-green-100"}
      `}
    >
      {/* LEFT */}
      <div className="flex items-center gap-4">
        <div
          className={`
            p-2 rounded-full
            ${isPaid ? "bg-red-100 text-red-600" : "bg-green-100 text-green-600"}
          `}
        >
          {isPaid ? <ArrowUpRight /> : <ArrowDownLeft />}
        </div>

        <div>
          <p className="font-semibold text-gray-900">
            {isPaid ? "You paid" : "You received"}
          </p>

          <p className="text-sm text-gray-600">
            {item.fromUser} ({item.fromEmail})
          </p>

          <p className="text-xs text-gray-500 mt-1">
            Group: {item.groupName}
          </p>
        </div>
      </div>

      {/* RIGHT */}
      <div className="text-right">
        <p
          className={`text-lg font-bold ${
            isPaid ? "text-red-600" : "text-green-600"
          }`}
        >
          ₹{item.amount}
        </p>

        <p className="text-xs text-gray-500 mt-1">
          {new Date(item.createdAt).toLocaleDateString()} •{" "}
          {new Date(item.createdAt).toLocaleTimeString()}
        </p>
      </div>
    </div>
  )
}
