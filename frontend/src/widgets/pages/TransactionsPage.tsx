import React, { useState, useMemo } from "react";
import {
  Box,
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  TextField,
  Typography,
  Table,
  TableBody,
  TableCell,
  TableHead,
  TableRow,
  Paper,
  MenuItem,
  Select,
  FormControl,
  InputLabel,
} from "@mui/material";
import { useSelector } from "react-redux";
import { useGetTransactionsByBranchQuery, useCreateTransactionMutation } from "@/features/transactions/api/transactionApi";
import { useGetBranchesQuery } from "@/features/branch/api/branchApi";

/* ================== TYPES ================== */

interface Expense {
  id: string | number;
  category: "company" | "transfers";
  date: string; // YYYY-MM-DD
  amount: number;
}

/* ================== DATA ================== */

const categories = [
  { value: "", label: "Все категории" },
  { value: "company", label: "Расходы компании" },
  { value: "transfers", label: "Переводы" },
];

const months = [
  { value: "", label: "Все месяцы" },
  { value: "01", label: "Январь" },
  { value: "02", label: "Февраль" },
  { value: "03", label: "Март" },
  { value: "04", label: "Апрель" },
  { value: "05", label: "Май" },
  { value: "06", label: "Июнь" },
  { value: "07", label: "Июль" },
  { value: "08", label: "Август" },
  { value: "09", label: "Сентябрь" },
  { value: "10", label: "Октябрь" },
  { value: "11", label: "Ноябрь" },
  { value: "12", label: "Декабрь" },
];

/* ================== COMPONENT ================== */

