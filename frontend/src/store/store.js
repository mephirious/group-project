import {configureStore} from "@reduxjs/toolkit";
import sidebarReducer from "./sidebarSlice";
import categoryReducer from "./categorySlice";
import brandReducer from "./brandsSlice";
import productReducer from "./productSlice";
import cartReducer from "./cartSlice";
import searchReducer from "./searchSlice";
import authReducer from "./authSlice";

const store = configureStore({
    reducer: {
        sidebar: sidebarReducer,
        category: categoryReducer,
        brand: brandReducer,
        product: productReducer,
        cart: cartReducer,
        search: searchReducer,
        auth: authReducer,
    }
});

export default store;