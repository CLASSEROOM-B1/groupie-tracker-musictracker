function initMap(){
    navigator.geolocation.getCurrentPosition(pos => { // Demande de localisation de l'utilisateur
        
        let crd = pos.coords
        // stockage de la lontitude et la latitude
        lat = crd.latitude;
        lng = crd.longitude
        
        console.log("Latitude : " + lat)
        console.log("Longitude : " + lng)
        
        let map = new google.maps.Map(document.getElementById("map"), { // Postion de l'utilisateur sur la carte
            zoom : 7,
            center : gps = {lat, lng},
        });
        
        const marker = new google.maps.Marker({ // Placement d'un marqueur a sa position.
            position: gps,
            map: map,
        });
        
        const geocoder = new google.maps.Geocoder(); // appel de la fonction de Google Maps
        codeAddress(geocoder, map, marker);
        
    })
    
}

function codeAddress(geocoder, resultsMap, marker) {
    
    const address = document.querySelector('select') 
    
    address.addEventListener("change", async function(event){  // recupération des valeurs 
        
        let addressValue = event.target.value
        
        // récupérer les données des locations de l'API avec un await pour éviter toute les requetes en même temps
        let res = await fetch(`/cities?groups=${addressValue}`)
        let cities = await res.json()
        console.log(cities)
        

        // affichage des concerts
        const divCities = document.getElementById("villes_artistes")
        while (divCities.firstChild){
            divCities.removeChild(divCities.firstChild);
        }
        
        for (let i = 0; i < cities.length; i++){
            
            let replaceUnderscore = cities[i].replaceAll("_", " ")
            let replace = replaceUnderscore.replace("-", " ")
            const capitalize = replace.replace(/(^\w{1})|(\s+\w{1})/g, letter => letter.toUpperCase());
            let city = document.createElement("p");
            city.textContent = capitalize
            divCities.append(city)
        }
        console.log(divCities)
        
        for (let i = 0; i < cities.length; i ++) {
            console.log(cities.length)
            
            geocoder.geocode({ address: cities[i]}, (results, status) => { // utilisation de geocoding avec les valeur récupérer (cities)
                
                if (status === "OK") {
                    resultsMap.setCenter(results[0].geometry.location); // placement de la position de la requete demandée
                    
                    marker.setPosition(results[0].geometry.location) // placement du marqueur non fonctionnel
                    
                } else {
                    alert("Error: " + status);
                }
            });
        }
    });
}