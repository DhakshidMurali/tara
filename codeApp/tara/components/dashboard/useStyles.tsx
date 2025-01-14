import { SxProps } from "@mui/system";
const departmentContainerListBoxStyle: SxProps = {
  borderRadius: 2,
  height: 72,
  width: "100%",
  display: "flex",
  justifyContent: "flex-start",
  alignItems: "center",
  marginTop: "8px",
};

const departmentContainerListBoxStackStyle: SxProps = {
  width: "100%",
  display: "flex",
  flexDirection: "row",
  justifyContent: "space-between",
};

const dashboardTextTypographStyle: SxProps = {
  color: "#f3def5",
  fontWeight: "bolder",
};

const deparmentListIconStyle: SxProps = {
  color: "secondary.main",
};

const employeeByDepartmentListBoxBarStyle: SxProps = {
  height: "8px",
  borderRadius: 2,
};

const employeeByDepartmentListBoxLabelTypographyStyle: SxProps = {
  color: "#f3def5",
  fontWeight: "bolder",
  // textWrap:"nowrap"
};

const employeeByDepartmentListBoxLabelBoxStyle: SxProps = {
  height: "12px",
  borderRadius: 2,
  width: "20%",
  marginRight: "8px",
};

const employeeByDepartmentListBoxBarBoxToolTipStyle: SxProps = {
  color: "#fff",
  fontSize: "0.8rem",
  padding: "4px 8px",
  borderRadius: "8px",
};

export const styles = {
  departmentContainerListBoxStyle,
  departmentContainerListBoxStackStyle,
  dashboardTextTypographStyle,
  deparmentListIconStyle,
  employeeByDepartmentListBoxBarStyle,
  employeeByDepartmentListBoxLabelTypographyStyle,
  employeeByDepartmentListBoxLabelBoxStyle,
  employeeByDepartmentListBoxBarBoxToolTipStyle,
};
