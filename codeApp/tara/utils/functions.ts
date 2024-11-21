const translations = {
  DevOps: "DevOps",
  "Web Devlopment": "Web Devlopment",
  "App Devlopment": "App Devlopment",
  "Internet of Things": "Internet of Things",
  "Cloud Computing": "Cloud Computing",
  "Netflix Blog Post": "Netflix Blog Post",
} as const;

export function addLabels<T extends { dataKey: keyof typeof translations }>(
  series: T[]
) {
  return series.map((item) => ({
    ...item,
    label: translations[item.dataKey],
  }));
}
