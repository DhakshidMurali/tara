"use client";

import { Close } from "@mui/icons-material";
import {
  Box,
  Button,
  Checkbox,
  Grid,
  IconButton,
  Stack,
  Typography,
} from "@mui/material";
import { styles } from "./useStyles";

export function AnalysisHeader() {
  return (
    <Grid item xs={4} sx={{ height: "20%" }}>
      <Stack direction={"row"} spacing={1} height={"100%"} marginLeft={1}>
        <Button sx={{ backgroundColor: "primary.light", padding: 2 }}>
          Change Domains
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
                  <IconButton
                    sx={{
                      padding: 0,
                      "& .css-55fg2l-MuiSvgIcon-root": {
                        height: "0.7em",
                        width: "0.7em",
                      },
                    }}
                  >
                    <Close sx={{ alignSelf: "center" }}></Close>
                  </IconButton>
                  <Typography
                    variant="body2"
                    sx={{
                      alignSelf: "center",
                      color: "primary.contrastText",
                      whiteSpace: "nowrap",
                      overflow: "hidden",
                      textOverflow: "ellipsis",
                      padding: 0.5,
                    }}
                  >
                    Cloud Compusdsdst..
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
      <Grid container sx={{ height: "100%", paddingRight: 1 }} paddingLeft={1}>
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
                    backgroundColor: "secondary.dark",
                    marginTop: "4px",
                  }}
                ></Box>
              );
            } else {
              return (
                <Box
                  sx={{
                    height: "16%",
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
          {[1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1].map(
            (data, index) => {
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
                      color: "rgb(46, 36,57)"[100],
                      "&.Mui-checked": {
                        color: "rgb(33,24,43)",
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
            }
          )}
        </Grid>
      </Grid>
    </Grid>
  );
}
