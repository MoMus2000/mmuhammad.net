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

async function fetchSPYRegime(){
    resp = await fetch("/api/v1/monitoring/spy/regime")
    resp = await resp.json()
    low = []
    med = []
    high = []
    timeStamp = []
    for(let i=0; i<resp.length; i++){
        if(resp[i][0] == "LOW_VOL_PROB_SPY"){
            low.push(resp[i][1])
            timeStamp.push(resp[i][2])
        }
        else if(resp[i][0] == "MED_VOL_PROB_SPY"){
            med.push(resp[i][1])
            timeStamp.push(resp[i][2])
        }
        else if(resp[i][0] == "HIGH_VOL_PROB_SPY"){
            high.push(resp[i][1])
            timeStamp.push(resp[i][2])
        }
    }
    return [low, high, med, timeStamp]
}

async function fetchCADRegime(){
    resp = await fetch("/api/v1/monitoring/cad_housing/regime")
    resp = await resp.json()
    low = []
    med = []
    high = []
    timeStamp = []
    for(let i=0; i<resp.length; i++){
        if(resp[i][0] == "LOW_VOL_PROB_XRE.TO"){
            low.push(resp[i][1])
            timeStamp.push(resp[i][2])
        }
        else if(resp[i][0] == "MED_VOL_PROB_XRE.TO"){
            med.push(resp[i][1])
            timeStamp.push(resp[i][2])
        }
        else if(resp[i][0] == "HIGH_VOL_PROB_XRE.TO"){
            high.push(resp[i][1])
            timeStamp.push(resp[i][2])
        }
    }
    return [low, high, med, timeStamp]
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


async function fetchWindsorApartmentRates(){
    resp = await fetch("/api/v1/monitoring/windsor/apartment")
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
        if(resp[i][0] == "TWO_BD_WINDSOR_APARTMENT_MEAN"){
            data.push(resp[i][1])
            timeStamp.push(resp[i][2])
        }
        else if(resp[i][0] == "TWO_BD_WINDSOR_APARTMENT_MAX"){
            dataMax.push(resp[i][1])
            timeStampMax.push(resp[i][2])
        }
        else if(resp[i][0] == "TWO_BD_WINDSOR_APARTMENT_MIN"){
            dataMin.push(resp[i][1])
            timeStampMin.push(resp[i][2])
        }
        else if(resp[i][0] == "TWO_BD_WINDSOR_APARTMENT_LIKELY_PRICE"){
            dataLikely.push(resp[i][1])
            timeStampLikely.push(resp[i][2])
        }
    }
    return [data, timeStamp, dataMax, timeStampMax, dataMin, timeStampMin, dataLikely, timeStampLikely]
}

async function fetchWindsorBasementRates(){
    resp = await fetch("/api/v1/monitoring/windsor/basement")
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
        if(resp[i][0] == "TWO_BD_WINDSOR_BASEMENT_MEAN"){
            data.push(resp[i][1])
            timeStamp.push(resp[i][2])
        }
        else if(resp[i][0] == "TWO_BD_WINDSOR_BASEMENT_MAX"){
            dataMax.push(resp[i][1])
            timeStampMax.push(resp[i][2])
        }
        else if(resp[i][0] == "TWO_BD_WINDSOR_BASEMENT_MIN"){
            dataMin.push(resp[i][1])
            timeStampMin.push(resp[i][2])
        }
        else if(resp[i][0] == "TWO_BD_WINDSOR_BASEMENT_LIKELY_PRICE"){
            dataLikely.push(resp[i][1])
            timeStampLikely.push(resp[i][2])
        }
    }
    return [data, timeStamp, dataMax, timeStampMax, dataMin, timeStampMin, dataLikely, timeStampLikely]
}

async function fetchHamiltonApartmentRates(){
    resp = await fetch("/api/v1/monitoring/hamilton/apartment")
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
        if(resp[i][0] == "TWO_BD_HAMILTON_APARTMENT_MEAN"){
            data.push(resp[i][1])
            timeStamp.push(resp[i][2])
        }
        else if(resp[i][0] == "TWO_BD_HAMILTON_APARTMENT_MAX"){
            dataMax.push(resp[i][1])
            timeStampMax.push(resp[i][2])
        }
        else if(resp[i][0] == "TWO_BD_HAMILTON_APARTMENT_MIN"){
            dataMin.push(resp[i][1])
            timeStampMin.push(resp[i][2])
        }
        else if(resp[i][0] == "TWO_BD_HAMILTON_APARTMENT_LIKELY_PRICE"){
            dataLikely.push(resp[i][1])
            timeStampLikely.push(resp[i][2])
        }
    }
    return [data, timeStamp, dataMax, timeStampMax, dataMin, timeStampMin, dataLikely, timeStampLikely]
}

