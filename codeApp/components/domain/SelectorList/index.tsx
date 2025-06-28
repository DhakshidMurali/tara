"use client";
import { DepartmentList } from "@/public/Sample/data";
import {
  KeyboardArrowLeftOutlined,
  KeyboardArrowRightOutlined
} from "@mui/icons-material";
import {
  Box,
  Button,
  Checkbox,
  Grid,
  Stack,
  Typography
} from "@mui/material";
import { useState } from "react";
import { styles } from "./useStyles";

export default function DomainSelectorList() {
  const departmentList = DepartmentList;
  const [expandedRow, setExpandedRow] = useState(-1);
  const handleClick = (index) => {
    expandedRow != -1 ? (index = -1) : null;
    setExpandedRow(index);
  };

  if (expandedRow != -1) {
    return (
      <Grid size={{ xs: 4 }} component="div">
        <Grid container spacing={1} sx={styles.departmentContainerBoxStyle}>
          <Grid component="div" size={{ xs: 12 }} >
            <Typography
              variant="h4"
              sx={styles.dashboardTextTypographStyle}
            >
              Domains
            </Typography>
          </Grid>
          <Grid component="div" size={{ xs: 12 }} >
            <Grid container sx={styles.departmentContainerListBoxStyle}>
              <Grid component={"div"} size={{ xs: 12 }}>
                <Stack
                  direction={"row"}
                  sx={styles.departmentContainerListBoxStackStyle}
                >
                  <Typography
                    variant="h6"
                    sx={styles.dashboardTextTypographStyle}
                    paddingLeft={2}
                  >
                    {DepartmentList[expandedRow].Domain}
                  </Typography>
                  <Button
                    sx={{ marginRight: 4 }}
                    onClick={() => handleClick(expandedRow)}
                  >
                    <KeyboardArrowLeftOutlined
                      sx={{
                        color: "primary.contrastText",
                      }}
                    ></KeyboardArrowLeftOutlined>
                  </Button>
                </Stack>
              </Grid>
              <Grid component={"div"} size={{ xs: 12 }} sx={{ height: "24rem" }}>
                <Box
                  flexWrap={"wrap"}
                  sx={styles.departmentContainerListBoxBoxExpandedStyle}
                >
                  {DepartmentList[expandedRow].Topics.map((data, index) => {
                    return (
                      <Box
                        key={index}
                        sx={styles.departmentContainerListBoxStackTopicBoxStyle}
                      >
                        <Typography sx={{ color: "primary.contrastText" }}>
                          {data}
                        </Typography>
                      </Box>
                    );
                  })}
                </Box>
              </Grid>
            </Grid>
          </Grid>
        </Grid>
      </Grid >
    );
  }
  return (
    <Grid size={{ xs: 4 }} component="div">
      <Grid container sx={styles.departmentContainerBoxStyle} >
        <Grid component="div" size={{ xs: 12 }} >
          <Typography
            variant="h4"
            sx={styles.dashboardTextTypographStyle}
          >
            Domains
          </Typography>
        </Grid>
        <Grid component="div" size={{ xs: 12 }} sx={{ height: "28rem" }}>
          <Box
            flexWrap={"wrap"}
            sx={styles.departmentContainerListBoxBoxExpandedStyle}
          >
            {departmentList.map((data, index) => {
              return (
                <Box sx={styles.departmentContainerListBoxStyle}>
                  <Stack
                    direction={"row"}
                    sx={styles.departmentContainerListBoxStackStyle}
                  >
                    <Box sx={{ display: "flex", flexDirection: "row" }}>
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
                        variant="h6"
                        sx={styles.dashboardTextTypographStyle}
                        paddingLeft={2}
                      >
                        {DepartmentList[index].Domain}
                      </Typography>
                    </Box>
                    <Button
                      sx={{ marginRight: 4 }}
                      onClick={() => handleClick(index)}
                    >
                      <KeyboardArrowRightOutlined
                        sx={{
                          color: "primary.contrastText",
                        }}
                      ></KeyboardArrowRightOutlined>
                    </Button>
                  </Stack>
                </Box>
              );
            })}
          </Box>
        </Grid>
      </Grid>
    </Grid >
  );
}
