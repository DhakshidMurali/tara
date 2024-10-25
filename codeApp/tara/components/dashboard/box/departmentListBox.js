"use client";
import { Avatar, Box, Button, Grid, Stack, Typography } from "@mui/material";
import { styles } from "../useStyles";
import { Domain, ExpandMore } from "@mui/icons-material";
import { useRouter } from "next/navigation";
import { DepartmentList } from "@/public/Sample/data";

export default function DepartmentListBox(props) {
  const router = useRouter();
  let isDropDownOpen = false;
  const departmentList = DepartmentList;
  const handleDropDownDepartmentList = () => {
    isDropDownOpen = !isDropDownOpen;
    router.refresh();
  };
  return (
    <Grid
      container
      sx={{
        justifyContent: "flex-start",
      }}
    >
      <Grid
        item
        xs={12}
        sx={{
          display: "flex",
          flexDirection: "column",
          justifyContent: "flex-start",
        }}
      >
        {departmentList.map((i, index) => {
          return (
            <Box sx={styles.departmentContainerListBoxStyle}>
              <Stack direction={"row"}>
                <Avatar>
                  <Domain sx={styles.deparmentListIconStyle}></Domain>
                </Avatar>
                <Typography
                  variant="h6"
                  sx={styles.dashboardTextTypographStyle}
                  paddingLeft={2}
                >
                  {departmentList[index]}
                </Typography>
                <Button
                  sx={{ margin: 0 }}
                  onClick={handleDropDownDepartmentList}
                >
                  <ExpandMore
                    sx={{ marginLeft: 0, color: "white" }}
                  ></ExpandMore>
                </Button>
              </Stack>
            </Box>
          );
        })}
      </Grid>
    </Grid>
  );
}

{
  /* <Box sx={styles.departmentContainerListBoxStyle}>
      <Stack direction={"row"}>
        <Avatar>
          <Domain sx={styles.deparmentListIconStyle}></Domain>
        </Avatar>
        <Typography
          variant="h6"
          sx={styles.dashboardTextTypographStyle}
          paddingLeft={2}
        >
          {isDropDownOpen
            ? "Computer Science Eng"
            : "Computer Science Engineering"}
        </Typography>
        <Button sx={{ margin: 0 }} onClick={handleDropDownDepartmentList}>
          <ExpandMore sx={{ marginLeft: 0, color: "white" }}></ExpandMore>
        </Button>
      </Stack>
    </Box> */
}
