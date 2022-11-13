window.onload = ()=>{
    charts = []
    stocksProportion = []
    stockTicker = []
    colors = []
    futureValues = []
    months = 0
    button = document.getElementById("crunch")
    button.addEventListener("click", ()=>{
    loader = document.getElementById("overlay")
    loader.style.display = "block"
    third = document.getElementById("third")
    third.style.display="none"
    for(i=0; i<charts.length; i++){
        charts[i].destroy()
        stocksProportion = []
        stockTicker = []
        colors = []
        futureValues = []
    }
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
            if(data.stats == 0) {
              alert("Make sure the defined tickers are correct.")
              loader.style.display = "none"
              return
            }
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
                  plugins: {
                    title: {
                        display: true,
                        text: 'Return - Risk - Sharpe Ratio'
                    }
                },
                  responsive: true,
                  maintainAspectRatio: false,
                }
            }));
            for(let i=0; i<data.stock.length; i++){
              stockTicker.push(data.stock[i][1])
              stocksProportion.push(parseFloat(data.stock[i][0]))
            }
            for(let i=0;i<data.stock.length;i++){
              colors.push('#'+Math.floor(Math.random()*16777215).toString(16));
            }
            console.log(stocksProportion, stockTicker)
            charts.push(new Chart(document.getElementById("data-stock"), {
                type: 'doughnut',
                data: {
                  labels: stockTicker,
                  datasets: [
                    {
                      label: "Portfolio Stats",
                      backgroundColor: colors,
                      data: stocksProportion
                    }
                  ]
                },
                options: {
                  legend: { display: false },
                  plugins: {
                    title: {
                        display: true,
                        text: 'Portfolio Allocation'
                    }
                },
                  responsive: true,
                  maintainAspectRatio: false,
                }
            }));
            if(timeValue == "1mo"){
              months = 1
            }
            else if(timeValue == "6mo"){
              months = 6
            }
            else if(timeValue == "1y"){
              months = 12
            }
            else if(timeValue == "5y"){
              months = 60
            }
            days = []
            let futureValue = parseFloat(amount.value);
            futureValues.push(futureValue)
            days.push(0)
            for ( i = 1; i <= months; i++ ) {
              futureValue += (futureValue + 0) * (parseFloat(data.returns));
              futureValues.push(futureValue)
              days.push(i)
            }
            charts.push(new Chart(document.getElementById("returns-plot"), {
              type: 'line',
              data: {
                labels: days,
                datasets: [
                  {
                    label: "Investment growth",
                    backgroundColor: colors,
                    data: futureValues
                  }
                ]
              },
              options: {
                legend: { display: false },
                plugins: {
                  title: {
                      display: true,
                      text: 'Portfolio Growth - Expected Returns'
                  }
              },
                responsive: true,
                maintainAspectRatio: false,
              }
          }));
            loader.style.display = "none"
            third.style.display="block"
        }
    )
})
}