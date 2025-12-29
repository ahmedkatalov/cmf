import React, { useState } from "react";
import { useLoginMutation } from "../api/authApi";
import { useNavigate, useLocation, Link } from "react-router-dom";
import Input from "@/shared/ui/Input";
import Button from "@/shared/ui/Button";

const LoginPage: React.FC = () => {
	const [email, setEmail] = useState("");
	const [password, setPassword] = useState("");
	const [login, { isLoading }] = useLoginMutation();
	const navigate = useNavigate();
	const location = useLocation();

	const from = (location.state as any)?.from?.pathname || "/";

	const handleSubmit = async (e: React.FormEvent) => {
		e.preventDefault();
		try {
			await login({ email, password }).unwrap();
			navigate(from, { replace: true });
		} catch (err) {
			console.error(err);
		}
	};

	return (
		<div className="min-h-screen flex items-center justify-center bg-gray-100">
			<div className="w-full max-w-md bg-white rounded-lg shadow-md p-6">
				<h1 className="text-2xl font-semibold text-gray-800 mb-4">Вход</h1>
				<form onSubmit={handleSubmit} className="space-y-4">
					<Input label="Email" type="email" value={email} onChange={(e) => setEmail(e.target.value)} required />
					<Input label="Пароль" type="password" value={password} onChange={(e) => setPassword(e.target.value)} required />
					<div className="flex items-center justify-between">
						<Button type="submit" disabled={isLoading} className="w-full">
							Войти
						</Button>
					</div>
				</form>
				<div className="mt-4 text-sm text-gray-600 text-center">
					Нет аккаунта? <Link to="/register" className="text-blue-600 hover:underline">Зарегистрироваться</Link>
				</div>
			</div>
		</div>
	);
};

export default LoginPage;
