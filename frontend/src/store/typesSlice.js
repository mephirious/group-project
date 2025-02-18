import { createAsyncThunk, createSlice} from "@reduxjs/toolkit";
import {BASE_URL} from "../utils/apiURL";
import {STATUS} from "../utils/status";

const initialState = {
    types: [],
    typesStatus: STATUS.IDLE,
    typeProducts: [],
    typeProductsStatus: STATUS.IDLE
}

const typeSlice = createSlice({
    name: 'type',
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder
        .addCase(fetchAsyncTypes.pending, (state, action) => {
            state.typesStatus = STATUS.LOADING;
        })

        .addCase(fetchAsyncTypes.fulfilled, (state, action) => {
            state.types = action.payload;
            state.typesStatus = STATUS.SUCCEEDED;
        })

        .addCase(fetchAsyncTypes.rejected, (state, action) => {
            state.typesStatus = STATUS.FAILED;
        })

        .addCase(fetchAsyncProductsOfType.pending, (state, action) => {
            state.typeProductsStatus = STATUS.LOADING;
        })

        .addCase(fetchAsyncProductsOfType.fulfilled, (state, action) => {
            state.typeProducts = action.payload;
            state.typeProductsStatus = STATUS.SUCCEEDED;
        })

        .addCase(fetchAsyncProductsOfType.rejected, (state, action) => {
            state.typeProductsStatus = STATUS.FAILED;
        })
    }
});

export const fetchAsyncTypes = createAsyncThunk('types/fetch', async() => {
    const response = await fetch(`${BASE_URL}products/types`);
    const data = await response.json();
    return data.map(type => type.type_name);
});

export const fetchAsyncProductsOfType = createAsyncThunk('type-products/fetch', async(type) => {
    const response = await fetch(`${BASE_URL}products/products?limit=100`);
    const data = await response.json();
    return data.filter(product => product.type === type);
});

export const getAllTypes = (state) => state.type.types;
export const getAllProductsByType = (state) => state.type.typeProducts;
export const getTypeProductsStatus = (state) => state.type.typeProductsStatus;
export default typeSlice.reducer;