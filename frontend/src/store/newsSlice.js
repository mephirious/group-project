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

// for getting the products list with limited numbers
export const fetchAsyncPosts = createAsyncThunk('posts/fetch', async(limit) => {
    // const response = await fetch(`${BASE_URL}blog/posts?limit=${limit}`);
    // const data = await response.json();
    return [
        {
            "id": "loremID123123",
            "title": "lorem ipsum",
            "description": "После покупки ноутубка, мы рекомендуем создать учетные записи для комфортного и эффективного использования.",
            "image": "https://picsum.photos/800/600",
            "created_at": "2021-09-01T00:00:00.000Z"
        },
        {
            "id": "loremID123124",
            "title": "lorem ipsum",
            "description": "После покупки ноутубка, мы рекомендуем создать учетные записи для комфортного и эффективного использования.",
             "image": "https://picsum.photos/800/610",
            "created_at": "2021-10-01T00:00:00.000Z"
        },
        {
            "id": "loremID123125",
            "title": "lorem ipsum",
            "description": "После покупки ноутубка, мы рекомендуем создать учетные записи для комфортного и эффективного использования.",
             "image": "https://picsum.photos/800/620",
            "created_at": "2025-02-11T07:30:29.000Z"
        },
        {
            "id": "loremID123126",
            "title": "lorem ipsum",
            "description": "После покупки ноутубка, мы рекомендуем создать учетные записи для комфортного и эффективного использования.",
            "image": "https://picsum.photos/800/630",
            "created_at": "2025-02-15T07:30:29.000Z"
        },
        {
            "id": "loremID123127",
            "title": "lorem ipsum",
            "description": "После покупки ноутубка, мы рекомендуем создать учетные записи для комфортного и эффективного использования.",
            "image": "https://picsum.photos/800/640",
            "created_at": "2025-02-16T07:30:29.000Z"
        },
        {
            "id": "loremID123128",
            "title": "lorem ipsum",
            "description": "После покупки ноутубка, мы рекомендуем создать учетные записи для комфортного и эффективного использования.",
            "image": "https://picsum.photos/800/650",
            "created_at": "2025-02-16T11:11:29.000Z"
        }
    ];
});

// getting the single product data also
export const fetchAsyncPostSingle = createAsyncThunk('post-single/fetch', async(id) => {
    // const response = await fetch(`${BASE_URL}blog/posts/${id}`);
    // const data = await response.json();
    return {
        "id": "loremID123123",
        "title": "lorem ipsum",
        "description": "lorem ipsum dolor sit amet, consectetur adipiscing elit. Integer nec odio. Praesent libero. Sed cursus ante dapibus diam. Sed nisi. Nulla quis sem at nibh elementum imperdiet.",
         "image": "https://picsum.photos/800/605",
            "created_at": "2025-02-16T11:11:29.000Z"
    };
});


export const getAllPosts = (state) => state.post.posts;
export const getAllPostsStatus = (state) => state.post.postsStatus;
export const getPostSingle = (state) => state.post.postSingle;
export const getSinglePostStatus = (state) => state.post.postSingleStatus;
export default postSlice.reducer;