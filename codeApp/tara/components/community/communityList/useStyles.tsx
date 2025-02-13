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
  marginTop: "1%",
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

const toolsContainerListBoxGridStyle: SxProps = {
  height: "100%",
  display: "flex",
  justifyContent: "space-between",
};

const toolsContainerListBoxGridTypograghyCommunityNameStyle: SxProps = {
  color: "#f3def5",
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
  paddingLeft: 1,
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
  toolsContainerListBoxStyle,
  toolsContainerListBoxGridStyle,
  toolsContainerListBoxGridTypograghyCommunityNameStyle,
  toolsContainerListBoxGridBoxStyle,
  toolsContainerListBoxGridBoxTypographyParticipantsStyle,
  toolsContainerListBoxGridIncreaseInLastMonthBoxStyle,
  toolsContainerListBoxGridIncreaseInLastMonthBoxTypograghyStyle,
};
