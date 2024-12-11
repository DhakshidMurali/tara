"use client";

import NavBar from "@/components/navBar/navBar.js";
import { PageviewRounded } from "@mui/icons-material";
import {
  Box,
  Button,
  Grid,
  Stack,
  TextField,
  Tooltip,
  Typography,
} from "@mui/material";
import { styles } from "./useStyles.tsx";

import { dayNames, month } from "@/utils/constants";
import SearchBox from "@/components/dashboard/box/searchBox.js";
import DepartmentListBox from "@/components/dashboard/box/departmentListBox.js";
import EmployeeByDepartmentListBox from "@/components/dashboard/box/employeeByDepartmentListBox.js";
import { DepartmentList } from "@/public/Sample/data.js";
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
            <Box sx={styles.toolsContainerBoxStyle}>
              <Grid container spacing={2} sx={{ flexWrap: "nowrap" }}>
                {Array.from({ length: 10 }).map((_, index) => (
                  <Grid item key={index}>
                    <Box sx={styles.toolsContainerListBoxStyle}>
                      <Stack sx={styles.toolsContainerListBoxStackStyle}>
                        <Box
                          sx={{
                            width: "2%",
                            height: "85%",
                            bgcolor: "yellow",
                            borderRadius: "4px",
                          }}
                        ></Box>
                        <Box
                          sx={{
                            width: "2%",
                            height: "85%",
                            bgcolor: "yellow",
                            borderRadius: "4px",
                          }}
                        ></Box>
                        <Grid></Grid>
                      </Stack>
                    </Box>
                  </Grid>
                ))}
              </Grid>
            </Box>
          </Grid>
          <Grid item xs={8}>
            <Box
              sx={styles.employerByDepartmentBoxStyle}
              padding={4}
              height={548}
            >
              <Stack spacing={4} sx={{ width: "100%" }}>
                <Typography
                  variant="h4"
                  sx={styles.dashboardTextTypographStyle}
                >
                  Employee By Department
                </Typography>
                <EmployeeByDepartmentListBox
                  departmentList={deparmentList}
                ></EmployeeByDepartmentListBox>
              </Stack>
            </Box>
          </Grid>
          <Grid item xs={4}>
            <Box sx={styles.departmentContainerBoxStyle}>
              <Stack spacing={0}>
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
