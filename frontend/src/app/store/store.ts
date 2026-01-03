import { configureStore } from "@reduxjs/toolkit";
import { baseApi } from "@/app/api/baseApi";
import authReducer from "@/features/auth/model/authSlice";
import branchReducer from "@/features/branch/model/branchSlice";
import transactionReducer from "@/features/transactions/model/transactionSlice";
import userReducer from "@/features/user/model/userSlice";

export const store = configureStore({
  reducer: {
    auth: authReducer,
    branch: branchReducer,
    transactions: transactionReducer,
    user: userReducer,
    [baseApi.reducerPath]: baseApi.reducer,
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware().concat(baseApi.middleware),
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
