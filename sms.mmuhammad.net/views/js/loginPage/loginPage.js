document.addEventListener("DOMContentLoaded", ()=>{
    button = document.getElementById("LoginSubmit")
    button.addEventListener("click", (e)=>{
        e.preventDefault()
        email = document.getElementById("email").value
        password = document.getElementById("password").value

        if(email != "" && password != ""){
            const payload = {
                email: email,
                password: password
            }
            resp = fetch("/login", {
                method: 'POST',
                body: JSON.stringify(payload)
            }).then((response) => 
            console.log(response))
        }
    })
});