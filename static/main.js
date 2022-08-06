(function() {
    $.get("/data", function(rawData) {
        const data = JSON.parse(rawData);

        const tempChart = new Chart(
            document.getElementById("tempChart"),
            {
                type: 'line',
                data: {
                    labels: data.map(x => new Date(x.Timestamp)),
                    datasets: [{
                        label: '温度',
                        data: data.map(x => x.Temp),
                        borderColor: 'rgb(255, 99, 132)',
                        yAxisID: 'y',
                        tension: 0.4,
                    },
                    {
                        label: '湿度',
                        data: data.map(x => x.Hum),
                        borderColor: 'rgb(54, 162, 235)',
                        yAxisID: 'y2',
                        tension: 0.4,
                    }],
                },
                options: {
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
                    labels: data.map(x => new Date(x.Timestamp)),
                    datasets: [{
                        label: '気圧',
                        data: data.map(x => x.Press),
                        borderColor: 'rgb(153, 102, 255)',
                        tension: 0.4,
                    }],
                },
                options: {
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