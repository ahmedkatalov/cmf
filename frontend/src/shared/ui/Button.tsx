import React from 'react'

type ButtonProps = React.ButtonHTMLAttributes<HTMLButtonElement> & {
	variant?: 'primary' | 'ghost'
}

const Button: React.FC<ButtonProps> = ({ children, className = '', variant = 'primary', ...props }) => {
	const base = 'px-4 py-2 rounded-md font-medium focus:outline-none focus:ring-2'
	const variants: Record<string, string> = {
		primary: 'bg-blue-600 text-white hover:bg-blue-700 focus:ring-blue-300',
		ghost: 'bg-white text-blue-600 border border-blue-200 hover:bg-blue-50 focus:ring-blue-200',
	}

	return (
		<button className={`${base} ${variants[variant]} ${className}`} {...props}>
			{children}
		</button>
	)
}

export default Button
