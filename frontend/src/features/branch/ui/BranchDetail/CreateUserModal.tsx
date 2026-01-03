import React from "react";
import {
  Modal,
  Box,
  Typography,
  TextField,
  Button,
  Stack,
  MenuItem,
  CircularProgress,
  IconButton,
} from "@mui/material";
import CloseIcon from "@mui/icons-material/Close";
import { useForm, Controller } from "react-hook-form";
import { useCreateUserMutation } from "@/features/user/api/userApi";
import type { SubmitHandler } from "react-hook-form";
import type { CreateUserFormValues } from "../../types/types";

type Props = {
  open: boolean;
  onClose: () => void;
  branchId: string;
};

const roles = [
  { value: "admin", label: "Admin" },
  { value: "user", label: "User" },
];

const CreateUserModal: React.FC<Props> = ({ open, onClose, branchId }) => {
  const [createUser, { isLoading }] = useCreateUserMutation();

  const {
    handleSubmit,
    control,
    formState: { errors },
  } = useForm<CreateUserFormValues>({
    defaultValues: { email: "", password: "", role: "user" },
  });

  const onSubmit: SubmitHandler<CreateUserFormValues> = async (data) => {
    await createUser({ branch_id: branchId, ...data }).unwrap();
  };

  return (
    <Modal open={open} onClose={onClose} closeAfterTransition>
      <Box
        sx={{
          position: "absolute" as const,
          top: "50%",
          left: "50%",
          transform: "translate(-50%, -50%)",
          width: 400,
          bgcolor: "background.paper",
          borderRadius: 2,
          boxShadow: 24,
          p: 4,
        }}
      >
        <Box
          display="flex"
          justifyContent="space-between"
          alignItems="center"
          mb={2}
        >
          <Typography variant="h6" fontWeight={600}>
            Создать пользователя
          </Typography>
          <IconButton onClick={onClose} size="small">
            <CloseIcon />
          </IconButton>
        </Box>

        <form onSubmit={handleSubmit(onSubmit)}>
          <Stack spacing={2}>
            <Controller
              name="email"
              control={control}
              rules={{
                required: "Email обязателен",
                pattern: {
                  value: /^[\w-\.]+@([\w-]+\.)+[\w-]{2,4}$/,
                  message: "Неверный email",
                },
              }}
              render={({ field }) => (
                <TextField
                  {...field}
                  label="Email"
                  fullWidth
                  error={!!errors.email}
                  helperText={errors.email?.message}
                />
              )}
            />

            <Controller
              name="password"
              control={control}
              rules={{ required: "Пароль обязателен" }}
              render={({ field }) => (
                <TextField
                  {...field}
                  type="password"
                  label="Пароль"
                  fullWidth
                  error={!!errors.password}
                  helperText={errors.password?.message}
                />
              )}
            />

            <Controller
              name="role"
              control={control}
              rules={{ required: "Роль обязательна" }}
              render={({ field }) => (
                <TextField
                  {...field}
                  select
                  label="Роль"
                  fullWidth
                  error={!!errors.role}
                  helperText={errors.role?.message}
                >
                  {roles.map((r) => (
                    <MenuItem key={r.value} value={r.value}>
                      {r.label}
                    </MenuItem>
                  ))}
                </TextField>
              )}
            />

            <Button
              type="submit"
              variant="contained"
              color="primary"
              fullWidth
              disabled={isLoading}
              startIcon={isLoading ? <CircularProgress size={20} /> : null}
            >
              Создать
            </Button>
          </Stack>
        </form>
      </Box>
    </Modal>
  );
};

export default CreateUserModal;
