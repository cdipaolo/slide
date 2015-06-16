(function () {
    'use strict';
    
    var currentSlide = 1;
    var slides = 0;
    
    var exists = function (url) {
        var http = new XMLHttpRequest();

        http.open('HEAD', url, false);
        http.send();

        return http.status !== 404;
    };
    
    $(document).ready(function () {
        /* start by looking for png files, then jpg */
        var elem = document.getElementsByClassName('slides')[0];

        if (exists('img/1.png')) {
            var url = 'img/1.png';
            var i = 1;
            for (; exists(url); i++) {
                var slide = document.createElement('img');
                var div = document.createElement('div');

                slide.setAttribute('src', url);
                slide.setAttribute('class', 'slide')
                div.setAttribute('class', 'slide-container');

                div.appendChild(slide);

                elem.appendChild(div);

                url = 'img/' + (i+1) + '.png';
                console.log("url: ", url);
            }
            slides = i-1;

        } else if (exists('img/1.jpg')) {
            var url = 'img/1.jpg';
            var i = 1;
            for (; exists(url); i++) {
                var slide = document.createElement('img');
                var div = document.createElement('div');

                slide.setAttribute('src', url);
                slide.setAttribute('class', 'slide')
                div.setAttribute('class', 'slide-container');

                div.appendChild(slide);

                elem.appendChild(div);

                url = 'img/' + (i+1) + '.jpg';
                console.log("url: ", url);
            }
            slides = i-1;

        } else {
            // add error message if nothing is matched
            var err = document.createElement('h1');
            err.appendChild(document.createTextNode('Error loading images!'));

            elem.appendChild(err);
        }
        
        
        var owl = $('.slides')
        
        owl.owlCarousel({
            navigation: false,
            items: 1,
            center: true
        });

        console.log('Owl!', owl);    
        
        var arrowCheck = function (e) {

            e = e || window.event;

            if (e.keyCode == '37') {
                // left arrow
                console.log('Clicked left arrow');
                
                owl.trigger('prev.owl.carousel');
            }
            else if (e.keyCode == '39' || e.keyCode == '32') {
                // right arrow or spacebar clicked
                console.log((e.keyCode == '39') ? 'Clicked right arrow' : 'Clicked spacebar');
                
                owl.trigger('next.owl.carousel');
            }
        }
        
        document.onkeydown = arrowCheck;
    });
}());