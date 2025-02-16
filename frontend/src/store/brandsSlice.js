import { createAsyncThunk, createSlice} from "@reduxjs/toolkit";
import {BASE_URL} from "../utils/apiURL";
import {STATUS} from "../utils/status";

const initialState = {
    brands: [],
    brandsStatus: STATUS.IDLE,
    brandProducts: [],
    brandProductsStatus: STATUS.IDLE
}

const brandSlice = createSlice({
    name: 'brand',
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder
        .addCase(fetchAsyncBrands.pending, (state, action) => {
            state.brandsStatus = STATUS.LOADING;
        })

        .addCase(fetchAsyncBrands.fulfilled, (state, action) => {
            state.brands = action.payload;
            state.brandsStatus = STATUS.SUCCEEDED;
        })

        .addCase(fetchAsyncBrands.rejected, (state, action) => {
            state.brandsStatus = STATUS.FAILED;
        })

        .addCase(fetchAsyncProductsOfBrand.pending, (state, action) => {
            state.brandProductsStatus = STATUS.LOADING;
        })

        .addCase(fetchAsyncProductsOfBrand.fulfilled, (state, action) => {
            state.brandProducts = action.payload;
            state.brandProductsStatus = STATUS.SUCCEEDED;
        })

        .addCase(fetchAsyncProductsOfBrand.rejected, (state, action) => {
            state.brandProductsStatus = STATUS.FAILED;
        })
    }
});

export const fetchAsyncBrands = createAsyncThunk('brands/fetch', async() => {
    const response = await fetch(`${BASE_URL}products/brands`);
    const data = await response.json();
    return data.map(brand => brand.brand_name);
});

export const fetchAsyncProductsOfBrand = createAsyncThunk('brand-products/fetch', async(brand) => {
    const response = await fetch(`${BASE_URL}products/products?limit=100`);
    const data = await response.json();
    return data.filter(product => product.brand === brand);
});

export const getAllBrands = (state) => state.brand.brands;
export const getAllProductsByBrand = (state) => state.brand.brandProducts;
export const getBrandProductsStatus = (state) => state.brand.brandProductsStatus;
export default brandSlice.reducer;