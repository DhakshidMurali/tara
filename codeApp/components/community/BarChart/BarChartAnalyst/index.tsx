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
import { styles } from "../useStyles";

export function AnalysisHeader() {
  return (
    <Grid size={{ xs: 4 }} sx={{ height: "20%" }}>
      <Stack direction={"row"} spacing={1} height={"100%"} marginLeft={1} sx={{
        // Learning : In the specific scope of styling, we are applying the changes
        "& > :first-child.MuiButtonBase-root:hover": {
          backgroundColor: "primary.light",
          fontWeight: "bold",
          padding: 1.7,
          zIndex: 1
        }
      }}>
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
              <Grid size={{ xs: 3.5 }} marginRight={1}>
                <Box
                  sx={
                    styles.toolsContainerListBoxGridSelectGridItemGridItemStyle
                  }
                >
                  <IconButton
                    sx={{
                      padding: 0,
                      zIndex: 0,
                      "& .css-55fg2l-MuiSvgIcon-root": {
                        height: "0.4em",
                        width: "0.2em",
                      },
                    }}
                  >
                    <Close sx={{ alignSelf: "center", color: "primary.contrastText" }}></Close>
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
    </Grid >
  );
}

export function AnalysisCommunitySelection() {
  return (
    <Grid size={{ xs: 12 }} sx={{ height: "80%" }}>
      <Grid container sx={{ height: "100%", paddingRight: 1 }} paddingLeft={1}>
        <Grid size={{ xs: 12 }}>
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
        <Grid size={{ xs: 0.3 }}>
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

          size={{ xs: 6.7 }}
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
