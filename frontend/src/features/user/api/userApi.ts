import { baseApi } from "@/app/api/baseApi";
import type { User, CreateUserRequest } from "../types/types";

export const userApi = baseApi.injectEndpoints({
  endpoints: (build) => ({
    createUser: build.mutation<User, CreateUserRequest>({
      query: (body) => ({
        url: "/users",
        method: "POST",
        body,
      }),
    }),
    getUsersByBranch: build.query<User[], string>({
      query: (branchId) => `/branches/${branchId}/users`,
    }),
  }),
});

export const { useCreateUserMutation, useGetUsersByBranchQuery } = userApi;
