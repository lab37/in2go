var compare = {                           // Declare compare object
  name: function(a, b) {                  // 用于比较名称
    a = a.replace(/^the /i, '');          // 对于英文的名称，移去开头的the
    b = b.replace(/^the /i, '');          // Remove The from start of parameter

    if (a < b) {                          // If value a is less than value b
      return -1;                          // Return -1
    } else {                              // Otherwise
      return a > b ? 1 : 0;               // If a is greater than b return 1 OR
    }                                     // if they are the same return 0
  },
  duration: function(a, b) {              // Add a method 用于比较时长
    a = a.split(':');                     // Split the time at the colon
    b = b.split(':');                     // Split the time at the colon

    a = Number(a[0]) * 60 + Number(a[1]); // Convert the time to seconds
    b = Number(b[0]) * 60 + Number(b[1]); // Convert the time to seconds

    return a - b;                         // Return a minus b
  },
  date: function(a, b) {                  // Add a method 用于比较日期
    a = new Date(a);                      // 转化为日期对像
    b = new Date(b);                      // New Date object to hold the date

    return a - b;                         // Return a minus b
  },
  iint: function(a, b) {                  // Add a method 用于比较日期
    a = parseInt(a);                      // 转化为日期对像
    b = parseInt(b);                      // New Date object to hold the date

    if (a < b) {                          // If value a is less than value b
      return -1;                          // Return -1
    } else {                              // Otherwise
      return a > b ? 1 : 0;               // If a is greater than b return 1 OR
    }                          // Return a minus b
  },
  ffloat: function(a, b) {                  // Add a method 用于比较日期
    a = parseFloat(a);                      // 转化为日期对像
    b = parseFloat(b);                      // New Date object to hold the date

    if (a < b) {                          // If value a is less than value b
      return -1;                          // Return -1
    } else {                              // Otherwise
      return a > b ? 1 : 0;               // If a is greater than b return 1 OR
    }                          // Return a minus b
  }
};




  var $table = $('.sortable');    //获取表格元素，并转化为jquery对象                 
       
  var $th = $table.find('th');      //获取所有的表头  
    // Store array containing rows
  $th.on('click', function() {//集中绑定表头事件
    var $tbody = $table.find('tbody');  
    var rows = $tbody.find('tr').toArray(); //把每一行对象放入数组        
    var $header = $(this);      //获取当前点击的表头            
    var order = $header.data('sort');       // 获取表头中 data-sort 属性的值
    var column;                             

    // If selected item has ascending or descending class, reverse contents
    if ($header.is('.ascending') || $header.is('.descending')) {  //如果有样式,说明排过序了，没必要再排一遍
      $header.toggleClass('ascending descending');    // 轮换这个目标的样式，
      $tbody.append(rows.reverse());                // 反转当前的排序即可,不需要再计算重排。
    } else {                                                                  
      $header.addClass('ascending');                // 如果没有样式，说明没排过序，需要计算。
      $header.siblings().removeClass('ascending descending'); //获取当前th的同胞，清除他们的排序标记
      if (compare.hasOwnProperty(order)) {  // 如果compare中有指定的适应于这类数据类型的排序方法，则采用。
        column = $th.index(this);         // 得到当前列在rows中的索引
        rows.sort(function(a, b) {               // 指定排序方法
          a = $(a).find('td').eq(column).text(); // 从row中提取用于比较的值
          b = $(b).find('td').eq(column).text(); // 从row中提取用于比较的值
          return compare[order](a, b);           // Call compare method
        });

        $tbody.append(rows);
      } else {
        column = $th.index(this);         // 得到当前列在rows中的索引
        rows.sort(function(a, b) {               // 指定排序方法
          a = $(a).find('td').eq(column).text(); // 从row中提取用于比较的值
          b = $(b).find('td').eq(column).text(); // 从row中提取用于比较的值
          return a>b?1:(a<b?-1:0);           // 比较
        });

        $tbody.append(rows);

      }
    }
  });
