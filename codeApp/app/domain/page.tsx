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
      <Grid size={{ xs: 2 }} component="div" sx={{ marginRight: 2 }}>
        <NavBar selected="domain" />
      </Grid>
      <Grid size={{ xs: 9.75 }} component="div">
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
          <Grid size={{ xs: 12 }} component="div">
            <Stack
              direction={"row"}
              display={"flex"}
              justifyContent={"space-between"}
              sx={{
                "& .MuiButtonBase-root:hover": {
                  backgroundColor: "primary.light",
                  fontWeight: "bold"
                }
              }}
            >
              <Typography variant="h3" sx={styles.dashboardTextTypographStyle}>
                Domain{" "}
              </Typography>
              <Button sx={styles.createButtonStyle}>Create</Button>
            </Stack>
          </Grid>
          <Grid size={{ xs: 12 }} component="div">
            <DashBoardList></DashBoardList>
          </Grid>
          <Grid size={{ xs: 8 }} component="div">
            <DashBoardBarChart></DashBoardBarChart>
          </Grid>
          <Grid size={{ xs: 4 }} component="div">
            <Box sx={styles.departmentContainerBoxStyle}>
              <Stack spacing={2}>
                <Typography
                  variant="h4"
                  sx={styles.dashboardTextTypographStyle}
                >
                  Domains
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
