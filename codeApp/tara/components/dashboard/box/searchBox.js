import { Box } from "@mui/material";
import { styled } from "@mui/material/styles";
const SearchBox = styled(Box)(({ theme, selected }) => ({
  margin: 16,
  marginTop: 8,
  padding: 16,
  borderRadius: 16,
  height: 72,
  backgroundColor: theme.palette.secondary.light,
  //   padding: 8,
}));

export default SearchBox;
