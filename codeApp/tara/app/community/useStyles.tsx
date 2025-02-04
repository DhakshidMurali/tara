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

const toolsContainerListBoxGridSelectGridItemStyle: SxProps = {
  display: "flex",
  justifyContent: "center",
  alignContent: "center",
};

const toolsContainerListBoxGridSelectGridItemTypographyStyle: SxProps = {
  color: "#f3def5",
  paddingLeft: 1,
  margin: "auto",
};

const toolsContainerListBoxGridSelectGridItemSplittingBoxStyle: SxProps = {
  width: "1.1%",
  backgroundColor: "blue",
  height: "100%",
  borderRadius: 2,
};
const toolsContainerListBoxGridSelectGridItemGridItemStyle: SxProps = {
  height: "90%",
  backgroundColor: "green",
  borderRadius: 2,
  display: "flex",
  justifyContent: "center",
  alignContent: "center",
};
const toolsContainerListBoxGridSelectGridItemGridItemStackStyle: SxProps = {
  height: "16%",
  backgroundColor: "orange",
  borderTopLeftRadius: "16px",
};
const toolsContainerListBoxGridSelectGridItemGridItemStackTypographyListStyle: SxProps =
  {
    color: "white",
    margin: "auto",
  };
const toolsContainerListBoxGridSelectGridItemGridItemStackTypographyCommunityListStyle: SxProps =
  {
    color: "white",
    paddingLeft: 0.1,
  };
const toolsContainerListBoxGridSelectGridItemGridItemStackCommunityListStyle: SxProps =
  {
    overflowX: "auto",
    borderRadius: 2,
    height: "100%",
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
export const styles = {
  dateTypographyStyle,
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
  toolsContainerListBoxGridSelectGridItemStyle,
  toolsContainerListBoxGridSelectGridItemTypographyStyle,
  toolsContainerListBoxGridSelectGridItemSplittingBoxStyle,
  toolsContainerListBoxGridSelectGridItemGridItemStyle,
  toolsContainerListBoxGridSelectGridItemGridItemStackStyle,
  toolsContainerListBoxGridSelectGridItemGridItemStackTypographyListStyle,
  toolsContainerListBoxGridSelectGridItemGridItemStackTypographyCommunityListStyle,
  toolsContainerListBoxGridSelectGridItemGridItemStackCommunityListStyle
};
