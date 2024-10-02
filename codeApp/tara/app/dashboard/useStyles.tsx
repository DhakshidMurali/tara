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
  backgroundColor: "primary.light",
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
  backgroundColor: "secondary.light",
  display: "flex",
  alignItems: "center",
  justifyContent: "center",
  borderRadius: 2,
};

const employerByDepartmentBoxStyle: SxProps = {
  borderRadius: 4,
  backgroundColor: "secondary.light",
};

const departmentContainerBoxStyle: SxProps = {
  backgroundColor: "secondary.light",
  height: 548,
  padding: 4,
  borderRadius: 4,
  marginRight: 4,
};

const dashboardTextTypographStyle: SxProps = {
  color: "#f3def5",
  fontWeight: "bolder",
};

const departmentContainerListDividerBoxStyle: SxProps = {
  height: 1.25,
  width: "70%",
  bgcolor: "secondary.main",
  alignSelf: "center",
};


export const styles = {
  searchTextFieldInputPropsStyle,
  dateTypographyStyle,
  createButtonStyle,
  toolsContainerBoxStyle,
  toolsContainerListBoxStyle,
  employerByDepartmentBoxStyle,
  departmentContainerBoxStyle,
  departmentContainerListDividerBoxStyle,
  dashboardTextTypographStyle
};
