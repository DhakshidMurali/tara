export type ToolListPayload = {
  toolName: string;
  subDomain: string;
  domainName: string;
  communicationCount: number;
};

export type ToolSelectorListPayload = {
  subDomain: {
    subDomainName: string;
    tool: string[];
  }[];
  domainName: string;
};

export type ToolStatsByLikesPayload = {
  toolName: string;
  likes: number;
};
