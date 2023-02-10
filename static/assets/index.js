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
  }, 2000);
};
loading();

var firstAlbumDateMinInput = document.getElementById('first-album-date-min');
var firstAlbumDateMinLabel = document.getElementById('first-album-date-min-label');
function firstAlbumDateMin() {
  firstAlbumDateMinLabel.innerHTML = "First album date (min): " + firstAlbumDateMinInput.value;
}
firstAlbumDateMinLabel.innerHTML = "First album date (min): " + firstAlbumDateMinInput.value;
var firstAlbumDateMaxInput = document.getElementById('first-album-date-max');
var firstAlbumDateMaxLabel = document.getElementById('first-album-date-max-label');
function firstAlbumDateMax() {
  firstAlbumDateMaxLabel.innerHTML = "First album date (max): " + firstAlbumDateMaxInput.value;
}
firstAlbumDateMaxLabel.innerHTML = "First album date (max): " + firstAlbumDateMaxInput.value;

var creationDateMinInput = document.getElementById('creation-date-min');
var creationDateMinLabel = document.getElementById('creation-date-min-label');
function creationDateMin() {
  creationDateMinLabel.innerHTML = "Creation date (min): " + creationDateMinInput.value;
}
creationDateMinLabel.innerHTML = "Creation date (min): " + creationDateMinInput.value;

var creationDateMaxInput = document.getElementById('creation-date-max');
var creationDateMaxLabel = document.getElementById('creation-date-max-label');
function creationDateMax() {
  creationDateMaxLabel.innerHTML = "Creation date (max): " + creationDateMaxInput.value;
}
creationDateMaxLabel.innerHTML = "Creation date (max): " + creationDateMaxInput.value;