"use client";
import NavBar from "@/components/common/navBar/navBar";
import SearchBox from "@/components/common/searchBox";
import { PageviewRounded } from "@mui/icons-material";
import { Box, Button, Grid, Stack, TextField, Typography } from "@mui/material";
import { styles } from "./useStyles";
import { dayNames, month } from "@/utils/constants";
import ToolList from "@/components/tool/toolList/toollist";

export default function Tool() {
  let time = new Date();
  return (
    <>
      <Grid item xs={2.15} sx={{ marginRight: 2 }}>
        <NavBar selected="tool"></NavBar>
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
        </SearchBox>{" "}
        <Grid container spacing={1}>
          {" "}
          <Grid item xs={12}>
            <Stack
              direction={"row"}
              display={"flex"}
              justifyContent={"space-between"}
            >
              <Typography variant="h3" sx={styles.toolTextTypographStyle}>
                Tool{" "}
              </Typography>
              <Button sx={styles.createButtonStyle}>Create</Button>
            </Stack>
          </Grid>
          <Grid item xs={12}>
            <ToolList></ToolList>
          </Grid>
        </Grid>
      </Grid>
    </>
  );
}
