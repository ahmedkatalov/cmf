import { createSlice, type PayloadAction } from "@reduxjs/toolkit";
import type { User } from "../types/types";

interface UserState {
  users: User[];
  selected: User | null;
}

const initialState: UserState = {
  users: [],
  selected: null,
};

const userSlice = createSlice({
  name: "user",
  initialState,
  reducers: {
    setUsers(state, action: PayloadAction<User[]>) {
      state.users = action.payload;
    },
    addUser(state, action: PayloadAction<User>) {
      state.users.push(action.payload);
    },
    updateUser(state, action: PayloadAction<User>) {
      state.users = state.users.map((u) => (u.id === action.payload.id ? action.payload : u));
      if (state.selected?.id === action.payload.id) state.selected = action.payload;
    },
    removeUser(state, action: PayloadAction<string>) {
      state.users = state.users.filter((u) => u.id !== action.payload);
      if (state.selected?.id === action.payload) state.selected = null;
    },
    setSelectedUser(state, action: PayloadAction<User | null>) {
      state.selected = action.payload;
    },
  },
});

export const { setUsers, addUser, updateUser, removeUser, setSelectedUser } = userSlice.actions;
export default userSlice.reducer;
