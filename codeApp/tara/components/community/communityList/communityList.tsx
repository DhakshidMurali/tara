"use client";

import { communityOverViewDetails } from "@/public/data/community";
import { Group, GroupAddRounded, Groups, Groups2Rounded, Groups3Rounded, GroupsRounded, GroupWorkRounded } from "@mui/icons-material";
import { Box, Grid, Typography } from "@mui/material";
import { styles } from "./useStyles";
import { getRandomColor } from "@/utils/functions";

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
            <Grid item xs={5}>
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
            <Grid item xs={5} marginRight={2}>
              <Box
                sx={styles.toolsContainerListBoxGridIncreaseInLastMonthBoxStyle}
              >
                <Typography
                  variant="subtitle1"
                  sx={
                    styles.toolsContainerListBoxGridIncreaseInLastMonthBoxTypograghyStyle
                  }
                >
                  {"+  " + data.increaseInLastMonth +" MoM"}
                </Typography>
              </Box>
            </Grid>
          </Grid>
        </Grid>
      ))}
    </Grid>
  );
}
