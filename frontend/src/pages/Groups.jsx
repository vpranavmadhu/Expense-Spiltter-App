import { useEffect, useState } from "react";
import api from "../api";
import { useNavigate } from "react-router-dom";

export default function Groups() {
  const [groups, setGroups] = useState([]);
  const [name, setName] = useState("");
  const [error, setError] = useState("");
  const [loading, setLoading] = useState(false);
  const navigate = useNavigate();

  const fetchGroups = async () => {
    const res = await api.get("/api/groups");
    setGroups(res.data);
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

  return (
    <div className="min-h-screen bg-gray-100 p-6">
      <div className="max-w-4xl mx-auto">
        {/* HEADER */}
        <div className="mb-8">
          <h1 className="text-3xl font-semibold text-gray-900">
            Your Groups
          </h1>
          <p className="text-gray-600 mt-1">
            Manage and track shared expenses easily
          </p>
        </div>

        {/* CREATE GROUP CARD */}
        <div className="bg-white rounded-xl shadow-sm p-6 mb-8">
          <h2 className="text-lg font-medium text-gray-900 mb-4">
            Create a new group
          </h2>

          <form onSubmit={handleCreateGroup} className="flex gap-4">
            <input
              className="flex-1 border border-gray-300 rounded-lg px-4 py-2
                         focus:outline-none focus:ring-2 focus:ring-blue-500
                         focus:border-blue-500 transition"
              placeholder="e.g. Goa Trip, Roommates, Office Lunch"
              value={name}
              onChange={(e) => setName(e.target.value)}
            />

            <button
              disabled={loading}
              className="bg-blue-600 text-white px-6 py-2 rounded-lg
                         hover:bg-blue-700 transition
                         disabled:opacity-50 disabled:cursor-not-allowed"
            >
              {loading ? "Creating..." : "Create"}
            </button>
          </form>

          {error && (
            <p className="text-red-500 text-sm mt-3">{error}</p>
          )}
        </div>

        {/* GROUP LIST */}
        <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
          {groups.map((g) => (
            <div
              key={g.id}
              className="bg-white rounded-xl shadow-sm p-5
                         hover:shadow-md transition cursor-pointer"
              onClick={() => navigate(`/groups/${g.id}`)}
            >
              <h3 className="text-lg font-medium text-gray-900">
                {g.name}
              </h3>
              <p className="text-sm text-gray-500 mt-1">
                Click to view details
              </p>
            </div>
          ))}
        </div>

        {groups.length === 0 && (
          <p className="text-gray-500 mt-6">
            You haven’t created or joined any groups yet.
          </p>
        )}
      </div>
    </div>
  );
}
