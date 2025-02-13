import { SxProps } from "@mui/material";

const dashboardTextTypographStyle: SxProps = {
  color: "#f3def5",
  fontWeight: "bolder",
  marginBottom: "8px",
};

const employeeByDepartmentListBoxBarBoxToolTipStyle: SxProps = {
  color: "#fff",
  fontSize: "0.8rem",
  padding: "4px 8px",
  borderRadius: "8px",
};
const employeeByDepartmentListBoxBarStyle: SxProps = {
  height: "8px",
  borderRadius: 2,
};

const employerByDepartmentBoxStyle: SxProps = {
  borderRadius: 4,
  backgroundColor: "secondary.light",
  overflowY: "auto",
  display: "flex",
};


const employeeByDepartmentListBoxLabelBoxStyle: SxProps = {
  height: "12px",
  borderRadius: 2,
  width: "15%",
  marginRight: "8px",
};

const employeeByDepartmentListBoxLabelTypographyStyle: SxProps = {
  color: "#f3def5",
  fontWeight: "bolder",
};
export const styles = {
  employerByDepartmentBoxStyle,
  dashboardTextTypographStyle,
  employeeByDepartmentListBoxBarBoxToolTipStyle,
  employeeByDepartmentListBoxBarStyle,
  employeeByDepartmentListBoxLabelBoxStyle,
  employeeByDepartmentListBoxLabelTypographyStyle
};