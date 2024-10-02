import { Box } from "@mui/material";
import { styled } from "@mui/material/styles";
const DividerBox = styled(Box)(({ theme }) => ({
    height:'16px',
    width:'16px',
    color:theme.palette.secondary.main
}));

export default DividerBox;
