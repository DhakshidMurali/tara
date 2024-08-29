import NavBar from "@/components/navBar/navBar";
import { Box, Grid } from "@mui/material";

export default function Hierarchy() {
  return (
    <>
      <Grid item xs={2.25}>
        <NavBar selected="hierarchy"></NavBar>
      </Grid>
      <MainBoard>
        <Grid item xs={9.75}>
          <Stack>
            <Stack></Stack>
            <Grid container spacing={2}>
              <Grid xs={8}>
                <Box sx={{ height: 240, backgroundColor: "green" }}></Box>
              </Grid>
              <Grid xs={4}>
                <Box sx={{ color: green, height: 240 }}></Box>
              </Grid>
              <Grid xs={8}>
                <Box sx={{ color: green, height: 240 }}></Box>
              </Grid>
              <Grid xs={4}>
                <Box sx={{ color: green, height: 240 }}></Box>
              </Grid>
            </Grid>
          </Stack>
        </Grid>
      </MainBoard>
    </>
  );
}
