async function fetchUSD(){
    resp = await fetch("/api/v1/monitoring/usopen")
    resp = await resp.json()
    data = []
    timeStamp = []
    for(let i=0; i<resp.length; i++){
        if(resp[i][0] == "CLOSE_USD"){
            data.push(resp[i][1])
            timeStamp.push(resp[i][2])
        }
    }
    return [data, timeStamp]
}

async function fetchSPY(){
    resp = await fetch("/api/v1/monitoring/spy")
    resp = await resp.json()
    data = []
    timeStamp = []
    for(let i=0; i<resp.length; i++){
        if(resp[i][0] == "CLOSE_SPY"){
            data.push(resp[i][1])
            timeStamp.push(resp[i][2])
        }
    }
    return [data, timeStamp]
}

async function fetchBasementRates(){
    resp = await fetch("/api/v1/monitoring/basement")
    resp = await resp.json()
    data = []
    dataMax = []
    timeStamp = []
    timeStampMax = []
    dataMin = []
    timeStampMin = []
    dataLikely = []
    timeStampLikely = []
    for(let i=0; i<resp.length; i++){
        if(resp[i][0] == "BASEMENT_MEAN"){
            data.push(resp[i][1])
            timeStamp.push(resp[i][2])
        }
        else if(resp[i][0] == "BASEMENT_MAX"){
            dataMax.push(resp[i][1])
            timeStampMax.push(resp[i][2])
        }
        else if(resp[i][0] == "BASEMENT_MIN"){
            dataMin.push(resp[i][1])
            timeStampMin.push(resp[i][2])
        }
        else if(resp[i][0] == "BASEMENT_LIKELY_PRICE"){
            dataLikely.push(resp[i][1])
            timeStampLikely.push(resp[i][2])
        }
    }
    return [data, timeStamp, dataMax, timeStampMax, dataMin, timeStampMin, dataLikely, timeStampLikely]
}

async function fetchApartmentRates(){
    resp = await fetch("/api/v1/monitoring/apartment")
    resp = await resp.json()
    data = []
    dataMax = []
    timeStampMax = []
    timeStamp = []
    dataMin = []
    timeStampMin = []
    dataLikely = []
    timeStampLikely = []
    for(let i=0; i<resp.length; i++){
        if(resp[i][0] == "APARTMENT_MEAN"){
            data.push(resp[i][1])
            timeStamp.push(resp[i][2])
        }
        else if(resp[i][0] == "APARTMENT_MAX"){
            dataMax.push(resp[i][1])
            timeStampMax.push(resp[i][2])
        }
        else if(resp[i][0] == "APARTMENT_MIN"){
            dataMin.push(resp[i][1])
            timeStampMin.push(resp[i][2])
        }
        else if(resp[i][0] == "APARTMENT_LIKELY_PRICE"){
            dataLikely.push(resp[i][1])
            timeStampLikely.push(resp[i][2])
        }
    }
    return [data, timeStamp, dataMax, timeStampMax, dataMin, timeStampMin, dataLikely, timeStampLikely]
}

async function fetchOilRates(){
    resp = await fetch("/api/v1/monitoring/oil")
    resp = await resp.json()
    
    brentCrude = []
    brentCrudeTS = []
    westTexas = []
    westTexasTS = []

    for(let i=0; i<resp.length; i++){
        if(resp[i][0] == "BRENTOIL"){
            brentCrude.push(resp[i][1])
            brentCrudeTS.push(resp[i][2].split(" ")[0])
        }
        else if(resp[i][0] == "WTIOIL"){
            westTexas.push(resp[i][1])
            westTexasTS.push(resp[i][2].split(" ")[0])
        }
    }
    return [brentCrude, brentCrudeTS, westTexasTS, westTexas]
}

