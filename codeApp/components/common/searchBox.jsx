import { Box } from "@mui/material";
import { styled } from "@mui/material/styles";
const SearchBox = styled(Box)(({ theme, selected }) => ({
  margin: "1rem",
  marginTop: "0.5rem",
  padding: "1rem",
  borderRadius: 16,
  height: "4rem",
  backgroundColor: theme.palette.secondary.light,
}));

export default SearchBox;
