import { SxProps } from "@mui/system";
const dateTypographyStyle: SxProps = {
  textAlign: "center",
  color: "primary.contrastText",
  fontWeight: "bolder",
  fontSize: 18,
  paddingLeft: 8,
  textWrap: "nowrap",
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
  gridContainerCommnityBox
};



