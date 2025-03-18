import { RefreshOutlined } from "@mui/icons-material";
import { Grid, IconButton, Stack, Typography } from "@mui/material";
import { BarChart } from "@mui/x-charts/BarChart";
import { axisClasses } from "@mui/x-charts/ChartsAxis";
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

export const dataset = [
    {
        london: 59,
        paris: 57,
        newYork: 86,
        seoul: 21,
        month: 'Jan1',
    },
    {
        london: 50,
        paris: 52,
        newYork: 78,
        seoul: 28,
        month: 'Feb1',
    },
    {
        london: 47,
        paris: 53,
        newYork: 106,
        seoul: 41,
        month: 'Mar1',
    },
    {
        london: 54,
        paris: 56,
        newYork: 92,
        seoul: 73,
        month: 'Apr1',
    },
    {
        london: 57,
        paris: 69,
        newYork: 92,
        seoul: 99,
        month: 'May1',
    },
    {
        london: 60,
        paris: 63,
        newYork: 103,
        seoul: 144,
        month: 'June1',
    },
    {
        london: 59,
        paris: 60,
        newYork: 105,
        seoul: 319,
        month: 'July1',
    },
    {
        london: 65,
        paris: 60,
        newYork: 106,
        seoul: 249,
        month: 'Aug1',
    },
    {
        london: 51,
        paris: 51,
        newYork: 95,
        seoul: 131,
        month: 'Sept1',
    },
    {
        london: 60,
        paris: 65,
        newYork: 97,
        seoul: 55,
        month: 'Oct1',
    },
    {
        london: 67,
        paris: 64,
        newYork: 76,
        seoul: 48,
        month: 'Nov1',
    },
    {
        london: 61,
        paris: 70,
        newYork: 103,
        seoul: 25,
        month: 'Dec1',
    },
    {
        london: 59,
        paris: 57,
        newYork: 86,
        seoul: 21,
        month: 'Jan',
    },
    {
        london: 50,
        paris: 52,
        newYork: 78,
        seoul: 28,
        month: 'Feb',
    },
    {
        london: 47,
        paris: 53,
        newYork: 106,
        seoul: 41,
        month: 'Mar',
    },
    {
        london: 54,
        paris: 56,
        newYork: 92,
        seoul: 73,
        month: 'Apr',
    },
    {
        london: 57,
        paris: 69,
        newYork: 92,
        seoul: 99,
        month: 'May',
    },
    {
        london: 60,
        paris: 63,
        newYork: 103,
        seoul: 144,
        month: 'June',
    },
    {
        london: 59,
        paris: 60,
        newYork: 105,
        seoul: 319,
        month: 'July',
    },
    {
        london: 65,
        paris: 60,
        newYork: 106,
        seoul: 249,
        month: 'Aug',
    },
    {
        london: 51,
        paris: 51,
        newYork: 95,
        seoul: 131,
        month: 'Sept',
    },
    {
        london: 60,
        paris: 65,
        newYork: 97,
        seoul: 55,
        month: 'Oct',
    },
    {
        london: 67,
        paris: 64,
        newYork: 76,
        seoul: 48,
        month: 'Nov',
    },
    {
        london: 61,
        paris: 70,
        newYork: 103,
        seoul: 25,
        month: 'Dec',
    },
];

export function valueFormatter(value: number | null) {
    return `${value}mm`;
}
export default function ToolBarChart() {
    return (
        <Grid container sx={{
            borderRadius: 4,
            backgroundColor: "secondary.light",
            overflowY: "auto",
            display: "flex",
        }} padding={4} height={548}>
            <Grid item xs={12} height={"20%"} >
                <Stack direction={"row"} justifyContent={"space-between"}>
                    <Typography variant="h4" sx={{
                        color: "primary.contrastText",
                        fontWeight: "bolder",
                        marginLeft: 1.5,
                    }}>
                        Find the Tools most liked
                    </Typography>
                    <IconButton
                        sx={{ padding: 1, bgcolor: "secondary.contrastText", marginRight: 4 }}
                    >
                        {" "}
                        <RefreshOutlined></RefreshOutlined>
                    </IconButton>
                </Stack>
            </Grid>
            <Grid item xs={12} height={"80%"} alignContent={"flex-end"}><BarChart
                dataset={dataset}
                xAxis={[
                    { scaleType: 'band', dataKey: 'month' }
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
    )
}