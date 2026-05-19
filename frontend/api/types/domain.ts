export type DomainSelectorListPayload = {
  domainName: string;
  subDomains: string[];
};

export type DomainListPayload = {
  domainName: string;
  toolsCount: number;
};

export type DeliveryStatsByDomain = {
  domainName: string;
  deliveryFormat: string;
  toolsCount: number;
};
export type DeliveryStatsFormatData = {
  domainName: string;
  deliveryFormat: {
    format: string;
    count: number;
    percent: number;
  }[];
  toolsCount: number;
};
