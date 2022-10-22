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
            })
            .then((response) => handleRequestResponse(response))
            .catch(error => 
                swal("Oops!", "Please enter valid login credentials !", "error")
                )
        }
    })

    forgotPassword = document.getElementById("forgotPassword")
    forgotPassword.addEventListener("click", (e)=>{
        e.preventDefault()
        swal("Please contact the admin to change your password.")
    })

    function handleRequestResponse(response){  
        if(response.status == 201){
            swal("Success", "You are being redirected to the Sms Terminal", "success")
            .then(()=>{
                window.location = "/sms"
            })
            return
        }
        else{
            swal("Oops!", "Please enter valid login credentials !", "error")
            return
        }
    }
});