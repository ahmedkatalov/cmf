import React from "react";
import { Link } from "react-router-dom";

const Dashboard: React.FC = () => {
	return (
		<div style={{ padding: 20 }}>
			<h1>Главная</h1>
			<p>Добро пожаловать в приложение.</p>
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
    )
};

export default Dashboard;
