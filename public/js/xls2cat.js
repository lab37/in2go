
var rABS = false;
var fx;
var files2cat;
var rsts = new Array();
$("#monitor").hide();


function catArr(xArr, yArrs) {

  let maps = new Array();
  for (let i = 1; i < xArr.length; i++) {
    let ct=0;
    for (let j = 0; j < yArrs[0].length; j++) {
      if (yArrs[0][j] == xArr[i]) {
        maps.push(j);
        break;
      }
      ct=ct+1;
    }
    if(ct==yArrs[0].length) {
      alert("在表格--("+xArr[0]+".xlsx)--中未找到对应表中指明的标题为--("+xArr[i]+")--的列！所有合并工作中止，请修改完成后刷新页面重新操作！一定要刷新页面再操作！！！");
    }
  }
  for (let k = 1; k < yArrs.length; k++) {

    let tmp = new Array();
    tmp.push(xArr[0]);
    for (let c = 0; c < maps.length; c++) {
      tmp.push(yArrs[k][maps[c]]);
    }
    rsts.push(tmp);
  }
  $("#" + xArr[0]).css("background-color", "green");
  $("#" + xArr[0]).append("---------合并成功，并入数据：" + (yArrs.length-1) +"条");
}



function xls2arr(file, fxarr) {
  let reader = new FileReader();
  reader.onerror = function(evt) {
    switch(evt.target.error.code) {
      case evt.target.error.NOT_FOUND_ERR:
        alert(file.name+'文件不存在！');
        break;
      case evt.target.error.NOT_READABLE_ERR:
        alert(file.name+'文件不可读！');
        break;
      case evt.target.error.ABORT_ERR:
        break;
      default:
        alert(file.name+'文件出错！');
    };
  }
  reader.onload = function (evt) {
    let data = evt.target.result;
    if (!rABS) data = new Uint8Array(data);
    let workbook = XLSX.read(data, { type: rABS ? 'binary' : 'array' });
    let sheetNames = workbook.SheetNames;
    let sheetName = sheetNames[0];//只取第一个sheet
    let yarr = XLSX.utils.sheet_to_json(workbook.Sheets[sheetName], { header: 1 });

    catArr(fxarr, yarr);


  };
  if (rABS) reader.readAsBinaryString(file); else reader.readAsArrayBuffer(file);
}
document.querySelector('#i2').onchange = function () {
  let files = this.files;
  fx = files[0];
}



document.querySelector('#i1').onchange = function () {
  // let newWB = XLSX.utils.book_new();
  files2cat = this.files;
};



document.querySelector('#i3').onclick = function () {
  let reader = new FileReader();
  reader.onerror = function(evt) {
  switch(evt.target.error.code) {
    case evt.target.error.NOT_FOUND_ERR:
      alert('对应表文件不存在！');
      break;
    case evt.target.error.NOT_READABLE_ERR:
      alert('对应表文件不可读！');
      break;
    case evt.target.error.ABORT_ERR:
      break;
    default:
      alert('读取对应表文件出错！');
  };
}
  reader.onload = function (evt) {
   rsts=[];
    let data = evt.target.result;
    if (!rABS) data = new Uint8Array(data);
    let workbook = XLSX.read(data, { type: rABS ? 'binary' : 'array' });
    let sheetNames = workbook.SheetNames;
    let sheetName = sheetNames[0];//只取第一个sheet
    let fxarrs = XLSX.utils.sheet_to_json(workbook.Sheets[sheetName], { header: 1 });

    $("#result").empty();

    for (let i in files2cat) {
      if (typeof files2cat[i] == "object") {
        $("#result").append("<li id=" + files2cat[i].name.substr(0, files2cat[i].name.indexOf(".")) + " >" + files2cat[i].name + "</li>");
        let ct=0;
        for (let y = 0; y < fxarrs.length; y++) {
          if (fxarrs[y][0] == files2cat[i].name.substr(0, files2cat[i].name.indexOf("."))) {
            xls2arr(files2cat[i], fxarrs[y]);
            break;
          }
          ct=ct+1
        }
        if(ct==fxarrs.length){
          alert("没有在对应表中找到文件--("+files2cat[i].name+")--的对应项！所有合并中止，请修改表格后刷新网页重新操作！ 一定要刷新页面再操作！！！");
          return;
        }
      
      }
    };
    $("#monitor").show();
  }
  if (rABS) reader.readAsBinaryString(fx); else reader.readAsArrayBuffer(fx);
};

document.querySelector('#i4').onclick = function () {
  var workbook = XLSX.utils.book_new();
  var sheet = XLSX.utils.aoa_to_sheet(rsts);
  XLSX.utils.book_append_sheet(workbook, sheet, "汇总");
  XLSX.writeFile(workbook, '汇总.xlsx');

}

var to_json = function to_json(workbook) {
  var result = {};
  workbook.SheetNames.forEach(function (sheetName) {
    var roa = XLSX.utils.sheet_to_json(workbook.Sheets[sheetName], { header: 1 });
    if (roa.length) result[sheetName] = roa;
  });
  return JSON.stringify(result, 2, 2);
};

    // var to_csv = function to_csv(workbook) {
    //   var result = [];
    //   workbook.SheetNames.forEach(function (sheetName) {
    //     var csv = XLSX.utils.sheet_to_csv(workbook.Sheets[sheetName]);
    //     if (csv.length) {
    //       result.push("SHEET: " + sheetName);
    //       result.push("");
    //       result.push(csv);
    //     }
    //   });
    //   return result.join("\n");
    // };

    // var to_fmla = function to_fmla(workbook) {
    //   var result = [];
    //   workbook.SheetNames.forEach(function (sheetName) {
    //     var formulae = XLSX.utils.get_formulae(workbook.Sheets[sheetName]);
    //     if (formulae.length) {
    //       result.push("SHEET: " + sheetName);
    //       result.push("");
    //       result.push(formulae.join("\n"));
    //     }
    //   });
    //   return result.join("\n");
    // };

    // var to_html = function to_html(workbook) {
    //   HTMLOUT.innerHTML = "";
    //   workbook.SheetNames.forEach(function (sheetName) {
    //     var htmlstr = XLSX.write(workbook, { sheet: sheetName, type: 'string', bookType: 'html' });
    //     HTMLOUT.innerHTML += htmlstr;
    //   });
    //   return "";
    // };
