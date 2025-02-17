import { createAsyncThunk, createSlice } from "@reduxjs/toolkit";
import { BASE_URL } from "../utils/apiURL";
import { STATUS } from "../utils/status";

const initialState = {
    posts: [],
    postsStatus: STATUS.IDLE,
    postSingle: [],
    postSingleStatus: STATUS.IDLE
}

const postSlice = createSlice({
    name: "post",
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder
        .addCase(fetchAsyncPosts.pending, (state, action) => {
            state.postsStatus = STATUS.LOADING;
        })

        .addCase(fetchAsyncPosts.fulfilled, (state, action) => {
            state.posts = action.payload;
            state.postsStatus = STATUS.SUCCEEDED;
        })
        
        .addCase(fetchAsyncPosts.rejected, (state, action) => {
            state.postsStatus = STATUS.FAILED
        })

        .addCase(fetchAsyncPostSingle.pending, (state, action) => {
            state.postSingleStatus = STATUS.LOADING;
        })

        .addCase(fetchAsyncPostSingle.fulfilled, (state, action) => {
            state.postSingle = action.payload;
            state.postSingleStatus = STATUS.SUCCEEDED;
        })

        .addCase(fetchAsyncPostSingle.rejected, (state, action) => {
            state.postSingleStatus = STATUS.FAILED;
        })
    }
});

export const fetchAsyncPosts = createAsyncThunk('posts/fetch', async(limit) => {
    const response = await fetch(`${BASE_URL}blogs/blog-posts?limit=${limit}`);
    const data = await response.json();
    return data;
});

export const fetchAsyncPostSingle = createAsyncThunk('post-single/fetch', async(id) => {
    const response = await fetch(`${BASE_URL}blogs/blog-posts/${id}`);
    const data = await response.json();
    return data;
});


export const getAllPosts = (state) => state.post.posts;
export const getAllPostsStatus = (state) => state.post.postsStatus;
export const getPostSingle = (state) => state.post.postSingle;
export const getSinglePostStatus = (state) => state.post.postSingleStatus;
export default postSlice.reducer;