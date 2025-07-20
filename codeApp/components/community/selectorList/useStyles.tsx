import { SxProps } from "@mui/material";

const toolsContainerListBoxGridSelectGridItemGridItemStackStyle: SxProps = {
    backgroundColor: "secondary.dark", borderTopLeftRadius: "16px",
};

const toolsContainerListBoxGridSelectGridItemGridItemStackTypographyListStyle: SxProps =
{
    color: "primary.contrastText",
    margin: "auto",
};

const toolsContainerListBoxGridSelectGridItemGridItemStackCommunityListStyle: SxProps =
{
    overflowX: "auto",
    borderRadius: 2,
    height: "inherit",
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

const toolsContainerListBoxGridSelectGridItemGridItemStackTypographyCommunityListStyle: SxProps =
{
    color: "primary.contrastText",
    margin: "auto 0",
};

export const styles = {
    toolsContainerListBoxGridSelectGridItemGridItemStackStyle,
    toolsContainerListBoxGridSelectGridItemGridItemStackTypographyListStyle,
    toolsContainerListBoxGridSelectGridItemGridItemStackCommunityListStyle,
    toolsContainerListBoxGridSelectGridItemGridItemStackTypographyCommunityListStyle
};