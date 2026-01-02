import { configureStore } from "@reduxjs/toolkit";
import { baseApi } from "@/app/api/baseApi";
import authReducer from "@/features/auth/model/authSlice";
import branchReducer from "@/features/branch/model/branchSlice";

export const store = configureStore({
  reducer: {
    auth: authReducer,
    branch: branchReducer,
    [baseApi.reducerPath]: baseApi.reducer,
  },
  middleware: (getDefaultMiddleware) =>
    getDefaultMiddleware().concat(baseApi.middleware),
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
