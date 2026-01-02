-- Категории (можно использовать для фильтров и аккуратных отчётов)
CREATE TABLE categories (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
  name TEXT NOT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  UNIQUE (organization_id, name)
);

-- Основная таблица операций (доходы + расходы)
CREATE TABLE transactions (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),

  organization_id UUID NOT NULL REFERENCES organizations(id) ON DELETE CASCADE,
  branch_id UUID NOT NULL REFERENCES branches(id) ON DELETE CASCADE,
  created_by UUID REFERENCES users(id) ON DELETE SET NULL,

  type VARCHAR(30) NOT NULL, -- income | expense_company | expense_people
  category_id UUID REFERENCES categories(id) ON DELETE SET NULL,

  amount BIGINT NOT NULL CHECK (amount > 0),
  occurred_at DATE NOT NULL,
  description TEXT,

  created_at TIMESTAMP DEFAULT NOW()
);

CREATE INDEX idx_tx_org_branch_date ON transactions(organization_id, branch_id, occurred_at);
CREATE INDEX idx_tx_type ON transactions(type);
CREATE INDEX idx_tx_category ON transactions(category_id);
