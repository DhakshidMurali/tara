"use client";

import { communityOverViewDetails } from "@/public/data/community";
import { DescriptionOutlined } from "@mui/icons-material";
import { Avatar, AvatarGroup, Box, Grid, Paper, Stack, Tooltip, Typography } from "@mui/material";
import { styles } from "./useStyles";
export default function ToolList() {
  return (
    <Grid size={{ xs: 12 }} component="div" spacing={2} height={"14rem"}>
      <Box sx={styles.toolsContainerBoxStyle}>
        {communityOverViewDetails.map((data, index) => (
          <Box>
            <Grid container sx={styles.toolsContainerListBoxGridStyle}>
              <Grid size={{ xs: 12 }} component="div">
                {" "}
                <Typography
                  variant="h6"
                  sx={styles.toolsContainerListBoxGridToolNameTypography}
                >
                  {data.communityName}
                </Typography>
              </Grid>
              <Grid size={{ xs: 4 }} component="div" marginLeft={1}>
                <Paper
                  elevation={2}
                  sx={{
                    padding: 0.5,
                    margin: 0.5,
                    display: "flex",
                    justifyContent: "center",
                    backgroundColor: "primary.light",
                  }}
                >
                  {" "}
                  <Stack direction={"row"} display={"flex"} justifyContent={"center"} alignItems={'center'}>
                    <DescriptionOutlined sx={{ fontSize: 16, marginRight: 0.2 }}></DescriptionOutlined>
                    <Typography
                      sx={
                        styles.toolsContainerListBoxGridCommunicationCountTypography
                      }
                    >
                      {data.increaseInLastMonth}
                    </Typography>
                  </Stack>
                </Paper>
              </Grid>
              <Grid size={{ xs: 4 }} component="div" marginRight={1}>
                <AvatarGroup max={3} spacing={12} sx={{ "& .MuiAvatar-root": { border: "2px solid black" } }} >
                  <Tooltip title="Web Development" ><Avatar sx={{ bgcolor: "rgb(239,241,255)" }} ><Typography sx={{ color: "secondary.main" }}>W</Typography></Avatar></Tooltip>
                  <Tooltip title="Deb Development"><Avatar sx={{ bgcolor: "rgb(239,241,255)" }}><Typography sx={{ color: "secondary.main" }}>D</Typography></Avatar></Tooltip>
                </AvatarGroup>
              </Grid>
            </Grid>
          </Box>
        ))}
      </Box>
    </Grid>
  );
}
