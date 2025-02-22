"use client";
import CommunityBarChart from "@/components/community/communityBarChart/communityBarChart";
import CommunityList from "@/components/community/communityList/communityList";
import SearchBox from "@/components/common/searchBox";
import NavBar from "@/components/common/navBar/navBar";
import { dayNames, month } from "@/utils/constants";
import { PageviewRounded } from "@mui/icons-material";
import { Box, Grid, Stack, TextField, Typography } from "@mui/material";
import { styles } from "./useStyles";

export default function Community() {
  let time = new Date();
  const label = { inputProps: { "aria-label": "Checkbox demo" } };
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
          <CommunityList></CommunityList>
          <Box sx={styles.gridContainerCommnityBox}>
            <CommunityBarChart></CommunityBarChart>
            <Grid item xs={12} >
              <Grid container>
                <Grid item xs={8}>
                  <Grid container>
                    {[1, 1, 1, 1].map((data) =>
                    (<Grid item xs={12} sx={{ height: "10%px", display: "flex", backgroundColor: "red" }}>
                      <Box sx={{ height: "100%", backgroundColor: "orange" }}>

                      </Box>
                    </Grid>))}
                  </Grid>
                </Grid>
                <Grid item xs={4}></Grid>
              </Grid>
            </Grid>
          </Box>
        </Grid>
      </Grid>
    </>
  );
}
