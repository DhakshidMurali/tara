"use client";
import NavBar from "@/components/common/navBar/navBar";
import CommunityBarChart from "@/components/community/BarChart";
import CommmunityHeader from "@/components/community/Header";
import CommunityList from "@/components/community/List";
import { CommunitySelectorList } from "@/components/community/selectorList";
import { Grid } from "@mui/material";

export default function Community() {
  let time = new Date();
  const label = { inputProps: { "aria-label": "Checkbox demo" } };
  return (
    <Grid container>
      <Grid size={{ xs: 2 }} component="div" sx={{ marginRight: 2 }}>
        <NavBar selected="community" />
      </Grid>
      <Grid size={{ xs: 9.75 }} component={"div"} >
        <Grid container spacing={1} sx={{ marginTop: 2 }}>
          <CommmunityHeader></CommmunityHeader>
          <CommunityList></CommunityList>
          <CommunityBarChart></CommunityBarChart>
          <CommunitySelectorList></CommunitySelectorList>
        </Grid>
      </Grid>
    </Grid>
  );
}
