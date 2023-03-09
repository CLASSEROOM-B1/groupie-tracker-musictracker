function initMap() {
  const map = new google.maps.Map(document.getElementById("map"), {
    zoom: 10,
    center: { lat: -33.9, lng: 151.2 },
  });
  
  setMarkers(map);
  }
  
  // Data for the markers consisting of a name, a LatLng and a zIndex for the
  // order in which these markers should display on top of each other.
  const beaches = [
  ["Bondi Beach", -33.890542, 151.274856, 4],
  ["Coogee Beach", -33.923036, 151.259052, 5],
  ["Cronulla Beach", -34.028249, 151.157507, 3],
  ["Manly Beach", -33.80010128657071, 151.28747820854187, 2],
  ["Maroubra Beach", -33.950198, 151.259302, 1],
  ];
  
  function setMarkers(map) {
  // Adds markers to the map.
  // Marker sizes are expressed as a Size of X,Y where the origin of the image
  // (0,0) is located in the top left of the image.
  // Origins, anchor positions and coordinates of the marker increase in the X
  // direction to the right and in the Y direction down.
  const image = {
    url: "https://developers.google.com/maps/documentation/javascript/examples/full/images/beachflag.png",
    // This marker is 20 pixels wide by 32 pixels high.
    size: new google.maps.Size(20, 32),
    // The origin for this image is (0, 0).
    origin: new google.maps.Point(0, 0),
    // The anchor for this image is the base of the flagpole at (0, 32).
    anchor: new google.maps.Point(0, 32),
  };
  // Shapes define the clickable region of the icon. The type defines an HTML
  // <area> element 'poly' which traces out a polygon as a series of X,Y points.
  // The final coordinate closes the poly by connecting to the first coordinate.
  const shape = {
    coords: [1, 1, 1, 20, 18, 20, 18, 1],
    type: "poly",
  };
  
  for (let i = 0; i < beaches.length; i++) {
    const beach = beaches[i];
  
    new google.maps.Marker({
      position: { lat: beach[1], lng: beach[2] },
      map,
      icon: image,
      shape: shape,
      title: beach[0],
      zIndex: beach[3],
    });
  }
  }
  window.initMap = initMap;

  var coords = [];
  var art = window.location.pathname.split('/')[1]
  
  $.post("/map",
      {
          id: art - 1
      },
      function (res) {
          coords = res
      }
  );
  
  ymaps.ready(init);
  function init() {
      var myMap = new ymaps.Map("map", {
          center: [37, 20],
          zoom: 2,
          controls: []
      })
      $.each(coords, function () {
          var dates = ""
          $.each(this.Dates, function () {
              dates += this + "<br>"
          })
          myMap.geoObjects.add(new ymaps.Placemark(this.Coords, {
              balloonContent: "<b>" + this.Place + "</b><br>" + dates
          }))
      })
  }

  var suggest = []
var pg = 1
updateList()

$('#search').on('keyup input paste keydown', function () {
    updateList()
})

$('#saveFilter').on('click', function () {
    updateList()
})

$('#clearFilters').on('click', function () {
    $('#modal input').val('');
    updateList()
})


function getPage(n) {
    pg = n
    updateList()
}

function updateList() {
    $.post("/api", {
        txt: $('#search').val(),
        page: pg,
        // sort: document.forms['filter'].elements['sort'].value,
        bycreation: $('#est_start').val() + '~' + $('#est_end').val(),
        bymembers: $('#mem_start').val() + '~' + $('#mem_end').val(),
        byalbum: $('#alb_start').val() + '~' + $('#alb_end').val(),
        byevent: $('#con_start').val() + '~' + $('#con_end').val(),
        bylocation: $('#loc_patern').val()
    },
        function (r) {

            // Making up paggination
            if (r.Found > 8) {
                $('#pag').empty()
                var l = r.Pages
                var p = r.Page
                var z = 1
                if (p != 1) {
                    z = p - 1
                }
                $('#pag').append('<li class="page-item"><a class="page-link" href="#" onclick=getPage(' + z + ')><span>&laquo;</span></a></li>')
                for (x = 1; x <= l; x++) {
                    cl = 'active'
                    if (x != p) {
                        cl = ''
                    }
                    $('#pag').append('<li class="page-item ' + cl + '"><a class="page-link" href="#" onclick=getPage(' + x + ')>' + x + '</a></li>')
                }
                k = l
                if (p != l) {
                    k = p + 1
                }
                $('#pag').append('<li class="page-item"><a class="page-link" href="#" onclick=getPage(' + k + ')><span>&raquo;</span></a></li>')
            } else {
                $('#pag').empty()
            }

            suggest = r.Suggestions
            $('#found').text(r.Found);
            // Appending cards template
            $('#cards').empty();
            $.each(r.List, function () {
                var tmpl = '<div class="col-sm-12 col-md-6 col-lg-4 col-xl-3 mb-4">' +
                    '<div class="card">' +
                    '<img class="card-img-top" src="' + this.Image + '">' +
                    '<div class="card-footer">' +
                    '<h5 class="card-title mt-3 mb-3">' + this.Name + '</h5>' +
                    '<div class="row">' +
                    '<div class="col-6">' +
                    '<p>Established</p>' +
                    '<p>Members</p>' +
                    '<p>First release</p>' +
                    '</div>' +
                    '<div class="col-6 rightflow">' +
                    '<p>' + this.Establish + '</p>' +
                    '<p>' + this.MembersNum + '</p>' +
                    '<p>' + this.Album + '</p>' +
                    '</div>' +
                    '</div>' +
                    '</div>' +
                    '<div class="card-footer d-flex justify-content-center">' +
                    '<i class="material-icons mr-3 ' + this.Match.name + '">people</i>' +
                    '<i class="material-icons mr-3 ' + this.Match.member + '">person</i>' +
                    '<i class="material-icons mr-3 ' + this.Match.location + '">location_on</i>' +
                    '<i class="material-icons mr-3 ' + this.Match.album + '">library_music</i>' +
                    '<i class="material-icons mr-3 ' + this.Match.establish + '">cake</i>' +
                    '<i class="material-icons ' + this.Match.event + '">event</i>' +
                    '</div>' +
                    '<a href="/' + this.Id + '"></a>' +
                    '</div>' +
                    '</div>';
                $('#cards').append(tmpl)
            })
        }
    )
}
