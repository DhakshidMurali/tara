"use client";

import { communityOverViewDetails } from "@/public/data/community";
import { Group } from "@mui/icons-material";
import { Box, Button, Grid, Stack, Typography } from "@mui/material";
import { styles } from "./useStyles";

export default function CommunityList() {
  return (
    <Grid item xs={12} sx={{ height: "25%" }}>
      <Grid container sx={{ height: "100%" }}>
        <Grid item xs={12}>
          <Stack
            direction={"row"}
            display={"flex"}
            justifyContent={"space-between"}
          >
            <Typography variant="h3" sx={styles.communityTextTypographStyle}>
              Community{" "}
            </Typography>
            <Button sx={styles.createButtonStyle}>Create</Button>
          </Stack>
        </Grid>
        <Grid item xs={12} sx={{ height: "90%" }}>
          <Box sx={styles.toolsContainerBoxStyle}>
            <Grid container spacing={2} sx={{ flexWrap: "nowrap" }}>
              {communityOverViewDetails.map((data, index) => (
                <Grid item key={index} sx={{ height: "100%" }}>
                  <Box sx={styles.toolsContainerListBoxStyle}>
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
                    ;
                  </Box>
                </Grid>
              ))}
            </Grid>
          </Box>
        </Grid>
      </Grid>
    </Grid>
  );
}
