// alert("hello");
document.addEventListener("DOMContentLoaded", async ()=>{
    contact = document.getElementById("contact")

    getStarted = document.getElementById("gsB1").addEventListener("click", ()=>{
        contact.scrollIntoView()
    })

    getStarted = document.getElementById("gsB2").addEventListener("click", ()=>{
        contact.scrollIntoView()
    })

    getStarted = document.getElementById("gsB3").addEventListener("click", ()=>{
        contact.scrollIntoView()
    })

    const submitButton = document.getElementById("ContactSubmit")
    submitButton.addEventListener("click", (evt)=>{
        evt.preventDefault()
        const name = document.getElementById("name").value
        const email = document.getElementById("email").value
        const subject = document.getElementById("subject").value
        const message = document.getElementById("message").value

        if(name.length == 0){
            swal("Oops!", "Please enter your name!", "error")
            return
        } 
        else if(email.length == 0){
            swal("Oops!", "Please enter your email!", "error")
            return
        } 
        else if(message.length == 0){
            swal("Oops!", "Please fill out your message", "error")
            return
        } 

        if(name != "" && email != "" && message!= ""){
            if(validateEmail(email) == false){
                swal("Oops!", "Please enter a proper email address!", "error")
                return
            }
            let payload = {
                name: name,
                email: email,
                subject: subject,
                message: message,
            }
            console.log(payload)
            req = fetch("/api/v1/landing/contact", {
                method: 'POST',
                body: JSON.stringify(payload)
            })
            swal("Information Sent!", "We will reach out to you shortly, thank you !", "success");
        }
    })

    const validateEmail = (email) => {
        const regexExp = /^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$/gi;
        return regexExp.test(
            email
        )
    }
});