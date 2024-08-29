"use client";
import AvatarBox from "@/components/box/avatarBox";
import ListBox from "@/components/box/listBox";
import SpacerBox from "@/components/box/spacerBox";
import {
  AccountCircleOutlined,
  DashboardRounded,
  Groups2,
  TimelineOutlined,
} from "@mui/icons-material";
import { Avatar, Button, Stack, Typography } from "@mui/material";
import Box from "@mui/material/Box";
import { green } from "@mui/material/colors";
import Paper from "@mui/material/Paper";
import { styled } from "@mui/material/styles";
import { useRouter } from "next/navigation";
import useStyles from "./useStyles";
const Sidebar = styled(Paper)(({ theme }) => ({
  height: "100vh",
  backgroundColor: theme.palette.primary.main,
  display: "flex",
  flexDirection: "column",
}));

export default function NavBar(props) {
  const { selected } = props;
  const style = useStyles();
  const route = useRouter();
  const handleNavigation = (currentPage) => {
    currentPage == "dashboard"
      ? route.push("/dashboard")
      : currentPage == "hierarchy"
      ? route.push("/hierarchy")
      : currentPage == "community"
      ? route.push("/community")
      : currentPage == "profile"
      ? route.push("/profile")
      : null;
  };
  return (
    <Sidebar>
      <Stack spacing={1}>
        <SpacerBox height={8}></SpacerBox>
        <Stack direction="row" alignItems="baseline" justifyContent="center">
          <Box
            component="img"
            src="/images/logo.jpg"
            sx={{ height: 80, width: 80, mixBlendMode: "multiply" }}
          ></Box>
          <Typography variant="h3" className={style.titleText}>
            Tara
          </Typography>
        </Stack>
        <SpacerBox height={16}></SpacerBox>
        <AvatarBox>
          <Avatar
            variant="rounded"
            sx={{ bgcolor: green[500], width: 112, height: 112 }}
          >
            AV
          </Avatar>
        </AvatarBox>
        <Typography variant="body1" className={style.centerBoxText}>
          Arjun Vijay
        </Typography>
        <Typography
          variant="h3"
          className={style.centerBoxTextFWLight}
          sx={{ margin: 0 }}
        >
          arjun.vijay@gmail.com
        </Typography>
        <SpacerBox height={24}></SpacerBox>
        <Button
          onClick={() => {
            handleNavigation("dashboard");
          }}
        >
          <ListBox selected={selected == "dashboard"}>
            <DashboardRounded
              color={selected == "dashboard" ? "primary" : "secondary"}
            ></DashboardRounded>
            <Typography variant="body1" className={style.centerBoxText}>
              Dashboard
            </Typography>
          </ListBox>
        </Button>
        <Button
          onClick={() => {
            handleNavigation("community");
          }}
        >
          <ListBox selected={selected == "community"}>
            <Groups2
              color={selected == "community" ? "primary" : "secondary"}
            ></Groups2>
            <Typography variant="body1" className={style.centerBoxText}>
              Community
            </Typography>
          </ListBox>
        </Button>
        <Button
          onClick={() => {
            handleNavigation("hierarchy");
          }}
        >
          {" "}
          <ListBox selected={selected == "hierarchy"}>
            <TimelineOutlined
              color={selected == "hierarchy" ? "primary" : "secondary"}
            ></TimelineOutlined>
            <Typography variant="body1" className={style.centerBoxText}>
              Hierarchy
            </Typography>
          </ListBox>
        </Button>
        <Button
          onClick={() => {
            handleNavigation("profile");
          }}
        >
          <ListBox selected={selected == "profile"}>
            <AccountCircleOutlined
              color={selected == "profile" ? "primary" : "secondary"}
            ></AccountCircleOutlined>
            <Typography variant="body1" className={style.centerBoxText}>
              Profile
            </Typography>
          </ListBox>
        </Button>
      </Stack>
    </Sidebar>
  );
}
