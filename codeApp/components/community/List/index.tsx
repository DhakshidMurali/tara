"use client";

import { communityOverViewDetails } from "@/public/data/community";
import { Group } from "@mui/icons-material";
import { Box, Grid, Typography } from "@mui/material";
import { styles } from "./useStyles";

export default function CommunityList() {
  return (
    <Grid container size={{ xs: 12 }} spacing={2} sx={{ height: "90rem", width: "120rem" }}>
      <Grid
        sx={{
          overflowX: 'auto',
          display: 'flex',
          flexWrap: 'nowrap',
          width: '100%',
          maxWidth: '100%',
          padding: 2,
          borderRadius: 2,
          '&::-webkit-scrollbar': {
            height: '8px',
          },
          '&::-webkit-scrollbar-track': {
            backgroundColor: 'secondary.main',
          },
          '&::-webkit-scrollbar-thumb': {
            backgroundColor: 'secondary.light',
            borderRadius: '8px',
          },
          '&::-webkit-scrollbar-thumb:hover': {
            background: '#555',
          },
        }}

      >
        {communityOverViewDetails.map((data) => (
          <Grid container sx={styles.toolsContainerListBoxGridStyle}>
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
        ))}
      </Grid>
    </Grid>
  );
}
