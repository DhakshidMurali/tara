"use client";
import CommunityBarChart from "@/components/community/communityBarChart/communityBarChart";
import CommunityList from "@/components/community/communityList/communityList";
import SearchBox from "@/components/common/searchBox";
import NavBar from "@/components/common/navBar/navBar";
import { dayNames, month } from "@/utils/constants";
import { PageviewRounded } from "@mui/icons-material";
import { Box, Button, Grid, Stack, TextField, Typography } from "@mui/material";
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
        <Grid container spacing={1} height={"80%"}>
          <Grid item xs={12}>
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
              <Typography variant="h3" sx={styles.communityTextTypographStyle}>
                Community{" "}
              </Typography>
              <Button sx={styles.createButtonStyle}>Create</Button>
            </Stack>
          </Grid>
          <Grid item xs={12}>
            <CommunityList></CommunityList>
          </Grid>
          <Grid item xs={12} sx={styles.gridContainerCommnityBox}>
            <CommunityBarChart></CommunityBarChart>
          </Grid>
        </Grid>
      </Grid>
    </>
  );
}
