import { createSlice, type PayloadAction } from "@reduxjs/toolkit";
import type { Branch } from "../types/types";

interface BranchState {
  branches: Branch[];
  selected: Branch | null;
}

const initialState: BranchState = {
  branches: [],
  selected: null,
};

const branchSlice = createSlice({
  name: "branch",
  initialState,
  reducers: {
    setBranches(state, action: PayloadAction<Branch[]>) {
      state.branches = action.payload;
    },
    addBranch(state, action: PayloadAction<Branch>) {
      state.branches.push(action.payload);
    },
    updateBranch(state, action: PayloadAction<Branch>) {
      state.branches = state.branches.map((b) => (b.id === action.payload.id ? action.payload : b));
      if (state.selected?.id === action.payload.id) state.selected = action.payload;
    },
    removeBranch(state, action: PayloadAction<string>) {
      state.branches = state.branches.filter((b) => b.id !== action.payload);
      if (state.selected?.id === action.payload) state.selected = null;
    },
    setSelected(state, action: PayloadAction<Branch | null>) {
      state.selected = action.payload;
    },
  },
});

export const { setBranches, addBranch, updateBranch, removeBranch, setSelected } = branchSlice.actions;
export default branchSlice.reducer;
