"use client";
import NavBar from "@/components/common/navBar/navBar";
import ToolHeader from "@/components/tool/Header";
import ToolBarChart from "@/components/tool/BarChart";
import ToolBarChartAnalysis from "@/components/tool/SelectorList";
import ToolList from "@/components/tool/List";
import { Grid } from "@mui/material";
import useSWR from "swr";
import { GET_TOOL_LIST_ENDPOINT, GET_TOOL_SELECTOR_LIST_ENDPOINT, getToolList, getToolSelectorList } from "@/api/services/toolApi";
import { ToolListPayload, ToolSelectorListPayload, ToolStatsByLikes } from "@/api/types/tool";

export default function Tool() {
  const { data: toolList, isLoading: toolListLoading, error: toolListError } = useSWR<ToolListPayload[]>
    (GET_TOOL_LIST_ENDPOINT,
      getToolList,
      { revalidateOnFocus: false });

  const { data: toolSelectorList, isLoading: toolSelectorListLoading, error: toolSelectorListError } = useSWR<ToolSelectorListPayload[]>
    (GET_TOOL_SELECTOR_LIST_ENDPOINT,
      getToolSelectorList,
      { revalidateOnFocus: false });

  const { data: toolStatsByLikesList, isLoading: toolStatsByLikesListLoading, error: toolStatsByLikesListError } = useSWR<ToolStatsByLikes[]>
    (GET_TOOL_SELECTOR_LIST_ENDPOINT,
      getToolSelectorList,
      { revalidateOnFocus: false });

  if (toolListLoading || toolSelectorListLoading || toolStatsByLikesListLoading) {
    return <></>
  }

  console.log('toolSelectorList', toolSelectorList)

  console.log('toolList', toolList)
  return (
    <Grid container>
      <Grid size={{ xs: 2 }} component="div" sx={{ marginRight: 2 }}>
        <NavBar selected="tool" />
      </Grid>
      <Grid size={{ xs: 9.75 }} component="div" sx={{ marginTop: 2 }}>
        <Grid container spacing={1}>
          <ToolHeader></ToolHeader>
          <ToolList toolListPayload={toolList}></ToolList>
          <ToolBarChart></ToolBarChart>
          <ToolBarChartAnalysis toolSelectorList={toolSelectorList}></ToolBarChartAnalysis>
        </Grid>
      </Grid>
    </Grid >
  );
}
