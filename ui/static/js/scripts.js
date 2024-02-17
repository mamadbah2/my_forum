// Enregistre la position de défilement lorsqu'un utilisateur fait défiler la page
window.addEventListener('scroll', function() {
    sessionStorage.setItem('scrollPosition', window.scrollY);
});

// Récupère la position de défilement enregistrée et rétablit la position au chargement de la page
window.addEventListener('load', function() {
    var scrollPosition = sessionStorage.getItem('scrollPosition');
    if (scrollPosition !== null) {
        window.scrollTo(0, scrollPosition);
        sessionStorage.removeItem('scrollPosition'); // Supprime la position de défilement après l'avoir restaurée
    }
});
