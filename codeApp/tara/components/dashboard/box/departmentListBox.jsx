"use client";
import { Avatar, Box, Button, Grid, Stack, Typography } from "@mui/material";
import { styles } from "../useStyles";
import {
  Domain,
  KeyboardArrowLeftOutlined,
  KeyboardArrowRightOutlined,
} from "@mui/icons-material";
import { DepartmentList } from "@/public/Sample/data";
import { useState } from "react";

export default function DepartmentListBox(props) {
  const departmentList = DepartmentList;
  const [expandedRow, setExpandedRow] = useState(-1);
  const handleClick = (index) => {
    expandedRow != -1 ? (index = -1) : null;
    setExpandedRow(index);
  };

  if (expandedRow != -1) {
    return (
      <Box sx={styles.departmentContainerListBoxStyle}>
        <Stack
          direction={"row"}
          sx={styles.departmentContainerListBoxStackStyle}
        >
          <Box sx={{ display: "flex", flexDirection: "row" }}>
            <Avatar>
              <Domain sx={styles.deparmentListIconStyle}></Domain>
            </Avatar>
            <Typography
              variant="h6"
              sx={styles.dashboardTextTypographStyle}
              paddingLeft={2}
            >
              {DepartmentList[expandedRow].Domain}
            </Typography>
          </Box>
          <Button
            sx={{ marginRight: 4 }}
            onClick={() => handleClick(expandedRow)}
          >
            <KeyboardArrowLeftOutlined
              sx={{
                color: "white",
              }}
            ></KeyboardArrowLeftOutlined>
          </Button>
        </Stack>
        <Stack
          direction={"row"}
          flexWrap={"wrap"}
          overflow={"scroll"}
          sx={styles.departmentContainerListBoxStackExpandedStyle}
        >
          {DepartmentList[expandedRow].Topics.map((data, index) => {
            return (
              <Box
                key={index}
                sx={styles.departmentContainerListBoxStackTopicBoxStyle}
              >
                <Typography variant="h6" sx={{ color: "primary.contrastText" }}>
                  {data}
                </Typography>
              </Box>
            );
          })}
        </Stack>
      </Box>
    );
  }
  return departmentList.map((data, index) => {
    return (
      <Box sx={styles.departmentContainerListBoxStyle}>
        <Stack
          direction={"row"}
          sx={styles.departmentContainerListBoxStackStyle}
        >
          <Box sx={{ display: "flex", flexDirection: "row" }}>
            <Avatar>
              <Domain sx={styles.deparmentListIconStyle}></Domain>
            </Avatar>
            <Typography
              variant="h6"
              sx={styles.dashboardTextTypographStyle}
              paddingLeft={2}
            >
              {DepartmentList[index].Domain}
            </Typography>
          </Box>
          <Button sx={{ marginRight: 4 }} onClick={() => handleClick(index)}>
            <KeyboardArrowRightOutlined
              sx={{
                color: "white",
              }}
            ></KeyboardArrowRightOutlined>
          </Button>
        </Stack>
      </Box>
    );
  });
}
