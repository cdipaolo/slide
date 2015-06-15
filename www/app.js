(function () {
    'use strict';
    
    var exists = function (url) {
        var http = new XMLHttpRequest();

        http.open('HEAD', url, false);
        http.send();

        return http.status !== 404;
    };
    
    $(document).ready(function () {
        $('.slides').slick({
            
        });
    });
    
    /* start by looking for png files, then jpg */
    var elem = document.getElementsByClassName('slides')[0];
    
    if (exists('img/1.png')) {
        var url = 'img/1.png';
        for (var i = 1; exists(url); i++) {
            var slide = document.createElement('img');
            var div = document.createElement('div');
            
            slide.setAttribute('src', url);
            div.appendChild(slide);
            
            elem.appendChild(div);
            
            url = 'img/' + (i+1) + '.png';
            console.log("url: ", url);
        }
        
    } else if (exists('img/1.jpg')) {
        var url = 'img/1.jpg';
        for (var i = 1; exists(url); i++) {
            var slide = document.createElement('img');
            slide.setAttribute('src', url);
            
            elem.appendChild(slide);
            
            url = 'img/' + (i+1) + '.jpg';
            console.log("url: ", url);
        }
        
    } else {
        // add error message if nothing is matched
        var err = document.createElement('h1');
        err.appendChild(document.createTextNode('Error loading images!'));
        
        elem.appendChild(err);
    }
}());