import { configureStore } from "@reduxjs/toolkit";
import baseReducer from "./Reducers/baseSlice.ts";

const store = configureStore({
    reducer: {
        base:baseReducer,
    },
});
export type AppStore = typeof store;
export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch

export default store;