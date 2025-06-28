"use client";
import NavBar from "@/components/common/navBar/navBar";
import ToolHeader from "@/components/tool/Header";
import ToolBarChart from "@/components/tool/toolBarChart/toolBarChart";
import ToolBarChartAnalysis from "@/components/tool/toolBarChart/toolBarChartAnalysis";
import ToolList from "@/components/tool/toolList/toollist";
import { Grid } from "@mui/material";

export default function Tool() {
  return (
    <Grid container>
      <Grid size={{ xs: 2 }} component="div" sx={{ marginRight: 2 }}>
        <NavBar selected="tool" />
      </Grid>
      <Grid size={{ xs: 9.75 }} component="div" sx={{ marginTop: 2 }}>
        <Grid container spacing={1}>
          <ToolHeader></ToolHeader>
          <ToolList></ToolList>
          <ToolBarChart></ToolBarChart>
          <ToolBarChartAnalysis></ToolBarChartAnalysis>
        </Grid>
      </Grid>
    </Grid >
  );
}
