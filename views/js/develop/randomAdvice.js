
async function fetchRandomAdvice(){
    resp = await fetch("https://api.adviceslip.com/advice")
    resp = await resp.json()
    return resp["slip"]["advice"]
}

async function setToTextBox(){
    resp = await fetchRandomAdvice()
    console.log(resp)
    textBox.innerHTML = `${resp}`
}
window.onload = function(){
    button = document.getElementById("button")
    textBox = document.getElementById("textBox")
    window.setInterval(setToTextBox, 8000)
}
