import React from 'react'
import { Link } from 'react-router-dom'

const Header: React.FC = () => {
  return (
    <header className="bg-white border-b shadow-sm">
      <div className="max-w-6xl mx-auto px-4 py-3 flex items-center justify-between">
        <Link to="/" className="text-lg font-semibold text-blue-600">CMF</Link>
        <nav className="space-x-4">
          <Link to="/income" className="text-gray-700 hover:text-blue-600">Доходы</Link>
          <Link to="/expenses" className="text-gray-700 hover:text-blue-600">Расходы</Link>
        </nav>
      </div>
    </header>
  )
}

export default Header
