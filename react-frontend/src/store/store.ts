import { configureStore } from "@reduxjs/toolkit";
import  rootReducer from './rootReducer';
import { appMiddleware } from "./middleware";

export const store = configureStore({
    reducer: rootReducer,
    middleware: appMiddleware,
})

export type RootState = ReturnType<typeof store.getState>
export type AppDispatch = typeof store.dispatch

export default store;