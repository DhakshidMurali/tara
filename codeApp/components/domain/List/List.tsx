"use client";

import { getRandomColor } from "@/utils/functions";
import { Box, Grid, Stack, Typography } from "@mui/material";
import { styles } from "./useStyles";

export function DomainList() {
  return (
    <Grid component={"div"} size={{ xs: 12 }} spacing={2} height={"14rem"}>
      <Box sx={styles.toolsContainerBoxStyle}>
        {Array.from({ length: 10 }).map((_, index) => (
          <Box>
            <Grid container sx={styles.toolsContainerListBoxStyle}>
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
                  <Grid size={{ xs: 12 }}>
                    <Typography
                      variant="h3"
                      sx={styles.toolsContainerListBoxCountTypographyStyle}
                    >
                      79
                    </Typography>
                  </Grid>
                  <Grid size={{ xs: 12 }}>
                    <Typography
                      variant="h6"
                      sx={styles.toolsContainerListBoxDomainTypographyStyle}
                    >
                      Cloud Computing
                    </Typography>
                  </Grid>
                </Grid>
              </Stack>
            </Grid>
          </Box>
        ))}
      </Box>
    </Grid >
  );
}
