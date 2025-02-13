import { getRandomColor } from "@/utils/functions";
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
  borderRadius: 2,
  "&::-webkit-scrollbar": {
    height: "8px",
  },
  "&::-webkit-scrollbar-track": {
    backgroundColor: "secondary.main",
  },
  "&::-webkit-scrollbar-thumb": {
    backgroundColor: "secondary.light",
    borderRadius: "8px",
  },
  "&::-webkit-scrollbar-thumb:hover": {
    background: "#555",
  },
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
  overflowY: "auto",
  display: "flex",
};

const employerByDepartmenyBoxStackStyle: SxProps = {
  overflowX: "auto",
  display: "flex",
};

const departmentContainerBoxStyle: SxProps = {
  backgroundColor: "secondary.light",
  height: 548,
  padding: 4,
  borderRadius: 4,
  marginRight: 4,
};

const dashboardTextTypographStyle: SxProps = {
  color:  "primary.contrastText",
  fontWeight: "bolder",
  marginBottom: "8px",
};

const departmentContainerListDividerBoxStyle: SxProps = {
  height: 1.25,
  width: "70%",
  bgcolor: "secondary.main",
  alignSelf: "center",
};

const toolsContainerListBoxStackStyle: SxProps = {
  height: "100%",
  width: "100%",
};

const toolsContainerListBoxBoxStyle: SxProps = {
  width: "2%",
  height: "70%",
  borderRadius: "4px",
  marginLeft: 2,
  marginRight: 1.5,
};

const toolsContainerListBoxCountTypographyStyle: SxProps = {
  color: "#f3def5",
  fontWeight: "bolder",
};
const toolsContainerListBoxDomainTypographyStyle: SxProps = {
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
  employerByDepartmenyBoxStackStyle,
  departmentContainerBoxStyle,
  departmentContainerListDividerBoxStyle,
  dashboardTextTypographStyle,
  toolsContainerListBoxStackStyle,
  toolsContainerListBoxBoxStyle,
  toolsContainerListBoxCountTypographyStyle,
  toolsContainerListBoxDomainTypographyStyle,
};
