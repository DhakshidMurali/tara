"use client";


import { lightPurple } from "@/utils/constants";
import {
  AccountCircleOutlined,
  DashboardRounded,
  Groups2,
  TimelineOutlined,
} from "@mui/icons-material";
import { Avatar, Button, Stack, Typography } from "@mui/material";
import Box from "@mui/material/Box";
import Paper from "@mui/material/Paper";
import { styled } from "@mui/material/styles";
import { useRouter } from "next/navigation";
import { styles } from "./useStyles";
import { SpacerBox, AvatarBox, ListBox } from "./boxComponent/avatarBox";
const Sidebar = styled(Paper)(({ theme }) => ({
  height: "100vh",
  backgroundColor: theme.palette.primary.main,
  display: "flex",
  flexDirection: "column",
}));

export default function NavBar(props) {
  const { selected } = props;
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
            sx={styles.imageLogo}
          ></Box>
          <Typography variant="h3" sx={styles.titleTextStyle}>
            Tara
          </Typography>
        </Stack>
        <SpacerBox height={16}></SpacerBox>
        <AvatarBox>
          <Avatar
            variant="rounded"
            sx={{ bgcolor: lightPurple, width: 112, height: 112 }}
          >
            AV
          </Avatar>
        </AvatarBox>
        <Typography variant="body1" sx={styles.userNameTextStyle}>
          Arjun Vijay
        </Typography>
        <Typography variant="h3" sx={styles.userEmailAddressTextStyle}>
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
            <Typography
              variant="body1"
              sx={
                selected == "dashboard"
                  ? styles.navbarSelectedMenuTextStyle
                  : styles.navbarMenuTextStyle
              }
            >
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
            <Typography variant="body1" sx={
              selected == "community"
                ? styles.navbarSelectedMenuTextStyle
                : styles.navbarMenuTextStyle
            }>
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
            <Typography variant="body1" sx={styles.navbarMenuTextStyle}>
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
            <Typography variant="body1" sx={styles.navbarMenuTextStyle}>
              Profile
            </Typography>
          </ListBox>
        </Button>
      </Stack>
    </Sidebar>
  );
}
