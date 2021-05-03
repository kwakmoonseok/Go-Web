$(function(){
    if (!window.EventSource) {
        alert("No EventSource!")
        return
    }
    var $chatLog = $('#chat-log')
    var $chatMsg = $('#chat-msg')

    var isBlank = function(string) {
        return string == null || string.trim() === "";
    };
    var username;

    while (isBlank(username)) {
        username = prompt("What's your name?")
        if (!isBlank(username)) {
            $('#user-name').html('<b>' + username + '</b>')
        }
    }

    $('#input-form').on('submit', function(e) {
        $.post('/messages', {
            msg: $chatMsg.val(),
            name: username
        })
        $chatMsg.val("")
        $chatMsg.focus()
        return false
    })

    var addMessage = function(data) {
        var text = ""
        if (!isBlank(data.name)) {
            text = '<strong>' + data.name + ':</strong>'
        }
        text += data.msg
        $chatLog.prepend('<div><span>' + text + '</span></div>')

    }
    addMessage({
        msg: 'hello',
        name: 'aaa'
    })

    addMessage({
        msg: 'hello2'
    })

    var es = new EventSource('/stream')
    es.onopen = function(e) {
        $.post('users/', {
            name: username
        })
    }
    es.onmessage = function(e) {
        var msg = JSON.parse(e.data)
        addMessage(msg)
    }
    window.onbeforeunload = function() {
        $.ajax({
            url: "/users?username=" + username,
            type: "DELETE"
        })
        es.close()
    }
})