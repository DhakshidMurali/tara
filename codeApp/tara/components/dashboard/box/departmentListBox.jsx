"use client";
import { Avatar, Box, Button, Grid, Stack, Typography } from "@mui/material";
import { styles } from "../useStyles";
import { Domain, ExpandLess, ExpandMore } from "@mui/icons-material";
import { DepartmentList } from "@/public/Sample/data";
import { useState } from "react";

export default function DepartmentListBox(props) {
  const departmentList = DepartmentList;

  const [expandedRow, setExpandedRow] = useState(-1);

  const handleClick = (index) => {
    expandedRow != -1 ? (index = -1) : null;
    setExpandedRow(index);
  };
  return departmentList.map((i, index) => {
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
              {departmentList[index]}
            </Typography>
          </Box>
          {expandedRow == index ? (
            <Button sx={{ marginRight: 4 }} onClick={() => handleClick(index)}>
              <ExpandLess
                sx={{
                  color: "white",
                }}
              ></ExpandLess>
            </Button>
          ) : (
            <Button sx={{ marginRight: 4 }} onClick={() => handleClick(index)}>
              <ExpandMore
                sx={{
                  color: "white",
                }}
              ></ExpandMore>
            </Button>
          )}
        </Stack>
      </Box>
    );
  });
}
