
$(document).ready(function() {
    page.parseHash();
    $(window).on('hashchange', page.parseHash);
});
