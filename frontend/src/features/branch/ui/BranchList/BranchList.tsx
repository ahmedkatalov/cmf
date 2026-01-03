import React from "react";
import type { Branch } from "../../types/types";
import {
  Table,
  TableBody,
  TableCell,
  TableContainer,
  TableHead,
  TableRow,
  Paper,
  IconButton,
} from "@mui/material";
import { useNavigate } from "react-router-dom";
import DeleteIcon from "@mui/icons-material/Delete";

type Props = {
  branches: Branch[];
  onDelete: (id: string) => void;
};

const BranchList: React.FC<Props> = ({ branches, onDelete }) => {
  const navigate = useNavigate();

  return (
    <TableContainer component={Paper} elevation={1}>
      <Table>
        <TableHead>
          <TableRow>
            <TableCell>Название</TableCell>
            <TableCell>Адрес</TableCell>
            <TableCell align="right">Действия</TableCell>
          </TableRow>
        </TableHead>

        <TableBody>
          {branches.map((b) => (
            <TableRow
              key={b.id}
              hover
              sx={{ cursor: "pointer" }}
              onClick={() => navigate(`/branches/${b.id}`)}
            >
              <TableCell>{b.name}</TableCell>
              <TableCell>{b.address || "-"}</TableCell>
              <TableCell align="right">
                <IconButton
                  color="error"
                  onClick={(e) => {
                    e.stopPropagation();
                    onDelete(b.id);
                  }}
                >
                  <DeleteIcon />
                </IconButton>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </TableContainer>
  );
};

export default BranchList;
