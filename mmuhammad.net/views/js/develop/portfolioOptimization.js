window.onload = ()=>{
    button = document.getElementById("crunch")
    button.addEventListener("click", ()=>{
    third = document.getElementById("third")
    tickers = document.getElementById("tickers")
    amount = document.getElementById("amount")
    time = document.getElementById("timeHorizon")
    timeValue = time.value;
    timeValueText = time.options[time.selectedIndex].text;
    console.log(timeValueText, timeValue, amount.value, tickers.value)
    
    fetch("/api/v1/optimize/crunch", {
        method: "POST",
        body: {
            tickerValue : tickers.value,
            amountValue: amount.value,
            timeHorizon : timeValue
        }
    })


    third.style.display="block"
})
}