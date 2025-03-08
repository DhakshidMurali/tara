"use client";

import { Dataset } from "@/public/Sample/data";
import { addLabels } from "@/utils/functions";
import { Box, Grid, IconButton, Typography } from "@mui/material";
import { BarChart } from "@mui/x-charts/BarChart";
import {
  AnalysisCommunitySelection,
  AnalysisHeader,
} from "./communityBarChartAnalysis";
import { styles } from "./useStyles";
import { RefreshOutlined, RefreshTwoTone } from "@mui/icons-material";

export default function CommunityBarChart() {
  return (
    <Grid container sx={{ height: "100%" }}>
      <Grid
        item
        xs={8}
        sx={{
          height: "20%",
          backgroundColor: "secondary.light",
          borderTopRightRadius: "16px",
          borderTopLeftRadius: "16px",
          display: "flex",
          flexDirection: "row",
          alignItems: "center",
          justifyContent: "space-between",
        }}
      >
        <Typography variant="h4" sx={styles.communityTextTypographStyle}>
          Find Community works for your
        </Typography>
        <IconButton
          sx={{ padding: 1, bgcolor: "secondary.contrastText", marginRight: 8 }}
        >
          {" "}
          <RefreshOutlined></RefreshOutlined>
        </IconButton>
      </Grid>
      <AnalysisHeader></AnalysisHeader>
      <Grid
        item
        xs={8}
        sx={{
          height: "80%",
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
            "& .MuiChartsAxis-bottom .MuiChartsAxis-tick": {
              stroke: "#f3def5",
            },
          }}
          tooltip={{ trigger: "item" }}
        />
      </Grid>
      <AnalysisCommunitySelection></AnalysisCommunitySelection>
    </Grid>
  );
}
