import { SxProps } from "@mui/system";
const departmentContainerListBoxStyle: SxProps = {
    borderRadius: 2,
    width: "100%",
    display: "flex",
    flexDirection: "column",
    justifyContent: "flex-start",
    alignItems: "center",
    height: "100%"
};

const departmentContainerListBoxStackStyle: SxProps = {
    width: "100%",
    display: "flex",
    flexDirection: "row",
    justifyContent: "space-between",
    margin: 1.5,
};

const departmentContainerListBoxBoxExpandedStyle: SxProps = {
    width: "100%",
    height: "380px",
    padding: 1,
    display: "flex",
    flexDirection: "row",
    justifyContent: "space-between",
    alignContent: "flex-start",
    overflowX: "auto",
    borderRadius: 2,
    alignItems: "flex-start",
    "&::-webkit-scrollbar": {
        height: "8px",
        width: "4px"
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
const departmentContainerListBoxStackTopicBoxStyle: SxProps = {
    backgroundColor: "secondary.main",
    padding: 1,
    borderRadius: 2,
    marginTop: 1.5
};

const dashboardTextTypographStyle: SxProps = {
    color: "#f3def5",
    fontWeight: "bolder",
};

const deparmentListIconStyle: SxProps = {
    color: "secondary.main",
};




export const styles = {
    departmentContainerListBoxStyle,
    departmentContainerListBoxStackStyle,
    departmentContainerListBoxBoxExpandedStyle,
    departmentContainerListBoxStackTopicBoxStyle,
    dashboardTextTypographStyle,
    deparmentListIconStyle,
};
