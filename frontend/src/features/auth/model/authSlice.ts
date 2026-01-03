import { createSlice, type PayloadAction } from "@reduxjs/toolkit";

type User = { id?: string; email?: string; name?: string; [key: string]: any } | null;

interface AuthState {
	user: User;
	token: string | null;
	isAuthenticated: boolean;
}

const STORAGE_KEY = "auth";

function loadFromStorage(): { user: User; token: string } | null {
	try {
		if (typeof window === "undefined") return null;
		const raw = localStorage.getItem(STORAGE_KEY);
		if (!raw) return null;
		return JSON.parse(raw);
	} catch (e) {
		return null;
	}
}

const persisted = loadFromStorage();

const initialState: AuthState = {
	user: persisted?.user ?? null,
	token: persisted?.token ?? null,
	isAuthenticated: persisted ? true : false,
};

const authSlice = createSlice({
	name: "auth",
	initialState,
	reducers: {
		setCredentials(state, action: PayloadAction<{ user: NonNullable<User>; token: string }>) {
			state.user = action.payload.user;
			state.token = action.payload.token;
			state.isAuthenticated = true;
			try {
				if (typeof window !== "undefined") {
					localStorage.setItem(STORAGE_KEY, JSON.stringify(state.token ? { user: state.user, token: state.token } : null	));
				}
			} catch (e) {
				// ignore storage errors
			}
		},
		logout(state) {
			state.user = null;
			state.token = null;
			state.isAuthenticated = false;
			try {
				if (typeof window !== "undefined") {
					localStorage.removeItem(STORAGE_KEY);
				}
			} catch (e) {
				// ignore
			}
		},
		updateUser(state, action: PayloadAction<NonNullable<User>>) {
			console.log(state, action)
			state.user = action.payload;
		},
	},
});

export const { setCredentials, logout, updateUser } = authSlice.actions;
export default authSlice.reducer;
