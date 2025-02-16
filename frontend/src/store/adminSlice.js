import { createAsyncThunk, createSlice } from '@reduxjs/toolkit';
import { BASE_URL } from '../utils/apiURL';
import { STATUS } from '../utils/status';

// Category Slice
export const createCategory = createAsyncThunk('admin/createCategory', async(category_name, { rejectWithValue }) => {
  const response = await fetch(`${BASE_URL}products/categories`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    credentials: 'include',
    body: JSON.stringify(category_name)
  });
  if (!response.ok) return rejectWithValue(await response.json());
  return response.json();
});

export const updateCategory = createAsyncThunk('admin/updateCategory', async({ id, data }, { rejectWithValue }) => {
  const response = await fetch(`${BASE_URL}products/categories/${id}`, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    credentials: 'include',
    body: JSON.stringify(data)
  });
  if (!response.ok) return rejectWithValue(await response.json());
  return response.json();
});

export const deleteCategory = createAsyncThunk('admin/deleteCategory', async(id, { rejectWithValue }) => {
  const response = await fetch(`${BASE_URL}products/categories/${id}`, {
    method: 'DELETE',
    credentials: 'include'
  });
  if (!response.ok) return rejectWithValue(await response.json());
  return { id };
});

// Brand Slice
export const createBrand = createAsyncThunk('admin/createBrand', async(brand_name, { rejectWithValue }) => {
  const response = await fetch(`${BASE_URL}products/brands`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    credentials: 'include',
    body: JSON.stringify(brand_name)
  });
  if (!response.ok) return rejectWithValue(await response.json());
  return response.json();
});

export const updateBrand = createAsyncThunk('admin/updateBrand', async({ id, data }, { rejectWithValue }) => {
  const response = await fetch(`${BASE_URL}products/brands/${id}`, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    credentials: 'include',
    body: JSON.stringify(data)
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

// Type Slice
export const createType = createAsyncThunk('admin/createType', async(type_name, { rejectWithValue }) => {
  const response = await fetch(`${BASE_URL}products/types`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    credentials: 'include',
    body: JSON.stringify(type_name)
  });
  if (!response.ok) return rejectWithValue(await response.json());
  return response.json();
});

export const updateType = createAsyncThunk('admin/updateType', async({ id, data }, { rejectWithValue }) => {
  const response = await fetch(`${BASE_URL}products/types/${id}`, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    credentials: 'include',
    body: JSON.stringify(data)
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

// Product Slice
export const createProduct = createAsyncThunk('admin/createProduct', async(product, { rejectWithValue }) => {
  const response = await fetch(`${BASE_URL}products/products`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    credentials: 'include',
    body: JSON.stringify(product)
  });
  if (!response.ok) return rejectWithValue(await response.json());
  return response.json();
});

export const updateProduct = createAsyncThunk('admin/updateProduct', async({ id, data }, { rejectWithValue }) => {
  const response = await fetch(`${BASE_URL}products/products/${id}`, {
    method: 'PUT',
    headers: { 'Content-Type': 'application/json' },
    credentials: 'include',
    body: JSON.stringify(data)
  });
  if (!response.ok) return rejectWithValue(await response.json());
  return response.json();
});

export const deleteProduct = createAsyncThunk('admin/deleteProduct', async(id, { rejectWithValue }) => {
  const response = await fetch(`${BASE_URL}products/products/${id}`, {
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
      .addCase(createCategory.pending, (state) => { state.status = STATUS.LOADING; })
      .addCase(createCategory.fulfilled, (state) => { state.status = STATUS.SUCCEEDED; })
      .addCase(createCategory.rejected, (state, action) => { state.status = STATUS.FAILED; state.error = action.payload; })
      .addCase(updateCategory.pending, (state) => { state.status = STATUS.LOADING; })
      .addCase(updateCategory.fulfilled, (state) => { state.status = STATUS.SUCCEEDED; })
      .addCase(updateCategory.rejected, (state, action) => { state.status = STATUS.FAILED; state.error = action.payload; })
      .addCase(deleteCategory.pending, (state) => { state.status = STATUS.LOADING; })
      .addCase(deleteCategory.fulfilled, (state) => { state.status = STATUS.SUCCEEDED; })
      .addCase(deleteCategory.rejected, (state, action) => { state.status = STATUS.FAILED; state.error = action.payload; })

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
