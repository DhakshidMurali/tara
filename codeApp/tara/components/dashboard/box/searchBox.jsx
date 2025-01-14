import { Box } from "@mui/material";
import { styled } from "@mui/material/styles";
const SearchBox = styled(Box)(({ theme, selected }) => ({
  margin: "1%",
  marginTop: "0.5%",
  padding: "1%",
  borderRadius: 16,
  height: "8%",
  backgroundColor: theme.palette.secondary.light,
  //   padding: 8,
}));

export default SearchBox;
