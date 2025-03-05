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
};

const gridContainerCommnityBox: SxProps = {
  overflowY: "auto",
  height: "70%",
  "&::-webkit-scrollbar": {
    height: "8px",
    width: "6px"
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

export const styles = {
  dateTypographyStyle,
  gridContainerCommnityBox,
  communityTextTypographStyle,
  createButtonStyle
};



