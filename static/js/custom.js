/* Rest functions */

/* globals */
var api = "https://localhost:9443"
/* 
  this is only for testing the cors implementation
  credentials should never be used in plain text and 
  not put in any javascript file except you want to make
  them public available 
  If you disabled basic auth you can just leave them blank
*/
var username = "testuser"
var password = "testpass"

/* startup */
$(document).ready(function() {
  $('#show-projects').click( function() { 
    GetProjects() 
  })
});



/* rest calls */
function GetProjects() {
  path = "/projects"
  id = $('#search').val()
  if (id !== "") { path="/project/" + id }
  $.ajax({
    headers: {
      "Authorization": "Basic " + btoa(username + ":" + password),
    },
    url: api + path,
    type: "GET",
    dataType: "json",
    crossDomain: true,
    withCredentials: true,
    success: function(data) {
      $('#showcase-projects').empty()
      $('#showcase-projects-thead').empty()
      $('#showcase-errors').empty()
      if ( typeof data.error !== "undefined" ) {
        $('#showcase-errors').append("<div class='alert alert-danger' role='alert'><strong>"+data.error+"</strong></div>")
        return
      }
      $('#showcase-projects-thead').append("<tr><th>ID</th><th>Name</th></tr>")
      $.each(data.projects, function (key, val) {
        $('#showcase-projects').append("<tr><th>"+val["id"]+"</th><td>"+val["name"]+"</td></tr>")

      })
    }
  });
}
