import { Box } from "@mui/material";
import { styled } from "@mui/material/styles";
const ListBox = styled(Box)(({ theme, selected }) => ({
  alignSelf: "center",
  padding: "16px 16px",
  backgroundColor: selected
    ? theme.palette.secondary.dark
    : theme.palette.primary.main,
  borderRadius: "16px",
  width: "80%",
  display:'flex',
  flexDirection:'row',
  alignContent:'baseline'
}));

export default ListBox;
