import React from "react";
import { Routes, Route } from "react-router-dom";
import LoginPage from "@/features/auth/ui/LoginPage";
import RegisterPage from "@/features/auth/ui/RegisterPage";
import Dashboard from "@/widgets/pages/app/Dashboard";
import Income from "@/widgets/pages/Income";
import Expenses from "@/widgets/pages/Expenses";
import MainLayout from "@/widgets/layout/MainLayout";
import ProtectedRoute from "./ProtectedRoute";

export const AppRoutes: React.FC = () => {
	return (
		<Routes>
			<Route
				path="/"
				element={
					<ProtectedRoute>
						<MainLayout>
							<Dashboard />
						</MainLayout>
					</ProtectedRoute>
				}
			/>
			<Route path="/login" element={<LoginPage />} />
			<Route path="/register" element={<RegisterPage />} />
			<Route
				path="/income"
				element={
					<ProtectedRoute>
						<MainLayout>
							<Income />
						</MainLayout>
					</ProtectedRoute>
				}
			/>
			<Route
				path="/expenses"
				element={
						<MainLayout>
							<Expenses />
						</MainLayout>
				}
			/>
		</Routes>
	);
};

export default AppRoutes;
