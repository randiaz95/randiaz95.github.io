$(document).ready(function() {
    setInterval ('cursorAnimation()', 800);
});


function cursorAnimation() {
    $('#cursor').animate({
        opacity: 0
    }, 'slow', 'swing').animate({
        opacity: 1
    }, 'slow', 'swing');
}
