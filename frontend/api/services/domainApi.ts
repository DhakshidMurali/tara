export const GET_DOMAIN_SELECTOR_LIST_ENDPOINT =
  "http://localhost:8080/domain/get/list/domain-details";
export const GET_DOMAIN_LIST_ENDPOINT =
  "http://localhost:8080/domain/get/list/count-tools-by-domains";
export const GET_DELIVERY_FORMAT_GROUP_BY_DOMAINS =
  "http://localhost:8080/domain/get/list/group-delivery-format-by-domains";

export const getDomainSelectorList = async () => {
  const response = await fetch(GET_DOMAIN_SELECTOR_LIST_ENDPOINT);
  if (!response.ok) throw new Error("Network response was not ok");
  return await response.json();
};

export const getDomainList = async () => {
  const response = await fetch(GET_DOMAIN_LIST_ENDPOINT);
  if (!response.ok) throw new Error("Network response was not ok");
  return await response.json();
};

export const getDeliveryFormatByDomains = async () => {
  const response = await fetch(GET_DELIVERY_FORMAT_GROUP_BY_DOMAINS);
  if (!response.ok) throw new Error("Network response was not ok");
  return await response.json();
};
