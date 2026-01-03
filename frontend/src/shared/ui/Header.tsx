import React from "react";
import { logout } from "@/features/auth/model/authSlice";
import { useDispatch } from "react-redux";
import { Link as RouterLink } from "react-router-dom";
import {
  AppBar,
  Toolbar,
  Typography,
  Button,
  Box,
} from "@mui/material";

const Header: React.FC = () => {
  const dispatch = useDispatch();

  const logoutHandler = () => {
    dispatch(logout());
  };

  return (
    <AppBar
      position="static"
      elevation={1}
      color="inherit"
    >
      <Toolbar sx={{ maxWidth: 1200, mx: "auto", width: "100%" }}>
        {/* Logo */}
        <Typography
          variant="h6"
          component={RouterLink}
          to="/"
          sx={{
            textDecoration: "none",
            color: "primary.main",
            fontWeight: 600,
            mr: 4,
          }}
        >
          CMF
        </Typography>

        {/* Navigation */}
        <Box sx={{ display: "flex", gap: 3, flexGrow: 1 }}>
          <Button
            component={RouterLink}
            to="/branches"
            color="inherit"
          >
            Филиалы
          </Button>

          <Button
            component={RouterLink}
            to="/transactions"
            color="inherit"
          >
            Транзакции
          </Button>
        </Box>

        {/* Actions */}
        <Button
          variant="contained"
          color="primary"
          onClick={logoutHandler}
        >
          Выход
        </Button>
      </Toolbar>
    </AppBar>
  );
};

export default Header;
