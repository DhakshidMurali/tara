"use client";

import { getRandomColor } from "@/utils/functions";
import { Box, Grid, Stack, Typography } from "@mui/material";
import { styles } from "./useStyles";

export function DashBoardList() {
  return (
    <Box sx={styles.toolsContainerBoxStyle}>
      <Grid container spacing={2} sx={{ flexWrap: "nowrap" }}>
        {Array.from({ length: 10 }).map((_, index) => (
          <Grid item key={index}>
            <Box sx={styles.toolsContainerListBoxStyle}>
              <Stack
                direction={"row"}
                justifyContent={"flex-start"}
                alignItems={"center"}
                sx={styles.toolsContainerListBoxStackStyle}
              >
                {/* Learning : To get sx props and inline styling alone applied, we do following way  */}
                <Box
                  sx={{
                    ...styles.toolsContainerListBoxBoxStyle,
                    bgcolor: getRandomColor(index),
                  }}
                ></Box>
                <Grid container paddingRight={"4px"}>
                  {/* Grid for Showing Count of tools group by Domain */}
                  <Grid item xs={12}>
                    <Typography
                      variant="h3"
                      sx={styles.toolsContainerListBoxCountTypographyStyle}
                    >
                      79
                    </Typography>
                  </Grid>
                  <Grid item xs={12}>
                    <Typography
                      variant="h6"
                      sx={styles.toolsContainerListBoxDomainTypographyStyle}
                    >
                      Cloud Computing
                    </Typography>
                  </Grid>
                </Grid>
              </Stack>
            </Box>
          </Grid>
        ))}
      </Grid>
    </Box>
  );
}
