function ping(){
    let reqBody = {
        "user":document.getElementById("recfrom").value
    }
    fetch("/getmsg",{
        method:"POST",
        headers:{
            "Content-Type":"application/json",
            "Accept":"application/json",
        },
        body:JSON.stringify(reqBody),
        keepalive:true
    })
    .then((response) =>{
        return response.text();   
    })
    .then(function(data){

        let content = JSON.parse(data)
        let msgs =content
        let obj = document.getElementById("recontainer");
        for(let i=0;i<msgs.length;i++){
            obj.innerHTML += '<div id="containerrec"> \
            <msg>\
                <div id="left">\
                    <p id="from">'+msgs[i].from+'</p>\
                </div>\
                <div id="right">\
                    <p id="msg">'+msgs[i].message+'<span id="time">'+msgs[i].time+'</span></p>\
                </div>\
            </msg>\
        </div>';
        }
        obj.scrollTop = obj.scrollHeight;
        doPing = true;
        // if(window.hasFoc)
    })
    .catch((err)=>{
        doPing = true;
    })
}
function sendmsg(){

    let data = {
        "to":document.getElementById("sendto").value,
        "from":document.getElementById("recfrom").value,
        "message":document.getElementById("message").value,
        "time":new Date().toLocaleString()
    };
    if(document.getElementById("recfrom").readOnly == false){
        window.alert("save your username.")
        return
    }
    if(data.to==""){
        window.alert("recipient id should be valid")
    }
    if( data.message==""){
        window.alert("message cannot be empty.");
        return
    }
    if(data.from==data.to){
        window.alert("cannot send message to yourself");
        return
    }
    fetch("/sendmsg",{
        method:"POST",
        headers:{
            "Content-Type":"application/json",
            "Accept":"application/json",
        },
        body:JSON.stringify(data)
    })
    .then((response)=>{
        let obj = document.getElementById("recontainer");
        obj.innerHTML += '<div id="containersend"> \
            <msg>\
                <div id="left">\
                    <p id="from">'+data["to"]+'</p>\
                </div>\
                <div id="right">\
                    <p id="msg">'+data["message"]+'<span id="time">'+data["time"]+'</span></p>\
                </div>\
            </msg>\
        </div>';
        obj.scrollTop = obj.scrollHeight;
    })
    document.getElementById("message").value = ""
}

var doPing = true

function pingTheServer(){
    ping();
    document.getElementById("recfrom").readOnly = true;
    document.getElementById("tickimg").style.visibility = "visible";
    setInterval(function(){
        if(doPing){
            ping();
            getUsersOnline();
            doPing = false;
        }
    },10000);
}

setInterval(function(){
    if( navigator.onLine == false){
        window.alert("you are offline");
    }
},5000);


function textTo(element){
    document.getElementById("sendto").value = element.innerText
}

// get online users

function getUsersOnline(){
    fetch("/usersonline",{
        method:"GET",
        headers:{
            "Accept" : "application/json",
        }
    })
    .then((response)=>{
        return response.text()
    })
    .then((data)=>{
        let users = JSON.parse(data);
        let obj = document.getElementById("onlineUsers");
        obj.innerHTML = '';
        (Object.keys(users)).forEach(element => {
            obj.innerHTML +='\
            <p id="user" onclick="textTo(this)"> \
                    <img src="static/images/online">'+element+'\
            </p> '
        });
    })
}