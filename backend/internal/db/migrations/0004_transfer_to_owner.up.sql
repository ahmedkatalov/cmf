-- Ничего в таблице менять не надо, потому что type VARCHAR
-- Но если ты хочешь строгую валидацию через CHECK — тогда нужно CHECK constraint

-- Сейчас CHECK нет на type, поэтому можно просто использовать новый type без миграции.
-- Но лучше добавить справочник типов или CHECK.

-- Если хочешь CHECK, можно так:

ALTER TABLE transactions
DROP CONSTRAINT IF EXISTS transactions_type_check;

ALTER TABLE transactions
ADD CONSTRAINT transactions_type_check
CHECK (type IN ('income', 'expense_company', 'expense_people', 'transfer_to_owner'));
