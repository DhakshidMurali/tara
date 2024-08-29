"use client";
import SearchBox from "@/components/box/searchBox";
import NavBar from "@/components/navBar/navBar";
import {
  BuildCircleOutlined,
  Domain,
  ExpandMore,
  PageviewRounded,
} from "@mui/icons-material";
import {
  Avatar,
  Box,
  Button,
  Grid,
  Stack,
  TextField,
  Typography,
} from "@mui/material";
import useStyles, {
  departmentsListBoxStyle,
  employerByDepartmentBoxStyle,
  listBoxStyle,
  listContainBoxStyle,
  searchTextFieldInputPropsStyle,
} from "./useStyles";

export default function DashBoard() {
  const style = useStyles();
  let time = new Date();
  const month = [
    "Jan",
    "Feb",
    "Mar",
    "Apr",
    "May",
    "Jun",
    "Jul",
    "Aug",
    "Sep",
    "Oct",
    "Nov",
    "Dec",
  ];
  var dayNames = ["Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"];
  return (
    <>
      <Grid item xs={2.15} sx={{ bgcolor: "green", marginRight: 2 }}>
        <NavBar selected="dashboard"></NavBar>
      </Grid>
      <Grid item xs={9.75}>
        <SearchBox>
          <Stack
            direction={"row"}
            alignItems={"center"}
            height={"100%"}
            spacing={2}
          >
            <PageviewRounded></PageviewRounded>
            <TextField
              id="standard-basic"
              variant="standard"
              placeholder="Search"
              fullWidth
              InputProps={{
                style: { color: "#f3def5", fontWeight: "bolder", fontSize: 16 },
              }}
            />
            <Typography
              className={style.centerBoxText}
              sx={{ textWrap: "nowrap" }}
            >
              {dayNames[time.getDay()] +
                ", " +
                month[time.getMonth()] +
                " " +
                time.getDate()}
            </Typography>
          </Stack>
        </SearchBox>

        <Stack
          direction={"row"}
          display={"flex"}
          justifyContent={"space-between"}
        >
          <Typography variant="h3" className={style.titleText}>
            Dashboard{" "}
          </Typography>
          <Button
            sx={{
              paddingRight: 2,
              marginRight: 4,
              backgroundColor: "green",
              display: "flex",
              alignItems: "center",
              justifyContent: "center",
              width: 128,
            }}
          >
            Create
          </Button>
        </Stack>
        <Grid container spacing={2}>
          {/* Horizontal Scroll Grid */}
          <Grid item xs={12}>
            <Box sx={listContainBoxStyle}>
              <Grid container spacing={2} sx={{ flexWrap: "nowrap" }}>
                {Array.from({ length: 10 }).map((_, index) => (
                  <Grid item key={index}>
                    <Box sx={listBoxStyle}>
                      <Typography variant="h6" color="white">
                        Item {index + 1}
                      </Typography>
                    </Box>
                  </Grid>
                ))}
              </Grid>
            </Box>
          </Grid>
          <Grid item xs={8}>
            <Box
              sx={employerByDepartmentBoxStyle}
              padding={4}
              bgcolor={"purple"}
              height={548}
            >
              <Stack spacing={4}>
                <Typography variant="h4" className={style.titleText}>
                  Employee By Department
                </Typography>
                <Typography variant="h6" className={style.titleText}>
                  Computer Science Eng
                </Typography>
                <Typography variant="h6" className={style.titleText}>
                  {" "}
                  Computer Science Eng
                </Typography>
                <Typography variant="h6" className={style.titleText}>
                  {" "}
                  Computer Science Eng
                </Typography>
                <Typography variant="h6" className={style.titleText}>
                  {" "}
                  Computer Science Eng
                </Typography>
              </Stack>
            </Box>
          </Grid>
          <Grid item xs={4}>
            <Box
              sx={{
                backgroundColor: "purple",
                height: 548,
                padding: 4,
                borderRadius: 4,
                marginRight: 4,
              }}
            >
              <Stack spacing={0}>
                <Typography variant="h4" className={style.titleText}>
                  Departments
                </Typography>
                <Box
                  sx={{
                    borderRadius: 2,
                    height: 72,
                    width: "90%",
                    display: "flex",
                    justifyContent: "center",
                    alignItems: "center",
                  }}
                >
                  <Stack direction={"row"}>
                    <Avatar>
                      <Domain></Domain>
                    </Avatar>
                    <Typography
                      variant="h6"
                      className={style.titleText}
                      paddingLeft={2}
                    >
                      Computer Science Eng
                    </Typography>
                    <Button sx={{ margin: 0 }}>
                      <ExpandMore
                        sx={{ marginLeft: 0, color: "white" }}
                      ></ExpandMore>
                    </Button>
                  </Stack>
                </Box>
                <Box
                  sx={{
                    height: 1.25,
                    width: "70%",
                    bgcolor: "red",
                    alignSelf: "center",
                  }}
                ></Box>
                <Typography variant="h6" className={style.titleText}>
                  {" "}
                  Computer Science Eng
                </Typography>
                <Typography variant="h6" className={style.titleText}>
                  {" "}
                  Computer Science Eng
                </Typography>
                <Typography variant="h6" className={style.titleText}>
                  {" "}
                  Computer Science Eng
                </Typography>
              </Stack>
            </Box>
          </Grid>
        </Grid>
      </Grid>
    </>
  );
}
