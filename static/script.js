//attach jQuery to html
const form = $("#user-input");
const list = $("#conversation_list");
// audio objects to support mp3
var audio = new Audio("/audio/sentAudio.mp3" ) ;
var audio2 = new Audio("/audio/recAudio.mp3");

form.keypress(function(event){
    if(event.keyCode != 13){ // ENTER
        return;
    }
    
    event.preventDefault(); // don't refresh the page.

    const userText = form.val(); // get the text from the input form
    form.val(" "); // wipes the text box.
    
    // before you send request, make sure the user input is valid i.e. not all empty.
    list.append("<figure><img id='avatar-left' src='/img/user-avatar.png' alt='useravatar'><figcaption>User</figcaption></figure><li  class='list-group-item  text-left list-group-item-success' id='leftList'>"+userText + "</li>");

    // GET/POST
    const queryParams = {"user-input" : userText }
    $.get("/chat", queryParams)
    
        .done(function(resp){
            audio2.play();//plays sent message audio
            const newItem = "<figure><img id='avatar-right' src='/img/eliza-avatar.png' alt='eliza-avatar'><figcaption id ='figRight'>Eliza</figcaption></figure><li  class='list-group-item text-right list-group-item-info' id='rightList'>"+ resp + "</li>";
            setTimeout(function(){
              
                list.append(newItem)
                audio.play();//plays recievedmessage audio
                //  need to keep the page scroling to the bottom as the the list size increases
                //  https://stackoverflow.com/questions/47425453/html-css-auto-scroll-page-to-bottom/47425655#47425655
                $("html, body").scrollTop($("body").height());
            }, 2000);//set timeout to give wait to response
        }).fail(function(){
           // display error message and red text box if connection fails
            const newItem = "<figure><img id='avatar-right' src='/img/eliza-avatar.png' alt='eliza-avatar'><figcaption id ='figRight'>Eliza</figcaption></figure><li  class='list-group-item text-right list-group-item-danger' id='rightList'>Sorry im not home right now come back later</li>";
            list.append(newItem);
        });
});