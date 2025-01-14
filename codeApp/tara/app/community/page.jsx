"use client";
import SearchBox from "@/components/dashboard/box/searchBox.jsx";
import NavBar from "@/components/navBar/navBar";
import { PageviewRounded } from "@mui/icons-material";
import { Box, Button, Grid, Stack, TextField, Typography } from "@mui/material";
import { styles } from "./useStyles";
import { dayNames, month } from "@/utils/constants";
import { BarChart } from "@mui/x-charts/BarChart";
import { Dataset } from "@/public/Sample/data";
import { axisClasses } from "@mui/x-charts/ChartsAxis";
import { addLabels } from "@/utils/functions";

export default function Community() {
  let time = new Date();

  return (
    <>
      <Grid item xs={2.15} sx={{ marginRight: 2 }}>
        <NavBar selected="community"></NavBar>
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
        <Grid container sx={{ height: "89%", paddingLeft: "8px" }}>
          <Grid item xs={12} sx={{ height: "5%", marginBottom: "0.5%" }}>
            <Stack
              direction={"row"}
              display={"flex"}
              justifyContent={"space-between"}
            >
              <Typography variant="h3" sx={styles.communityTextTypographStyle}>
                Community{" "}
              </Typography>
              <Button sx={styles.createButtonStyle}>Create</Button>
            </Stack>
          </Grid>
          <Grid item xs={12} sx={{ height: "19.5%" }}>
            <Box sx={styles.toolsContainerBoxStyle}>
              <Grid container spacing={2} sx={{ flexWrap: "nowrap" }}>
                {Array.from({ length: 10 }).map((_, index) => (
                  <Grid item key={index}>
                    <Box sx={styles.toolsContainerListBoxStyle}>
                      <Typography variant="h6" color="white">
                        Item {index + 1}
                      </Typography>
                    </Box>
                  </Grid>
                ))}
              </Grid>
            </Box>
          </Grid>
          <Grid item xs={8} sx={{ height: "40%" }}>
            <Box sx={{ height: "100%" }}>
              <Grid container sx={{ height: "100%" }}>
                <Grid item xs={12} sx={{ height: "20%" }}>
                  <Typography
                    variant="h4"
                    sx={styles.communityTextTypographStyle}
                  >
                    Find Community works for your
                  </Typography>
                </Grid>
                <Grid item xs={12} sx={{ height: "80%" }}>
                  <BarChart
                    dataset={Dataset}
                    xAxis={[
                      {
                        scaleType: "band",
                        dataKey: "month",
                      },
                    ]}
                    series={addLabels([
                      { dataKey: "DevOps", stack: "assets" },
                      { dataKey: "Web Devlopment", stack: "assets" },
                      { dataKey: "App Devlopment", stack: "liability" },
                      {
                        dataKey: "Internet of Things",
                        stack: "liability",
                      },
                      { dataKey: "Cloud Computing", stack: "equity" },
                      { dataKey: "DevOps", stack: "assets" },
                      { dataKey: "Web Devlopment", stack: "assets" },
                      { dataKey: "App Devlopment", stack: "liability" },
                      {
                        dataKey: "Internet of Things",
                        stack: "liability",
                      },
                      { dataKey: "Cloud Computing", stack: "equity" },
                    ])}
                    slotProps={{
                      legend: { hidden: true },
                    }}
                    sx={{
                      /* Left Axis Label Font Styling */
                      "& .MuiChartsAxis-left .MuiChartsAxis-tickLabel": {
                        strokeWidth: "1",
                        fill: "#f3def5",
                      },
                      /* Bottom Label Font Styling */
                      "& .MuiChartsAxis-bottom .MuiChartsAxis-tickLabel": {
                        strokeWidth: "1",
                        fill: "#f3def5",
                        fontSize: "24px",
                      },
                      /* X axis Line Styling */
                      "& .MuiChartsAxis-bottom .MuiChartsAxis-line": {
                        stroke: "#f3def5",
                        strokeWidth: 2.4,
                      },
                      /* Y Axis Line Styling  */
                      "& .MuiChartsAxis-left .MuiChartsAxis-line": {
                        stroke: "#f3def5",
                        strokeWidth: 2.4,
                      },
                    }}
                    tooltip={{ trigger: "item" }}
                  />
                </Grid>
              </Grid>
            </Box>
          </Grid>
          <Grid item xs={4} sx={{ height: "40%" }}>
            <Box></Box>
          </Grid>
          <Grid
            item
            xs={12}
            sx={{ backgroundColor: "greenyellow", height: "35%" }}
          >
            <Box></Box>
          </Grid>
        </Grid>
      </Grid>
    </>
  );
}
