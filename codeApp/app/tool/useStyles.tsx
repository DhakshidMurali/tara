import { SxProps } from "@mui/material";

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
const toolTextTypographStyle: SxProps = {
  color: "primary.contrastText",
  fontWeight: "bolder",
  marginBottom: "8px",
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
export const styles = {
  searchTextFieldInputPropsStyle,
  dateTypographyStyle,
  toolTextTypographStyle,
  createButtonStyle,
};
