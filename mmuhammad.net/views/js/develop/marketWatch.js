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

async function fetchSteelRates(){
    resp = await fetch("/api/v1/monitoring/steel")
    resp = await resp.json()
    china = []
    turkeyRebar = []
    turkeyScrap = []
    timeStamp = []
    for(let i=0; i<resp.length; i++){
        if(resp[i][0] == "CHINA_HOT_ROLL"){
            china.push(resp[i][1])
            timeStamp.push(resp[i][2].split(" ")[0])
        }
        else if(resp[i][0] == "TURKEY_REBAR"){
            turkeyRebar.push(resp[i][1])
        }
        else if(resp[i][0] == "TURKEY_SC"){
            turkeyScrap.push(resp[i][1])
        }
    }
    return [china, timeStamp, turkeyRebar, turkeyScrap]
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
    steelRates = await fetchSteelRates()
    oilRates = await fetchOilRates()
    basementRates = await fetchBasementRates()
    apartmentRates = await fetchApartmentRates()

    console.log("DATAPOINTS", usdRates[0])
    console.log("DP2", steelRates)

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


    const steelctx = document.getElementById('steel').getContext('2d');
    const steelChart = new Chart(steelctx, {
        type: 'line',
        data: {
            labels: steelRates[1],
            datasets: [{
                label: '# China Hot Roll (PKR price per ounce)',
                data: steelRates[0],
                backgroundColor: [
                    'rgba(28, 213, 200, 0.2)',
                ],
                borderColor: [
                    'rgba(28, 213, 200, 0.2)',
                ],
                borderWidth: 2
            },
            {
            label: '# Turkey Rebar (PKR price per ounce)',
            data: steelRates[2],
            backgroundColor: [
                'rgba(255, 10, 0, 0.2)',
            ],
            borderColor: [
                'rgba(255, 10, 0, 0.2)',
            ],
            borderWidth: 2
            },
            {
            label: '# Turkey Scrap (PKR price per ounce)',
            data: steelRates[3],
            backgroundColor: [
                'rgba(0, 10, 255, 0.2)',
            ],
            borderColor: [
                'rgba(0, 10, 255, 0.2)',
            ],
            borderWidth: 2
            }
            ]
        },
        options: {
            plugins: {
            title: {
                display: true,
                text: 'Steel rate per ounce (PKR)',
                font: {
                    size: 18
                }
            },
            },
            responsive: true,
            maintainAspectRatio: false,
            scales: {
                y: {
                    beginAtZero: true
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