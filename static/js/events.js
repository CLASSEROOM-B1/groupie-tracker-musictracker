//Sélection de la barre de recherche sur la page html
const form = document.getElementById('formSearch') 

// Redirection vers le groupe choisi
function displayGroup(ID){
    // Ajout d'un form pour la redirection vers l'ID du groupe
    form.setAttribute('action', `/events?oneGroup=${ID}`)
    form.submit()
}