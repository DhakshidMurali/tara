"use client";

import { ToolListPayload } from "@/api/types/tool";
import { DescriptionOutlined } from "@mui/icons-material";
import { Avatar, AvatarGroup, Box, Grid, Paper, Stack, Tooltip, Typography } from "@mui/material";
import { styles } from "./useStyles";

type Props = {
  toolListPayload: ToolListPayload[]
}
export default function ToolList(props: Props) {
  const { toolListPayload } = props
  return (
    <Grid size={{ xs: 12 }} component="div" spacing={2} height={"14rem"}>
      <Box sx={styles.toolsContainerBoxStyle}>
        {toolListPayload.map((data, index) => (
          <Box key={index}>
            <Grid container sx={styles.toolsContainerListBoxGridStyle}>
              <Grid size={{ xs: 12 }} component="div">
                {" "}
                <Typography
                  variant="h6"
                  sx={styles.toolsContainerListBoxGridToolNameTypography}
                >
                  {data.toolName}
                </Typography>
              </Grid>
              <Grid size={{ xs: 4 }} component="div" marginLeft={1}>
                <Paper
                  elevation={2}
                  sx={{
                    padding: 0.5,
                    margin: 0.5,
                    display: "flex",
                    justifyContent: "center",
                    backgroundColor: "primary.light",
                  }}
                >
                  {" "}
                  <Stack direction={"row"} display={"flex"} justifyContent={"center"} alignItems={'center'}>
                    <DescriptionOutlined sx={{ fontSize: 16, marginRight: 0.2 }}></DescriptionOutlined>
                    <Typography
                      sx={
                        styles.toolsContainerListBoxGridCommunicationCountTypography
                      }
                    >
                      {data.communicationCount}
                    </Typography>
                  </Stack>
                </Paper>
              </Grid>
              <Grid size={{ xs: 4 }} component="div" marginRight={1}>
                <AvatarGroup max={3} spacing={12} sx={{ "& .MuiAvatar-root": { border: "2px solid black" } }} >
                  <Tooltip title={data.domainName} ><Avatar sx={{ bgcolor: "rgb(239,241,255)" }} ><Typography sx={{ color: "secondary.main" }}> {data.domainName.substring(0, 1)}</Typography></Avatar></Tooltip>
                  <Tooltip title={data.subDomain}><Avatar sx={{ bgcolor: "rgb(239,241,255)" }}><Typography sx={{ color: "secondary.main" }}> {data.subDomain.substring(0, 1)}</Typography></Avatar></Tooltip>
                </AvatarGroup>
              </Grid>
            </Grid>
          </Box>
        ))}
      </Box>
    </Grid>
  );
}
