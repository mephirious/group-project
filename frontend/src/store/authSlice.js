import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';
import { BASE_URL } from '../utils/apiURL';

const initialState = {
    loginStatus: 'idle',
    loginError: null,
    registerStatus: 'idle',
    registerError: null,
    authStatus: 'idle',
    authError: null,
    user: null,
};

export const loginAsync = createAsyncThunk('auth/login', async ({ email, password }) => {
    const response = await fetch(`${BASE_URL}auth/api/v1/login`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, password }),
        credentials: 'include',
    });
    if (!response.ok) {
        throw new Error('Login failed');
    }
    return await response.json();
});

export const registerAsync = createAsyncThunk('auth/register', async ({ email, password, confirmPassword }) => {
    const response = await fetch(`${BASE_URL}auth/api/v1/register`, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({ email, password, confirmPassword }),
        credentials: 'include',
    });
    if (!response.ok) {
        throw new Error('Registration failed');
    }
    return await response.json();
});

export const validateTokenAsync = createAsyncThunk('auth/validateToken', async () => {
    const response = await fetch(`${BASE_URL}auth/api/v1/validate-token`, {
        method: 'GET',
        credentials: 'include',
    });
    if (!response.ok) {
        throw new Error('Token validation failed');
    }
    return await response.json();
});

export const refreshTokenAsync = createAsyncThunk('auth/refreshToken', async () => {
    const response = await fetch(`${BASE_URL}auth/api/v1/refresh`, {
        method: 'GET',
        credentials: 'include',
    });
    if (!response.ok) {
        throw new Error('Token refresh failed');
    }
    return await response.json();
});

export const logoutAsync = createAsyncThunk('auth/logout', async () => {
    const response = await fetch(`${BASE_URL}auth/api/v1/logout`, {
        method: 'GET',
        credentials: 'include',
    });
    if (!response.ok) {
        throw new Error('Logout failed');
    }
    return await response.json();
});

export const verifyAuth = createAsyncThunk('auth/verifyAuth', async (_, { dispatch, rejectWithValue }) => {
    try {
        const validated = await dispatch(validateTokenAsync()).unwrap();
        return validated;
    } catch (error) {
        try {
            await dispatch(refreshTokenAsync()).unwrap();
            const validatedAfterRefresh = await dispatch(validateTokenAsync()).unwrap();
            return validatedAfterRefresh;
        } catch (refreshError) {
            return rejectWithValue(refreshError);
        }
    }
});

const authSlice = createSlice({
    name: 'auth',
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder
            .addCase(loginAsync.pending, (state) => {
                state.loginStatus = 'loading';
                state.loginError = null;
            })
            .addCase(loginAsync.fulfilled, (state, action) => {
                state.loginStatus = 'succeeded';
                state.user = action.payload;
            })
            .addCase(loginAsync.rejected, (state, action) => {
                state.loginStatus = 'failed';
                state.loginError = action.error.message;
            })
            .addCase(registerAsync.pending, (state) => {
                state.registerStatus = 'loading';
                state.registerError = null;
            })
            .addCase(registerAsync.fulfilled, (state) => {
                state.registerStatus = 'succeeded';
            })
            .addCase(registerAsync.rejected, (state, action) => {
                state.registerStatus = 'failed';
                state.registerError = action.error.message;
            })
            .addCase(validateTokenAsync.pending, (state) => {
                state.authStatus = 'loading';
                state.authError = null;
            })
            .addCase(validateTokenAsync.fulfilled, (state, action) => {
                state.authStatus = 'succeeded';
                state.user = action.payload;
            })
            .addCase(validateTokenAsync.rejected, (state, action) => {
                state.authStatus = 'failed';
                state.authError = action.error.message;
            })
            .addCase(refreshTokenAsync.pending, (state) => {
                state.authStatus = 'loading';
                state.authError = null;
            })
            .addCase(refreshTokenAsync.fulfilled, (state) => {
                state.authStatus = 'succeeded';
            })
            .addCase(refreshTokenAsync.rejected, (state, action) => {
                state.authStatus = 'failed';
                state.authError = action.error.message;
            })
            .addCase(logoutAsync.pending, (state) => {
                state.authStatus = 'loading';
                state.authError = null;
            })
            .addCase(logoutAsync.fulfilled, (state) => {
                state.authStatus = 'idle';
                state.user = null;
            })
            .addCase(logoutAsync.rejected, (state, action) => {
                state.authStatus = 'failed';
                state.authError = action.error.message;
            })
            .addCase(verifyAuth.pending, (state) => {
                state.authStatus = 'loading';
                state.authError = null;
            })
            .addCase(verifyAuth.fulfilled, (state, action) => {
                state.authStatus = 'succeeded';
                state.user = action.payload;
            })
            .addCase(verifyAuth.rejected, (state, action) => {
                state.authStatus = 'failed';
                state.authError = action.payload || action.error.message;
                state.user = null;
            });
    },
});

export const selectLoginStatus = (state) => state.auth.loginStatus;
export const selectLoginError = (state) => state.auth.loginError;
export const selectRegisterStatus = (state) => state.auth.registerStatus;
export const selectRegisterError = (state) => state.auth.registerError;
export const selectAuthStatus = (state) => state.auth.authStatus;
export const selectAuthError = (state) => state.auth.authError;

export default authSlice.reducer;
