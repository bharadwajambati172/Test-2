for(var i=0;i<=5;i++){
    (function(index) {
        setTimeout(function() {        
            console.log(index)
        }, i * 1000);
    })(i);
}