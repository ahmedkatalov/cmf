import React, { useState } from "react";
import { useGetBranchesQuery, useCreateBranchMutation, useDeleteBranchMutation } from "@/features/branch/api/branchApi";
import { Button } from "@mui/material";
import BranchList from "./BranchList.tsx";
import AddBranchModal from "./AddBranchModal.tsx";
import Loading from "@/shared/ui/Loading.tsx";

const BranchComponent: React.FC = () => {
  const { data: branches, isLoading, isError, refetch } = useGetBranchesQuery();
  const [createBranch] = useCreateBranchMutation();
  const [deleteBranch] = useDeleteBranchMutation();
  const [isOpen, setIsOpen] = useState(false);

  const handleCreate = async (payload: any) => {
    try {
      await createBranch(payload).unwrap();
      setIsOpen(false);
      refetch();
    } catch (e) {
      console.error(e);
    }
  };

  const handleDelete = async (id: string) => {
    try {
      await deleteBranch(id).unwrap();
      refetch();
    } catch (e) {
      console.error(e);
    }
  };

  return (
    <div className="p-6">
      <div className="flex items-center justify-between mb-4">
        <h2 className="text-2xl font-semibold">Филиалы</h2>
        <Button variant="contained" color="primary" onClick={() => setIsOpen(true)}>Добавить филиал</Button>
      </div>

      {isLoading && <Loading />}
      {isError && <div>Error loading branches</div>}

      {branches && <BranchList branches={branches} onDelete={handleDelete} />}

      <AddBranchModal open={isOpen} onClose={() => setIsOpen(false)} onCreate={handleCreate} />
    </div>
  );
};

export default BranchComponent;
