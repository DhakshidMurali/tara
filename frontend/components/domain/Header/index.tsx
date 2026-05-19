import { Button, Grid, Stack, Typography } from "@mui/material";
import { styles } from "./useStyles";

export function DomainHeader() {
    return (<Grid size={{ xs: 12 }} component="div">
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
            <Typography variant="h3" sx={styles.domainHeaderContainerSx}>
                Domain{" "}
            </Typography>
            <Button sx={styles.createButtonSx}>Create</Button>
        </Stack>
    </Grid>);
}