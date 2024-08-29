"use client";
import { ThemeProvider } from "@emotion/react";
import { Grid } from "@mui/material";
import "../styles/theme";
import theme from "../styles/theme";
import "./globals.css";

export default function RootLayout({ children }) {
  return (
    <ThemeProvider theme={theme}>
      <html>
        <body>
          <Grid container sx={{ height: "100vh", width: "100vw" }}>
            {children}
          </Grid>
        </body>
      </html>
    </ThemeProvider>
  );
}
