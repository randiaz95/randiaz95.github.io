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



$(document).ready(function(){
    columnChart();
    
    function columnChart(){
        var item = $('.chart', '.column-chart').find('.item'),
        itemWidth = 100 / item.length;
        item.css('width', itemWidth + '%');
        
        $('.column-chart').find('.item-progress').each(function(){
            var itemProgress = $(this),
            itemProgressHeight = $(this).parent().height() * ($(this).data('percent') / 100);
            itemProgress.css('height', itemProgressHeight);
        });
    };
});