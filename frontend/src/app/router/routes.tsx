import React from "react";
import { Routes, Route } from "react-router-dom";
import LoginPage from "@/pages/LoginPage";
import RegisterPage from "@/pages/RegisterPage";
import Dashboard from "@/pages/app/Dashboard";
import TransactionsPage from "@/pages/TransactionsPage";
import MainLayout from "@/widgets/layout/MainLayout";
import ProtectedRoute from "./ProtectedRoute";
import BranchPage from "@/pages/BranchPage";
import BranchDetailPage from "@/pages/BranchDetailPage";

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
				path="/transactions"
				element={
					<ProtectedRoute>
						<MainLayout>
							<TransactionsPage />
						</MainLayout>
					</ProtectedRoute>
				}
			/>
			<Route
				path="/branches"
				element={
					<ProtectedRoute>
						<MainLayout>
							<BranchPage />
						</MainLayout>
					</ProtectedRoute>
				}
			/>
			<Route
				path="/branches/:id"
				element={
					<ProtectedRoute>
						<MainLayout>
							<BranchDetailPage />
						</MainLayout>
					</ProtectedRoute>
				}
			/>
		</Routes>
	);
};

export default AppRoutes;
