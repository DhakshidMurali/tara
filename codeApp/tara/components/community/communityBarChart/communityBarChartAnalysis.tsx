"use client";

import { Close } from "@mui/icons-material";
import {
  Grid,
  Stack,
  Button,
  Typography,
  Box,
  IconButton,
  Checkbox,
} from "@mui/material";
import { styles } from "./useStyles";
import { purple } from "@mui/material/colors";

export function AnalysisHeader() {
  return (
    <Grid item xs={4} sx={{ height: "20%" }}>
      <Stack direction={"row"} spacing={1} height={"100%"}>
        <Button>
          <Typography
            variant="h6"
            sx={styles.toolsContainerListBoxGridSelectGridItemTypographyStyle}
            textAlign="center"
          >
            Analysis
          </Typography>
        </Button>

        <Box
          sx={styles.toolsContainerListBoxGridSelectGridItemSplittingBoxStyle}
        ></Box>
        <Grid
          container
          sx={{
            height: "100%",
            width: "90%",
          }}
        >
          {[1, 1, 1, 1, 1].map(() => {
            return (
              <Grid item xs={3.5} marginRight={1}>
                <Box
                  sx={
                    styles.toolsContainerListBoxGridSelectGridItemGridItemStyle
                  }
                >
                  <IconButton sx={{ padding: 0 }}>
                    <Close sx={{ alignSelf: "center" }}></Close>
                  </IconButton>
                  <Typography
                    variant="body2"
                    sx={{ alignSelf: "center", color: "white" }}
                  >
                    Cloud-Neo..
                  </Typography>
                </Box>
              </Grid>
            );
          })}
        </Grid>
      </Stack>
    </Grid>
  );
}

export function AnalysisCommunitySelection() {
  return (
    <Grid item xs={4} sx={{ height: "80%" }}>
      <Grid container sx={{ height: "100%" }}>
        <Grid item xs={5}>
          {[1, 1, 1, 1, 1].map((data, index) => {
            return (
              <Stack
                direction={"row"}
                sx={
                  styles.toolsContainerListBoxGridSelectGridItemGridItemStackStyle
                }
                marginTop={0.5}
              >
                <Typography
                  variant="subtitle1"
                  sx={
                    styles.toolsContainerListBoxGridSelectGridItemGridItemStackTypographyListStyle
                  }
                >
                  Cloud Computing and DevOps
                </Typography>
              </Stack>
            );
          })}
        </Grid>
        <Grid item xs={0.3}>
          {[1, 1, 1, 1, 1].map((data, index) => {
            if (index == 1) {
              return (
                <Box
                  sx={{
                    height: "16%",
                    backgroundColor: "orange",
                    marginTop: "4px",
                  }}
                ></Box>
              );
            }
          })}
        </Grid>
        <Grid
          item
          xs={6.7}
          sx={
            styles.toolsContainerListBoxGridSelectGridItemGridItemStackCommunityListStyle
          }
        >
          {[1, 1, 1, 1, 1, 1, 1, 1].map((data, index) => {
            return (
              <Stack
                direction={"row"}
                marginTop={1}
                sx={{
                  display: "flex",

                  alignmentBaseline: "baseline",
                }}
              >
                <Checkbox
                  sx={{
                    "& .MuiSvgIcon-root": { fontSize: 20 },
                    color: purple[800],
                    "&.Mui-checked": {
                      color: purple[600],
                    },
                  }}
                />
                <Typography
                  variant="subtitle1"
                  sx={
                    styles.toolsContainerListBoxGridSelectGridItemGridItemStackTypographyCommunityListStyle
                  }
                >
                  Cloud Computing and DevOps
                </Typography>
              </Stack>
            );
          })}
        </Grid>
      </Grid>
    </Grid>
  );
}
