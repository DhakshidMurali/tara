import { Avatar, Box, Button, Stack, Typography } from "@mui/material";
import { styles } from "../useStyles";
import { Domain, ExpandMore } from "@mui/icons-material";
depar

export default function DepartmentListBox(props) {
  return (
    <Box sx={styles.departmentContainerListBoxStyle}>
      <Stack direction={"row"}>
        <Avatar>
          <Domain sx={deparmentListIconStyle}></Domain>
        </Avatar>
        <Typography
          variant="h6"
          sx={styles.dashboardTextTypographStyle}
          paddingLeft={2}
        >
          Computer Science Eng
        </Typography>
        <Button sx={{ margin: 0 }}>
          <ExpandMore sx={{ marginLeft: 0, color: "white" }}></ExpandMore>
        </Button>
      </Stack>
    </Box>
  );
}
