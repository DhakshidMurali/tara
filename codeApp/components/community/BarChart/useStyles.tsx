import { SxProps } from "@mui/system";

const toolsContainerListBoxGridSelectGridItemTypographyStyle: SxProps = {
  color: "#f3def5",
  paddingLeft: 1,
  margin: "auto",
};

const toolsContainerListBoxGridSelectGridItemSplittingBoxStyle: SxProps = {
  width: "1.1%",
  backgroundColor: "rgb(174,199,243)",
  borderRadius: 2,
};
const toolsContainerListBoxGridSelectGridItemGridItemStyle: SxProps = {
  backgroundColor: "secondary.dark",
  borderRadius: 2,
  width: "100%",
  display: "flex",
  justifyContent: "space-evenly",
  padding: 0.3,
  ":hover": {
    width: "200px",
    transform: "scale(1.05)",
    transition: "all 0.1s ease"
  }
};
const toolsContainerListBoxGridSelectGridItemGridItemStackStyle: SxProps = {
  height: "16%",
  backgroundColor: "secondary.dark",
  borderTopLeftRadius: "16px",
};
const toolsContainerListBoxGridSelectGridItemGridItemStackTypographyListStyle: SxProps =
{
  color: "primary.contrastText",
  margin: "auto",
};
const toolsContainerListBoxGridSelectGridItemGridItemStackTypographyCommunityListStyle: SxProps =
{
  color: "primary.contrastText",
  margin: "auto 0",
};
const toolsContainerListBoxGridSelectGridItemGridItemStackCommunityListStyle: SxProps =
{
  backgroundColor: "secondary.dark",
  overflowX: "auto",
  borderRadius: 2,
  height: "380px",
  "&::-webkit-scrollbar": {
    height: "8px",
    width: "6px",
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
  marginLeft: 4,
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
