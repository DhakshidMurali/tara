"use client";

import NavBar from "@/components/common/navBar/navBar";
import { PageviewRounded } from "@mui/icons-material";
import { Box, Button, Grid, Stack, TextField, Typography } from "@mui/material";
import { styles } from "./useStyles";

import SearchBox from "@/components/common/searchBox.jsx";
import DashBoardBarChart from "@/components/dashboard/dashboardBarChart/dashboardBarChart";
import DashboardDeptMenu from "@/components/dashboard/dashboardDeptMenu/dashboardDeptMenu";
import { DashBoardList } from "@/components/dashboard/dashBoardList/dashboardList";
import { dayNames, month } from "@/utils/constants";

export default function DashBoard() {
  let time = new Date();
  return (
    <>
      <Grid item xs={2.15} sx={{ marginRight: 2 }}>
        <NavBar selected="dashboard"></NavBar>
      </Grid>
      <Grid item xs={9.75}>
        <SearchBox>
          <Stack
            direction={"row"}
            alignItems={"center"}
            height={"100%"}
            spacing={2}
          >
            <PageviewRounded></PageviewRounded>
            <TextField
              id="standard-basic"
              variant="standard"
              placeholder="Search"
              fullWidth
              InputProps={{
                style: { color: "#f3def5", fontWeight: "bolder", fontSize: 16 },
              }}
            />
            <Typography sx={styles.dateTypographyStyle}>
              {dayNames[time.getDay()] +
                ", " +
                month[time.getMonth()] +
                " " +
                time.getDate()}
            </Typography>
          </Stack>
        </SearchBox>
        <Grid container spacing={1}>
          <Grid item xs={12} sx={{}}>
            <Stack
              direction={"row"}
              display={"flex"}
              justifyContent={"space-between"}
            >
              <Typography variant="h3" sx={styles.dashboardTextTypographStyle}>
                Dashboard{" "}
              </Typography>
              <Button sx={styles.createButtonStyle}>Create</Button>
            </Stack>
          </Grid>
          <Grid item xs={12}>
            <DashBoardList></DashBoardList>
          </Grid>
          <Grid item xs={8}>
            <DashBoardBarChart></DashBoardBarChart>
          </Grid>
          <Grid item xs={4}>
            <Box sx={styles.departmentContainerBoxStyle}>
              <Stack spacing={2}>
                <Typography
                  variant="h4"
                  sx={styles.dashboardTextTypographStyle}
                >
                  Departments
                </Typography>
                <DashboardDeptMenu></DashboardDeptMenu>
              </Stack>
            </Box>
          </Grid>
        </Grid>
      </Grid>
    </>
  );
}
