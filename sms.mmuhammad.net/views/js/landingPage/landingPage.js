// Like $('a'), gets all the <a> elements in the document.
var aElements = document.getElementsByTagName('a');
console.log(aElements)
// Create one function object instead of one per <a> element.
// The calling convention is the same for jQuery as for regular JS.
function preventDefaultListener(e) { 
    elem = e.srcElement.href.split("#")[1]
    document.getElementById(elem).scrollIntoView()
    e.preventDefault(); 
}
// For each a element,
for (var i = 0, n = aElements.length; i < n; ++i) {
  // register the listener to be fired on click.
  aElements[i].addEventListener('click', preventDefaultListener);
}