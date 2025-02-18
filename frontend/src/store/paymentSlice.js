import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import { BASE_URL } from "../utils/apiURL";
import { STATUS } from "../utils/status";

const initialState = {
  checkoutUrl: null,
  status: STATUS.IDLE,
};

export const createCheckoutSession = createAsyncThunk(
  "payment/createCheckoutSession",
  async (paymentItems, { rejectWithValue }) => {
    const response = await fetch(`${BASE_URL}payment/create-checkout-session`, {
      method: "POST",
      credentials: "include",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(paymentItems),
    });
    if (!response.ok) return rejectWithValue(await response.json());
    const data = await response.json();
    return data;
  }
);

const paymentSlice = createSlice({
  name: "payment",
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(createCheckoutSession.pending, (state) => {
        state.status = STATUS.LOADING;
      })
      .addCase(createCheckoutSession.fulfilled, (state, action) => {
        state.checkoutUrl = action.payload.url;
        state.status = STATUS.SUCCEEDED;
      })
      .addCase(createCheckoutSession.rejected, (state) => {
        state.status = STATUS.FAILED;
      });
  },
});

export const getCheckoutUrl = (state) => state.payment.checkoutUrl;
export const getPaymentStatus = (state) => state.payment.status;
export default paymentSlice.reducer;
