export interface Branch {
  id: string;
  organization_id: string;
  name: string;
  address?: string;
  [key: string]: any;
}

export interface CreateBranchDto {
 organization_id: string;
  name: string;
  address?: string;
}