const TransactionsPage: React.FC = () => {
  const [open, setOpen] = useState(false);

  const [newExpense, setNewExpense] = useState({
    category: "" as Expense["category"] | "",
    date: "",
    amount: 0,
  });

  const [filterYear, setFilterYear] = useState("");
  const [filterMonth, setFilterMonth] = useState("");
  const [filterCategory, setFilterCategory] = useState("");

  const selectedBranchFromStore = useSelector((s: any) => s.branch?.selected);
  const { data: branches } = useGetBranchesQuery();
  const branch = selectedBranchFromStore ?? branches?.[0];
  const orgID = branch?.organization_id ?? "";
  const branchID = branch?.id ?? "";

  const { data: transactions = [], isLoading: txLoading, isError: txError, error: txErrorObj } = useGetTransactionsByBranchQuery(
    { orgID, branchID },
    { skip: !orgID || !branchID }
  );

  const [createTransaction] = useCreateTransactionMutation();

  /* ================== FILTERS ================== */

  const filteredExpenses = useMemo(() => {
    // map fetched transactions to local shape for filtering/display
    const mapped = transactions.map((t: any) => ({
      id: t.id,
      category: t.type === "expense_company" ? "company" : "transfers",
      date: t.occurred_at?.slice(0, 10) ?? t.created_at?.slice(0, 10) ?? "",
      amount: t.amount,
    })) as Expense[];

    return mapped.filter((e) => {
      if (filterYear && !e.date.startsWith(filterYear)) return false;
      if (filterMonth && e.date.slice(5, 7) !== filterMonth) return false;
      if (filterCategory && e.category !== filterCategory) return false;
      return true;
    });
  }, [transactions, filterYear, filterMonth, filterCategory]);

  const total = useMemo(
    () => filteredExpenses.reduce((acc, curr) => acc + curr.amount, 0),
    [filteredExpenses]
  );

  const years = Array.from(new Set((transactions || []).map((t: any) => (t.occurred_at || "").slice(0, 4)))).filter(Boolean);

  /* ================== HANDLERS ================== */

  const handleAddExpense = async () => {
    if (!newExpense.category || !newExpense.date || !newExpense.amount) return;
    if (!orgID || !branchID) return;

    const txType = newExpense.category === "company" ? "expense_company" : "expense_people";

    try {
      await createTransaction({
        organization_id: orgID,
        branch_id: branchID,
        type: txType,
        amount: newExpense.amount,
        occurred_at: newExpense.date,
        description: null,
      }).unwrap?.();
    } catch (err) {
      // ignore error here — could show notification
    }

    setNewExpense({ category: "", date: "", amount: 0 });
    setOpen(false);
  };

  /* ================== RENDER ================== */

  return (
    <Box p={4} bgcolor="#f5f5f5" minHeight="100vh">
      <Typography variant="h4" fontWeight="bold" mb={4}>
        Расходы
      </Typography>

      {/* ===== Summary + Filters ===== */}
      <Paper sx={{ p: 4, mb: 6, borderRadius: 3 }}>
        <Box display="flex" justifyContent="space-between" flexWrap="wrap" gap={2}>
          <Box>
            <Typography color="text.secondary">Общие расходы</Typography>
            <Typography variant="h5" color="error" fontWeight="bold">
              {total.toLocaleString()} ₽
            </Typography>
          </Box>

          <Box display="flex" gap={2} flexWrap="wrap">
            <FormControl size="small" className="w-25">
              <InputLabel>Год</InputLabel>
              <Select
                value={filterYear}
                label="Год"
                onChange={(e) => setFilterYear(e.target.value)}
              >
                <MenuItem value="">Все годы</MenuItem>
                {years.map((y) => (
                  <MenuItem key={y} value={y}>
                    {y}
                  </MenuItem>
                ))}
              </Select>
            </FormControl>

            <FormControl size="small" className="w-25">
              <InputLabel>Месяц</InputLabel>
              <Select
                value={filterMonth}
                label="Месяц"
                onChange={(e) => setFilterMonth(e.target.value)}
              >
                {months.map((m) => (
                  <MenuItem key={m.value} value={m.value}>
                    {m.label}
                  </MenuItem>
                ))}
              </Select>
            </FormControl>

            <FormControl size="small" className="w-40">
              <InputLabel>Категория</InputLabel>
              <Select
                value={filterCategory}
                label="Категория"
                onChange={(e) => setFilterCategory(e.target.value)}
              >
                {categories.map((c) => (
                  <MenuItem key={c.value} value={c.value}>
                    {c.label}
                  </MenuItem>
                ))}
              </Select>
            </FormControl>

            <Button
              variant="contained"
              onClick={() => setOpen(true)}
              sx={{ borderRadius: 2, textTransform: "none", height: 40 }}
            >
              Добавить расход
            </Button>
          </Box>
        </Box>
      </Paper>

      {/* ===== Table ===== */}
      <Paper sx={{ borderRadius: 3, overflow: "hidden" }}>
        <Table>
          <TableHead>
            <TableRow sx={{ bgcolor: "#f0f0f0" }}>
              <TableCell>Дата</TableCell>
              <TableCell>Категория</TableCell>
              <TableCell align="right">Сумма</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {filteredExpenses.map((e) => (
              <TableRow key={e.id} hover>
                <TableCell>{e.date}</TableCell>
                <TableCell>
                  {categories.find((c) => c.value === e.category)?.label}
                </TableCell>
                <TableCell align="right" sx={{ color: "red", fontWeight: 500 }}>
                  {e.amount.toLocaleString()} ₽
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </Paper>

      {/* ===== Modal ===== */}
      <Dialog open={open} onClose={() => setOpen(false)} fullWidth maxWidth="sm">
        <DialogTitle>Добавить расход</DialogTitle>
        <DialogContent sx={{ display: "flex", flexDirection: "column", gap: 2 }}>
          <FormControl fullWidth>
            <InputLabel>Категория</InputLabel>
            <Select
              value={newExpense.category}
              label="Категория"
              onChange={(e) =>
                setNewExpense((prev) => ({
                  ...prev,
                  category: e.target.value as Expense["category"],
                }))
              }
            >
              {categories
                .filter((c) => c.value)
                .map((c) => (
                  <MenuItem key={c.value} value={c.value}>
                    {c.label}
                  </MenuItem>
                ))}
            </Select>
          </FormControl>

          <TextField
            label="Дата"
            type="date"
            value={newExpense.date}
            onChange={(e) =>
              setNewExpense((prev) => ({ ...prev, date: e.target.value }))
            }
            InputLabelProps={{ shrink: true }}
          />

          <TextField
            label="Сумма"
            type="number"
            value={newExpense.amount}
            onChange={(e) =>
              setNewExpense((prev) => ({
                ...prev,
                amount: Number(e.target.value),
              }))
            }
          />
        </DialogContent>

        <DialogActions>
          <Button onClick={() => setOpen(false)}>Отмена</Button>
          <Button variant="contained" onClick={handleAddExpense}>
            Добавить
          </Button>
        </DialogActions>
      </Dialog>
    </Box>
  );
};

export default TransactionsPage;
