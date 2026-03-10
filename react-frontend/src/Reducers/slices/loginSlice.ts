// src/features/userSlice.ts
import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';
import { loginService } from '../../services/loginService';

interface UserState {
  data: any | null; // In a real app, replace 'any' with a User interface
  loading: boolean;
  error: string | null;
}

const initialState: UserState = {
  data: null,
  loading: false,
  error: null,
};
interface LoginCredentials {
  username?: string;
  password?: string;
}
export const loginUser = createAsyncThunk(
  'user/login',
  async (credentials: LoginCredentials, thunkAPI) => {
    try {
      // You need a specific login function, not getUserById
      // Pass the whole object: { username: "...", password: "..." }
      return await loginService(credentials);
    } catch (error: any) {
      return thunkAPI.rejectWithValue(error.message);
    }
  }
);

const loginSlice = createSlice({
  name: 'login', // Updated name to reflect purpose
  initialState,
  reducers: {
    logout: (state) => {
      state.data = null;
      state.loading = false;
      state.error = null;
    },
  },
  extraReducers: (builder) => {
    builder
      // 1. Pending: User clicked login, waiting for response
      .addCase(loginUser.pending, (state) => {
        state.loading = true;
        state.error = null;
      })
      // 2. Fulfilled: Login successful, store user data/token
      .addCase(loginUser.fulfilled, (state, action) => {
        state.loading = false;
        state.data = action.payload; // Store the user object or token
      })
      // 3. Rejected: Login failed (wrong password, server down)
      .addCase(loginUser.rejected, (state, action) => {
        state.loading = false;
        state.error = action.payload as string;
      });
  },
});

export const { logout } = loginSlice.actions;
export default loginSlice.reducer;
