import { SxProps } from "@mui/system";
const departmentContainerListBoxStyle: SxProps = {
  borderRadius: 2,
  height: 72,
  width: "90%",
  display: "flex",
  justifyContent: "center",
  alignItems: "center",
};

const dashboardTextTypographStyle: SxProps = {
  color: "#f3def5",
  fontWeight: "bolder",
};

const deparmentListIconStyle:SxProps={
    color:"secondary.main",
}


export const styles={
    departmentContainerListBoxStyle,
    dashboardTextTypographStyle,
    deparmentListIconStyle
}