async function fetchHamiltonBasementRates(){
    resp = await fetch("/api/v1/monitoring/hamilton/basement")
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
        if(resp[i][0] == "TWO_BD_HAMILTON_BASEMENT_MEAN"){
            data.push(resp[i][1])
            timeStamp.push(resp[i][2])
        }
        else if(resp[i][0] == "TWO_BD_HAMILTON_BASEMENT_MAX"){
            dataMax.push(resp[i][1])
            timeStampMax.push(resp[i][2])
        }
        else if(resp[i][0] == "TWO_BD_HAMILTON_BASEMENT_MIN"){
            dataMin.push(resp[i][1])
            timeStampMin.push(resp[i][2])
        }
        else if(resp[i][0] == "TWO_BD_HAMILTON_BASEMENT_LIKELY_PRICE"){
            dataLikely.push(resp[i][1])
            timeStampLikely.push(resp[i][2])
        }
    }
    return [data, timeStamp, dataMax, timeStampMax, dataMin, timeStampMin, dataLikely, timeStampLikely]
}

async function fetchDurhamApartmentRates(){
    resp = await fetch("/api/v1/monitoring/durham/apartment")
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
        if(resp[i][0] == "TWO_BD_DURHAM_APARTMENT_MEAN"){
            data.push(resp[i][1])
            timeStamp.push(resp[i][2])
        }
        else if(resp[i][0] == "TWO_BD_DURHAM_APARTMENT_MAX"){
            dataMax.push(resp[i][1])
            timeStampMax.push(resp[i][2])
        }
        else if(resp[i][0] == "TWO_BD_DURHAM_APARTMENT_MIN"){
            dataMin.push(resp[i][1])
            timeStampMin.push(resp[i][2])
        }
        else if(resp[i][0] == "TWO_BD_DURHAM_APARTMENT_LIKELY_PRICE"){
            dataLikely.push(resp[i][1])
            timeStampLikely.push(resp[i][2])
        }
    }
    return [data, timeStamp, dataMax, timeStampMax, dataMin, timeStampMin, dataLikely, timeStampLikely]
}

async function fetchDurhamBasementRates(){
    resp = await fetch("/api/v1/monitoring/durham/basement")
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
        if(resp[i][0] == "TWO_BD_DURHAM_APARTMENT_MEAN"){
            data.push(resp[i][1])
            timeStamp.push(resp[i][2])
        }
        else if(resp[i][0] == "TWO_BD_DURHAM_APARTMENT_MAX"){
            dataMax.push(resp[i][1])
            timeStampMax.push(resp[i][2])
        }
        else if(resp[i][0] == "TWO_BD_DURHAM_APARTMENT_MIN"){
            dataMin.push(resp[i][1])
            timeStampMin.push(resp[i][2])
        }
        else if(resp[i][0] == "TWO_BD_DURHAM_APARTMENT_LIKELY_PRICE"){
            dataLikely.push(resp[i][1])
            timeStampLikely.push(resp[i][2])
        }
    }
    return [data, timeStamp, dataMax, timeStampMax, dataMin, timeStampMin, dataLikely, timeStampLikely]
}

async function fetchStCatharinesApartmentRates(){
    resp = await fetch("/api/v1/monitoring/stcatharines/apartment")
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
        if(resp[i][0] == "TWO_BD_ST_CATHARINES_APARTMENT_MEAN"){
            data.push(resp[i][1])
            timeStamp.push(resp[i][2])
        }
        else if(resp[i][0] == "TWO_BD_ST_CATHARINES_APARTMENT_MAX"){
            dataMax.push(resp[i][1])
            timeStampMax.push(resp[i][2])
        }
        else if(resp[i][0] == "TWO_BD_ST_CATHARINES_APARTMENT_MIN"){
            dataMin.push(resp[i][1])
            timeStampMin.push(resp[i][2])
        }
        else if(resp[i][0] == "TWO_BD_ST_CATHARINES_APARTMENT_LIKELY_PRICE"){
            dataLikely.push(resp[i][1])
            timeStampLikely.push(resp[i][2])
        }
    }
    return [data, timeStamp, dataMax, timeStampMax, dataMin, timeStampMin, dataLikely, timeStampLikely]
}

