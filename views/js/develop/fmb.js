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
function validateUpload(input) {
    const fileSize = input.files[0].size / 1024 / 1024; // in MiB
    if (fileSize > 20) {
      alert('File size exceeds 20 MiB');
      clearFileInput("File")
    }
    const fileName = input.files[0].name.split('.').pop()
    console.log(fileName)
    if(fileName != "xlsx" && fileName != "xls"){
        alert("Only excel files can be provided")
        clearFileInput("File")
    }
}

function clearFileInput(id) { 
    var oldInput = document.getElementById(id); 
    var newInput = document.createElement("input"); 

    newInput.type = "file"; 
    newInput.id = oldInput.id; 
    newInput.name = oldInput.name; 
    newInput.className = oldInput.className; 
    newInput.style.cssText = oldInput.style.cssText; 
    newInput.onchange = oldInput.onchange

    oldInput.parentNode.replaceChild(newInput, oldInput); 
}

