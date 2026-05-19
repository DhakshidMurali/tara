"use client";

import { DeliveryStatsByDomain, DeliveryStatsFormatData } from "@/api/types/domain";
import { DomainDeliveryFormatLegend } from "@/public/data/dashboard";
import { RefreshOutlined } from "@mui/icons-material";
import { Box, Grid, IconButton, Stack, Tooltip, Typography } from "@mui/material";
import { styles } from "./useStyles";

type Props = {
  deliveryStatsByDomain: DeliveryStatsByDomain[]
}

export default function DomainBarChart(props: Props) {
  const { deliveryStatsByDomain } = props

  const domainList = [];
  let deliveryStatsFormatData: DeliveryStatsFormatData[] = []

  deliveryStatsByDomain.map((i) => {
    if (!domainList.includes(i.domainName)) {
      domainList.push(i.domainName)
      deliveryStatsFormatData.push({
        domainName: i.domainName,
        deliveryFormat:
          [{
            format: i.deliveryFormat,
            count: i.toolsCount,
            percent: 100
          }]
        ,
        toolsCount: i.toolsCount
      })
    } else {
      deliveryStatsFormatData.map((data) => {
        if (i.domainName == data.domainName) {
          data.toolsCount = data.toolsCount + i.toolsCount
          data.deliveryFormat.push({
            format: i.deliveryFormat,
            count: i.toolsCount,
            percent: 0
          })
          data.deliveryFormat.map((deliveryFormat) => {
            deliveryFormat.percent = Math.round((deliveryFormat.count / data.toolsCount) * 100);
          })
        }
      })
    }
  })


  console.log('deliveryStatsFormatData', deliveryStatsFormatData)

  const domainDeliveryFormatLegend =
    DomainDeliveryFormatLegend;


  return (
    <Grid size={{ xs: 8 }} component="div">
      <Box sx={styles.employerByDepartmentBoxStyle} padding={4} height={"38rem"}>
        <Grid container spacing={4}>
          <Grid size={{ xs: 12 }} component="div">
            <Stack direction={"row"} sx={{ justifyContent: "space-between" }}>
              <Typography variant="h4" sx={styles.dashboardTextTypographStyle}>
                Tools By Domain
              </Typography>
              <IconButton
                sx={{
                  padding: 1,
                  bgcolor: "secondary.contrastText",
                  marginRight: 4,
                }}
              >
                <RefreshOutlined></RefreshOutlined>
              </IconButton>
            </Stack>
          </Grid>
          <Grid size={{ xs: 12 }} component="div">
            <Grid container>
              {/* Creating List of Department Name */}
              <Grid size={{ xs: 3 }} component="div">
                <Stack direction={"column"} spacing={5}>
                  {domainList.map((i, index) => (
                    <Typography variant="h6" sx={styles.dashboardTextTypographStyle}>
                      {i}
                    </Typography>
                  ))}
                </Stack>
              </Grid>
              {/* Creating List of Department Employee Bar */}
              <Grid size={{ xs: 9 }} component="div">
                <Stack
                  direction={"column"}
                  spacing={5}
                  padding={0.5}
                >
                  {/* Creating List of Department Employee Bar in a row */}

                  {deliveryStatsFormatData.map((domainDeliveryStats, index) => {
                    return ((<Stack
                      direction={"row"}
                      spacing={1}
                      padding={1.5}
                    >
                      {domainDeliveryFormatLegend.map((deliveryFormat, index) => {
                        const matchDeliveryFormat = domainDeliveryStats.deliveryFormat.filter((data) => data.format == deliveryFormat.LabelName)
                        return (<Tooltip
                          title={`${matchDeliveryFormat[0].count} [${domainDeliveryStats.toolsCount}]`}
                          slotProps={{
                            tooltip: {
                              sx: {
                                ...styles.employeeByDepartmentListBoxBarBoxToolTipStyle,
                                backgroundColor: `${deliveryFormat.Color}`,
                              },
                            },
                          }}
                          placement="top-start"
                          arrow
                        >
                          <Box
                            sx={{
                              ...styles.employeeByDepartmentListBoxBarStyle,
                              backgroundColor: `${deliveryFormat.Color}`,
                            }}
                            width={`${matchDeliveryFormat[0].percent}%`}
                          ></Box>
                        </Tooltip>)
                      })}
                    </Stack>))
                  }
                  )}
                </Stack>
              </Grid>
            </Grid>
          </Grid>
          <Stack direction={"row"} spacing={4} sx={{ width: "100%" }}>
            {domainDeliveryFormatLegend.map((i, index) => (
              <Stack
                direction={"row"}
                alignItems={"baseline"}
                sx={{ width: "25%" }}
              >
                <Box
                  sx={{
                    ...styles.employeeByDepartmentListBoxLabelBoxStyle,
                    backgroundColor: i["Color"],
                  }}
                ></Box>
                <Typography
                  variant="h6"
                  sx={styles.employeeByDepartmentListBoxLabelTypographyStyle}
                >
                  {i["LabelName"]}
                </Typography>
              </Stack>
            ))}
          </Stack>
        </Grid>
      </Box >
    </Grid >
  );
}