async function fetchStCatharinesBasementRates(){
    resp = await fetch("/api/v1/monitoring/stcatharines/basement")
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
        if(resp[i][0] == "TWO_BD_ST_CATHARINES_BASEMENT_MEAN"){
            data.push(resp[i][1])
            timeStamp.push(resp[i][2])
        }
        else if(resp[i][0] == "TWO_BD_ST_CATHARINES_BASEMENT_MAX"){
            dataMax.push(resp[i][1])
            timeStampMax.push(resp[i][2])
        }
        else if(resp[i][0] == "TWO_BD_ST_CATHARINES_BASEMENT_MIN"){
            dataMin.push(resp[i][1])
            timeStampMin.push(resp[i][2])
        }
        else if(resp[i][0] == "TWO_BD_ST_CATHARINES_BASEMENT_LIKELY_PRICE"){
            dataLikely.push(resp[i][1])
            timeStampLikely.push(resp[i][2])
        }
    }
    return [data, timeStamp, dataMax, timeStampMax, dataMin, timeStampMin, dataLikely, timeStampLikely]
}


function calculatePercentageChange(array){
    let result = [0]
    for(let i=0;i<array.length;i++){
        if(i>0){
            let increase = array[i] - array[i-1]
            let percentage = (increase/array[i])*100
            result.push(percentage)
        }
    }
    return result
}

function backgroundColor(array, normalColor, hightlightIndex, colorChange){
    let colorArray = []
    for(let i=0;i<array.length; i++){
        if(hightlightIndex.includes(i)){
            colorArray.push(colorChange)
        }
        else{
            colorArray.push(normalColor)
        }
    }
    if(colorArray[0] == colorChange){
        colorArray[0] = normalColor
    }
    return colorArray
}

