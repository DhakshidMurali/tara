import { SxProps } from "@mui/material";

const toolsContainerBoxStyle: SxProps = {
  overflowX: "auto",
  display: "flex",
  padding: 1,
  borderRadius: 2,
  flexWrap: "nowrap",
  height: "inherit",
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
  height: "12rem",
  width: "16rem",
  backgroundColor: "secondary.light",
  display: "flex",
  alignItems: "center",
  justifyContent: "center",
  borderRadius: 2,
  marginLeft: 1,
  marginRight: 1
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
  toolsContainerBoxStyle,
  toolsContainerListBoxStyle,
  toolsContainerListBoxStackStyle,
  toolsContainerListBoxBoxStyle,
  toolsContainerListBoxCountTypographyStyle,
  toolsContainerListBoxDomainTypographyStyle,
};
