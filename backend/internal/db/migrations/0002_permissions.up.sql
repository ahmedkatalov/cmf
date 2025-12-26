-- Справочник разрешений (permission codes)
CREATE TABLE permissions (
  code TEXT PRIMARY KEY,
  description TEXT NOT NULL
);

-- Какие permissions входят в роль по умолчанию
CREATE TABLE role_permissions (
  role VARCHAR(30) NOT NULL,
  permission_code TEXT NOT NULL REFERENCES permissions(code) ON DELETE CASCADE,
  PRIMARY KEY (role, permission_code)
);

-- Индивидуальные permissions пользователю (добавить/забрать)
CREATE TABLE user_permissions (
  user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  permission_code TEXT NOT NULL REFERENCES permissions(code) ON DELETE CASCADE,
  PRIMARY KEY (user_id, permission_code)
);

-- 🚀 Заполним базовые permissions
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

-- ✅ Привяжем permissions к ролям (базовые правила)
-- OWNER: всё
INSERT INTO role_permissions (role, permission_code)
SELECT 'owner', code FROM permissions;

-- ADMIN: всё (как owner)
INSERT INTO role_permissions (role, permission_code)
SELECT 'admin', code FROM permissions;

-- MANAGER: операционка точки + клиенты + договоры + платежи + сотрудники в своей точке
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

-- ACCOUNTANT: финансы + отчёты точки, без клиентов и договоров
INSERT INTO role_permissions (role, permission_code) VALUES
('accountant', 'payments:view'),
('accountant', 'transactions:create'),
('accountant', 'transactions:view'),
('accountant', 'summary:view');

-- SECURITY (СБ): поиск клиентов + просмотр карточек + внесение платежа, без создания клиентов и без summary
INSERT INTO role_permissions (role, permission_code) VALUES
('security', 'clients:view'),
('security', 'clients:search'),
('security', 'payments:create'),
('security', 'payments:view'),
('security', 'contracts:view');

-- EMPLOYEE (кассир/обычный): внесение платежа + просмотр клиентов (опционально)
INSERT INTO role_permissions (role, permission_code) VALUES
('employee', 'payments:create'),
('employee', 'payments:view'),
('employee', 'clients:search'),
('employee', 'clients:view');
