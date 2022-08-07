(function() {
    $.get("/data", function(rawData) {
        const data = JSON.parse(rawData);

        const sampledDatum = largestTriangleThreeBuckets(
            data,
            1440,
            'Timestamp',
            'Temp'
        );

        const tempChart = new Chart(
            document.getElementById("tempChart"),
            {
                type: 'line',
                data: {
                    labels: sampledDatum.map(p => new Date(p.Timestamp)),
                    datasets: [{
                        label: '温度',
                        data: sampledDatum.map(p => p.Temp),
                        borderColor: 'rgb(255, 99, 132)',
                        yAxisID: 'y',
                    },
                    {
                        label: '湿度',
                        data: sampledDatum.map(p => p.Hum),
                        borderColor: 'rgb(54, 162, 235)',
                        yAxisID: 'y2',
                    }
                ],
                },
                options: {
                    animation: false,

                    pointRadius: 0,
                    scales: {
                        y: {
                            type: 'linear',
                            display: true,
                            position: 'left',
                        },
                        y2: {
                            type: 'linear',
                            display: true,
                            position: 'right',
                        },
                        x: {
                            type: 'time',
                        },
                    },
                },
            }
        );

        const pressChart = new Chart(
            document.getElementById("pressChart"),
            {
                type: 'line',
                data: {
                    labels: sampledDatum.map(x => new Date(x.Timestamp)),
                    datasets: [{
                        label: '気圧',
                        data: sampledDatum.map(x => x.Press),
                        borderColor: 'rgb(153, 102, 255)',
                    }],
                },
                options: {
                    animation: false,

                    pointRadius: 0,
                    scales: {
                        x: {
                            type: 'time',
                        },
                    },
                },
            }
        );
    });
}())