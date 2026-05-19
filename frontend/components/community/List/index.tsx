"use client";

import { communityOverViewDetails } from "@/public/data/community";
import { Group } from "@mui/icons-material";
import { Box, Grid, Typography } from "@mui/material";
import { styles } from "./useStyles";

export default function CommunityList() {
  return (
    <Grid container size={{ xs: 12 }} spacing={2} sx={styles.toolsContainerBoxStyle}>
      {communityOverViewDetails.map((data, index) => (
        <Box>
          <Grid container sx={styles.toolsContainerListBoxGridStyle} key={index}>
            <Grid size={{ xs: 12 }}>
              <Typography
                variant="h6"
                sx={
                  styles.toolsContainerListBoxGridTypograghyCommunityNameStyle
                }
              >
                {data.communityName}
              </Typography>
            </Grid>
            <Grid size={{ xs: 5 }}>
              <Box sx={styles.toolsContainerListBoxGridBoxStyle}>
                <Group sx={{ color: "primary.contrastText" }}></Group>
                <Typography
                  variant="subtitle1"
                  sx={{
                    ...styles.toolsContainerListBoxGridBoxTypographyParticipantsStyle,
                  }}
                >
                  {" "}
                  {data.totalParticipants}
                </Typography>
              </Box>
            </Grid>
            <Grid size={{ xs: 5 }} marginRight={2}>
              <Box
                sx={styles.toolsContainerListBoxGridIncreaseInLastMonthBoxStyle}
              >
                <Typography
                  variant="subtitle1"
                  sx={
                    styles.toolsContainerListBoxGridIncreaseInLastMonthBoxTypograghyStyle
                  }
                >
                  {"+  " + data.increaseInLastMonth + " MoM"}
                </Typography>
              </Box>
            </Grid>
          </Grid>
        </Box>
      ))}
    </Grid>

  );
}
