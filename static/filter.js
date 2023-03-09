function show(cb, filterID) {
    var checkBox = document.getElementById(cb);
    
    var input = document.getElementById(filterID);
  
   
    if (checkBox.checked == true){
      input.style.display = "block";
    } else {
      input.style.display = "none";
    }
}