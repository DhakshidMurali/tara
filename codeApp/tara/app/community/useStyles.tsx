import { SxProps } from "@mui/system";
const dateTypographyStyle: SxProps = {
  textAlign: "center",
  color: "primary.contrastText",
  fontWeight: "bolder",
  fontSize: 18,
  paddingLeft: 8,
  textWrap: "nowrap",
};
const communityTextTypographStyle: SxProps = {
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
  "& :hover": {
    backgroundColor: "red"
  }
};

const gridContainerCommnityBox: SxProps = {
  height: "70%",
};

export const styles = {
  dateTypographyStyle,
  gridContainerCommnityBox,
  communityTextTypographStyle,
  createButtonStyle,
};
