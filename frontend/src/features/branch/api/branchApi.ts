import { baseApi } from "@/app/api/baseApi";
import type { Branch, CreateBranchDto } from "../types/types";

export const branchApi = baseApi.injectEndpoints({
  endpoints: (build) => ({
    getBranches: build.query<Branch[], void>({
      query: () => ({ url: "/branches", method: "GET" }),
      providesTags: ["Branch"],
    }),
    getBranch: build.query<Branch, string>({
      query: (id) => ({ url: `/branches/${id}`, method: "GET" }),
      providesTags: ["Branch"],
    }),
    createBranch: build.mutation<Branch, CreateBranchDto>({
      query: (body) => ({ url: "/branches", method: "POST", body }),
      invalidatesTags: ["Branch"],
    }),
    updateBranch: build.mutation<Branch, { id: string; body: Partial<CreateBranchDto> }>({
      query: ({ id, body }) => ({ url: `/branches/${id}`, method: "PUT", body }),
      invalidatesTags: ["Branch"],
    }),
    deleteBranch: build.mutation<{ ok: boolean }, string>({
      query: (id) => ({ url: `/branches/${id}`, method: "DELETE" }),
      invalidatesTags: ["Branch"],
    }),
  }),
  overrideExisting: false,
});

export const {
  useGetBranchesQuery,
  useGetBranchQuery,
  useCreateBranchMutation,
  useUpdateBranchMutation,
  useDeleteBranchMutation,
} = branchApi;
