import { useEffect, useState } from "react"
import axios from "axios"
import { ArrowUpRight, ArrowDownLeft, Wallet } from "lucide-react"
import PaymentHistoryCard from "../components/PaymentHistoryCard"

export default function PaymentHistory() {
  const [history, setHistory] = useState([])
  const [loading, setLoading] = useState(true)

  const fetchHistory = async () => {
    try {
      console.log("entered api callig");
      
      const res = await axios.get(
        "http://localhost:8080/api/payments/history",
        { withCredentials: true }
      )
      setHistory(res.data)
    } catch (err) {
      console.error("Failed to load history", err)
    } finally {
      setLoading(false)
    }
  }
   console.log("history",history);

  useEffect(() => {
    fetchHistory()
  }, [])

  if (loading) {
    return (
      <div className="flex justify-center items-center h-screen text-gray-500">
        Loading payment history...
      </div>
    )
  }

  if (history == null) {
    return (
      <div className="flex flex-col items-center justify-center h-screen text-gray-400">
        <Wallet className="w-10 h-10 mb-3" />
        <p>No payment history yet</p>
      </div>
    )
  }  

  return (
    <div className="max-w-4xl mx-auto p-6 space-y-5">
      <h1 className="text-3xl font-bold text-gray-900">
        Payment History
      </h1>

      {history.map((item, idx) => (
        <PaymentHistoryCard key={idx} item={item} />
      ))}
    </div>
  )
}
