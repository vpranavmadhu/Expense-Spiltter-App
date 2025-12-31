import { createSlice } from "@reduxjs/toolkit";

const authSlice = createSlice({
  name: "auth",
  initialState: {
    user: null,
    loading: true,
  },
  reducers: {
    setUser: (state, action) => {
      state.user = action.payload;
      state.loading = false;
    },
    logout: (state) => {
      state.user = null;
      state.loading = false;
    },
    stopLoading: (state) => {
      state.loading = false;
    }
  },
});

export const { setUser, logout, stopLoading } = authSlice.actions;
export default authSlice.reducer;