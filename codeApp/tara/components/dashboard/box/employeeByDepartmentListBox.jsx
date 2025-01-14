import { Box, Stack, Tooltip, Typography } from "@mui/material";
import { styles } from "../useStyles";
import { DashBoardPageEmployeeByDepartmentBoxLabelData } from "@/public/data/const";

export default function EmployeeByDepartmentListBox(props) {
  const { departmentList } = props;
  const dashBoardPageEmployeeByDepartmentBoxLabelData =
    DashBoardPageEmployeeByDepartmentBoxLabelData;
  console.log(dashBoardPageEmployeeByDepartmentBoxLabelData);
  return (
    <Stack direction={"row"} spacing={2} sx={{ width: "100%" }}>
      {/* Creating List of Department Name */}
      <Stack direction={"column"} spacing={5}>
        {departmentList.map((i, index) => (
          <Typography variant="h6" sx={styles.dashboardTextTypographStyle}>
            {i}
          </Typography>
        ))}
      </Stack>
      {/* Creating List of Department Employee Bar */}
      <Stack
        direction={"column"}
        spacing={5}
        padding={0.5}
        sx={{ width: "70%" }}
      >
        {/* Creating List of Department Employee Bar in a row */}

        {departmentList.map((i, index) => (
          <Stack
            direction={"row"}
            sx={{ width: "100%" }}
            spacing={1}
            padding={1.5}
          >
            <Tooltip
              title="1000 [100000]"
              componentsProps={{
                tooltip: {
                  sx: {
                    ...styles.employeeByDepartmentListBoxBarBoxToolTipStyle,
                    backgroundColor: "rgb(141,100,200)",
                  },
                },
              }}
              placement="top-start"
              arrow
            >
              <Box
                sx={{
                  ...styles.employeeByDepartmentListBoxBarStyle,
                  backgroundColor: "rgb(141,100,200)",
                }}
                width={"30%"}
              ></Box>
            </Tooltip>
            <Tooltip
              title="100 [100000]"
              componentsProps={{
                tooltip: {
                  sx: {
                    ...styles.employeeByDepartmentListBoxBarBoxToolTipStyle,
                    backgroundColor: "rgb(228,172,99)",
                  },
                },
              }}
              placement="top"
              arrow
            >
              <Box
                sx={{
                  ...styles.employeeByDepartmentListBoxBarStyle,
                  backgroundColor: "rgb(228,172,99)",
                }}
                width={"20%"}
              ></Box>
            </Tooltip>
            <Tooltip
              title="1000 [100000]"
              componentsProps={{
                tooltip: {
                  sx: {
                    ...styles.employeeByDepartmentListBoxBarBoxToolTipStyle,
                    backgroundColor: "rgb(54,157,63)",
                  },
                },
              }}
              placement="top"
              arrow
            >
              <Box
                sx={{
                  ...styles.employeeByDepartmentListBoxBarStyle,
                  backgroundColor: "rgb(54,157,63)",
                }}
                width={"35%"}
              ></Box>
            </Tooltip>
            <Tooltip
              title="1000 [100000]"
              componentsProps={{
                tooltip: {
                  sx: {
                    ...styles.employeeByDepartmentListBoxBarBoxToolTipStyle,
                    backgroundColor: "rgb(174,199,243)",
                  },
                },
              }}
              placement="top-end"
              arrow
            >
              <Box
                sx={{
                  ...styles.employeeByDepartmentListBoxBarStyle,
                  backgroundColor: "rgb(174,199,243)",
                }}
                width={"15%"}
              ></Box>
            </Tooltip>
          </Stack>
        ))}
        {/* Creating row of Legend for Employee By Department Color Bar */}
        <Stack direction={"row"} spacing={4} sx={{ width: "100%" }}>
          {dashBoardPageEmployeeByDepartmentBoxLabelData.map((i, index) => (
            <Stack
              direction={"row"}
              alignItems={"baseline"}
              sx={{ width: "25%" }}
            >
              <Box
                sx={{
                  ...styles.employeeByDepartmentListBoxLabelBoxStyle,
                  backgroundColor: i["Color"],
                }}
              ></Box>
              <Typography
                variant="h6"
                sx={styles.employeeByDepartmentListBoxLabelTypographyStyle}
              >
                {i["LabelName"]}
              </Typography>
            </Stack>
          ))}
        </Stack>
      </Stack>
    </Stack>
  );
}
