import { createSlice } from "@reduxjs/toolkit";

const fetchComparisonFromLocalStorage = () => {
    let comparison = localStorage.getItem('comparison');
    if(comparison) {
        return JSON.parse(comparison);
    } else {
        return [];
    }
}

const storeComparisonInLocalStorage = (data) => {
    localStorage.setItem('comparison', JSON.stringify(data));
}

const initialState = {
    comparisonList: fetchComparisonFromLocalStorage(),
    isComparisonMessageOn: false
}

const comparisonSlice = createSlice({
    name: 'comparison',
    initialState,
    reducers: {
        addToComparison: (state, action) => {
            const isItemInComparison = state.comparisonList.find(
                item => item.id === action.payload.id
            );

            if(!isItemInComparison) {
                state.comparisonList.push(action.payload);
                storeComparisonInLocalStorage(state.comparisonList);
            }
        },

        removeFromComparison: (state, action) => {
            const tempComparison = state.comparisonList.filter(
                item => item.id !== action.payload
            );
            state.comparisonList = tempComparison;
            storeComparisonInLocalStorage(state.comparisonList);
        },

        clearComparison: (state) => {
            state.comparisonList = [];
            storeComparisonInLocalStorage(state.comparisonList);
        },

        setComparisonMessageOn: (state) => {
            state.isComparisonMessageOn = true;
        },

        setComparisonMessageOff: (state) => {
            state.isComparisonMessageOn = false;
        }
    }
});

export const { 
    addToComparison, 
    removeFromComparison, 
    clearComparison,
    setComparisonMessageOn,
    setComparisonMessageOff
} = comparisonSlice.actions;

export const getComparisonList = (state) => state.comparison.comparisonList;
export const getComparisonCount = (state) => state.comparison.comparisonList.length;
export const getComparisonMessageStatus = (state) => state.comparison.isComparisonMessageOn;

export default comparisonSlice.reducer;