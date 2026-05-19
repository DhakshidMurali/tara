import { SxProps } from "@mui/material";

const toolTextTypographStyle: SxProps = {
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

export const styles = {
    toolTextTypographStyle,
    createButtonStyle
};
