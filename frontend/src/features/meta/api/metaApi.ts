import { baseApi } from "@/app/api/baseApi";
import type { TxType } from "../types/types";

export const metaApi = baseApi.injectEndpoints({
  endpoints: (build) => ({
    getTransactionTypes: build.query<TxType[], void>({
      query: () => "/meta/transaction-types",
      keepUnusedDataFor: 60, // кэшируем на 1 минуту
    }),
  }),
});

export const { useGetTransactionTypesQuery } = metaApi;
