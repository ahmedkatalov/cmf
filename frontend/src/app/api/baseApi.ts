import { createApi, fetchBaseQuery } from "@reduxjs/toolkit/query/react";
import { ENV } from "../config/env";
import type { RootState } from "../store/store";

export const baseApi = createApi({
	reducerPath: "baseApi",
	baseQuery: fetchBaseQuery({
		baseUrl: ENV.API_URL,
		prepareHeaders: (headers, { getState }) => {
			const state = getState() as RootState;
			const token = state?.auth?.token;
			if (token) headers.set("Authorization", `Bearer ${token}`);
			return headers;
		},
	}),
	tagTypes: ["Auth"],
	endpoints: () => ({}),
});
