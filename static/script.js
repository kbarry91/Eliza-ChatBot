const form = $("#user-input");
const list = $("#conversation_list");

form.keypress(function(event){
    if(event.keyCode != 13){ // ENTER
        return;
    }

    event.preventDefault(); // don't refresh the page.

    const userText = form.val(); // get the text from the input form
    form.val(" "); // wipes the text box.
    
    // before you send request, make sure the user input is valid i.e. not all empty.
    list.append("<li class='list-group-item list-group-item-success'>"+"User : " + userText + "</li>");

    // GET/POST
    const queryParams = {"user-input" : userText }
    $.get("/chat", queryParams)
        .done(function(resp){
            const newItem = "<li class='list-group-item list-group-item-info'>"+"ELiza : " + resp + "</li>";
            setTimeout(function(){
                list.append(newItem)
            }, 1000);//set timeout to give wait to response
        }).fail(function(){
            const newItem = "<li class='list-group-item list-group-item-danger' >Sorry I'm not home right now.</li>";
            list.append(newItem);
        });
});