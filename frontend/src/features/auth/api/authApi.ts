import { baseApi } from "@/app/api/baseApi";
import { setCredentials } from "@/features/auth/model/authSlice";
import type { AuthToken } from "../types/types";

export const authApi = baseApi.injectEndpoints({
	endpoints: (build) => ({
		login: build.mutation<{ token: string; user: any }, { email: string; password: string }>({
			query: (credentials) => ({ url: "/auth/login", method: "POST", body: credentials }),
			async onQueryStarted(_, { dispatch, queryFulfilled }) {
				try {
					const { data } = await queryFulfilled;
					if (data) dispatch(setCredentials(data));
				} catch (e) {
					// ignore
				}
			},
		}),
		register: build.mutation<{ token: string; user: any }, { email: string; password: string; organization_name?: string }>({
			query: (body) => ({ url: "/auth/register-root", method: "POST", body }),
			async onQueryStarted(_, { dispatch, queryFulfilled }) {
				try {
					const { data } = await queryFulfilled;
					if (data) dispatch(setCredentials(data));
				} catch (e) {
					// ignore
				}
			},
		}),
		me: build.query<AuthToken, void>({
			query: () => {
				return { url: "/auth/me" }
			},
		}),
	}),
	overrideExisting: false,
});

export const { useLoginMutation, useRegisterMutation, useMeQuery } = authApi;
