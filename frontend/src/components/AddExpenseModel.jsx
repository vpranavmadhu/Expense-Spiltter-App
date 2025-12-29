import { useState } from "react"
import axios from "axios"

export default function AddExpenseModal({
    groupId,
    members,
    onClose,
    onAdded,
}) {
    const [title, setTitle] = useState("")
    const [amount, setAmount] = useState("")
    const [splitType, setSplitType] = useState("equal")
    const [splits, setSplits] = useState({})
    const [loading, setLoading] = useState(false)

    const handleSplitChange = (user_id, value) => {
        setSplits({
            ...splits,
            [user_id]: Number(value),
        })
    }

    const submit = async () => {
        if (!title || !amount) {
            alert("Title and amount required")
            return
        }

        let splitPayload = []

        if (splitType === "custom") {
            const total = Object.values(splits).reduce((a, b) => a + b, 0)

            if (total !== Number(amount)) {
                alert("Split amounts must equal total amount")
                return
            }

            splitPayload = Object.entries(splits).map(([user_id, amt]) => ({
                user_id: Number(user_id),
                amount: amt,
            }))
        }

        try {
            setLoading(true)
            await axios.post(
                "http://localhost:8080/api/createexpense",
                {
                    group_id: Number(groupId),
                    title,
                    amount: Number(amount),
                    splits: splitType === "custom" ? splitPayload : [],
                },
                { withCredentials: true }
            )

            onAdded()
            onClose()
        } catch (err) {
            alert(err.response?.data?.error || "Failed to add expense")
        } finally {
            setLoading(false)
        }
    }

    return (
        <div className="fixed inset-0 bg-black/40 flex items-center justify-center z-50">
            <div className="bg-white rounded-lg w-full max-w-md p-6">
                <h2 className="text-lg font-semibold mb-4">Add Expense</h2>

                <input
                    type="text"
                    placeholder="Expense title"
                    value={title}
                    onChange={(e) => setTitle(e.target.value)}
                    className="w-full border rounded px-3 py-2 mb-3"
                />

                <input
                    type="number"
                    placeholder="Amount"
                    value={amount}
                    onChange={(e) => setAmount(e.target.value)}
                    className="w-full border rounded px-3 py-2 mb-4"
                />

                {/* Split type */}
                <div className="mb-4">
                    <label className="mr-4">
                        <input
                            type="radio"
                            value="equal"
                            checked={splitType === "equal"}
                            onChange={() => setSplitType("equal")}
                        />{" "}
                        Equal
                    </label>

                    <label>
                        <input
                            type="radio"
                            value="custom"
                            checked={splitType === "custom"}
                            onChange={() => setSplitType("custom")}
                        />{" "}
                        Custom
                    </label>
                </div>

                {/* Custom splits */}
                {splitType === "custom" && (
                    <div className="space-y-2 mb-4">
                        {members.map((m) => (
                            <div
                                key={m.id}
                                className="flex justify-between items-center"
                            >
                                <span>{m.username}</span>
                                <input
                                    type="number"
                                    className="border rounded px-2 py-1 w-24"
                                    onChange={(e) =>
                                        handleSplitChange(m.id, e.target.value)
                                    }
                                />
                            </div>
                        ))}
                    </div>
                )}

                <div className="flex justify-end gap-2">
                    <button onClick={onClose} className="border px-4 py-2 rounded">
                        Cancel
                    </button>
                    <button
                        onClick={submit}
                        disabled={loading}
                        className="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700 disabled:opacity-50"
                    >
                        Add
                    </button>
                </div>
            </div>
        </div>
    )
}
