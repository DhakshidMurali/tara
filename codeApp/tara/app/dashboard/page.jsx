"use client";

import NavBar from "@/components/navBar/navBar.js";
import { PageviewRounded } from "@mui/icons-material";
import { Box, Button, Grid, Stack, TextField, Typography } from "@mui/material";
import { styles } from "./useStyles.tsx";

import SearchBox from "@/components/common/searchBox.jsx";
import DepartmentListBox from "@/components/dashboard/box/departmentListBox.jsx";
import EmployeeByDepartmentListBox from "@/components/dashboard/box/employeeByDepartmentListBox.jsx";
import { DashBoardList } from "@/components/dashboard/dashBoardList/dashboardList";
import { DepartmentList } from "@/public/Sample/data.js";
import { dayNames, month } from "@/utils/constants";

export default function DashBoard() {
  let time = new Date();
  const deparmentList = DepartmentList;
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
        <Grid container spacing={2}>
          {/* Horizontal Scroll Grid */}
          <Grid item xs={12}>
            <DashBoardList></DashBoardList>
          </Grid>
          <Grid item xs={8}>
            
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
                <DepartmentListBox></DepartmentListBox>
              </Stack>
            </Box>
          </Grid>
        </Grid>
      </Grid>
    </>
  );
}
