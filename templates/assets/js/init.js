// initialize necessary materialize components
(function($){
    $(function(){
        $('.sidenav').sidenav();
        
        $('.modal').modal({
            dismissible: false
        });
            
        $('.dropdown-trigger').dropdown();
        
        $('select').formSelect();
        
        $('input#input_text, textarea#textarea2').characterCounter();
    });
})(jQuery);