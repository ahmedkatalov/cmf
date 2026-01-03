import { useParams } from "react-router-dom";
import { skipToken } from "@reduxjs/toolkit/query";
import { useGetBranchQuery } from "../../api/branchApi";
import Loading from "@/shared/ui/Loading";
import {
  Box,
  Paper,
  Typography,
  Divider,
  Button,
} from "@mui/material";
import { useState } from "react";
import CreateUserModal from "./CreateUserModal";

const BranchDetailComponent = () => {
  const { id } = useParams<{ id: string }>();

  const {
    data: branch,
    isLoading,
    isError,
  } = useGetBranchQuery(id ?? skipToken);

  const [isModalOpen, setIsModalOpen] = useState(false);

  if (isLoading) return <Loading />;
  if (isError) return <Typography color="error">Ошибка загрузки филиала</Typography>;
  if (!branch) return null;

return (
    <Box sx={{ p: 4 }}>
      <Typography variant="h5" fontWeight={600} mb={2}>
        Филиал
      </Typography>

      <Paper sx={{ p: 3, mb: 2 }}>
        <Typography variant="h6" fontWeight={500}>
          {branch.name}
        </Typography>
        <Divider sx={{ my: 2 }} />
        <Typography color="text.secondary">{branch.address ?? "Адрес не указан"}</Typography>
      </Paper>

      <Button variant="contained" onClick={() => setIsModalOpen(true)}>
        Добавить пользователя
      </Button>

      <CreateUserModal
        open={isModalOpen}
        onClose={() => setIsModalOpen(false)}
        branchId={id as string}
      />
    </Box>
  );
};

export default BranchDetailComponent;
