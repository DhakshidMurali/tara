import { HighlightOff } from "@mui/icons-material";
import { Avatar, Box, Button, Checkbox, Divider, Grid, Stack, TextField, Tooltip, Typography } from "@mui/material";
import { styles } from './useStyles';
import { ToolSelectorListPayload } from "@/api/types/tool";
import { useState } from "react";

type Props = {
    toolSelectorList: ToolSelectorListPayload[]
}
export default function SelectorList(props: Props) {
    const { toolSelectorList } = props

    const [selectedIndex, setSelectedIndex] = useState({
        domainIndex: 0,
        subDomainIndex: 0
    });
    return (
        <Grid size={{ xs: 4 }} component="div">
            <Grid container spacing={1}>
                <Grid size={{ xs: 12 }} component="div">
                    <Stack direction={"row"} spacing={1} sx={{
                        "& .MuiFormControl-root .MuiTextField-root": { color: "white" }, "& .MuiButtonBase-root:hover": {
                            backgroundColor: "primary.light",
                            fontWeight: "bold",
                            padding: 1.7
                        }
                    }} marginLeft={1}>
                        <Button sx={{ backgroundColor: "primary.light", padding: 2 }}>
                            Change Domains
                        </Button>
                        <TextField
                            id="outlined-basic"
                            label="Index to Remove" variant="filled"
                            sx={{
                                "& .MuiInputBase-root": {
                                    borderBottom: "2px solid rgb(239,241,255)",
                                    color: "rgb(239,241,255)"
                                },
                                "& .MuiFormLabel-root": { color: "rgb(239,241,255)" },
                            }} />
                        <Button sx={{ padding: 2, backgroundColor: "secondary.light" }}>
                            <HighlightOff sx={{ color: "rgb(239,241,255)" }} ></HighlightOff>
                        </Button>
                    </Stack>
                </Grid>
                <Grid size={{ xs: 12 }} component="div">
                    <Stack direction={"row"} sx={{
                        "& hr": { borderColor: "white" }, flexWrap: "wrap",
                        "& .MuiAvatar-root": { width: "25px", height: "25px", fontSize: "0.95rem", backgroundColor: "rgb(239,241,255)" },
                        border: "2px solid rgb(239,241,255)"
                    }} padding={1}>
                        {[1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1].map(() => <>
                            <Tooltip title="0 : Appication Tools " sx={{ margin: 0.5 }}>
                                <Avatar>
                                    <Typography sx={{ color: "secondary.main" }}>A</Typography>
                                </Avatar>
                            </Tooltip>
                            <Divider sx={{ color: "rgb(239,241,255)" }} orientation="vertical" variant="middle" flexItem />
                        </>)
                        }
                    </Stack>
                </Grid>
                <Grid size={{ xs: 12 }} component="div">
                    <Grid container sx={{ paddingRight: 1 }} paddingLeft={1} height={"28rem"}>
                        <Grid size={{ xs: 5 }} component="div" sx={styles.toolsContainerBoxStyle}>
                            {toolSelectorList.map((domain, domainIndex) => {
                                return (domain.subDomain.map((subDomain, subDomainIndex) => {
                                    return (
                                        <Box
                                            sx={
                                                {
                                                    backgroundColor: "secondary.dark", borderTopLeftRadius: "16px",
                                                }
                                            }
                                            marginTop={0.5}
                                            height={"4rem"}
                                        >
                                            <Button onClick={() => {
                                                setSelectedIndex({
                                                    domainIndex: domainIndex,
                                                    subDomainIndex: subDomainIndex
                                                })
                                            }}>
                                                <Typography
                                                    variant="subtitle1"
                                                    sx={
                                                        {
                                                            color: "primary.contrastText",
                                                            display: "flex"
                                                        }}
                                                >
                                                    {subDomain.subDomainName}
                                                </Typography>
                                            </Button>
                                        </Box>
                                    );
                                }))
                            })}
                        </Grid>
                        <Grid size={{ xs: 0.3 }} component="div">
                            {toolSelectorList[selectedIndex.domainIndex].subDomain[selectedIndex.subDomainIndex].tool.map((data, index) => {
                                if (index == 1) {
                                    return (
                                        <Box
                                            sx={{
                                                height: "4rem",
                                                backgroundColor: "secondary.dark",
                                                marginTop: "4px",
                                            }}
                                        ></Box>
                                    );
                                } else {
                                    return (
                                        <Box
                                            sx={{
                                                marginTop: "4px",
                                            }}
                                        ></Box>
                                    );
                                }
                            })}
                        </Grid>
                        <Grid
                            size={{ xs: 6.7 }} component="div"
                            sx={styles.toolsContainerBoxStyle}>
                            {toolSelectorList[selectedIndex.domainIndex].subDomain[selectedIndex.subDomainIndex].tool.map(
                                (tool, index) => {
                                    return (
                                        <Stack
                                            direction={"row"}
                                            marginTop={1}
                                            sx={{
                                                backgroundColor: "secondary.dark",
                                                marginTop: "4px",
                                                padding: 0.5
                                            }}
                                        >
                                            <Checkbox
                                                sx={{
                                                    "& .MuiSvgIcon-root": { fontSize: 20 },
                                                    color: "rgb(46, 36,57)"[100],
                                                    "&.Mui-checked": {
                                                        color: "rgb(33,24,43)",
                                                    },
                                                }}
                                            />
                                            <Typography
                                                variant="subtitle1"
                                                sx={
                                                    {
                                                        color: "primary.contrastText",
                                                        margin: "auto"
                                                    }
                                                }
                                            >
                                                {tool}
                                            </Typography>
                                        </Stack>
                                    );
                                }
                            )}
                        </Grid>
                    </Grid>
                </Grid>
            </Grid >
        </Grid >
    );
}