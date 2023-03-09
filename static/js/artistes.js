// Création fonction submit du form des filtres
let form = document.getElementById('formFilter')
function formSubmit(){
    form.submit()
}

// Affichage années Creation Date  + ajout de l'évènement formSubmit

// Récupération Début du filtre Création Date
let startCD = document.getElementById("startCD")
let fromCD = document.getElementById("fromCD")

fromCD.innerHTML = startCD.value 

startCD.addEventListener("change", function(e){
    formSubmit()
})
startCD.addEventListener('input', function func(e){
    fromCD.innerHTML = e.target.value
})

// Récupération fin du filtre CRéation Date
let endCD = document.getElementById("endCD") 
let toCD = document.getElementById("toCD")
toCD.innerHTML = endCD.value
endCD.addEventListener("change", function(e){
    formSubmit()
})
endCD.addEventListener('input', function func(e){
    toCD.innerHTML = e.target.value
})


// Affichage année First Album + ajout de l'évènement formSubmit

// Récupération Début du filtre First Album
let startFA = document.getElementById("startFA") 
let fromFA = document.getElementById("fromFA")
fromFA.innerHTML = startFA.value
startFA.addEventListener("change", function(e){
    formSubmit()
})
startFA.addEventListener('input', function func(e){
    fromFA.innerHTML = e.target.value
})

// Récupération FIN du filtre First Album
let endFA = document.getElementById("endFA") 
let toFA = document.getElementById("toFA")
toFA.innerHTML = endFA.value
endFA.addEventListener("change", function(e){
    formSubmit()
})
endFA.addEventListener('input', function func(e){
    toFA.innerHTML = e.target.value
})

// Ajout évènement checkbox formsSubmit
let checkbox = document.querySelectorAll('.checkbox')
checkbox.forEach(function(checkX){
    // Assignation de l'évènement à la checkbox
    checkX.addEventListener('click', function(e){
        formSubmit()
    })
})

// Submit des filtres avec la pagination
function goToPage(x){
    form.setAttribute('action', `/artistes?page=${x}`)
    form.submit()
}