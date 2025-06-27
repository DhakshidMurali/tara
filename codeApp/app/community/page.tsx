"use client";
import NavBar from "@/components/common/navBar/navBar";
import CommunityBarChart from "@/components/community/BarChart";
import CommmunityHeader from "@/components/community/Header";
import CommunityList from "@/components/community/List";
import { Grid } from "@mui/material";

export default function Community() {
  let time = new Date();
  const label = { inputProps: { "aria-label": "Checkbox demo" } };
  return (
    <>
      <Grid size={{ xs: 2 }} component="div" sx={{ marginRight: 2 }}>
        <NavBar selected="community" />
      </Grid>
      <Grid container size={{ xs: 9.75 }} spacing={1} sx={{ marginTop: 2 }}>
        <CommmunityHeader></CommmunityHeader>
        <CommunityList></CommunityList>
        {/* <CommunityBarChart></CommunityBarChart> */}
      </Grid>
    </>
  );
}
