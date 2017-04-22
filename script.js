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


$("#demo").jChart({

name: "Attributes",

headers: ["Reliability","Sales","Leadership","Quantitative","Communication","Creative"],

values: [250000,478000,88000,429000,423000],
footers: [100000,200000,300000,400000,500000],
colors: ["#1000ff","#006eff","#00b6ff","#00fff6","#00ff90"]

});




<script language=" JavaScript" >
<!-- 
function LoadOnce() 
{ 
window.location.reload(); 
} 
//-->
</script>
