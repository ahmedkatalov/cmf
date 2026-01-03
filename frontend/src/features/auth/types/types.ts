export interface AuthClaims {
	user_id: string;
	org_id: string;
	branch_id: string;
	role: string;
	email: string;
	exp: number;
}

export interface AuthToken {
	claims: AuthClaims;
}