"use client";

import NavBar from "@/components/common/navBar/navBar";
import { Grid } from "@mui/material";

import DomainBarChart from "@/components/domain/BarChart";
import { DomainHeader } from "@/components/domain/Header";
import { DomainList } from "@/components/domain/List/List";
import DomainSelectorList from "@/components/domain/SelectorList";

export default function Domain() {
  return (
    <Grid container>
      <Grid size={{ xs: 2 }} component="div" sx={{ marginRight: 2 }}>
        <NavBar selected="domain" />
      </Grid>
      <Grid size={{ xs: 9.75 }} component="div" sx={{ marginTop: 2 }}>
        <Grid container spacing={1}>
          <DomainHeader></DomainHeader>
          <DomainList></DomainList>
          <DomainBarChart></DomainBarChart>
          <DomainSelectorList></DomainSelectorList>
        </Grid>
      </Grid>
    </Grid>
  );
}
