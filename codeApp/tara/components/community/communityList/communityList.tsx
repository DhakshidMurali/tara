"use client";

import { communityOverViewDetails } from "@/public/data/community";
import { Group } from "@mui/icons-material";
import { Box, Grid, Typography } from "@mui/material";
import { styles } from "./useStyles";

export default function CommunityList() {
  return (
    <Grid container spacing={2} sx={styles.toolsContainerBoxStyle}>
      {communityOverViewDetails.map((data, index) => (
        <Grid item key={index}>
          <Grid container sx={styles.toolsContainerListBoxGridStyle}>
            <Grid item xs={12}>
              <Typography
                variant="h6"
                sx={
                  styles.toolsContainerListBoxGridTypograghyCommunityNameStyle
                }
              >
                {data.communityName}
              </Typography>
            </Grid>
            <Grid item xs={6}>
              <Box sx={styles.toolsContainerListBoxGridBoxStyle}>
                <Group sx={{ color: "white" }}></Group>
                <Typography
                  variant="subtitle1"
                  sx={
                    styles.toolsContainerListBoxGridBoxTypographyParticipantsStyle
                  }
                >
                  {" "}
                  {data.totalParticipants}
                </Typography>
              </Box>
            </Grid>
            <Grid item xs={3} marginRight={2}>
              <Box
                sx={
                  styles.toolsContainerListBoxGridIncreaseInLastMonthBoxStyle
                }
              >
                <Typography
                  variant="subtitle1"
                  sx={
                    styles.toolsContainerListBoxGridIncreaseInLastMonthBoxTypograghyStyle
                  }
                >
                  {"+  " + data.increaseInLastMonth}
                </Typography>
              </Box>
            </Grid>
          </Grid>
        </Grid>
      ))}
    </Grid>
  );
}
