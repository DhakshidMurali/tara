import { SxProps } from "@mui/material";

const toolsContainerBoxStyle: SxProps = {
    overflowY: "auto",
    borderRadius: 2,
    flexWrap: "nowrap",
    height: "inherit",
    "&::-webkit-scrollbar": {
        height: "2px",
        width: "2px"
    },
    "&::-webkit-scrollbar-track": {
        backgroundColor: "secondary.main",
    },
    "&::-webkit-scrollbar-thumb": {
        backgroundColor: "secondary.light",
        borderRadius: "1px",
    },
    "&::-webkit-scrollbar-thumb:hover": {
        background: "#555",
    },
};

export const styles = {
    toolsContainerBoxStyle
};