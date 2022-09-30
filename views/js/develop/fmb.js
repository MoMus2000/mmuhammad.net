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
            try{
                resp = await fetch(api, {
                    method: "POST",
                    body: JSON.stringify({
                        SenderName: senderName.value,
                        SenderPhone: senderPhone.value,
                        TextMessage: senderMessage.value
                    })
                })
                if(resp.status != 201) swal("Oops!", `Something went wrong, contact admin`, "error")
                else swal("Request Complete!", "Check test phone number for message!", "success");
            }
            catch(err){
                swal("Oops!", `Something went wrong, contact admin`, "error")
            }
        }
        if(senderName.value == "") swal("Oops!", `Please enter the sender name !`, "error");
        else if(senderPhone.value == "") swal("Oops!", `Please enter the test phone number !`, "error");
        else if(senderMessage.value == "") swal("Oops!", `Please enter the message !`, "error");
    })
    submitButton = document.getElementById("submitButton")
    submitButton.addEventListener("click", async ()=>{
        overlay = document.getElementById("overlay")
        overlay.style.display = "block"
        const formData  = new FormData();
        senderPhone = document.getElementById("SenderPhone").value
        senderName = document.getElementById("SenderName").value
        textMessage = document.getElementById("TextMessage").value
        file = document.getElementById("File").files[0]
        formData.append("SenderName", senderName);
        formData.append("SenderPhone", senderPhone);
        formData.append("TextMessage", textMessage);
        formData.append("File", file);
        try{
            const response = await fetch("/fmb/upload", {
                method: 'POST',
                body: formData
            });
            overlay.style.display = "none"
            let status = response.status
            console.log(status)
            if(status == 500){
                swal("Oops!", `Are you sure the excel file is of the expected format?`, "error")
            }
            else if(status != 201)  swal("Oops!", `Something went wrong, contact admin`, "error")
            else swal("Request Complete!", "Messages have been sent!", "success");
        }
        catch(err){
            swal("Oops!", `Something went wrong, contact admin`, "error")
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

