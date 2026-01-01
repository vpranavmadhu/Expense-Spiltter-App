import { useEffect, useState } from "react";
import api from "../api";
import { useNavigate } from "react-router-dom";
import { Trash2, Users, Calendar, User } from "lucide-react";
import { toast } from "react-toastify";

export default function Groups() {
  const [groups, setGroups] = useState([]);
  const [name, setName] = useState("");
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  const fetchGroups = async () => {
    try {
      const res = await api.get("/api/groups");
      setGroups(res.data || []);
    } catch (err) {
      console.error("Failed to fetch groups");
    }
  };

  useEffect(() => {
    fetchGroups();
  }, []);

  const handleCreateGroup = async (e) => {
    e.preventDefault();
    setError("");

    if (!name.trim()) {
      setError("Group name is required");
      return;
    }

    try {
      setLoading(true);
      await api.post("/api/creategroup", { name });
      setName("");
      fetchGroups();
    } catch {
      setError("Failed to create group");
    } finally {
      setLoading(false);
    }
  };

  const handleDeleteGroup = async (e, groupId) => {
    e.stopPropagation();
    try {
      await api.delete(`/api/groups/${groupId}`);
      fetchGroups();
    } catch (err) {
      const errorMsg = err.response?.data?.error;

      if (errorMsg === "unauthorized") {
        toast(`Oops.. U didn't created this account`);
      } else {
        toast("Failed: internal error");
      }

    }
  };

  return (
    <div className="min-h-screen bg-slate-200 p-6 lg:p-12">
      <div className="max-w-5xl mx-auto">
        <div className="mb-12">
          <h1 className="text-5xl font-black text-slate-900 tracking-tight mb-2">
            Your Workspace
          </h1>
          <p className="text-slate-400 font-bold uppercase tracking-widest text-xs">
            Manage and track shared expenses easily
          </p>
        </div>

        <div className="bg-white rounded-4xl border border-slate-100 shadow-sm p-8 mb-12">
          <h2 className="text-xl font-black text-slate-900 mb-6">
            Start a new group
          </h2>

          <form onSubmit={handleCreateGroup} className="flex flex-col md:flex-row gap-4">
            <input
              className="flex-1 bg-slate-50 border-none rounded-2xl px-6 py-4 font-bold text-slate-700
                         focus:outline-none focus:ring-2 focus:ring-purple-500 transition placeholder:text-slate-400"
              placeholder="e.g. Goa Trip, Roommates, Office Lunch"
              value={name}
              onChange={(e) => setName(e.target.value)}
            />

            <button
              disabled={loading}
              className="bg-[#7c3aed] text-white px-8 py-4 rounded-2xl font-black text-xs uppercase tracking-widest
                         hover:bg-[#6d28d9] shadow-lg shadow-purple-100 transition-all active:scale-95
                         disabled:opacity-50 disabled:cursor-not-allowed"
            >
              {loading ? "Creating..." : "Create Group"}
            </button>
          </form>

          {error && (
            <p className="text-rose-500 text-xs font-bold mt-4 bg-rose-50 inline-block px-3 py-1 rounded-lg">
              {error}
            </p>
          )}
        </div>

        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          {groups.map((g) => (
            <div
              key={g.id}
              className="group bg-white rounded-4xl border border-slate-100 shadow-sm p-8
                         hover:shadow-xl hover:-translate-y-1 transition-all duration-300 cursor-pointer relative"
              onClick={() => navigate(`/groups/${g.id}`)}
            >
              <div className="flex justify-between items-start mb-4">
                <h3 className="text-2xl font-black text-slate-800 group-hover:text-purple-600 transition-colors">
                  {g.name}
                </h3>

                <button
                  onClick={(e) => handleDeleteGroup(e, g.id)}
                  className="p-2 rounded-xl text-slate-300 hover:bg-rose-50 hover:text-rose-500 transition-colors"
                  title="Delete Group"
                >
                  <Trash2 className="w-5 h-5" />
                </button>
              </div>

              <div className="flex flex-wrap gap-3 mt-6">
                <div className="flex items-center gap-1.5 bg-slate-50 px-3 py-1.5 rounded-lg border border-slate-100">
                  <User className="w-3 h-3 text-slate-400" />
                  <span className="text-[10px] font-bold text-slate-500 uppercase tracking-wide">
                    {g.creator_name || "You"}
                  </span>
                </div>

                <div className="flex items-center gap-1.5 bg-slate-50 px-3 py-1.5 rounded-lg border border-slate-100">
                  <Calendar className="w-3 h-3 text-slate-400" />
                  <span className="text-[10px] font-bold text-slate-500 uppercase tracking-wide">
                    {g.created_at ? new Date(g.created_at).toLocaleDateString() : "Recently"}
                  </span>
                </div>

                <div className="flex items-center gap-1.5 bg-purple-50 px-3 py-1.5 rounded-lg border border-purple-100">
                  <Users className="w-3 h-3 text-purple-500" />
                  <span className="text-[10px] font-black text-purple-600 uppercase tracking-wide">
                    {g.members ? g.members.length : 1} Members
                  </span>
                </div>
              </div>

              <p className="text-[10px] font-bold text-slate-300 uppercase tracking-widest mt-6 flex items-center gap-1">
                View Details &rarr;
              </p>
            </div>
          ))}
        </div>

        {groups.length === 0 && (
          <div className="text-center py-16 opacity-50">
            <p className="text-slate-600 font-bold uppercase tracking-widest text-sm">
              No groups yet. Create one to get started.
            </p>
          </div>
        )}
      </div>
    </div>
  );
}