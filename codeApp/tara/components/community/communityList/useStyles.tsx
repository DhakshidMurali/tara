import { SxProps } from "@mui/system";

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

const toolsContainerBoxStyle: SxProps = {
  overflowX: "auto",
  display: "flex",
  borderRadius: 2,
  flexWrap: "nowrap",
  padding: 2,
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
  height: 168,
  width: 232,
  justifyContent: "space-between",
};

const toolsContainerListBoxGridTypograghyCommunityNameStyle: SxProps = {
  color: "primary.contrastText",
  paddingLeft: 1,
  paddingTop: 0.5,
};

const toolsContainerListBoxGridBoxStyle: SxProps = {
  display: "flex",
  flexDirection: "row",
  paddingLeft: 1,
};
const toolsContainerListBoxGridBoxTypographyParticipantsStyle: SxProps = {
  display: "flex",
  flexDirection: "row",
  fontWeight: "bold",
  paddingLeft: 1,
  color: "primary.contrastText",
};

const toolsContainerListBoxGridIncreaseInLastMonthBoxStyle: SxProps = {
  backgroundColor: "rgb(177,248,180)",
  display: "flex",
  justifyContent: "center",
  alignContent: "center",
  borderRadius: 3,
};

const toolsContainerListBoxGridIncreaseInLastMonthBoxTypograghyStyle: SxProps =
  {
    color: "rgb(33,24,43)",
    justifyContent: "flex-end",
    fontWeight: "bolder",
  };

export const styles = {
  communityTextTypographStyle,
  createButtonStyle,
  toolsContainerBoxStyle,
  toolsContainerListBoxGridStyle,
  toolsContainerListBoxGridTypograghyCommunityNameStyle,
  toolsContainerListBoxGridBoxStyle,
  toolsContainerListBoxGridBoxTypographyParticipantsStyle,
  toolsContainerListBoxGridIncreaseInLastMonthBoxStyle,
  toolsContainerListBoxGridIncreaseInLastMonthBoxTypograghyStyle,
};
