document.addEventListener("DOMContentLoaded", async ()=>{
    async function getTotalBalance(){
        let resp = await fetch("/api/v1/sms/totalPrice")
        resp = await resp.json()
        return resp
    }

    async function getTotalMessageLength(){
        let resp = await fetch("/api/v1/sms/totalMsg")
        resp = await resp.json()
        console.log(resp)
        return resp
    }

    async function getTodayPrice(){
        let resp = await fetch("/api/v1/sms/todayPrice")
        resp = await resp.json()
        console.log(resp)
        return resp
    }

    async function getTodayMessageLength(){
        let resp = await fetch("/api/v1/sms/todayMsg")
        resp = await resp.json()
        console.log(resp)
        return resp
    }

    function createGauge(value, id, title, end){
        var data = [
        {
            type: "indicator",
            mode: "gauge+number+delta",
            value: value,
            title: { text: title, font: { size: 24 } },
            gauge: {
            axis: { range: [null, end], tickwidth: 1, tickcolor: "darkblue" },
            bar: { color: "black" },
            bgcolor: "white",
            borderwidth: 2,
            bordercolor: "gray",
            steps: [
                { range: [0, 5], color: "#FF9999" },
                { range: [6, 15], color: "#FFFF99" },
                { range: [16, end], color: "#CCFFCC" }
            ],
            threshold: {
                line: { color: "red", width: 4 },
                thickness: 0.75,
                value: 490
            }
            }
        }
        ];
          
        var layout = {
        width: 325,
        height: 350,
        // autosize : true,
        // margin: { t: 25, r: 25, l: 25, b: 25 },
        // automargin: true,
        font: { color: "darkblue", family: "Arial" }
        };

        var config = {responsive: true};
          
        Plotly.newPlot(id, data, layout, config);
    }

    // const balance = await getBalance()
    // const msgLength = await getMessageLength()

    const balance = 10.0
    const msgLength = 450

    const ops = {
        // color configs
        colorStart: "#6fadcf",
        colorStop: void 0,
        gradientType: 0,
        strokeColor: "#e0e0e0",
        generateGradient: true,
        percentColors: [[0.0, "#ff0000" ], [0.50, "#f9c802"], [1.0, "#a9d70b"] ],
        // customize pointer
        pointer: {
          length: 0.8,
          strokeWidth: 0.035,
          iconScale: 1.0
        },
        // static labels
        staticLabels: {
          font: "10px sans-serif",
          labels: [5, 15, 20, 30],
          fractionDigits: 0
        },
        // static zones
        staticZones: null,
        // render ticks
        renderTicks: {
          divisions: 5,
          divWidth: 1.1,
          divLength: 0.7,
          divColor: "#333333",
          subDivisions: 3,
          subLength: 0.5,
          subWidth: 0.6,
          subColor: "#666666"
        },
        // the span of the gauge arc
        angle: 0.15,
        // line thickness
        lineWidth: 0.44,
        // radius scale
        radiusScale: 1.0,
        // font size
        fontSize: 40,
        // if false, max value increases automatically if value > maxValue
        limitMax: false,
        // if true, the min value of the gauge will be fixed
        limitMin: false,
        // High resolution support
        highDpiSupport: true
    };

    function createSpeedoMeter(gaugeId, textId, ops, value, max, concatText){

        var target = document.getElementById(gaugeId); 
        var gauge = new Gauge(target).setOptions(ops);
        textVal = document.getElementById(textId)
        textVal.innerHTML = concatText+value
        gauge.maxValue = max;
        gauge.setMinValue(0); 
        gauge.set(value);

        gauge.animationSpeed = 32
    }


    const balanceTotal = await getTotalBalance()
    const msgTotal = await getTotalMessageLength()
    const balanceToday = await getTodayPrice()
    const msgToday = await getTodayMessageLength()

    // createSpeedoMeter("Balance", "BalanceText", ops, balance['balance'], 20, "Balance $ ")
    createSpeedoMeter("Balance", "BalanceText", ops, balanceTotal["Data"], 20, "Balance $ ")

    createSpeedoMeter("BalanceToday", "BalanceTodayText", ops, balanceToday["Data"], 20, "Balance $ ")

    // createSpeedoMeter("Message", "MessageText", ops, msgLength['length'], 1200, "Messages Sent :")
    createSpeedoMeter("Message", "MessageText", ops, msgToday["Data"], 1200, "Messages Sent :")
    createSpeedoMeter("TotalMessage", "TotalMessageText", ops, msgTotal["Data"], 1200, "Messages Sent :")

})