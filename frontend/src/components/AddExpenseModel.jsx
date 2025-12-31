import { useState, useEffect } from "react"
import axios from "axios"

export default function AddExpenseModal({ groupId, members = [], onClose, onAdded }) {
    
    const safeMembers = members || []
    const [title, setTitle] = useState("")
    const [amount, setAmount] = useState("")
    const [splitType, setSplitType] = useState("equal")
    const [splits, setSplits] = useState({})
    const [loading, setLoading] = useState(false)
    const [selectedIds, setSelectedIds] = useState(() => {
        return safeMembers.map(m => m.id)
    })

    useEffect(() => {
        if (safeMembers.length > 0) {
            setSelectedIds(prev => {
                if (prev.length === 0) return safeMembers.map(m => m.id)
                return prev
            })
        }
    }, [safeMembers.length])

    const activeSplitsSum = Object.entries(splits)
        .filter(([id]) => selectedIds.includes(Number(id)))
        .reduce((sum, [_, val]) => sum + (Number(val) || 0), 0);

    const toggleMember = (id) => {
        if (selectedIds.includes(id)) {
            if (selectedIds.length > 1) setSelectedIds(selectedIds.filter(mid => mid !== id))
        } else {
            setSelectedIds([...selectedIds, id])
        }
    }

    const selectAll = () => setSelectedIds(safeMembers.map(m => m.id))

    const handleSplitChange = (user_id, value) => {
        const val = Math.max(0, Number(value));
        setSplits({ ...splits, [user_id]: val })
    }

    const submit = async () => {
        if (!title || !amount) return alert("Title and amount required")

        let splitPayload = []
        const totalAmount = Number(amount);

        if (splitType === "custom") {
            if (Math.abs(activeSplitsSum - totalAmount) > 0.01) {
                return alert(`Total split must be exactly ₹${totalAmount}. Current: ₹${activeSplitsSum}`)
            }
            splitPayload = selectedIds.map(id => ({
                user_id: id,
                amount: splits[id] || 0
            }))
        } else {
            const share = totalAmount / selectedIds.length
            splitPayload = selectedIds.map(id => ({
                user_id: id,
                amount: Number(share.toFixed(2))
            }))
        }

        try {
            setLoading(true)
            await axios.post("http://localhost:8080/api/createexpense", {
                group_id: Number(groupId),
                title,
                amount: totalAmount,
                splits: splitPayload,
            }, { withCredentials: true })
            onAdded(); onClose();
        } catch (err) {
            alert(err.response?.data?.error || "Failed")
        } finally {
            setLoading(false)
        }
    }

    return (
        <div className="fixed inset-0 bg-slate-900/40 backdrop-blur-sm flex items-center justify-center z-100 p-4">
            <div className="bg-white rounded-[2.5rem] w-full max-w-lg p-8 shadow-2xl overflow-hidden border border-white/20 animate-in fade-in zoom-in duration-200">

                <div className="flex justify-between items-center mb-6">
                    <div>
                        <h2 className="text-xl font-black text-slate-900">Add Expense</h2>
                        <p className="text-[10px] font-bold text-slate-400 uppercase tracking-widest">New Transaction</p>
                    </div>
                    <button onClick={onClose} className="w-8 h-8 flex items-center justify-center rounded-full bg-slate-50 text-slate-400 hover:text-rose-500 transition-colors">✕</button>
                </div>

                <div className="space-y-5">
                    <div className="grid grid-cols-2 gap-4">
                        <div className="space-y-1">
                            <label className="text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1">Title</label>
                            <input
                                type="text"
                                placeholder="Dinner?"
                                value={title}
                                onChange={(e) => setTitle(e.target.value)}
                                className="w-full bg-slate-50 border-none rounded-xl px-4 py-3 text-sm font-bold text-slate-800 focus:ring-2 focus:ring-purple-500 outline-none"
                            />
                        </div>
                        <div className="space-y-1">
                            <label className="text-[10px] font-black text-slate-400 uppercase tracking-widest ml-1">Amount</label>
                            <input
                                type="number"
                                min="0"
                                placeholder="0.00"
                                value={amount}
                                onChange={(e) => setAmount(e.target.value)}
                                className="w-full bg-slate-50 border-none rounded-xl px-4 py-3 text-sm font-black text-slate-900 focus:ring-2 focus:ring-purple-500 outline-none [appearance:textfield] [&::-webkit-outer-spin-button]:appearance-none [&::-webkit-inner-spin-button]:appearance-none"
                            />
                        </div>
                    </div>

                    <div className="space-y-2">
                        <div className="flex justify-between items-end px-1">
                            <label className="text-[10px] font-black text-slate-400 uppercase tracking-widest">Split with</label>
                            <button onClick={selectAll} className="text-[9px] font-black text-purple-600 hover:underline">SELECT ALL</button>
                        </div>
                        <div className="flex flex-wrap gap-1.5 p-1 max-h-20 overflow-y-auto custom-scrollbar">
                            {safeMembers.length === 0 && <p className="text-xs text-slate-300 italic">Loading members...</p>}

                            {safeMembers.map(m => {
                                const active = selectedIds.includes(m.id)
                                return (
                                    <button
                                        key={m.id}
                                        onClick={() => toggleMember(m.id)}
                                        className={`px-3 py-1.5 rounded-lg text-[11px] font-bold transition-all border ${active ? 'bg-slate-900 border-slate-900 text-white shadow-md shadow-slate-200' : 'bg-white border-slate-100 text-slate-400'
                                            }`}
                                    >
                                        {m.username}
                                    </button>
                                )
                            })}
                        </div>
                    </div>

                    <div className="bg-slate-50 rounded-2xl p-4 space-y-4 border border-slate-100/50">
                        <div className="flex bg-white/50 p-1 rounded-xl border border-slate-100">
                            {['equal', 'custom'].map(type => (
                                <button
                                    key={type}
                                    onClick={() => setSplitType(type)}
                                    className={`flex-1 py-1.5 text-[10px] font-black uppercase tracking-widest rounded-lg transition-all ${splitType === type ? 'bg-white shadow-sm text-purple-600' : 'text-slate-400'}`}
                                >
                                    {type}
                                </button>
                            ))}
                        </div>

                        <div className="max-h-32 overflow-y-auto pr-1 custom-scrollbar">
                            {splitType === "equal" ? (
                                <div className="py-2 flex justify-between items-center">
                                    <span className="text-xs font-bold text-slate-500 italic">Each of {selectedIds.length} pays</span>
                                    <span className="text-xl font-black text-purple-600">₹{amount ? (Number(amount) / Math.max(1, selectedIds.length)).toFixed(2) : "0"}</span>
                                </div>
                            ) : (
                                <div className="space-y-2">
                                    {safeMembers.filter(m => selectedIds.includes(m.id)).map((m) => (
                                        <div key={m.id} className="flex justify-between items-center bg-white p-2 rounded-xl border border-slate-100">
                                            <span className="text-[11px] font-bold text-slate-600 ml-1">{m.username}</span>
                                            <div className="relative">
                                                <span className="absolute left-3 top-1.5 text-[10px] font-bold text-slate-300">₹</span>
                                                <input
                                                    type="number"
                                                    min="0"
                                                    value={splits[m.id] || ""}
                                                    placeholder="0"
                                                    className="bg-slate-50 border-none rounded-lg px-6 py-1 text-xs font-black w-24 text-right focus:ring-1 focus:ring-purple-500 outline-none [appearance:textfield] [&::-webkit-outer-spin-button]:appearance-none [&::-webkit-inner-spin-button]:appearance-none"
                                                    onChange={(e) => handleSplitChange(m.id, e.target.value)}
                                                />
                                            </div>
                                        </div>
                                    ))}
                                    <div className="pt-2 flex justify-between px-2 text-[10px] font-bold uppercase">
                                        <span className="text-slate-400">Total Split: ₹{activeSplitsSum}</span>
                                        <span className={Math.abs(activeSplitsSum - Number(amount)) < 0.01 ? "text-emerald-500" : "text-rose-500"}>
                                            Remaining: ₹{(Number(amount) - activeSplitsSum).toFixed(2)}
                                        </span>
                                    </div>
                                </div>
                            )}
                        </div>
                    </div>
                </div>

                <div className="flex gap-3 mt-8">
                    <button onClick={onClose} className="flex-1 py-3 text-[11px] font-black text-slate-400 uppercase tracking-widest">Cancel</button>
                    <button
                        onClick={submit}
                        disabled={loading}
                        className="flex-2 bg-[#7c3aed] hover:bg-[#6d28d9] text-white py-3 rounded-xl font-black text-[11px] uppercase tracking-widest shadow-lg shadow-purple-100 active:scale-95 disabled:opacity-50 transition-all"
                    >
                        {loading ? "Adding..." : "Confirm Expense"}
                    </button>
                </div>
            </div>
        </div>
    )
}