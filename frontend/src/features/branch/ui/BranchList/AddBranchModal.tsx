import React, { useState } from "react";
import Input from "@/shared/ui/Input";
import { Button } from "@mui/material";
import type { CreateBranchDto } from "../../types/types";

type Props = {
  open: boolean;
  onClose: () => void;
  onCreate: (payload: CreateBranchDto) => void;
};

const AddBranchModal: React.FC<Props> = ({ open, onClose, onCreate }) => {
  const [name, setName] = useState("");
  const [address, setAddress] = useState("");

  if (!open) return null;

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    onCreate({ name, address });
  };

  return (
    <div className="fixed inset-0 bg-black/40 flex items-center justify-center">
      <div className="bg-white rounded p-6 w-full max-w-md">
        <h3 className="text-lg font-semibold mb-4">Добавить филиал</h3>
        <form onSubmit={handleSubmit} className="space-y-3">
          <Input label="Имя" value={name} onChange={(e) => setName(e.target.value)} />
          <Input label="Адрес" value={address} onChange={(e) => setAddress(e.target.value)} />
          <div className="flex items-center justify-end gap-2 mt-4">
            <Button variant="outlined" type="button" onClick={onClose}>
              Отмена
            </Button>
            <Button variant="contained" type="submit">Создать</Button>
          </div>
        </form>
      </div>
    </div>
  );
};

export default AddBranchModal;
