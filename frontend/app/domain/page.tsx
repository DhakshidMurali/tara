"use client";

import NavBar from "@/components/common/navBar/navBar";
import { Grid } from "@mui/material";

import { GET_DELIVERY_FORMAT_GROUP_BY_DOMAINS, GET_DOMAIN_LIST_ENDPOINT, GET_DOMAIN_SELECTOR_LIST_ENDPOINT, getDeliveryFormatByDomains, getDomainList, getDomainSelectorList } from "@/api/services/domainApi";
import { DeliveryStatsByDomain, DomainListPayload, DomainSelectorListPayload } from "@/api/types/domain";
import DomainBarChart from "@/components/domain/BarChart";
import { DomainHeader } from "@/components/domain/Header";
import { DomainList } from "@/components/domain/List/List";
import DomainSelectorList from "@/components/domain/SelectorList";
import useSWR from "swr";

export default function Domain() {
  const { data: domainSelectorList, isLoading: domainSelectorListLoading, error: domainSelectorListError } = useSWR<DomainSelectorListPayload[]>
    (GET_DOMAIN_SELECTOR_LIST_ENDPOINT,
      getDomainSelectorList,
      { revalidateOnFocus: false });

  const { data: domainList, isLoading: domainListLoading, error: domainListError } = useSWR<DomainListPayload[]>
    (GET_DOMAIN_LIST_ENDPOINT,
      getDomainList,
      { revalidateOnFocus: false });

  const { data: deliveryStatsByDomain, isLoading: deliveryStatsByDomainLoading, error: DeliveryStatsByDomainError } = useSWR<DeliveryStatsByDomain[]>
    (GET_DELIVERY_FORMAT_GROUP_BY_DOMAINS,
      getDeliveryFormatByDomains,
      { revalidateOnFocus: false });

  if (domainSelectorListLoading || domainListLoading || deliveryStatsByDomainLoading) {
    return <></>
  }
  console.log('deliveryStatsByDomain', deliveryStatsByDomain)
  return (
    <Grid container>
      <Grid size={{ xs: 2 }} component="div" sx={{ marginRight: 2 }}>
        <NavBar selected="domain" />
      </Grid>
      <Grid size={{ xs: 9.75 }} component="div" sx={{ marginTop: 2 }}>
        <Grid container spacing={1}>
          <DomainHeader></DomainHeader>
          <DomainList domainList={domainList}></DomainList>
          <DomainBarChart deliveryStatsByDomain={deliveryStatsByDomain}></DomainBarChart>
          <DomainSelectorList domainSelectorList={domainSelectorList}></DomainSelectorList>
        </Grid>
      </Grid>
    </Grid>
  );
}
