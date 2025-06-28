"use client";

import { Dataset } from "@/public/Sample/data";
import { addLabels } from "@/utils/functions";
import { RefreshOutlined } from "@mui/icons-material";
import { Grid, IconButton, Stack, Typography } from "@mui/material";
import { BarChart } from "@mui/x-charts/BarChart";
import {
  AnalysisCommunitySelection,
  AnalysisHeader,
} from "./BarChartAnalyst/index";
import { styles } from "./useStyles";

export default function CommunityBarChart() {
  return (
    <Grid container size={{ xs: 12 }}>
      <Grid
        size={{ xs: 8 }}
        sx={{
          backgroundColor: "secondary.light",
          borderTopRightRadius: "16px",
          borderTopLeftRadius: "16px",
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
          justifyContent: "space-between",
        }}
      >
        <Stack direction={"row"}>
          <Typography variant="h4" sx={styles.communityTextTypographStyle}>
            Find Community works for your
          </Typography>
          <IconButton
            sx={{
              padding: 1, bgcolor: "secondary.contrastText", marginRight: 8, "&   .MuiButtonBase-root:hover": {
                backgroundColor: "primary.light",
                fontWeight: "bold",
                padding: 1.7,
                zIndex: 1
              }
            }}
          >
            {" "}
            <RefreshOutlined></RefreshOutlined>
          </IconButton>
        </Stack>
        <Grid
          size={{ xs: 12 }}
          sx={{
            backgroundColor: "secondary.light",
            borderBottomRightRadius: "16px",
            borderBottomLeftRadius: "16px",
            paddingLeft: "8px",
          }}
        >
          <BarChart
            dataset={Dataset}
            xAxis={[
              {
                scaleType: "band",
                dataKey: "month",
                tickPlacement: "middle",
                tickLabelPlacement: "middle",
              },
            ]}
            series={addLabels([
              { dataKey: "DevOps Communication", stack: "assets" },
              { dataKey: "Web Devlopment Communication", stack: "liability" },
              { dataKey: "App Devlopment Communication", stack: "equity" },
              {
                dataKey: "Internet of Things Communication",
                stack: "assets",
              },
              { dataKey: "Cloud Computing Communication", stack: "liability" },
            ])}
            // slotProps={{
            //   legend: { hidden: true },
            // }}
            sx={{
              /* Left Axis Label Font Styling */
              "& .MuiChartsAxis-left .MuiChartsAxis-tickLabel": {
                strokeWidth: "1",
                fill: "#f3def5",
              },
              /* Bottom Label Font Styling */
              "& .MuiChartsAxis-bottom .MuiChartsAxis-tickLabel": {
                strokeWidth: "1",
                fill: "#f3def5",
                fontSize: "24px",
              },
              /* X axis Line Styling */
              "& .MuiChartsAxis-bottom .MuiChartsAxis-line": {
                stroke: "#f3def5",
                strokeWidth: 2.4,
              },
              /* Y Axis Line Styling  */
              "& .MuiChartsAxis-left .MuiChartsAxis-line": {
                stroke: "#f3def5",
                strokeWidth: 2.4,
              },
              "& .MuiChartsAxis-bottom .MuiChartsAxis-tick": {
                stroke: "#f3def5",
              },
            }}
          // tooltip={{ trigger: "item" }}
          />
        </Grid>
      </Grid>
      <Grid size={{ xs: 2 }}>
        {/* <AnalysisHeader></AnalysisHeader>
        <AnalysisCommunitySelection></AnalysisCommunitySelection> */}
      </Grid>
    </Grid>
  );
}
