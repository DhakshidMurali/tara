import { Button, Grid, Stack, Typography } from "@mui/material";
import { styles } from "./useStyles";
export default function CommmunityHeader() {
    return (<Grid size={{ xs: 12 }}>
        <Stack
            direction={"row"}
            display={"flex"}
            justifyContent={"space-between"}
            sx={{
                "& .MuiButtonBase-root:hover": {
                    backgroundColor: "primary.light",
                    fontWeight: "bold"
                }
            }}
        >
            <Typography variant="h3" sx={styles.communityTitleSx}>
                Community{" "}
            </Typography>
            <Button sx={styles.communityCreateButSx}>Create</Button>
        </Stack>
    </Grid>);
}