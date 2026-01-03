import type { CreateUserRequest } from "@/features/user/types/types";

export interface Branch {
  id: string;
  organization_id: string;
  name: string;
  address?: string;
  [key: string]: any;
}

export interface CreateBranchDto {
  name: string;
  address?: string;
}

export type CreateUserFormValues = Omit<CreateUserRequest, "branch_id">;

