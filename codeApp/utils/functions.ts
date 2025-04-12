import { Colors } from "./constants";

const translations = {
  "DevOps Communication": "DevOps Communication",
  "Web Devlopment Communication": "Web Devlopment Communication",
  "App Devlopment Communication": "App Devlopment Communication",
  "Internet of Things Communication": "Internet of Things Communication",
  "Cloud Computing Communication": "Cloud Computing Communication",
  "Netflix Blog Post Communication": "Netflix Blog Post Communication",
} as const;

export function addLabels<T extends { dataKey: keyof typeof translations }>(
  series: T[]
) {
  return series.map((item) => ({
    ...item,
    label: translations[item.dataKey],
  }));
}

export function getRandomColor(index) {
  return Colors[index%4];
}
