import React from 'react'
import { Navigate, Route, Routes } from 'react-router-dom'
import Register from './pages/Register'
import Login from './pages/Login'
import Groups from './pages/Groups'
import Navbar from './components/Navbar'
import GroupDetail from './pages/GroupDetail'

export const Router = ({ user, setUser }) => {
  const handleLogout = () => {
    setUser(null);
  };

  return (
    <Routes>
      <Route
        path="/login"
        element={user ? <Navigate to="/groups" /> : <Login setUser={setUser} />}
      />

      <Route
        path="/register"
        element={
          user ? <Navigate to="/groups" /> : <Register setUser={setUser} />
        }
      />

      <Route
        path="/groups"
        element={
          user ? (
            <>
              <Navbar user={user} onLogout={handleLogout} />
              <Groups />
            </>
          ) : (
            <Navigate to="/login" />
          )
        }
      />
    </Routes>
  );
}
