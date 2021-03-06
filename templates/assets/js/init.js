// initialize necessary materialize components
(function($){
    $(function(){
        // side bar navigation (start.html)
        $('.sidenav').sidenav();
        
        // modal (1 is from start.html)
        $('.modal').modal();
        $('#modal1').modal({
            dismissible: false
        });
        $('#modalWait').modal({
            dismissible: false
        });
        
        // card tab (final.html)
        $('.tabs').tabs();
        
        // dropdown (settings in start.html)
        $('.dropdown-trigger').dropdown();
        
        // i think this is from the login in form 
        $('select').formSelect();
        
        // text input character count (login.html)
        $('input#input_text, textarea#textarea2').characterCounter();
    });
})(jQuery);