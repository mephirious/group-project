import { createAsyncThunk, createSlice } from '@reduxjs/toolkit';
import { BASE_URL } from '../utils/apiURL';
import { STATUS } from '../utils/status';

export const createBrand = createAsyncThunk('admin/createBrand', async(brandData, { rejectWithValue }) => {
  const response = await fetch(`${BASE_URL}products/brands`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    credentials: 'include',
    body: JSON.stringify(brandData)
  });
  if (!response.ok) return rejectWithValue(await response.json());
  return response.json();
});

export const updateBrand = createAsyncThunk('admin/updateBrand', async({ id, brandData }, { rejectWithValue }) => {
  const response = await fetch(`${BASE_URL}products/brands/${id}`, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    credentials: 'include',
    body: JSON.stringify(brandData)
  });
  if (!response.ok) return rejectWithValue(await response.json());
  return response.json();
});

export const deleteBrand = createAsyncThunk('admin/deleteBrand', async(id, { rejectWithValue }) => {
  const response = await fetch(`${BASE_URL}products/brands/${id}`, {
    method: 'DELETE',
    credentials: 'include'
  });
  if (!response.ok) return rejectWithValue(await response.json());
  return { id };
});

export const createType = createAsyncThunk('admin/createType', async(typeData, { rejectWithValue }) => {
  const response = await fetch(`${BASE_URL}products/types`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    credentials: 'include',
    body: JSON.stringify(typeData)
  });
  if (!response.ok) return rejectWithValue(await response.json());
  return response.json();
});

export const updateType = createAsyncThunk('admin/updateType', async({ id, typeData }, { rejectWithValue }) => {
  const response = await fetch(`${BASE_URL}products/types/${id}`, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    credentials: 'include',
    body: JSON.stringify(typeData)
  });
  if (!response.ok) return rejectWithValue(await response.json());
  return response.json();
});

export const deleteType = createAsyncThunk('admin/deleteType', async(id, { rejectWithValue }) => {
  const response = await fetch(`${BASE_URL}products/types/${id}`, {
    method: 'DELETE',
    credentials: 'include'
  });
  if (!response.ok) return rejectWithValue(await response.json());
  return { id };
});

const adminSlice = createSlice({
  name: 'admin',
  initialState: {
    status: STATUS.IDLE,
    error: null
  },
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(createBrand.pending, (state) => { state.status = STATUS.LOADING; })
      .addCase(createBrand.fulfilled, (state) => { state.status = STATUS.SUCCEEDED; })
      .addCase(createBrand.rejected, (state, action) => { state.status = STATUS.FAILED; state.error = action.payload; })
      .addCase(updateBrand.pending, (state) => { state.status = STATUS.LOADING; })
      .addCase(updateBrand.fulfilled, (state) => { state.status = STATUS.SUCCEEDED; })
      .addCase(updateBrand.rejected, (state, action) => { state.status = STATUS.FAILED; state.error = action.payload; })
      .addCase(deleteBrand.pending, (state) => { state.status = STATUS.LOADING; })
      .addCase(deleteBrand.fulfilled, (state) => { state.status = STATUS.SUCCEEDED; })
      .addCase(deleteBrand.rejected, (state, action) => { state.status = STATUS.FAILED; state.error = action.payload; })
      .addCase(createType.pending, (state) => { state.status = STATUS.LOADING; })
      .addCase(createType.fulfilled, (state) => { state.status = STATUS.SUCCEEDED; })
      .addCase(createType.rejected, (state, action) => { state.status = STATUS.FAILED; state.error = action.payload; })
      .addCase(updateType.pending, (state) => { state.status = STATUS.LOADING; })
      .addCase(updateType.fulfilled, (state) => { state.status = STATUS.SUCCEEDED; })
      .addCase(updateType.rejected, (state, action) => { state.status = STATUS.FAILED; state.error = action.payload; })
      .addCase(deleteType.pending, (state) => { state.status = STATUS.LOADING; })
      .addCase(deleteType.fulfilled, (state) => { state.status = STATUS.SUCCEEDED; })
      .addCase(deleteType.rejected, (state, action) => { state.status = STATUS.FAILED; state.error = action.payload; });
  }
});

export default adminSlice.reducer;
