import { createSlice } from "@reduxjs/toolkit";
import type { RootState } from "../Store";
interface BaseState { 
    value: number;
}

const initialState: BaseState = {
    value: 0,
}
export const baseSlice = createSlice({
    name: "base",
    initialState,
    reducers: {
        increment: (state) => {
            state.value += 1;
        },
        decrement: (state) => {
            state.value -= 1;
        },
        reset: (state) => {
            state.value = 0;
        },
     },
});

export const { increment, decrement, reset } = baseSlice.actions
export const selectCount = (state: RootState) => state.base.value

export default baseSlice.reducer