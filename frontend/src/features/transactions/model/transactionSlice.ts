import { createSlice, type PayloadAction } from "@reduxjs/toolkit";
import type { Transaction } from "../types/types";

interface TransactionState {
  transactions: Transaction[];
  selected: Transaction | null;
}

const initialState: TransactionState = {
  transactions: [],
  selected: null,
};

const transactionSlice = createSlice({
  name: "transaction",
  initialState,
  reducers: {
    setTransactions(state, action: PayloadAction<Transaction[]>) {
      state.transactions = action.payload;
    },
    addTransaction(state, action: PayloadAction<Transaction>) {
      state.transactions.push(action.payload);
    },
    updateTransaction(state, action: PayloadAction<Transaction>) {
      state.transactions = state.transactions.map((t) => (t.id === action.payload.id ? action.payload : t));
      if (state.selected?.id === action.payload.id) state.selected = action.payload;
    },
    removeTransaction(state, action: PayloadAction<string>) {
      state.transactions = state.transactions.filter((t) => t.id !== action.payload);
      if (state.selected?.id === action.payload) state.selected = null;
    },
    setSelected(state, action: PayloadAction<Transaction | null>) {
      state.selected = action.payload;
    },
  },
});

export const { setTransactions, addTransaction, updateTransaction, removeTransaction, setSelected } = transactionSlice.actions;
export default transactionSlice.reducer;
