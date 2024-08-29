import { Box } from "@mui/material";
import { styled } from "@mui/material/styles";
const SpacerBox = styled(Box)(({ theme, height }) => ({
  height:height||0,
  width:'inherit'
}));

export default SpacerBox;
