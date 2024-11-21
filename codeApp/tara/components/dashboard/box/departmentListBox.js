"use client";
<<<<<<< HEAD
import {
  Avatar,
  Box,
  Button,
  Collapse,
  Grid,
  Stack,
  Typography,
} from "@mui/material";
=======
import { Avatar, Box, Button, Grid, Stack, Typography } from "@mui/material";
>>>>>>> dd64d5ba35af4a29e638d731b832199630c785c6
import { styles } from "../useStyles";
import { Domain, ExpandMore } from "@mui/icons-material";
import { useRouter } from "next/navigation";
import { DepartmentList } from "@/public/Sample/data";
<<<<<<< HEAD
import { useState } from "react";

export default function DepartmentListBox(props) {
  const router = useRouter();
  const departmentList = DepartmentList;
  const [selectedDropDown, setSelectedDropDownOpen] = useState();
  const handleDropDownDepartmentList = (value) => {
    if (value == selectedDropDown) {
      setSelectedDropDownOpen(-1);
      return;
    }
    setSelectedDropDownOpen(value);
=======

export default function DepartmentListBox(props) {
  const router = useRouter();
  let isDropDownOpen = false;
  const departmentList = DepartmentList;
  const handleDropDownDepartmentList = () => {
    isDropDownOpen = !isDropDownOpen;
    router.refresh();
>>>>>>> dd64d5ba35af4a29e638d731b832199630c785c6
  };
  return (
    <Grid
      container
      sx={{
        justifyContent: "flex-start",
      }}
    >
      <Grid
        item
        xs={12}
        sx={{
          display: "flex",
          flexDirection: "column",
          justifyContent: "flex-start",
        }}
      >
        {departmentList.map((i, index) => {
          return (
            <Box sx={styles.departmentContainerListBoxStyle}>
<<<<<<< HEAD
              <Grid
                container
                spacing={2}
                sx={{
                  justifyContent: "flex-start",
                }}
              >
                <Grid item xs={12}>
                  <Stack direction={"row"}>
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
                    <Button
                      sx={{ margin: 0 }}
                      onClick={() => {
                        handleDropDownDepartmentList(index);
                      }}
                    >
                      <ExpandMore
                        sx={{ marginLeft: 0, color: "white" }}
                      ></ExpandMore>
                    </Button>
                  </Stack>
                </Grid>
                {selectedDropDown == index ? (
                  <Collapse
                    in={selectedDropDown == index}
                    timeout="auto"
                    orientation="vertical"
                    unmountOnExit
                  >
                    <Grid
                      item
                      xs={12}
                      sx={{
                        padding: 0,
                        display: "flex",
                        flexDirection: "column",
                        justifyContent: "center",
                      }}
                    >
                      <Stack direction={"row"}>
                        <Typography
                          variant="h6"
                          sx={styles.dashboardTextTypographStyle}
                          paddingLeft={2}
                        >
                          {selectedDropDown == index
                            ? departmentList[index]
                            : "Hello"}
                        </Typography>
                      </Stack>
                    </Grid>
                    <Grid item xs={12} sx={{ padding: 0 }}>
                      <Stack direction={"row"}>
                        <Typography
                          variant="h6"
                          sx={styles.dashboardTextTypographStyle}
                          paddingLeft={2}
                        >
                          {selectedDropDown == index
                            ? departmentList[index]
                            : "Hello"}
                        </Typography>
                      </Stack>
                    </Grid>
                  </Collapse>
                ) : (
                  <></>
                )}
              </Grid>
=======
              <Stack direction={"row"}>
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
                <Button
                  sx={{ margin: 0 }}
                  onClick={handleDropDownDepartmentList}
                >
                  <ExpandMore
                    sx={{ marginLeft: 0, color: "white" }}
                  ></ExpandMore>
                </Button>
              </Stack>
>>>>>>> dd64d5ba35af4a29e638d731b832199630c785c6
            </Box>
          );
        })}
      </Grid>
    </Grid>
  );
}

{
  /* <Box sx={styles.departmentContainerListBoxStyle}>
      <Stack direction={"row"}>
        <Avatar>
          <Domain sx={styles.deparmentListIconStyle}></Domain>
        </Avatar>
        <Typography
          variant="h6"
          sx={styles.dashboardTextTypographStyle}
          paddingLeft={2}
        >
          {isDropDownOpen
            ? "Computer Science Eng"
            : "Computer Science Engineering"}
        </Typography>
        <Button sx={{ margin: 0 }} onClick={handleDropDownDepartmentList}>
          <ExpandMore sx={{ marginLeft: 0, color: "white" }}></ExpandMore>
        </Button>
      </Stack>
    </Box> */
}
