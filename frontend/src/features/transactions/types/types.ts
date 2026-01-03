export interface Transaction {
  id: string;
  organization_id: string;
  branch_id: string;
  created_by?: string | null;
  type: string;
  category_id?: string | null;
  amount: number;
  occurred_at: string; // ISO date string
  description?: string | null;
  created_at: string; // ISO date string
  [key: string]: any;
}

export interface CreateTransactionDto {
  organization_id: string;
  branch_id: string;
  created_by?: string | null;
  type: string;
  category_id?: string | null;
  amount: number;
  occurred_at: string; // ISO date string
  description?: string | null;
}
