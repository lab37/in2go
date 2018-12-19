function getTarget(e) {
  if (!e) {
    e = window.event;
  }
  return e.target || e.srcElement;
}

function checkInput(element) {
  var isNull = false;
  var elReal = element.value.replace(/\s+/g,"");
  if (elReal.length < 1) {            // If username too short
   isNull = true;                             // Clear msg
  }
  return isNull;
};

function itemDone(target) {
  chapter = target.getAttribute('title');
  if (!chapter) {
    return
  }
  $("#article").empty();
  $('#article').load('htmlturn?dest=' + chapter);
  
}

$(".tog").click(function(){
  itemDone(this);
});

function getCustomerName() {
  let rsts = new Array();
  $.ajaxSettings.async = false;
  $.getJSON("getcustomername",function(result){
    $.each(result, function(i, field){
      rsts.push(field.CstmName); 
    });
  });
  $.ajaxSettings.async = true;
  return rsts;
}

function getAllCcId() {
  let rsts = new Array();
  $.ajaxSettings.async = false;
  $.getJSON("getallccid",function(result){
    $.each(result, function(i, field){
      rsts.push(field.CcId); 
    });
  });
  $.ajaxSettings.async = true;
  return rsts;
}

function getProductNS() {
  let rsts = new Array();
  $.ajaxSettings.async = false;
  $.getJSON("getproductns",function(result){
    $.each(result, function(i, field){
      rsts.push(field); 
    });
  });
  $.ajaxSettings.async = true;
  return rsts;
}

$(function() {
  $( "#book_nav" ).accordion({
    collapsible: true,
    heightStyle: "content"
  });
});

function table2xlsx(type,elementID, fn, dl) {
	var elt = document.getElementById(elementID);
	var wb = XLSX.utils.table_to_book(elt, {sheet:"Sheet JS"});
	return dl ?
		XLSX.write(wb, {bookType:type, bookSST:true, type: 'base64'}) :
		XLSX.writeFile(wb, fn || ('test.' + (type || 'xlsx')));
}