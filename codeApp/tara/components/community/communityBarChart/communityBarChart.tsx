"use client";

import { Dataset } from "@/public/Sample/data";
import { addLabels } from "@/utils/functions";
import { Box, Grid, Typography } from "@mui/material";
import { BarChart } from "@mui/x-charts/BarChart";
import {
  AnalysisCommunitySelection,
  AnalysisHeader,
} from "./communityBarChartAnalysis";
import { styles } from "./useStyles";

export default function CommunityBarChart() {
  return (
    <Grid container sx={{ height: "65%" }}>
      <Grid item xs={12} sx={{ height: "100%" }} >
        <Grid container sx={{ height: "100%" }}>
          <Grid item xs={8} sx={{ height: "20%" }}>
            <Typography variant="h4" sx={styles.communityTextTypographStyle}>
              Find Community works for your
            </Typography>
          </Grid>
          {/* Grid Item xs={4} */}
          <AnalysisHeader></AnalysisHeader>
          <Grid item xs={8} sx={{ height: "80%" }}>
            <BarChart
              dataset={Dataset}
              xAxis={[
                {
                  scaleType: "band",
                  dataKey: "month",
                },
              ]}
              series={addLabels([
                { dataKey: "DevOps", stack: "assets" },
                { dataKey: "Web Devlopment", stack: "assets" },
                { dataKey: "App Devlopment", stack: "liability" },
                {
                  dataKey: "Internet of Things",
                  stack: "liability",
                },
                { dataKey: "Cloud Computing", stack: "equity" },
                { dataKey: "DevOps", stack: "assets" },
                { dataKey: "Web Devlopment", stack: "assets" },
                { dataKey: "App Devlopment", stack: "liability" },
                {
                  dataKey: "Internet of Things",
                  stack: "liability",
                },
                { dataKey: "Cloud Computing", stack: "equity" },
              ])}
              slotProps={{
                legend: { hidden: true },
              }}
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
              }}
              tooltip={{ trigger: "item" }}
            />
          </Grid>
          <AnalysisCommunitySelection></AnalysisCommunitySelection>
        </Grid>
      </Grid>
    </Grid>
  );
}
