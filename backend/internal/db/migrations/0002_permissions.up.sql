-- 0002_permissions.up.sql
DROP TABLE IF EXISTS user_permissions CASCADE;
DROP TABLE IF EXISTS role_permissions CASCADE;
DROP TABLE IF EXISTS permissions CASCADE;

-- Permission codes
CREATE TABLE permissions (
  code TEXT PRIMARY KEY,
  description TEXT NOT NULL
);

-- Default permissions per role
CREATE TABLE role_permissions (
  role VARCHAR(30) NOT NULL,
  permission_code TEXT NOT NULL REFERENCES permissions(code) ON DELETE CASCADE,
  PRIMARY KEY (role, permission_code)
);

-- Individual user permissions
CREATE TABLE user_permissions (
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  permission_code TEXT NOT NULL REFERENCES permissions(code) ON DELETE CASCADE,
  PRIMARY KEY (user_id, permission_code)
);

-- Fill permissions
INSERT INTO permissions (code, description) VALUES
('branches:create', 'Создание точек/филиалов'),
('branches:view', 'Просмотр списка точек/филиалов'),

('users:create', 'Создание сотрудников'),
('users:view', 'Просмотр сотрудников'),

('clients:create', 'Создание клиентов'),
('clients:edit', 'Редактирование клиентов'),
('clients:view', 'Просмотр карточек клиентов'),
('clients:search', 'Поиск клиентов'),

('contracts:create', 'Создание договоров рассрочки'),
('contracts:edit', 'Редактирование договоров'),
('contracts:view', 'Просмотр договоров'),
('contracts:status_change', 'Изменение статуса договора'),

('payments:create', 'Внесение платежа'),
('payments:cancel', 'Отмена платежа'),
('payments:view', 'Просмотр платежей'),

('transactions:create', 'Добавление расходов/операций'),
('transactions:view', 'Просмотр расходов/операций'),

('summary:view', 'Просмотр финансового отчёта по точке'),
('summary:view_all', 'Просмотр отчёта по всем точкам организации');

-- OWNER: all
INSERT INTO role_permissions (role, permission_code)
SELECT 'owner', code FROM permissions;

-- ADMIN: all
INSERT INTO role_permissions (role, permission_code)
SELECT 'admin', code FROM permissions;

-- MANAGER
INSERT INTO role_permissions (role, permission_code) VALUES
('manager', 'users:create'),
('manager', 'users:view'),

('manager', 'clients:create'),
('manager', 'clients:edit'),
('manager', 'clients:view'),
('manager', 'clients:search'),

('manager', 'contracts:create'),
('manager', 'contracts:edit'),
('manager', 'contracts:view'),
('manager', 'contracts:status_change'),

('manager', 'payments:create'),
('manager', 'payments:view'),

('manager', 'transactions:create'),
('manager', 'transactions:view'),

('manager', 'summary:view');

-- ACCOUNTANT
INSERT INTO role_permissions (role, permission_code) VALUES
('accountant', 'payments:view'),
('accountant', 'transactions:create'),
('accountant', 'transactions:view'),
('accountant', 'summary:view');

-- SECURITY
INSERT INTO role_permissions (role, permission_code) VALUES
('security', 'clients:view'),
('security', 'clients:search'),
('security', 'payments:create'),
('security', 'payments:view'),
('security', 'contracts:view');

-- EMPLOYEE
INSERT INTO role_permissions (role, permission_code) VALUES
('employee', 'payments:create'),
('employee', 'payments:view'),
('employee', 'clients:search'),
('employee', 'clients:view');
