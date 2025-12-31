import { useEffect } from "react";
import { BrowserRouter } from "react-router-dom";
import { Provider, useDispatch, useSelector } from "react-redux";
import { store } from "./store/store";
import { setUser, stopLoading } from "./store/authSlice";
import { Router } from "./Router";
import api from "./api";

function AppContent() {
  const dispatch = useDispatch();
  const { loading } = useSelector((state) => state.auth);

  useEffect(() => {
    api.get("/api/me")
      .then((res) => {
        dispatch(setUser(res.data));
      })
      .catch(() => {
        dispatch(stopLoading());
      });
  }, [dispatch]);

  if (loading) return <div className="p-10 text-center font-bold text-slate-400">LOADING...</div>;

  return <Router />;
}

export default function App() {
  return (
    <Provider store={store}>
      <BrowserRouter>
        <AppContent />
      </BrowserRouter>
    </Provider>
  );
}