/*jshint esversion: 6 */
const TYPE_MSG = 0;
const TYPE_GETID = 1;
const TYPE_SETID = 2;
const TYPE_NEWUSER = 3;
const MSGTMPL = {
    "type": 0,
    "version": 0,
    "data": {
        "sender_id": "system",
        "receiver_id": "222.76.251.164:54386",
        "content": "Welcome",
        "time_stamp": "2019-08-28T17:00:42.128094409+08:00",
        "msg_version": 0
    }
};
var sock = null;
var wsuri = "ws://hn.chenjr.top:60012";
var myid = null;
var myname = null;

var lastmsg = null;
var messages = [];
myname = names[Math.round(Math.random() * 100)];
var example1 = new Vue({
    el: '#app',
    data: {
        items: messages,
        myname: myname
    }
});




function handleMessage(e) {

    d = JSON.parse(e.data);
    lastmsg = e.data;
    // console.log("message received: " + e.data);
    switch (d.type) {
        case TYPE_SETID:
            myid = d.message.receiver_id;
            console.log("Set id " + myid);
            break;

        case TYPE_MSG:

            if (d.message.sender_id == myid) {
                d.message.sender_id = example1.myname;
            }
            if (d.message.receiver_id == myid) {
                d.message.receiver_id = example1.myname;
            }
            messages.push(d);
            console.log(`[${d.message.sender_id}]->[${d.message.receiver_id}] : ${d.message.content}`);
            break;
        default:
            break;
    }

}

function connectSocket(wsuri) {
    if (wsuri == null) {
        wsuri = document.getElementById('wsuri').value;
    }

    sock = new WebSocket(wsuri);

    sock.onopen = function() {
        console.log("connected to " + wsuri);
    };

    sock.onclose = function(e) {
        console.log("connection closed (" + e.code + ")");
    };

    sock.onmessage = handleMessage;
}




function NewMessage(content, to) {
    var d = new Date();

    var msg = {
        "type": 0,
        "version": 0,
        "message": {
            "sender_id": myid,
            "receiver_id": to,
            "content": content,
            "time_stamp": d.toISOString(),
            "msg_version": 0
        }
    };
    msg.message.sender_id = myid;
    msg.message.content = content;
    msg.message.receiver_id = to;
    return msg;
}

function send() {
    var content = document.getElementById('message').value;
    var msg = NewMessage(content, myid);
    sock.send(JSON.stringify(msg));
}