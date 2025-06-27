"use client";
import NavBar from "@/components/common/navBar/navBar";
import SearchBox from "@/components/common/searchBox";
import { PageviewRounded } from "@mui/icons-material";
import { Box, Button, Grid, Stack, TextField, Typography } from "@mui/material";
import { styles } from "./useStyles";
import { dayNames, month } from "@/utils/constants";
import ToolList from "@/components/tool/toolList/toollist";
import ToolBarChart from "@/components/tool/toolBarChart/toolBarChart";
import ToolBarChartAnalysis from "@/components/tool/toolBarChart/toolBarChartAnalysis";

export default function Tool() {
  let time = new Date();
  return (
    <>
      <Grid size={{ xs: 2 }} component="div" sx={{ marginRight: 2 }}>
        <NavBar selected="tool" />
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
        </SearchBox>{" "}
        <Grid container spacing={1}>
          {" "}
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
              <Typography variant="h3" sx={styles.toolTextTypographStyle}>
                Tool{" "}
              </Typography>
              <Button sx={styles.createButtonStyle}>Create</Button>
            </Stack>
          </Grid>
          <Grid size={{ xs: 12 }} component="div">
            <ToolList></ToolList>
          </Grid>
          <Grid size={{ xs: 8 }} component="div"><ToolBarChart></ToolBarChart></Grid>
          <Grid size={{ xs: 4 }} component="div"><ToolBarChartAnalysis></ToolBarChartAnalysis></Grid>
        </Grid>
      </Grid>
    </>
  );
}
