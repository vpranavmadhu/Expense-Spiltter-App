import React from 'react';

export default function SettlementSuggestions({ balances, members, currentUserId }) {
  const myBalance = balances[currentUserId] || 0;
  const iOweMoney = myBalance < 0;

  const otherBalances = Object.entries(balances || {})
    .filter(([userId]) => Number(userId) !== currentUserId);

  const getUsername = (id) => {
    const member = members.find(m => m.id === Number(id));
    return member ? member.username : `User ${id}`;
  };

  const handleQuickSettleAll = () => {
    console.log("Triggering bulk settlement for all debts...");
    alert(`Initiating payment for ₹${Math.abs(myBalance).toFixed(2)}. This will clear all your pending shares.`);
  };

  return (
    <div className="bg-white border border-gray-100 rounded-[2.5rem] p-8 shadow-sm sticky top-10">

      <div className={`p-6 rounded-4xl transition-all duration-500 ${iOweMoney ? 'bg-rose-50 border border-rose-100' : 'bg-emerald-50 border border-emerald-100'}`}>
        <h2 className="text-[10px] font-black text-gray-400 uppercase tracking-[0.2em] mb-2">Your Standing</h2>
        <div className={`text-4xl font-black mb-1 ${iOweMoney ? 'text-rose-600' : 'text-emerald-600'}`}>
          ₹{Math.abs(myBalance).toFixed(2)}
        </div>
        <p className="text-[10px] font-bold uppercase tracking-tight text-slate-500">
          {iOweMoney ? "Total amount you need to pay" : "Total amount owed to you"}
        </p>

        {iOweMoney && (
          <button
            onClick={handleQuickSettleAll}
            className="mt-6 w-full bg-slate-900 hover:bg-black text-white py-4 rounded-2xl text-[11px] font-black uppercase tracking-widest shadow-xl shadow-rose-100 transition-all active:scale-95 flex items-center justify-center gap-2"
          >
            ⚡ Quick Settle All
          </button>
        )}
      </div>

      <div className="mt-10">
        <h2 className="text-[10px] font-black text-gray-400 uppercase tracking-[0.2em] mb-6">Member Balances</h2>

        <div className="space-y-6">
          {otherBalances.length === 0 ? (
            <p className="text-xs text-gray-400 italic">No activity yet.</p>
          ) : (
            otherBalances.map(([userId, amount]) => {
              const userIsOwed = amount > 0;
              return (
                <div key={userId} className="flex justify-between items-center group">
                  <div>
                    <p className="text-sm font-extrabold text-slate-800 uppercase tracking-tight">
                      {getUsername(userId)}
                    </p>
                    <p className={`text-[9px] font-black uppercase ${userIsOwed ? 'text-emerald-500' : 'text-rose-400'}`}>
                      {userIsOwed ? 'Owed by group' : 'Owes group'}
                    </p>
                  </div>
                  <div className={`text-lg font-black ${userIsOwed ? 'text-emerald-500' : 'text-slate-300'}`}>
                    ₹{Math.abs(amount).toFixed(2)}
                  </div>
                </div>
              );
            })
          )}
        </div>
      </div>

      <div className="mt-10 pt-6 border-t border-slate-50">
        <div className="flex items-start gap-3 opacity-60">
          <div className="mt-1 w-2 h-2 rounded-full bg-blue-400"></div>
          <p className="text-[10px] text-slate-500 font-bold leading-relaxed italic">
            Settlements are calculated based on your total net share across all expenses in this group.
          </p>
        </div>
      </div>
    </div>
  );
}