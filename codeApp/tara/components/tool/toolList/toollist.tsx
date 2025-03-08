"use client";

import { Box, Grid, Paper, Typography } from "@mui/material";
import { styles } from "./useStyles";
import { communityOverViewDetails } from "@/public/data/community";

export default function ToolList() {
  return (
    <Grid container spacing={2} sx={styles.toolsContainerBoxStyle}>
      {communityOverViewDetails.map((data, index) => (
        <Grid item>
          <Grid container sx={styles.toolsContainerListBoxGridStyle}>
            <Grid item xs={12}>
              {" "}
              <Typography
                variant="h6"
                sx={styles.toolsContainerListBoxGridToolNameTypography}
              >
                {data.communityName}
              </Typography>
            </Grid>
            <Grid item xs={4}>
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
                <Typography
                  sx={
                    styles.toolsContainerListBoxGridCommunicationCountTypography
                  }
                >
                  {data.increaseInLastMonth}
                </Typography>
              </Paper>
            </Grid>
            <Grid item xs={8}>
              {" "}
              <Typography>{data.totalParticipants}</Typography>
            </Grid>
          </Grid>
        </Grid>
      ))}
    </Grid>
  );
}