async function prepareCharts(){
    usdRates = await fetchUSD()
    oilRates = await fetchOilRates()
    basementRates = await fetchBasementRates()
    apartmentRates = await fetchApartmentRates()
    spyRates = await fetchSPY()

    console.log("DATAPOINTS", usdRates[0])

    console.log(basementRates[4][0])
    console.log(basementRates[0][0])
    
    const ctx = document.getElementById('usd').getContext('2d');
    const myChart = new Chart(ctx, {
        type: 'line',
        data: {
            labels: usdRates[1],
            datasets: [{
                label: '# USD TO PKR',
                data: usdRates[0],
                backgroundColor: [
                    'rgba(75, 192, 192, 0.2)',
                ],
                borderColor: [
                    'rgba(75, 192, 192, 1)',
                ],
                borderWidth: 1
            }]
        },
        options: {
            plugins: {
            title: {
                display: true,
                text: '$USD To PKR rate',
                font: {
                    size: 18
                }
            },
            },
            responsive: true,
            maintainAspectRatio: false,
            scales: {
                y: {
                    beginAtZero: false
                }
            }
        }
    });

    const spyCtx = document.getElementById('spy').getContext('2d');
    const spyChart = new Chart(spyCtx, {
        type: 'line',
        data: {
            labels: spyRates[1],
            datasets: [{
                label: '# S&P 500',
                data: spyRates[0],
                backgroundColor: [
                    'rgba(75, 192, 192, 0.2)',
                ],
                borderColor: [
                    'rgba(75, 192, 192, 1)',
                ],
                borderWidth: 1
            }]
        },
        options: {
            plugins: {
            title: {
                display: true,
                text: 'S&P 500 (SPY ETF)',
                font: {
                    size: 18
                }
            },
            },
            responsive: true,
            maintainAspectRatio: false,
            scales: {
                y: {
                    beginAtZero: false
                }
            }
        }
    });

    const oilctx = document.getElementById('oil').getContext('2d');
    const oilChart = new Chart(oilctx, {
        type: 'line',
        data: {
            labels: oilRates[1],
            datasets: [{
                label: '# Brent Crude ($ per barrel)',
                data: oilRates[0],
                backgroundColor: [
                    'rgba(178, 0, 0, 0.2)',
                ],
                borderColor: [
                    'rgba(178, 0, 0, 0.2)',
                ],
                borderWidth: 2
                
            },
            {
                label: '# WTI ($ per barrel)',
                data: oilRates[3],
                backgroundColor: [
                    'rgba(139,69,19, 0.2)',
                ],
                borderColor: [
                    'rgba(139,69,19 0.2)',
                ],
                borderWidth: 2
            }
            
            ]},
        options: {
            plugins: {
            title: {
                display: true,
                text: 'Price of Oil per barrel ($USD)',
                font: {
                    size: 18
                }
            },
            },
            responsive: true,
            maintainAspectRatio: false,
            scales: {
                y: {
                    beginAtZero: false
                }
            }
        }
    });

    chart = new Chart(document.getElementById('basement').getContext('2d'), {
    type: 'bar',
    data: {
        labels: basementRates[1],
        datasets: [
            {
                type:'line',
                label: 'BASEMENT_LIKELY_PRICE',
                data: basementRates[6],
                backgroundColor: 'rgb(255, 159, 64)',
                borderColor: 'rgb(255, 159, 64)',
                borderWidth: 4
            },
            {
                type:'line',
                label: 'APARTMENT_LIKELY_PRICE',
                data: apartmentRates[6],
                backgroundColor: 'rgb(255, 255, 120)',
                borderColor: 'rgb(255, 255, 120)',
                borderWidth: 4
            },
            {
                label: 'BASEMENT_MIN',
                data: basementRates[4],
                backgroundColor: 'rgb(255, 99, 132, 0.6)',
                stack: 'Stack 0',
            },
            {
                label: 'BASEMENT_MEAN',
                data: basementRates[0],
                backgroundColor: 'rgb(54, 162, 235, 0.6)',
                stack: 'Stack 0',
            },
            {
                label: 'BASEMENT_MAX',
                data: basementRates[2],
                backgroundColor: 'rgb(75, 192, 192, 0.6)',
                stack: 'Stack 0',
            },
            {
                label: 'APARTMENT_MIN',
                data: apartmentRates[4],
                backgroundColor: 'rgb(255, 99, 132, 0.6)',
                stack: 'Stack 1',
            },
            {
                label: 'APARTMENT_MEAN',
                data: apartmentRates[0],
                backgroundColor: 'rgb(54, 162, 235, 0.6)',
                stack: 'Stack 1',
            },
            {
                label: 'APARTMENT_MAX',
                data: apartmentRates[2],
                backgroundColor: 'rgb(75, 192, 192, 0.6)',
                stack: 'Stack 1',
            },
        ]
    },
    options: {
        plugins: {
            title: {
                display: true,
                text: 'GTA rentals (Basement & Apartments)',
                font: {
                    size: 18
                }
            },
        },
        responsive: true,
        maintainAspectRatio: false,
        scales: {
            x: {
                stacked: true,
            },
            y: {
                stacked: true
            }
        }
    }
    });
}
prepareCharts()