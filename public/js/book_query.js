function getTarget(e) {                          
  if (!e) {                                      
    e = window.event;                            
  }
  return e.target || e.srcElement;           
}

function itemDone(e) {
  var target;        
  target = getTarget(e);                        
  chapter = target.getAttribute('title');
  var xhr = new XMLHttpRequest();                 
  xhr.onload = function() {                      
    if(xhr.status === 200) {
    document.getElementById('article').innerHTML = xhr.responseText;
    }
  };
  xhr.open('GET', 'htmlturn?dest=' + chapter, true);
  xhr.send(null);
  
    
  if (e.preventDefault) {                        
    e.preventDefault();                         
  } else {                                       
    e.returnValue = false;                      
  }
}

var el = document.getElementById('book_nav');
if (el.addEventListener) {                       
  el.addEventListener('click', function(e) {     
    itemDone(e);                                 
  }, false);                                     
} else {                                        
  el.attachEvent('onclick', function(e) {       
    itemDone(e);                                
  });
}