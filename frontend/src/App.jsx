import { BrowserRouter} from "react-router-dom";
import { useEffect, useState } from "react";
import api from "./api";
import { Router } from "./Router";

function App() {
  const [user, setUser] = useState(null);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    api
      .get("/api/me")
      .then((res) => setUser(res.data))
      .catch(() => setUser(null))
      .finally(() => setLoading(false));
  }, []);

  if (loading) return <div className="p-4">Loading...</div>;

  return (
    <BrowserRouter>
      <Router user={user} setUser={setUser} />
    </BrowserRouter>
  );
}

export default App;
