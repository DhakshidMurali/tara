import { SxProps } from "@mui/system";
const departmentContainerListBoxStyle: SxProps = {
  borderRadius: 2,
  width: "100%",
  display: "flex",
  flexDirection: "column",
  justifyContent: "flex-start",
  alignItems: "center",
  height:"100%"
};

const departmentContainerListBoxStackStyle: SxProps = {
  width: "100%",
  display: "flex",
  flexDirection: "row",
  justifyContent: "space-between",
};

const departmentContainerListBoxStackExpandedStyle: SxProps = {
  width: "100%",
  height: "400px",
  display: "flex",
  flexDirection: "row",
  justifyContent: "space-between",
};
const departmentContainerListBoxStackTopicBoxStyle: SxProps = {
  backgroundColor: "secondary.main",
  padding: 1,
  borderRadius: 2,
  marginTop: 1.5,
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
  departmentContainerListBoxStackExpandedStyle,
  departmentContainerListBoxStackTopicBoxStyle,
  dashboardTextTypographStyle,
  deparmentListIconStyle,
  employeeByDepartmentListBoxBarStyle,
  employeeByDepartmentListBoxLabelTypographyStyle,
  employeeByDepartmentListBoxLabelBoxStyle,
  employeeByDepartmentListBoxBarBoxToolTipStyle,
};
