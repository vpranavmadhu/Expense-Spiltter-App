import React from 'react';
import { Navigate, Route, Routes } from 'react-router-dom';
import { useSelector, useDispatch } from 'react-redux';
import { logout } from './store/authSlice';
import Register from './pages/Register';
import Login from './pages/Login';
import Groups from './pages/Groups';
import Navbar from './components/Navbar';
import GroupDetail from './pages/GroupDetail';
import PaymentHistory from './pages/PaymentHistory';
import api from './api';

export const Router = () => {
  const dispatch = useDispatch();
  const { user } = useSelector((state) => state.auth);
  const handleLogout = async () => {
    try {
      await api.post("/auth/logout");
    } catch (err) {
      console.error("Logout API failed", err);
    } finally {
      dispatch(logout());
    }
  };

  return (
    <Routes>
      <Route
        path="/login"
        element={user ? <Navigate to="/groups" /> : <Login />}
      />

      <Route
        path="/register"
        element={user ? <Navigate to="/groups" /> : <Register />}
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

      <Route
        path="/groups/:groupId"
        element={
          user ? (
            <>
              <Navbar user={user} onLogout={handleLogout} />
              <GroupDetail />
            </>
          ) : (
            <Navigate to="/login" />
          )
        }
      />

      <Route
        path="/paymenthistory"
        element={
          user ? (
            <>
              <Navbar user={user} onLogout={handleLogout} />
              <PaymentHistory />
            </>
          ) : (
            <Navigate to="/login" />
          )
        }
      />
      
      <Route path="*" element={<Navigate to="/groups" />} />
    </Routes>
  );
};