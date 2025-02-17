import {configureStore} from "@reduxjs/toolkit";
import sidebarReducer from "./sidebarSlice";
import categoryReducer from "./categorySlice";
import brandReducer from "./brandsSlice";
import typeReducer from "./typesSlice";
import productReducer from "./productSlice";
import cartReducer from "./cartSlice";
import searchReducer from "./searchSlice";
import authReducer from "./authSlice";
import postReducer from "./blogSlice";
import comparisonReducer from './comparisonSlice';

const store = configureStore({
    reducer: {
        sidebar: sidebarReducer,
        category: categoryReducer,
        brand: brandReducer,
        type: typeReducer,
        product: productReducer,
        cart: cartReducer,
        search: searchReducer,
        auth: authReducer,
        post: postReducer,
        comparison: comparisonReducer,
    }
});

export default store;