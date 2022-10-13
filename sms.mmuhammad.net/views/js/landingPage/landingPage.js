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

        if(name != "" && email != "email" && message!= ""){
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
        }
    })

    function writeYear(){
        const date = new Date();
        document.getElementById("copyright").innerHTML = `© ${date.getFullYear()} copyright all right reserved`
        }
        // Like $('a'), gets all the <a> elements in the document.
        var aElements = document.getElementsByTagName('a');
        // Create one function object instead of one per <a> element.
        // The calling convention is the same for jQuery as for regular JS.
        function preventDefaultListener(e) { 
            if(e.srcElement.href.includes("#")){
                elem = e.srcElement.href.split("#")[1]
                document.getElementById(elem).scrollIntoView()
                e.preventDefault(); 
            }
        }
        // For each a element,
        for (var i = 0, n = aElements.length; i < n; ++i) {
        // register the listener to be fired on click.
        aElements[i].addEventListener('click', preventDefaultListener);
        }
        writeYear()
});