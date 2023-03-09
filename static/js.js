document.getElementById('sch').addEventListener('input', function () {
   
    var inp = document.getElementById('sch').value.split(' // ');
    
    if (inp[1] == "Artist") {
      document.getElementById('sType').selectedIndex = 0;
      document.getElementById('sch').value = inp[0];
    };
    if (inp[1] == "Member") {
        document.getElementById('sType').selectedIndex = 1;
        document.getElementById('sch').value = inp[0];
    };
    if (inp[1] == "Location") {
        document.getElementById('sType').selectedIndex = 2;
        document.getElementById('sch').value = inp[0];
    };
    if (inp[1] == "First Album") {
        document.getElementById('sType').selectedIndex = 3;
        document.getElementById('sch').value = inp[0];
    };
    if (inp[1] == "Creation Date") {
        document.getElementById('sType').selectedIndex = 4;
        document.getElementById('sch').value = inp[0];
    };
  });