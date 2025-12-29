import React from "react";
import { useSelector } from "react-redux";
import { Navigate, useLocation } from "react-router-dom";
import type { RootState } from "../store/store";

interface Props {
	children: React.ReactElement;
}

export const ProtectedRoute: React.FC<Props> = ({ children }) => {
	const isAuthenticated = useSelector((s: RootState) => s.auth.isAuthenticated);
	const location = useLocation();
	if (!isAuthenticated) {
		return <Navigate to="/login" state={{ from: location }} replace />;
	}
	return children;
};

export default ProtectedRoute;
