document.addEventListener("DOMContentLoaded", ()=>{
    socket = new WebSocket("ws://localhost:3000/api/v1/msg/ws")

    socket.onopen = async ()=>{
        console.log("sucessfully connected to the socket")
        let jsonData = {}
        jsonData["action"] = "username"
        jsonData["username"] = this.value
        socket.send(JSON.stringify(jsonData))

        let recentMessages = await fetch("/api/v1/msg/recent")
        recentMessages = await recentMessages.json()
        output = document.getElementById("output")
        for(let i=0; i<Object.keys(recentMessages).length; i++){
            output.innerHTML += `<strong>${recentMessages[i][0]}</strong>: ${recentMessages[i][1]}`+"<br>"
        }
        
    }

    socket.onclose = ()=>{
        console.log("connection closed")
    }

    socket.onerror = ()=>{
        console.log("there was an error")
    }

    socket.onmessage = msg =>{
        let data = JSON.parse(msg.data)
        if(data.action == "list_users"){
            number = parseInt(data.message)
            numUsers = document.getElementById("online")
            numUsers.innerHTML = `Currently ${number} user(s) on the msg board ..`
        }
        if(data.action == "list_message"){
            chatMessage = data.message
            output = document.getElementById("output")
            output.innerHTML += chatMessage+"<br>"
        }
    }

    username = document.getElementById("username")
    username.addEventListener("change", ()=>{
        let jsonData = {}
        jsonData["action"] = "username"
        jsonData["username"] = this.value
        socket.send(JSON.stringify(jsonData))
    })

    messageBox = document.getElementById("message")
    messageBox.addEventListener("keydown", function(event){
        if(event.code == "Enter"){
            if(!socket){
                return false
            }
            event.preventDefault()
            event.stopPropagation()
            sendMessage()
        }
    })

    sendButton = document.getElementById("sendBtn")
    sendButton.addEventListener("click", sendMessage)

    function sendMessage(){
        let jsonData = {}
        jsonData["action"] = "broadcast"
        jsonData["message"] = document.getElementById("message").value
        jsonData["username"] = username.value
        console.log(jsonData)
        socket.send(JSON.stringify(jsonData))
        document.getElementById("message").value = ""
    }
})