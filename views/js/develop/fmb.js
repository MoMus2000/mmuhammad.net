document.addEventListener("DOMContentLoaded", ()=>{
    statusButton = document.getElementById("statusCheck")
    statusButton.addEventListener("click", async (e)=>{
        e.preventDefault()
        senderName = document.getElementById("SenderName")
        console.log(senderName.value)
        senderPhone = document.getElementById("SenderPhone")
        console.log(senderPhone.value)
        senderMessage = document.getElementById("TextMessage")
        if(senderPhone.value != "" && senderName.value != "" && senderMessage != ""){
            console.log("Sending ...")
            let api = `/api/v1/twilio/statusCheck`
            resp = await fetch(api, {
                method: "POST",
                body: JSON.stringify({
                    SenderName: senderName.value,
                    SenderPhone: senderPhone.value,
                    TextMessage: senderMessage.value
                })
            })
        }
    })
})