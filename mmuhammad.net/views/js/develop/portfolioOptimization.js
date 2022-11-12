window.onload = ()=>{
    charts = []
    button = document.getElementById("crunch")
    button.addEventListener("click", ()=>{
    for(i=0; i<charts.length; i++){
        charts[i].destroy()
    }
    third = document.getElementById("third")
    tickers = document.getElementById("tickers")
    amount = document.getElementById("amount")
    time = document.getElementById("timeHorizon")
    timeValue = time.value;
    timeValueText = time.options[time.selectedIndex].text;
    console.log(timeValueText, timeValue, amount.value, tickers.value)
    fetch("/api/v1/optimize/crunch", {
        method: "POST",
        body: JSON.stringify(
                {
                "tickerValue" : tickers.value,
                "amountValue": amount.value,
                "timeHorizon" : timeValue
                }
            )
    }).then((response) => {return response.json()}).then(
        (data) => {
            data = JSON.parse(data)
            console.log(data.returns, data.volatility, data.sharpe)
            charts.push(new Chart(document.getElementById("data"), {
                type: 'bar',
                data: {
                  labels: ["Returns", "Risk", "Sharpe Ratio"],
                  datasets: [
                    {
                      label: "Portfolio Stats",
                      backgroundColor: ["#3e95cd", "#8e5ea2","#3cba9f"],
                      data: [data.returns, data.volatility, data.sharpe]
                    }
                  ]
                },
                options: {
                  legend: { display: false },
                  title: {
                    display: true,
                  },
                  responsive: true,
                  maintainAspectRatio: false,
                }
            }));

            charts.push(new Chart(document.getElementById("data-stock"), {
                type: 'bar',
                data: {
                  labels: ["Returns", "Risk", "Sharpe Ratio"],
                  datasets: [
                    {
                      label: "Portfolio Stats",
                      backgroundColor: ["#3e95cd", "#8e5ea2","#3cba9f"],
                      data: [data.returns, data.volatility, data.sharpe]
                    }
                  ]
                },
                options: {
                  legend: { display: false },
                  title: {
                    display: true,
                  },
                  responsive: true,
                  maintainAspectRatio: false,
                }
            }));

            third.style.display="block"
        }
    )
})
}