import React from "react";
import { useRegisterMutation } from "../api/authApi";
import { useNavigate, Navigate, Link as RouterLink } from "react-router-dom";
import { useSelector } from "react-redux";
import type { RootState } from "@/app/store/store";
import { useForm, Controller } from "react-hook-form";
import type { SubmitHandler } from "react-hook-form";
import {
  Box,
  Paper,
  Typography,
  TextField,
  Button,
  CircularProgress,
  Stack,
  Link,
} from "@mui/material";

type FormValues = {
  organizationName: string;
  email: string;
  password: string;
};

const RegisterComponent: React.FC = () => {
  const isAuthenticated = useSelector((s: RootState) => s.auth.isAuthenticated);
  const [register, { isLoading }] = useRegisterMutation();
  const navigate = useNavigate();

  const {
    handleSubmit,
    control,
    formState: { errors },
  } = useForm<FormValues>({
    defaultValues: {
      organizationName: "",
      email: "",
      password: "",
    },
  });

  const onSubmit: SubmitHandler<FormValues> = async (data) => {
    try {
      await register({
        email: data.email,
        password: data.password,
        organization_name: data.organizationName,
      }).unwrap();
      navigate("/", { replace: true });
    } catch (err) {
      console.error(err);
    }
  };

  if (isAuthenticated) return <Navigate to="/" replace />;

  return (
    <Box
      sx={{
        minHeight: "100vh",
        display: "flex",
        alignItems: "center",
        justifyContent: "center",
        bgcolor: "grey.100",
        p: 2,
      }}
    >
      <Paper sx={{ p: 4, maxWidth: 400, width: "100%", boxShadow: 3 }}>
        <Typography variant="h5" fontWeight={600} mb={3} textAlign="center">
          Регистрация
        </Typography>

        <form onSubmit={handleSubmit(onSubmit)}>
          <Stack spacing={2}>
            {/* Organization Name */}
            <Controller
              name="organizationName"
              control={control}
              rules={{ required: "Имя организации обязательно" }}
              render={({ field }) => (
                <TextField
                  {...field}
                  label="Имя организации"
                  fullWidth
                  error={!!errors.organizationName}
                  helperText={errors.organizationName?.message}
                />
              )}
            />

            {/* Email */}
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
                  type="email"
                  fullWidth
                  error={!!errors.email}
                  helperText={errors.email?.message}
                />
              )}
            />

            {/* Password */}
            <Controller
              name="password"
              control={control}
              rules={{ required: "Пароль обязателен" }}
              render={({ field }) => (
                <TextField
                  {...field}
                  label="Пароль"
                  type="password"
                  fullWidth
                  error={!!errors.password}
                  helperText={errors.password?.message}
                />
              )}
            />

            {/* Submit Button */}
            <Button
              type="submit"
              variant="contained"
              color="primary"
              fullWidth
              disabled={isLoading}
              startIcon={isLoading ? <CircularProgress size={20} /> : null}
            >
              Зарегистрироваться
            </Button>
          </Stack>
        </form>

        <Box mt={2} textAlign="center">
          <Typography variant="body2" color="textSecondary">
            Уже есть аккаунт?{" "}
            <Link component={RouterLink} to="/login">
              Войти
            </Link>
          </Typography>
        </Box>
      </Paper>
    </Box>
  );
};

export default RegisterComponent;