function lowestContigousSum(array){
    let minEndingHere = Infinity
    let minSoFar = Infinity
    let lastIndex = 0
    let startIndex = 0
    let endIndex = 0
    let result = []

    for(let i=0;i<array.length;i++){
        if(minEndingHere > 0){
            minEndingHere = array[i]
            lastIndex = i
        }
        else{
            minEndingHere += array[i]
        }
        if(minSoFar > minEndingHere){
            minSoFar = minEndingHere
            startIndex = lastIndex
            endIndex = i
        }
    }
    for(let j=startIndex; j<=endIndex; j++){
        result.push(j)
    }
    return result
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
    overlay = document.getElementById("overlay")
    overlay.style.display = "block"
    usdRates = await fetchUSD()
    oilRates = await fetchOilRates()
    basementRates = await fetchBasementRates()
    apartmentRates = await fetchApartmentRates()
    spyRates = await fetchSPY()
    spyRegime = await fetchSPYRegime()
    cadRegime = await fetchCADRegime()  
    windsorApartmentRates = await fetchWindsorApartmentRates()
    windsorBasementRates = await fetchWindsorBasementRates()
    hamiltonApartmentRates = await fetchHamiltonApartmentRates()
    hamiltonBasementRates = await fetchHamiltonBasementRates()
    stcatharinesApartmentRates = await fetchStCatharinesApartmentRates()
    stcatharinesBasementRates = await fetchStCatharinesBasementRates()
    durhamApartmentRates = await fetchDurhamApartmentRates()
    durhamBasementRates = await fetchDurhamBasementRates()

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


    const spyRegimeCtx = document.getElementById('spy-regime').getContext('2d');
    const spyRegimeChart = new Chart(spyRegimeCtx, {
        type: 'line',
        data: {
            labels: Array.from(new Set(spyRegime[3])),
            datasets: [
                {
                label: '# Downturn Probability SPY',
                data: spyRegime[2],
                backgroundColor: [
                    'rgba(192, 0, 0, 0.2)',
                ],
                borderColor: [
                    'rgba(192, 0, 0, 1)',
                ],
                borderWidth: 1
            },
            {
                label: '# Med Risk SPY',
                data: spyRegime[1],
                backgroundColor: [
                    'rgba(192, 192, 0, 0.2)',
                ],
                borderColor: [
                    'rgba(192, 192, 0, 1)',
                ],
                borderWidth: 1
            },
            {
                label: '# Upturn Probability SPY',
                data: spyRegime[0],
                backgroundColor: [
                    'rgba(0, 192, 0, 0.2)',
                ],
                borderColor: [
                    'rgba(0, 192, 0, 1)',
                ],
                borderWidth: 1
            },
            {
                label: '# Downturn Probability CAN REIT',
                data: cadRegime[2],
                backgroundColor: [
                    'rgba(192, 0, 0, 0.2)',
                ],
                borderColor: [
                    'rgba(192, 0, 0, 1)',
                ],
                borderWidth: 1
            },
            {
                label: '# Med Risk CAN REIT',
                data: cadRegime[1],
                backgroundColor: [
                    'rgba(192, 192, 0, 0.2)',
                ],
                borderColor: [
                    'rgba(192, 192, 0, 1)',
                ],
                borderWidth: 1
            },
            {
                label: '# Upturn Probability CAN REIT',
                data: cadRegime[0],
                backgroundColor: [
                    'rgba(0, 192, 0, 0.2)',
                ],
                borderColor: [
                    'rgba(0, 192, 0, 1)',
                ],
                borderWidth: 1
            }
        ]
        },
        options: {
            plugins: {
            title: {
                display: true,
                text: 'Markov Regime Switch Model',
                font: {
                    size: 18
                },
                // link: "https://mmuhammad.net"
            },
            },
            responsive: true,
            maintainAspectRatio: false,
            scales: {
                y: {
                    beginAtZero: false
                }
            },
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
        },
    }
    });

    let test = lowestContigousSum(calculatePercentageChange(apartmentRates[0]))
    let colors = backgroundColor(apartmentRates[0], 'rgb(54, 162, 235, 0.6)', test)

    chart = new Chart(document.getElementById('basement-percent').getContext('2d'), {
        type: 'bar',
        data: {
            labels: basementRates[1],
            datasets: [
                {
                    type:'line',
                    label: 'BASEMENT_LIKELY_PRICE',
                    data: calculatePercentageChange(basementRates[6]),
                    backgroundColor: 'rgb(255, 159, 64)',
                    borderColor: 'rgb(255, 159, 64)',
                    borderWidth: 4
                },
                {
                    type:'line',
                    label: 'APARTMENT_LIKELY_PRICE',
                    data: calculatePercentageChange(apartmentRates[6]),
                    backgroundColor: 'rgb(255, 255, 120)',
                    borderColor: 'rgb(255, 255, 120)',
                    borderWidth: 4
                },
                {
                    label: 'BASEMENT_MEAN',
                    data: calculatePercentageChange(basementRates[0]),
                    backgroundColor: backgroundColor(basementRates[0], 'rgb(54, 162, 235, 0.6)', lowestContigousSum(calculatePercentageChange(basementRates[0])), "rgb(252, 0, 0, 0.5)"),
                    stack: 'Stack 0',
                },
                {
                    label: 'APARTMENT_MEAN',
                    data: calculatePercentageChange(apartmentRates[0]),
                    backgroundColor: backgroundColor(apartmentRates[0], 'rgb(54, 162, 235, 0.6)', lowestContigousSum(calculatePercentageChange(apartmentRates[0])), "rgb(252, 0, 0, 0.9)"),
                    stack: 'Stack 1',
                },
            ]
        },
        options: {
            plugins: {
                title: {
                    display: true,
                    text: 'GTA rentals (Basement & Apartments) % CHG',
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

        durhamChart = new Chart(document.getElementById('durham').getContext('2d'), {
            type: 'bar',
            data: {
                labels: durhamApartmentRates[1],
                datasets: [
                    {
                        type:'line',
                        label: 'BASEMENT_LIKELY_PRICE',
                        data: durhamBasementRates[6],
                        backgroundColor: 'rgb(255, 159, 64)',
                        borderColor: 'rgb(255, 159, 64)',
                        borderWidth: 4
                    },
                    {
                        type:'line',
                        label: 'APARTMENT_LIKELY_PRICE',
                        data: durhamApartmentRates[6],
                        backgroundColor: 'rgb(255, 255, 120)',
                        borderColor: 'rgb(255, 255, 120)',
                        borderWidth: 4
                    },
                    {
                        label: 'BASEMENT_MIN',
                        data: durhamBasementRates[4],
                        backgroundColor: 'rgb(255, 99, 132, 0.6)',
                        stack: 'Stack 0',
                    },
                    {
                        label: 'BASEMENT_MEAN',
                        data: durhamBasementRates[0],
                        backgroundColor: 'rgb(54, 162, 235, 0.6)',
                        stack: 'Stack 0',
                    },
                    {
                        label: 'BASEMENT_MAX',
                        data: durhamBasementRates[2],
                        backgroundColor: 'rgb(75, 192, 192, 0.6)',
                        stack: 'Stack 0',
                    },
                    {
                        label: 'APARTMENT_MIN',
                        data: durhamApartmentRates[4],
                        backgroundColor: 'rgb(255, 99, 132, 0.6)',
                        stack: 'Stack 1',
                    },
                    {
                        label: 'APARTMENT_MEAN',
                        data: durhamApartmentRates[0],
                        backgroundColor: 'rgb(54, 162, 235, 0.6)',
                        stack: 'Stack 1',
                    },
                    {
                        label: 'APARTMENT_MAX',
                        data: durhamApartmentRates[2],
                        backgroundColor: 'rgb(75, 192, 192, 0.6)',
                        stack: 'Stack 1',
                    },
                ]
            },
            options: {
                plugins: {
                    title: {
                        display: true,
                        text: 'Durham Region rentals (Basement & Apartments)',
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
                },
            }
            });
        
            let test2 = lowestContigousSum(calculatePercentageChange(durhamApartmentRates[0]))
            let colors2 = backgroundColor(durhamApartmentRates[0], 'rgb(54, 162, 235, 0.6)', test2)        
            chart = new Chart(document.getElementById('durham-percent').getContext('2d'), {
                type: 'bar',
                data: {
                    labels: durhamApartmentRates[1],
                    datasets: [
                        {
                            type:'line',
                            label: 'BASEMENT_LIKELY_PRICE',
                            data: calculatePercentageChange(durhamBasementRates[6]),
                            backgroundColor: 'rgb(255, 159, 64)',
                            borderColor: 'rgb(255, 159, 64)',
                            borderWidth: 4
                        },
                        {
                            type:'line',
                            label: 'APARTMENT_LIKELY_PRICE',
                            data: calculatePercentageChange(durhamApartmentRates[6]),
                            backgroundColor: 'rgb(255, 255, 120)',
                            borderColor: 'rgb(255, 255, 120)',
                            borderWidth: 4
                        },
                        {
                            label: 'BASEMENT_MEAN',
                            data: calculatePercentageChange(durhamBasementRates[0]),
                            backgroundColor: backgroundColor(durhamApartmentRates[0], 'rgb(54, 162, 235, 0.6)', lowestContigousSum(calculatePercentageChange(durhamBasementRates[0])), "rgb(252, 0, 0, 0.5)"),
                            stack: 'Stack 0',
                        },
                        {
                            label: 'APARTMENT_MEAN',
                            data: calculatePercentageChange(durhamApartmentRates[0]),
                            backgroundColor: backgroundColor(durhamApartmentRates[0], 'rgb(54, 162, 235, 0.6)', lowestContigousSum(calculatePercentageChange(durhamApartmentRates[0])), "rgb(252, 0, 0, 0.9)"),
                            stack: 'Stack 1',
                        },
                    ]
                },
                options: {
                    plugins: {
                        title: {
                            display: true,
                            text: 'Durham Region rentals (Basement & Apartments) % CHG',
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
        
    windsorChart = new Chart(document.getElementById('windsor').getContext('2d'), {
        type: 'bar',
        data: {
            labels: windsorApartmentRates[1],
            datasets: [
                {
                    type:'line',
                    label: 'BASEMENT_LIKELY_PRICE',
                    data: windsorBasementRates[6],
                    backgroundColor: 'rgb(255, 159, 64)',
                    borderColor: 'rgb(255, 159, 64)',
                    borderWidth: 4
                },
                {
                    type:'line',
                    label: 'APARTMENT_LIKELY_PRICE',
                    data: windsorApartmentRates[6],
                    backgroundColor: 'rgb(255, 255, 120)',
                    borderColor: 'rgb(255, 255, 120)',
                    borderWidth: 4
                },
                {
                    label: 'BASEMENT_MIN',
                    data: windsorBasementRates[4],
                    backgroundColor: 'rgb(255, 99, 132, 0.6)',
                    stack: 'Stack 0',
                },
                {
                    label: 'BASEMENT_MEAN',
                    data: windsorBasementRates[0],
                    backgroundColor: 'rgb(54, 162, 235, 0.6)',
                    stack: 'Stack 0',
                },
                {
                    label: 'BASEMENT_MAX',
                    data: windsorBasementRates[2],
                    backgroundColor: 'rgb(75, 192, 192, 0.6)',
                    stack: 'Stack 0',
                },
                {
                    label: 'APARTMENT_MIN',
                    data: windsorApartmentRates[4],
                    backgroundColor: 'rgb(255, 99, 132, 0.6)',
                    stack: 'Stack 1',
                },
                {
                    label: 'APARTMENT_MEAN',
                    data: windsorApartmentRates[0],
                    backgroundColor: 'rgb(54, 162, 235, 0.6)',
                    stack: 'Stack 1',
                },
                {
                    label: 'APARTMENT_MAX',
                    data: windsorApartmentRates[2],
                    backgroundColor: 'rgb(75, 192, 192, 0.6)',
                    stack: 'Stack 1',
                },
            ]
        },
        options: {
            plugins: {
                title: {
                    display: true,
                    text: 'Windsor Region rentals (Basement & Apartments)',
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
            },
        }
        });

        let test3 = lowestContigousSum(calculatePercentageChange(windsorApartmentRates[0]))
        let colors3 = backgroundColor(windsorApartmentRates[0], 'rgb(54, 162, 235, 0.6)', test3)        
        chart = new Chart(document.getElementById('windsor-percent').getContext('2d'), {
            type: 'bar',
            data: {
                labels: windsorApartmentRates[1],
                datasets: [
                    {
                        type:'line',
                        label: 'BASEMENT_LIKELY_PRICE',
                        data: calculatePercentageChange(windsorBasementRates[6]),
                        backgroundColor: 'rgb(255, 159, 64)',
                        borderColor: 'rgb(255, 159, 64)',
                        borderWidth: 4
                    },
                    {
                        type:'line',
                        label: 'APARTMENT_LIKELY_PRICE',
                        data: calculatePercentageChange(windsorApartmentRates[6]),
                        backgroundColor: 'rgb(255, 255, 120)',
                        borderColor: 'rgb(255, 255, 120)',
                        borderWidth: 4
                    },
                    {
                        label: 'BASEMENT_MEAN',
                        data: calculatePercentageChange(windsorBasementRates[0]),
                        backgroundColor: backgroundColor(windsorBasementRates[0], 'rgb(54, 162, 235, 0.6)', lowestContigousSum(calculatePercentageChange(windsorBasementRates[0])), "rgb(252, 0, 0, 0.5)"),
                        stack: 'Stack 0',
                    },
                    {
                        label: 'APARTMENT_MEAN',
                        data: calculatePercentageChange(windsorApartmentRates[0]),
                        backgroundColor: backgroundColor(windsorApartmentRates[0], 'rgb(54, 162, 235, 0.6)', lowestContigousSum(calculatePercentageChange(windsorApartmentRates[0])), "rgb(252, 0, 0, 0.9)"),
                        stack: 'Stack 1',
                    },
                ]
            },
            options: {
                plugins: {
                    title: {
                        display: true,
                        text: 'Windsor Region rentals (Basement & Apartments) % CHG',
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
    
        hamiltonChart = new Chart(document.getElementById('hamilton').getContext('2d'), {
            type: 'bar',
            data: {
                labels: hamiltonApartmentRates[1],
                datasets: [
                    {
                        type:'line',
                        label: 'BASEMENT_LIKELY_PRICE',
                        data: hamiltonBasementRates[6],
                        backgroundColor: 'rgb(255, 159, 64)',
                        borderColor: 'rgb(255, 159, 64)',
                        borderWidth: 4
                    },
                    {
                        type:'line',
                        label: 'APARTMENT_LIKELY_PRICE',
                        data: hamiltonApartmentRates[6],
                        backgroundColor: 'rgb(255, 255, 120)',
                        borderColor: 'rgb(255, 255, 120)',
                        borderWidth: 4
                    },
                    {
                        label: 'BASEMENT_MIN',
                        data: hamiltonBasementRates[4],
                        backgroundColor: 'rgb(255, 99, 132, 0.6)',
                        stack: 'Stack 0',
                    },
                    {
                        label: 'BASEMENT_MEAN',
                        data: hamiltonBasementRates[0],
                        backgroundColor: 'rgb(54, 162, 235, 0.6)',
                        stack: 'Stack 0',
                    },
                    {
                        label: 'BASEMENT_MAX',
                        data: hamiltonBasementRates[2],
                        backgroundColor: 'rgb(75, 192, 192, 0.6)',
                        stack: 'Stack 0',
                    },
                    {
                        label: 'APARTMENT_MIN',
                        data: hamiltonApartmentRates[4],
                        backgroundColor: 'rgb(255, 99, 132, 0.6)',
                        stack: 'Stack 1',
                    },
                    {
                        label: 'APARTMENT_MEAN',
                        data: hamiltonApartmentRates[0],
                        backgroundColor: 'rgb(54, 162, 235, 0.6)',
                        stack: 'Stack 1',
                    },
                    {
                        label: 'APARTMENT_MAX',
                        data: hamiltonApartmentRates[2],
                        backgroundColor: 'rgb(75, 192, 192, 0.6)',
                        stack: 'Stack 1',
                    },
                ]
            },
            options: {
                plugins: {
                    title: {
                        display: true,
                        text: 'Hamilton Region rentals (Basement & Apartments)',
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
                },
            }
            });
    
            let test4 = lowestContigousSum(calculatePercentageChange(hamiltonApartmentRates[0]))
            let colors4 = backgroundColor(hamiltonApartmentRates[0], 'rgb(54, 162, 235, 0.6)', test4)        
            chart = new Chart(document.getElementById('hamilton-percent').getContext('2d'), {
                type: 'bar',
                data: {
                    labels: hamiltonApartmentRates[1],
                    datasets: [
                        {
                            type:'line',
                            label: 'BASEMENT_LIKELY_PRICE',
                            data: calculatePercentageChange(hamiltonBasementRates[6]),
                            backgroundColor: 'rgb(255, 159, 64)',
                            borderColor: 'rgb(255, 159, 64)',
                            borderWidth: 4
                        },
                        {
                            type:'line',
                            label: 'APARTMENT_LIKELY_PRICE',
                            data: calculatePercentageChange(hamiltonApartmentRates[6]),
                            backgroundColor: 'rgb(255, 255, 120)',
                            borderColor: 'rgb(255, 255, 120)',
                            borderWidth: 4
                        },
                        {
                            label: 'BASEMENT_MEAN',
                            data: calculatePercentageChange(hamiltonBasementRates[0]),
                            backgroundColor: backgroundColor(hamiltonBasementRates[0], 'rgb(54, 162, 235, 0.6)', lowestContigousSum(calculatePercentageChange(hamiltonBasementRates[0])), "rgb(252, 0, 0, 0.5)"),
                            stack: 'Stack 0',
                        },
                        {
                            label: 'APARTMENT_MEAN',
                            data: calculatePercentageChange(hamiltonApartmentRates[0]),
                            backgroundColor: backgroundColor(hamiltonApartmentRates[0], 'rgb(54, 162, 235, 0.6)', lowestContigousSum(calculatePercentageChange(hamiltonApartmentRates[0])), "rgb(252, 0, 0, 0.9)"),
                            stack: 'Stack 1',
                        },
                    ]
                },
                options: {
                    plugins: {
                        title: {
                            display: true,
                            text: 'Hamilton Region rentals (Basement & Apartments) % CHG',
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
    
            stcatharinesChart = new Chart(document.getElementById('stcatharines').getContext('2d'), {
                type: 'bar',
                data: {
                    labels: stcatharinesApartmentRates[1],
                    datasets: [
                        {
                            type:'line',
                            label: 'BASEMENT_LIKELY_PRICE',
                            data: stcatharinesBasementRates[6],
                            backgroundColor: 'rgb(255, 159, 64)',
                            borderColor: 'rgb(255, 159, 64)',
                            borderWidth: 4
                        },
                        {
                            type:'line',
                            label: 'APARTMENT_LIKELY_PRICE',
                            data: stcatharinesApartmentRates[6],
                            backgroundColor: 'rgb(255, 255, 120)',
                            borderColor: 'rgb(255, 255, 120)',
                            borderWidth: 4
                        },
                        {
                            label: 'BASEMENT_MIN',
                            data: stcatharinesBasementRates[4],
                            backgroundColor: 'rgb(255, 99, 132, 0.6)',
                            stack: 'Stack 0',
                        },
                        {
                            label: 'BASEMENT_MEAN',
                            data: stcatharinesBasementRates[0],
                            backgroundColor: 'rgb(54, 162, 235, 0.6)',
                            stack: 'Stack 0',
                        },
                        {
                            label: 'BASEMENT_MAX',
                            data: stcatharinesBasementRates[2],
                            backgroundColor: 'rgb(75, 192, 192, 0.6)',
                            stack: 'Stack 0',
                        },
                        {
                            label: 'APARTMENT_MIN',
                            data: stcatharinesApartmentRates[4],
                            backgroundColor: 'rgb(255, 99, 132, 0.6)',
                            stack: 'Stack 1',
                        },
                        {
                            label: 'APARTMENT_MEAN',
                            data: stcatharinesApartmentRates[0],
                            backgroundColor: 'rgb(54, 162, 235, 0.6)',
                            stack: 'Stack 1',
                        },
                        {
                            label: 'APARTMENT_MAX',
                            data: stcatharinesApartmentRates[2],
                            backgroundColor: 'rgb(75, 192, 192, 0.6)',
                            stack: 'Stack 1',
                        },
                    ]
                },
                options: {
                    plugins: {
                        title: {
                            display: true,
                            text: 'St Catharines Region rentals (Basement & Apartments)',
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
                    },
                }
                });
        
                let test5 = lowestContigousSum(calculatePercentageChange(stcatharinesApartmentRates[0]))
                let colors5 = backgroundColor(stcatharinesApartmentRates[0], 'rgb(54, 162, 235, 0.6)', test5)        
                chart = new Chart(document.getElementById('stcatharines-percent').getContext('2d'), {
                    type: 'bar',
                    data: {
                        labels: stcatharinesApartmentRates[1],
                        datasets: [
                            {
                                type:'line',
                                label: 'BASEMENT_LIKELY_PRICE',
                                data: calculatePercentageChange(stcatharinesBasementRates[6]),
                                backgroundColor: 'rgb(255, 159, 64)',
                                borderColor: 'rgb(255, 159, 64)',
                                borderWidth: 4
                            },
                            {
                                type:'line',
                                label: 'APARTMENT_LIKELY_PRICE',
                                data: calculatePercentageChange(stcatharinesApartmentRates[6]),
                                backgroundColor: 'rgb(255, 255, 120)',
                                borderColor: 'rgb(255, 255, 120)',
                                borderWidth: 4
                            },
                            {
                                label: 'BASEMENT_MEAN',
                                data: calculatePercentageChange(stcatharinesBasementRates[0]),
                                backgroundColor: backgroundColor(stcatharinesBasementRates[0], 'rgb(54, 162, 235, 0.6)', lowestContigousSum(calculatePercentageChange(stcatharinesBasementRates[0])), "rgb(252, 0, 0, 0.5)"),
                                stack: 'Stack 0',
                            },
                            {
                                label: 'APARTMENT_MEAN',
                                data: calculatePercentageChange(stcatharinesApartmentRates[0]),
                                backgroundColor: backgroundColor(stcatharinesApartmentRates[0], 'rgb(54, 162, 235, 0.6)', lowestContigousSum(calculatePercentageChange(stcatharinesApartmentRates[0])), "rgb(252, 0, 0, 0.9)"),
                                stack: 'Stack 1',
                            },
                        ]
                    },
                    options: {
                        plugins: {
                            title: {
                                display: true,
                                text: 'St Catharines Region rentals (Basement & Apartments) % CHG',
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
            
    overlay.style.display = "none"
}
prepareCharts()