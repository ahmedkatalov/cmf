import React from "react";
import { Link } from "react-router-dom";
import {  useSelector } from "react-redux";

const Dashboard: React.FC = () => {
	const { auth } = useSelector((state: any) => state);

	return (
		<div style={{ padding: 20 }}>	
			<h1>Главная</h1>
			<p>Добро пожаловать в приложение.</p>

			<div style={{ marginBottom: 16 }}>
				<strong>Auth state:</strong>
				<pre style={{ whiteSpace: "pre-wrap", background: "#f5f5f5", padding: 8 }}>
					{JSON.stringify(auth, null, 2)}
				</pre>
				{auth?.isAuthenticated && (
					<p>Signed in as: {auth.user?.email ?? "(no email)"}</p>
				)}
			</div>

			<nav>
				<ul>
					<li>
						<Link to="/income">Доходы</Link>
					</li>
					<li>
						<Link to="/expenses">Расходы</Link>
					</li>
					<li>
						<Link to="/login">Войти</Link>
					</li>
					<li>
						<Link to="/register">Регистрация</Link>
					</li>
				</ul>
			</nav>
		</div>
	);
};

export default Dashboard;
