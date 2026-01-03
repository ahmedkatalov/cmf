import { baseApi } from "@/app/api/baseApi";
import type { Transaction, CreateTransactionDto } from "../types/types";

export const transactionApi = baseApi.injectEndpoints({
  endpoints: (build) => ({
    getTransactionsByBranch: build.query<Transaction[], { orgID: string; branchID: string }>({
      query: () => ({
        url: "/transactions/",
        method: "GET",
      }),
      providesTags: ["Transaction"],
    }),
    createTransaction: build.mutation<Transaction, CreateTransactionDto>({
      query: (body) => ({ url: "/transactions/", method: "POST", body }),
      invalidatesTags: ["Transaction"],
    }),
  }),
  overrideExisting: false,
});

export const { useGetTransactionsByBranchQuery, useCreateTransactionMutation } = transactionApi;
