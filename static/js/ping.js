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
    document.getElementById("recfrom").readOnly = true
    setInterval(function(){
        if(doPing){
            ping();
            doPing = false;
        }
    },10000);
}

setInterval(function(){
    if( navigator.onLine == false){
        window.alert("you are offline");
    }
},5000);