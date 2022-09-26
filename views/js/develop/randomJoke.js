
async function fetchRandomJoke(){
    resp = await fetch("https://icanhazdadjoke.com/",{
        headers: {
            'Accept': 'application/json'
        }
    })
    resp = await resp.json()
    return resp["joke"]
}

async function setToTextBox(){
    resp = await fetchRandomJoke()
    console.log(resp)
    textBox.innerHTML = `${resp}`
}
window.onload = function(){
    button = document.getElementById("button")
    textBox = document.getElementById("textBox")
    window.setInterval(setToTextBox, 8000)
}
