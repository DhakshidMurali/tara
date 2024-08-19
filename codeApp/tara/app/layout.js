"use client";
import "./globals.css";
import "../styles/theme";
import { ThemeProvider } from "@emotion/react";
import theme from "../styles/theme";

export default function RootLayout({ children }) {
  return (
    <ThemeProvider theme={theme}>
      <html>
        <body>{children}</body>
      </html>
    </ThemeProvider>
  );
}
