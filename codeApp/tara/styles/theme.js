"use Client"
import { createTheme } from '@mui/material/styles';

const theme = createTheme({
  palette: {
    primary: {
      main: '#2E236C',  
    },
    secondary: {
      main: '#17153B',
      light:'#C8ACD6'
    },
    background: {
      paper: '#ffffff', 
    },
  },
  typography: {
    fontFamily:'IBM Plex Sans'
  },
});

export default theme;


// color : 17153B , 2E236C , 433D8B , C8ACD6