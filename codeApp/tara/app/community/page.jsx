"use client";
import SearchBox from "@/components/dashboard/box/searchBox.jsx";
import NavBar from "@/components/navBar/navBar";
import {
  Close,
  ClosedCaptionOffOutlined,
  CloseFullscreenOutlined,
  Forward,
  Group,
  PageviewRounded,
  RemoveCircleOutlineRounded,
  RemoveFromQueueOutlined,
} from "@mui/icons-material";
import {
  Box,
  Button,
  Grid,
  IconButton,
  Paper,
  Stack,
  TextField,
  Typography,
} from "@mui/material";
import { styles } from "./useStyles";
import { dayNames, month } from "@/utils/constants";
import { BarChart } from "@mui/x-charts/BarChart";
import { Dataset } from "@/public/Sample/data";
import { addLabels } from "@/utils/functions";
import { communityOverViewDetails } from "@/public/data/community";

export default function Community() {
  let time = new Date();
  console.log(communityOverViewDetails);
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
          <Grid item xs={12} sx={{ height: "22.5%" }}>
            <Box sx={styles.toolsContainerBoxStyle}>
              <Grid container spacing={2} sx={{ flexWrap: "nowrap" }}>
                {communityOverViewDetails.map((data, index) => (
                  <Grid item key={index} sx={{ height: "100%" }}>
                    <Box sx={styles.toolsContainerListBoxStyle}>
                      <Grid
                        container
                        sx={{
                          height: "100%",
                          display: "flex",
                          justifyContent: "space-between",
                        }}
                      >
                        <Grid item xs={12}>
                          <Typography
                            variant="h6"
                            sx={{
                              color: "#f3def5",
                              paddingLeft: 1,
                              paddingTop: 0.5,
                            }}
                          >
                            {" "}
                            {data.communityName}
                          </Typography>
                        </Grid>
                        <Grid item xs={6}>
                          <Box
                            sx={{
                              display: "flex",
                              flexDirection: "row",
                              paddingLeft: 1,
                            }}
                          >
                            <Group sx={{ color: "white" }}></Group>
                            <Typography
                              variant="subtitle1"
                              sx={{
                                color: "#f3def5",
                                paddingLeft: 1,
                              }}
                            >
                              {" "}
                              {data.totalParticipants}
                            </Typography>
                          </Box>
                        </Grid>
                        <Grid item xs={3} marginRight={2}>
                          <Box
                            sx={{
                              backgroundColor: "rgb(177,248,180)",
                              display: "flex",
                              justifyContent: "center",
                              alignContent: "center",
                              borderRadius: 3,
                            }}
                          >
                            <Typography
                              variant="subtitle1"
                              sx={{
                                color: "rgb(33,24,43)",
                                justifyContent: "flex-end",
                                fontWeight: "bolder  ",
                              }}
                            >
                              {"+  " + data.increaseInLastMonth}
                            </Typography>
                          </Box>
                        </Grid>
                      </Grid>
                      ;
                    </Box>
                  </Grid>
                ))}
              </Grid>
            </Box>
          </Grid>
          <Grid item xs={12} sx={{ height: "36%" }}>
            <Box sx={{ height: "100%" }}>
              <Grid container sx={{ height: "100%" }}>
                <Grid item xs={8} sx={{ height: "20%" }}>
                  <Typography
                    variant="h4"
                    sx={styles.communityTextTypographStyle}
                  >
                    Find Community works for your
                  </Typography>
                </Grid>
                <Grid item xs={4} sx={{ height: "20%" }}>
                  <Stack direction={"row"} spacing={1} height={"100%"}>
                    <Box
                      sx={{
                        display: "flex",
                        justifyContent: "center",
                        alignContent: "center",
                      }}
                    >
                      <Typography
                        variant="h6"
                        sx={{
                          color: "#f3def5",
                          paddingLeft: 1,
                          margin: "auto",
                        }}
                        textAlign="center"
                      >
                        Select 5 for study case
                      </Typography>
                    </Box>
                    <Box
                      sx={{
                        width: "1.1%",
                        backgroundColor: "blue",
                        height: "100%",
                        borderRadius: 2,
                      }}
                    ></Box>
                    <Grid
                      container
                      sx={{
                        height: "100%",
                        width: "90%",
                      }}
                    >
                      {[1, 1, 1, 1, 1].map(() => {
                        return (
                          <Grid item xs={3.5} marginRight={1}>
                            <Box
                              variant="elevation"
                              sx={{
                                height: "90%",
                                backgroundColor: "green",
                                borderRadius: 2,
                                display: "flex",
                                justifyContent: "center",
                                alignContent: "center",
                              }}
                            >
                              <IconButton sx={{ padding: 0 }}>
                                <Close sx={{ alignSelf: "center" }}></Close>
                              </IconButton>
                              <Typography
                                variant="body2"
                                sx={{ alignSelf: "center", color: "white" }}
                              >
                                Cloud-Neo..
                              </Typography>
                            </Box>
                          </Grid>
                        );
                      })}
                    </Grid>
                  </Stack>
                </Grid>
                <Grid item xs={8} sx={{ height: "80%" }}>
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
                <Grid item xs={4} sx={{ height: "80%" }}>
                  <Grid container sx={{ height: "100%" }}>
                    <Grid item xs={5}>
                      {[1, 1, 1, 1, 1].map(() => {
                        return (
                          <Stack
                            direction={"column"}
                            sx={{
                              height: "18%",
                              display: "flex",
                              backgroundColor: "orange",
                              justifyContent: "center",
                              alignContent: "center",
                              textAlign: "center",
                              borderTopLeftRadius: "16px",
                            }}
                            marginTop={0.5}
                          >
                            <Typography
                              variant="subtitle1"
                              sx={{
                                color: "white",
                                overflow: "hidden",
                                textOverflow: "ellipsis",
                                whiteSpace: "nowrap",
                              }}
                            >
                              Cloud Computing and DevOps
                            </Typography>
                          </Stack>
                        );
                      })}
                    </Grid>
                    <Grid item xs={7}>
                      {[1, 1, 1, 1, 1].map((data, index) => {
                        return (
                          <Stack direction={"column"}>
                            <Typography
                              variant="subtitle1"
                              sx={{
                                color: "white",
                                paddingLeft: 2,
                                paddingTop: 1,
                                paddingRight: 2,
                              }}
                            >
                              Cloud Computing and DevOps
                            </Typography>
                          </Stack>
                        );
                      })}
                    </Grid>
                  </Grid>
                </Grid>
              </Grid>
            </Box>
          </Grid>
          <Grid item xs={12} sx={{ backgroundColor: "yellow", height: "35%" }}>
            <Box></Box>
          </Grid>
        </Grid>
      </Grid>
    </>
  );
}
