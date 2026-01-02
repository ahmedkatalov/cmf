import React, { useEffect } from "react";
import { useDispatch, useSelector } from "react-redux";
import { Navigate, useLocation } from "react-router-dom";
import type { RootState } from "../store/store";
import { useMeQuery } from "@/features/auth/api/authApi";
import { updateUser } from "@/features/auth/model/authSlice";

interface Props {
	children: React.ReactElement;
}

export const ProtectedRoute: React.FC<Props> = ({ children }) => {
	const { data, isLoading, isSuccess } = useMeQuery();
	const dispatch = useDispatch();

	useEffect(() => {
	if (isSuccess && data) {
		dispatch(updateUser(data.claims));
	}
	}, [isSuccess, data, dispatch]);

	const isAuthenticated = useSelector((s: RootState) => s.auth.isAuthenticated);
	const location = useLocation();
	if(isLoading) {
		return <div>Loading...</div>;
	}

	if (!isAuthenticated) {
		return <Navigate to="/login" state={{ from: location }} replace />;
	}
	
	return children;
};

export default ProtectedRoute;
