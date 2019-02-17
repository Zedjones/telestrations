// initialize necessary materialize components
(function($){
    $(function(){
        $('.sidenav').sidenav();
        
        $('.modal').modal({
            dismissible: false
        });
            
        $('.dropdown-trigger').dropdown();
        
        $('select').formSelect();
    });
})(jQuery);