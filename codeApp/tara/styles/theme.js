"use Client"
import { createTheme } from '@mui/material/styles';

const theme = createTheme({
  palette: {
    primary: {
      main: 'rgb(39,14,95)',  // Blue Nav Bar
      light:'rgb(177,248,180)', // Green Selected Nav Bar
      contrastText:'#f3def5' // white color
    },
    secondary: {
      main: 'rgb(17,8,27)', // Background Color
      light:'rgb(33,24,43)' // Dark Purple
    },
  },
  typography: {
    fontFamily:'IBM Plex Sans'
  },
});

export default theme;


// color : 17153B , 2E236C , 433D8B , C8ACD6