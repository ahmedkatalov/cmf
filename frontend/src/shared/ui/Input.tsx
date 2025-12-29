import React from 'react'

type InputProps = React.InputHTMLAttributes<HTMLInputElement> & {
	label?: string
}

const Input: React.FC<InputProps> = ({ label, className = '', ...props }) => {
	return (
		<label className="block">
			{label && <span className="text-sm font-medium mb-1 block text-gray-700">{label}</span>}
			<input
				className={
					'w-full px-4 py-2 border border-gray-300 rounded-md focus:outline-none focus:ring-2 focus:ring-blue-400 bg-white ' +
					className
				}
				{...props}
			/>
		</label>
	)
}

export default Input
