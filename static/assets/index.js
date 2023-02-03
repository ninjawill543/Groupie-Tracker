function openNav() {
  document.getElementById("mySidenav").style.width = "300px";
}
function closeNav() {
  document.getElementById("mySidenav").style.width = "0";
}

function loading() {
  setTimeout(function () {
    document.getElementById("loading").style.display = "block";
    document.getElementById("dots-flow").style.display = "none";
  }, 3000);
};
loading();