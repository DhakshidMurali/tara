import { SxProps } from "@mui/system";
const dateTypographyStyle: SxProps = {
  textAlign: "center",
  color: "#f3def5",
  fontWeight: "bolder",
  fontSize: 18,
  paddingLeft: 8,
  textWrap: "nowrap",
};

const communityTextTypographStyle: SxProps = {
  color: "#f3def5",
  fontWeight: "bolder",
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
  marginTop:"1%",
  borderRadius: 2,
  height: "80%",
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
  height: "90%",
  width: 232,
  backgroundColor: "secondary.light",
  display: "flex",
  alignItems: "center",
  justifyContent: "center",
  borderRadius: 2,
};
export const styles = {
  dateTypographyStyle,
  communityTextTypographStyle,
  createButtonStyle,
  toolsContainerBoxStyle,
  toolsContainerListBoxStyle,
};
