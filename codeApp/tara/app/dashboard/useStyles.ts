import { SxProps } from "@mui/system";
/*

MUI component accept an sx as parameter - We can assign constant (sxprops object which have css style properties) to it

*/

const searchTextFieldInputPropsStyle: SxProps = {
  color: "#f3def5",
  fontWeight: "bolder",
  fontSize: 16,
};

const dateTypographyStyle: SxProps = {
  textAlign: "center",
  color: "#f3def5",
  fontWeight: "bolder",
  fontSize: 18,
  paddingLeft: 8,
  textWrap: "nowrap",
};

const createButtonStyle: SxProps = {
  paddingRight: 2,
  marginRight: 4,
  backgroundColor: "green",
  display: "flex",
  alignItems: "center",
  justifyContent: "center",
  width: 128,
};

const toolsContainerBoxStyle: SxProps = {
  overflowX: "auto",
  display: "flex",
  padding: 2,
};
const toolsContainerListBoxStyle: SxProps = {
  height: 168,
  width: 232,
  backgroundColor: "primary.main",
  display: "flex",
  alignItems: "center",
  justifyContent: "center",
  borderRadius: 2,
};

const employerByDepartmentBoxStyle: SxProps = {
  borderRadius: 4,
};

const departmentContainerBoxStyle: SxProps = {
  backgroundColor: "purple",
  height: 548,
  padding: 4,
  borderRadius: 4,
  marginRight: 4,
};

const departmentContainerListBoxStyle: SxProps = {
  borderRadius: 2,
  height: 72,
  width: "90%",
  display: "flex",
  justifyContent: "center",
  alignItems: "center",
};

const departmentContainerListDividerBoxStyle: SxProps = {
  height: 1.25,
  width: "70%",
  bgcolor: "red",
  alignSelf: "center",
};

const dashboardTextTypographStyle: SxProps = {
  color: "#f3def5",
  fontWeight: "bolder",
};

export const styles = {
  searchTextFieldInputPropsStyle,
  dateTypographyStyle,
  createButtonStyle,
  toolsContainerBoxStyle,
  toolsContainerListBoxStyle,
  employerByDepartmentBoxStyle,
  departmentContainerBoxStyle,
  departmentContainerListBoxStyle,
  departmentContainerListDividerBoxStyle,
  dashboardTextTypographStyle,
};
