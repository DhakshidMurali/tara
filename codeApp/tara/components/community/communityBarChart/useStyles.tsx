import { SxProps } from "@mui/system";

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
const communityTextTypographStyle: SxProps = {
  color: "primary.contrastText",
  fontWeight: "bolder",
};
export const styles = {
  communityTextTypographStyle,
  toolsContainerListBoxGridSelectGridItemTypographyStyle,
  toolsContainerListBoxGridSelectGridItemSplittingBoxStyle,
  toolsContainerListBoxGridSelectGridItemGridItemStyle,
  toolsContainerListBoxGridSelectGridItemGridItemStackStyle,
  toolsContainerListBoxGridSelectGridItemGridItemStackTypographyListStyle,
  toolsContainerListBoxGridSelectGridItemGridItemStackTypographyCommunityListStyle,
  toolsContainerListBoxGridSelectGridItemGridItemStackCommunityListStyle,
};
