import { Box, BoxProps } from "@mui/material";
import { styled } from "@mui/material/styles";


interface ExtendedBoxProps extends BoxProps {
  selected?: boolean;
  height?: number
}

export const AvatarBox = styled(Box)(({ theme }) => ({
  alignSelf: "center",
}));

export const DividerBox = styled(Box)(({ theme }) => ({
  height: "16px",
  width: "16px",
  color: theme.palette.secondary.main,
}));

export const ListBox = styled(Box)<ExtendedBoxProps>(({ theme, selected }) => ({
  alignSelf: "center",
  padding: "16px 16px",
  backgroundColor: selected
    ? theme.palette.primary.light
    : theme.palette.primary.main,
  borderRadius: "16px",
  width: "80%",
  display: "flex",
  flexDirection: "row",
  alignContent: "baseline",
}));

export const SpacerBox = styled(Box)<ExtendedBoxProps>(({ height }) => ({
  height: height || 0,
  width: "inherit",
}));

export const TitleBox = styled(Box)(() => ({
  alignSelf: "center",
  padding: "32px 0px",
}));
