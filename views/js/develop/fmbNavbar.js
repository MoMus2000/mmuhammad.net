let toggle = 0
document.addEventListener("DOMContentLoaded", ()=>{
    hamburgerButton = document.getElementById("FmbNavbarButton").addEventListener("click", ()=>{
        mainNav = document.getElementById("mainNav")
        if(toggle % 2 == 0){
            let a = document.createElement('a')
            a.id = "toggled"
            a.className = 'nav-item nav-link'
            a.href="/fmb/upload"
            a.innerHTML = 'Sign Out';
            mainNav.appendChild(a)
            toggle += 1   
            console.log(a)
            console.log(mainNav)
        }
        else{
            anchorTag = document.getElementById("toggled")
            if(anchorTag != undefined || anchorTag != null){
                anchorTag.remove()
            }
            toggle+=1
        }
    })
})

window.onresize = ()=>{
    anchorTag = document.getElementById("toggled")
    if(anchorTag != undefined || anchorTag != null){
        anchorTag.remove()
    }
}