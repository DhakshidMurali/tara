import { SxProps } from "@mui/system";

const titleTextStyle: SxProps = {
  textAlign: "center",
  color: "primary.contrastText",
  fontWeight: 800,
  alignSelf: "center",
};

const userNameTextStyle: SxProps = {
  textAlign: "center",
  color: "primary.contrastText",
  fontWeight: "bolder",
  fontSize: 18,
  margin: 0,
};

const userEmailAddressTextStyle: SxProps = {
  textAlign: "center",
  color: "primary.contrastText",
  fontWeight: "light",
  fontSize: 20,
  margin: 0,
};

const navbarMenuTextStyle: SxProps = {
  textAlign: "center",
  color: "primary.contrastText",
  fontWeight: "bolder",
  fontSize: 18,
  margin: 0,
  paddingLeft: 2,
};

const navbarSelectedMenuTextStyle: SxProps = {
  textAlign: "center",
  color: "primary.main",
  fontWeight: "bolder",
  fontSize: 18,
  paddingLeft: 2,
};

const image: SxProps = {
  height: "100",
  width: " ",
};
const imageLogo: SxProps = {
  height: 80,
  width: 80,
  mixBlendMode: "multiply",
};
export const styles = {
  userNameTextStyle,
  navbarSelectedMenuTextStyle,
  navbarMenuTextStyle,
  titleTextStyle,
  userEmailAddressTextStyle,
  image,
  imageLogo,
};
