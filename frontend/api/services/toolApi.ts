export const GET_TOOL_LIST_ENDPOINT =
  "http://localhost:8080/tool/get/list/count-tools-by-communication";
export const GET_TOOL_SELECTOR_LIST_ENDPOINT =
  "http://localhost:8080/tool/get/list/tools-by-domain-details";
export const getToolList = async () => {
  const response = await fetch(GET_TOOL_LIST_ENDPOINT);
  if (!response.ok) throw new Error("Network response was not ok");
  return await response.json();
};

export const getToolSelectorList = async () => {
  const response = await fetch(GET_TOOL_SELECTOR_LIST_ENDPOINT);
  if (!response.ok) throw new Error("Network response was not ok");
  return await response.json();
};
