import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import { BASE_URL } from "../utils/apiURL";
import { STATUS } from "../utils/status";

const initialState = {
  reviews: [],
  reviewsStatus: STATUS.IDLE,
  reviewProducts: [],
  reviewProductsStatus: STATUS.IDLE,
  reviewCreateStatus: STATUS.IDLE,
};

export const fetchAsyncReviews = createAsyncThunk('reviews/fetch', async () => {
  const response = await fetch(`${BASE_URL}reviews/reviews`);
  const data = await response.json();
  return data;
});

export const fetchAsyncReviewById = createAsyncThunk('review-id/fetch', async (id) => {
  const response = await fetch(`${BASE_URL}reviews/reviews/${id}`);
  const data = await response.json();
  return data;
});

export const fetchAsyncReviewsOfUser = createAsyncThunk('reviews-user/fetch', async (user_id) => {
  const response = await fetch(`${BASE_URL}reviews/reviews/customer/${user_id}`);
  const data = await response.json();
  return data;
});

export const fetchAsyncReviewsOfProduct = createAsyncThunk('reviews-product/fetch', async (product_id) => {
  const response = await fetch(`${BASE_URL}reviews/reviews/product/${product_id}`);
  const data = await response.json();
  return data;
});

export const createAsyncReview = createAsyncThunk('reviews/create', async (reviewData) => {
  const response = await fetch(`${BASE_URL}reviews/reviews`, {
    method: "POST",
    credentials: "include",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(reviewData),
  });
  if (!response.ok) {
    throw new Error("Failed to create review");
  }
  return await response.json();
});

const reviewSlice = createSlice({
  name: 'review',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(fetchAsyncReviews.pending, (state) => {
        state.reviewsStatus = STATUS.LOADING;
      })
      .addCase(fetchAsyncReviews.fulfilled, (state, action) => {
        state.reviews = action.payload;
        state.reviewsStatus = STATUS.SUCCEEDED;
      })
      .addCase(fetchAsyncReviews.rejected, (state) => {
        state.reviewsStatus = STATUS.FAILED;
      })
      .addCase(fetchAsyncReviewById.pending, (state) => {
        state.reviewProductsStatus = STATUS.LOADING;
      })
      .addCase(fetchAsyncReviewById.fulfilled, (state, action) => {
        state.reviewProducts = action.payload;
        state.reviewProductsStatus = STATUS.SUCCEEDED;
      })
      .addCase(fetchAsyncReviewById.rejected, (state) => {
        state.reviewProductsStatus = STATUS.FAILED;
      })
      .addCase(fetchAsyncReviewsOfUser.pending, (state) => {
        state.reviewProductsStatus = STATUS.LOADING;
      })
      .addCase(fetchAsyncReviewsOfUser.fulfilled, (state, action) => {
        state.reviewProducts = action.payload;
        state.reviewProductsStatus = STATUS.SUCCEEDED;
      })
      .addCase(fetchAsyncReviewsOfUser.rejected, (state) => {
        state.reviewProductsStatus = STATUS.FAILED;
      })
      .addCase(fetchAsyncReviewsOfProduct.pending, (state) => {
        state.reviewProductsStatus = STATUS.LOADING;
      })
      .addCase(fetchAsyncReviewsOfProduct.fulfilled, (state, action) => {
        state.reviewProducts = action.payload;
        state.reviewProductsStatus = STATUS.SUCCEEDED;
      })
      .addCase(fetchAsyncReviewsOfProduct.rejected, (state) => {
        state.reviewProductsStatus = STATUS.FAILED;
      })
      .addCase(createAsyncReview.pending, (state) => {
        state.reviewCreateStatus = STATUS.LOADING;
      })
      .addCase(createAsyncReview.fulfilled, (state, action) => {
        state.reviewCreateStatus = STATUS.SUCCEEDED;
        if (state.reviewProducts && state.reviewProducts.reviews) {
          state.reviewProducts.reviews.unshift(action.payload);
        }
      })
      .addCase(createAsyncReview.rejected, (state) => {
        state.reviewCreateStatus = STATUS.FAILED;
      });
  }
});

export const getAllReviews = (state) => state.review.reviews;
export const getAllProductsByReview = (state) => state.review.reviewProducts;
export const getReviewProductsStatus = (state) => state.review.reviewProductsStatus;
export default reviewSlice.reducer;
