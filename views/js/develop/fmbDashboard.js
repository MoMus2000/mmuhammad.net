document.addEventListener("DOMContentLoaded", async ()=>{
    async function getBalance(){
        let resp = await fetch("/api/v1/twilio/balanceCheck")
        resp = await resp.json()
        resp = JSON.parse(resp)
        return resp
    }

    // Move this function to the upload page and have it run on every keystroke
    function checkSegment(){
        document.getElementById("segmentCheck").addEventListener('click', ()=>{
            text = document.getElementById("TextMessage")
            console.log(text.value)
            const segmentedMessage = new SegmentedMessage(text.value);
            alert(segmentedMessage.segmentsCount)
        })
    }

    async function getMessageLength(){
        let resp = await fetch("/api/v1/twilio/getMessageLength")
        resp = await resp.json()
        console.log(resp)
        resp = JSON.parse(resp)
        return resp
    }

    function createGauge(value, id, title){
        var data = [
        {
            type: "indicator",
            mode: "gauge+number+delta",
            value: value,
            title: { text: title, font: { size: 24 } },
            gauge: {
            axis: { range: [null, 100], tickwidth: 1, tickcolor: "darkblue" },
            bar: { color: "black" },
            bgcolor: "white",
            borderwidth: 2,
            bordercolor: "gray",
            steps: [
                { range: [0, 5], color: "#FF9999" },
                { range: [6, 15], color: "#FFFF99" },
                { range: [16, 100], color: "#CCFFCC" }
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
    createGauge(balance['balance'], "myDiv", "Twilio Account Balance ($usd)")
    createGauge(msgLength['length'], "myDiv2", "Total Messages Sent Today")
})