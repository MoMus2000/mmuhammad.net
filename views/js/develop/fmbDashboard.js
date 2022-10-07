document.addEventListener("DOMContentLoaded", async ()=>{
    async function getBalance(){
        let resp = await fetch("/api/v1/twilio/balanceCheck")
        resp = await resp.json()
        return resp
    }

    async function getMessageLength(){
        let resp = await fetch("/api/v1/twilio/getMessageLength")
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
        width: 375,
        height: 350,
        margin: { t: 25, r: 25, l: 25, b: 25 },
        font: { color: "darkblue", family: "Arial" }
        };
          
        Plotly.newPlot(id, data, layout);
    }

    const balance = await getBalance()
    const msgLength = await getMessageLength()
    createGauge(balance['balance'], "BalanceChart", "Twilio Account Balance ($usd)", 100)
    createGauge(msgLength['length'], "MessageChart", "Messages Sent Today", 1000)
})