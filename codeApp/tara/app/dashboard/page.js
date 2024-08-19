"use client";
import * as React from "react";
import { styled } from "@mui/material/styles";
import Box from "@mui/material/Box";
import Paper from "@mui/material/Paper";
import Grid from "@mui/material/Grid";
import { Avatar, Divider, Stack, Typography } from "@mui/material";
import { green } from "@mui/material/colors";
import AvatarBox from "@/components/box/avatarBox";
import TitleBox from "@/components/box/titleBox";
import ListBox from "@/components/box/listBox";
import DividerBox from "@/components/box/dividerBox";

const Sidebar = styled(Paper)(({ theme }) => ({
  height: "100vh",
  backgroundColor: theme.palette.primary.main,
  display: "flex",
  flexDirection: "column",
}));

const Content = styled(Box)(({ theme }) => ({
  padding: theme.spacing(2),
  backgroundColor: theme.palette.background.paper,
  height: "100vh",
}));

export default function DashBoard() {
  return (
    <Box sx={{ flexGrow: 1 }}>
      <Grid container>
        <Grid item xs={2.25}>
          <Sidebar>
            <Stack spacing={1}>
              <TitleBox>
                <Typography variant="h3" align="center" sx={{ color: "white", fontWeight: 600 }}>
                  Tara
                </Typography>
              </TitleBox>
              <AvatarBox>
                <Avatar
                  variant="rounded"
                  sx={{ bgcolor: green[500], width: 80, height: 80 }}
                >
                  AV
                </Avatar>
              </AvatarBox>
              <Typography
                variant="body1"
                textAlign="center"
                sx={{ color: "#f3def5", fontWeight: 600 }}
              >
                Arjun Vijay
              </Typography>
              <Typography
                variant="body1"
                textAlign="center"
                style={{ margin: 0 }}
                sx={{ color: "#f3def5", fontWeight: 600 }}
              >
                arjun.vijay@gmail.com
              </Typography>
              <DividerBox></DividerBox>
              <ListBox>
                <Typography
                  variant="body1"
                  textAlign="center"
                  sx={{ color: "#f3def5", fontWeight: 600 }}
                >
                  Dashboard
                </Typography>
              </ListBox>
              <ListBox>
                <Typography
                  variant="body1"
                  sx={{ color: "#f3def5", fontWeight: 600 }}
                >
                  Community
                </Typography>
              </ListBox>
              <ListBox>
                <Typography
                  variant="body1"
                  sx={{ color: "#f3def5", fontWeight: 600 }}
                >
                  Hierarchy
                </Typography>
              </ListBox>
              <ListBox>
                <Typography
                  variant="body1"
                  sx={{ color: "white", fontWeight: 600 }}
                >
                  Profile
                </Typography>
              </ListBox>
            </Stack>
          </Sidebar>
        </Grid>
        <Grid item xs={9.75}>
          <Content>
            <Typography variant="h4">Main Content</Typography>
            <Typography variant="body1">
              This is the main content area.
            </Typography>
          </Content>
        </Grid>
      </Grid>
    </Box>
  );
}
