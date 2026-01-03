export type User = {
  id: string;
  branch_id: string;
  email: string;
  role: "admin" | "user";
  created_at: string;
};

export type CreateUserRequest = {
  branch_id: string; // автоматический
  email: string;
  password: string;
  role: "admin" | "user";
};