function initialize() {
  var labels = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ';
  var labelIndex = 0;
  var mapOptions = {
    center: {lat: -33.8688, lng: 151.2195},
    zoom: 13,
    scrollwheel: false
  };
  var current;
  var map = new google.maps.Map(document.getElementById('map'), mapOptions);

  var input = document.getElementById('pac-input');

  var autocomplete = new google.maps.places.Autocomplete(input);
  autocomplete.bindTo('bounds', map);

  map.controls[google.maps.ControlPosition.TOP_LEFT].push(input);

  var infowindow = new google.maps.InfoWindow();
  var marker = new google.maps.Marker({
    map: map
  });
  google.maps.event.addListener(marker, 'click', function() {
    infowindow.open(map, marker);
  });

  google.maps.event.addListener(autocomplete, 'place_changed', function() {
    infowindow.close();
    var place = autocomplete.getPlace();
    if (!place.geometry) {
      current = undefined;
      return;
    }

    if (place.geometry.viewport) {
      map.fitBounds(place.geometry.viewport);
    } else {
      map.setCenter(place.geometry.location);
      map.setZoom(17);
    }

    // Set the position of the marker using the place ID and location.
    marker.setPlace(/** @type {!google.maps.Place} */ ({
      placeId: place.place_id,
      location: place.geometry.location
    }));
    marker.setVisible(true);
    current = place;
    console.log(place.place_id);
    var photoOptions = {'maxWidth': 175, 'maxHeight': 125};
    var photo = place.photos? '<img src="' + place.photos[0].getUrl(photoOptions) + '" />' : '';
     infowindow.setContent('<div class="preview"><p>' + place.name + '</p>' + photo +
        '<dl>' +
          '<dt>Address:</dt>' + '<dd>' + place.formatted_address +'</dd>' +
          '<dt>Website:</dt>' + '<dd><a href="place.website" target="_blank">' + place.website +'</a></dd>' +
        '</dl>' +
      '</div>');
    infowindow.open(map, marker);
  });

  var list = document.getElementById("list");
  function addToList(place) {
    var photoOptions = {'maxWidth': 175, 'maxHeight': 125};
    var photo = place.photos? '<img src="' + place.photos[0].getUrl(photoOptions) + '" />' : '';
    var li = document.createElement('li');
    li.innerHTML = photo + '<p>' + place.name + '</p>' + '<i>' + place.formatted_address + '</i>';
    list.appendChild(li); 
  }

  var addPlace = document.getElementById('add-place');

  addPlace.addEventListener('click', function(e) {
    e.preventDefault();

    if(current) {
      new google.maps.Marker({
        position: current.geometry.location,
        label: labels[labelIndex++ % labels.length],
        map: map
      });

      addToList(current);
    }
  });
}

// Run the initialize function when the window has finished loading.
google.maps.event.addDomListener(window, 'load', initialize);