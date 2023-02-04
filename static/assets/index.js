function openNav() {
  document.getElementById("mySidenav").style.width = "360px";
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

var firstAlbumDateInput = document.getElementById('first-album-date');
var firstAlbumDateLabel = document.getElementById('first-album-date-label');
function firstAlbumDate() {
  firstAlbumDateLabel.innerHTML = "First album date: " + firstAlbumDateInput.value;
}
firstAlbumDateLabel.innerHTML = "First album date: " + firstAlbumDateInput.value;

var creationDateInput = document.getElementById('creation-date');
var creationDateLabel = document.getElementById('creation-date-label');
function creationDate() {
  creationDateLabel.innerHTML = "Creation date: " + creationDateInput.value;
}
creationDateLabel.innerHTML = "Creation date: " + creationDateInput.value;