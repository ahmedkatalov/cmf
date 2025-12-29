import { createSlice, type PayloadAction } from "@reduxjs/toolkit";

type User = { id?: string; email?: string; name?: string; [key: string]: any } | null;

interface AuthState {
	user: User;
	token: string | null;
	isAuthenticated: boolean;
}

const initialState: AuthState = {
	user: null,
	token: null,
	isAuthenticated: false,
};

const authSlice = createSlice({
	name: "auth",
	initialState,
	reducers: {
		setCredentials(
			state,
			action: PayloadAction<{ user: NonNullable<User>; token: string }>
		) {
			state.user = action.payload.user;
			state.token = action.payload.token;
			state.isAuthenticated = true;
		},
		logout(state) {
			state.user = null;
			state.token = null;
			state.isAuthenticated = false;
		},
		updateUser(state, action: PayloadAction<NonNullable<User>>) {
			state.user = action.payload;
		},
	},
});

export const { setCredentials, logout, updateUser } = authSlice.actions;
export default authSlice.reducer;
