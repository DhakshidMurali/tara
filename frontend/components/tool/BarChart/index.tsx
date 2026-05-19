import { ToolStatsByLikesPayload } from "@/api/types/tool";
import { RefreshOutlined } from "@mui/icons-material";
import { Box, Grid, IconButton, Stack, Typography } from "@mui/material";
import { BarChart } from "@mui/x-charts/BarChart";
import { axisClasses } from "@mui/x-charts/ChartsAxis";

type Props = {
    toolsStatsByLikesList: ToolStatsByLikesPayload
}
const chartSetting = {
    yAxis: [

    ],
    series: [{ dataKey: 'seoul', label: 'Bookmarks', valueFormatter }],
    sx: {
        [`& .${axisClasses.directionY} .${axisClasses.label}`]: {
            transform: 'translateX(-10px)',
        },
    },
};

export function valueFormatter(value: number | null) {
    return `${value}mm`;
}
export default function ToolBarChart(props: Props) {
    const { toolsStatsByLikesList } = props
    return (
        <Grid size={{ xs: 8 }} component="div">
            <Box height={"38rem"} padding={4} sx={{
                borderRadius: 4,
                backgroundColor: "secondary.light",
            }}>
                <Grid container spacing={2} >
                    <Grid size={{ xs: 12 }} component="div" >
                        <Stack direction={"row"} justifyContent={"space-between"} >
                            <Typography variant="h4" sx={{
                                color: "primary.contrastText",
                                fontWeight: "bolder",
                            }}>
                                Find the Tools most liked
                            </Typography>
                            <IconButton
                                sx={{ padding: 1, bgcolor: "secondary.contrastText" }}
                            >
                                {" "}
                                <RefreshOutlined></RefreshOutlined>
                            </IconButton>
                        </Stack>
                    </Grid>
                    <Grid size={{ xs: 12 }} component="div" height={"32rem"} >
                        <BarChart
                            dataset={toolsStatsByLikesList.likes}
                            xAxis={[
                                { scaleType: 'band', dataKey: 'toolName' }
                            ]}
                            {...chartSetting}
                            sx={{
                                "& .MuiChartsAxis-left .MuiChartsAxis-line": {
                                    strokeWidth: "1",
                                    stroke: "white"
                                },
                                "& .MuiChartsAxis-bottom .MuiChartsAxis-line": {
                                    strokeWidth: "1",
                                    stroke: "white"
                                },
                                "& .MuiChartsAxis-bottom .MuiChartsAxis-tickLabel": {
                                    fill: 'white'
                                },
                                "& .MuiChartsAxis-left .MuiChartsAxis-tickLabel": {
                                    fill: 'white'
                                },
                                "& .MuiChartsAxis-left .MuiChartsAxis-label": {
                                    fill: 'white'
                                }
                            }}
                        /></Grid>
                </Grid>
            </Box >
        </Grid >
    )
}