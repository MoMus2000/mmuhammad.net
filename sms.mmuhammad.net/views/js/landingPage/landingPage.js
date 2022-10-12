var anchors = document.getElementsByTagName('a');
 for(i=0, len=anchors.length; i<len; i++){
     anchors[i].addEventListener('click', function(e){e.preventDefault();});
 }