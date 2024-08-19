import { Box } from "@mui/material";
import { styled } from "@mui/material/styles";
const ListBox = styled(Box)(({ theme }) => ({
  alignSelf:"center",
  padding:'8px 24px',
  backgroundColor:theme.palette.secondary.main,
  borderRadius:'16px',
  width:'70%',
}));

export default ListBox;
