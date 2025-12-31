import React, { useState } from "react";
import { useRegisterMutation } from "../api/authApi";
import { useNavigate, Link, Navigate } from "react-router-dom";
import { useSelector } from "react-redux";
import type { RootState } from "@/app/store/store";
import Input from "@/shared/ui/Input";
import Button from "@/shared/ui/Button";

const RegisterPage: React.FC = () => {
	const isAuthenticated = useSelector((s: RootState) => s.auth.isAuthenticated);
	if (isAuthenticated) return <Navigate to="/" replace />;
	const [email, setEmail] = useState("");
	const [password, setPassword] = useState("");
	const [organizationName, setOrganizationName] = useState("");
	const [register, { isLoading }] = useRegisterMutation();
	const navigate = useNavigate();

	const handleSubmit = async (e: React.FormEvent) => {
		e.preventDefault();
		try {
			await register({ email, password, organization_name: organizationName }).unwrap();
			navigate("/", { replace: true });
		} catch (err) {
			console.error(err);
		}
	};

	return (
		<div className="min-h-screen flex items-center justify-center bg-gray-100">
			<div className="w-full max-w-md bg-white rounded-lg shadow-md p-6">
				<h1 className="text-2xl font-semibold text-gray-800 mb-4">Регистрация</h1>
				<form onSubmit={handleSubmit} className="space-y-4">
					<Input label="Имя организации" value={organizationName} onChange={(e) => setOrganizationName(e.target.value)} required />
					<Input label="Email" type="email" value={email} onChange={(e) => setEmail(e.target.value)} required />
					<Input label="Пароль" type="password" value={password} onChange={(e) => setPassword(e.target.value)} required />
					<Button type="submit" disabled={isLoading} className="w-full">
						Зарегистрироваться
					</Button>
				</form>
				<div className="mt-4 text-sm text-gray-600 text-center">
					Уже есть аккаунт? <Link to="/login" className="text-blue-600 hover:underline">Войти</Link>
				</div>
			</div>
		</div>
	);
};

export default RegisterPage;
