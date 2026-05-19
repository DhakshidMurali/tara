import { SxProps } from "@mui/material";

const toolsContainerBoxStyle: SxProps = {
  overflowX: "auto",
  display: "flex",
  borderRadius: 2,
  flexWrap: "nowrap",
  padding: 1,
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
const toolsContainerListBoxGridStyle: SxProps = {
  backgroundColor: "secondary.light",
  display: "flex",
  alignItems: "center",
  borderRadius: 2,
  height: "12rem",
  width: "16rem",
  justifyContent: "space-between",
  marginLeft: 1,
  marginRight: 1
};

const toolsContainerListBoxGridCommunicationCountTypography: SxProps = {
  fontSize: "18px",
  color: "secondary.main",
  fontWeight: "bold"
};
const toolsContainerListBoxGridToolNameTypography: SxProps = {
  color: "primary.contrastText",
  paddingLeft: 1,
  paddingTop: 0.5,
};

export const styles = {
  toolsContainerBoxStyle,
  toolsContainerListBoxGridStyle,
  toolsContainerListBoxGridCommunicationCountTypography,
  toolsContainerListBoxGridToolNameTypography,
};
