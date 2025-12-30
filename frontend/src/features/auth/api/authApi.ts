import { baseApi } from "../../../app/api/baseApi";
import { setCredentials } from "../../auth/model/authSlice";

export const authApi = baseApi.injectEndpoints({
	endpoints: (build) => ({
		login: build.mutation<{ token: string; user: any }, { email: string; password: string }>({
			query: (credentials) => ({ url: "/api/auth/login", method: "POST", body: credentials }),
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
			query: (body) => ({ url: "/api/auth/register-root", method: "POST", body }),
			async onQueryStarted(_, { dispatch, queryFulfilled }) {
				try {
					const { data } = await queryFulfilled;
					if (data) dispatch(setCredentials(data));
				} catch (e) {
					// ignore
				}
			},
		}),
		me: build.query<any, void>({
			query: () => ({ url: "/auth/me" }),
		}),
	}),
	overrideExisting: false,
});

export const { useLoginMutation, useRegisterMutation, useMeQuery } = authApi;
