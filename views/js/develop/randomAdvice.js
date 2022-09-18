
async function fetchRandomAdvice(){
    resp = await fetch("https://api.adviceslip.com/advice")
    resp = await resp.json()
    return resp["slip"]["advice"]
}
window.addEventListener('load', function() {
    FastClick.attach(document.body);
}, false);

window.onload = function(){
    button = document.getElementById("button")
    textBox = document.getElementById("textBox")
    preTag = document.getElementById("advice")
    button.addEventListener('click', async function(){
        resp = await fetchRandomAdvice()
        console.log(resp)
        textBox.innerHTML = `${resp}`
        preTag.click()
    })
}
