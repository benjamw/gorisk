
if ( typeof handler === 'undefined' ) {
    var handler = {};
}

handler.home = {
    handle: function() {
        handlebarsData = {
            "pageHeaderH1": "Home",
            "pageHeaderSmall": "Welcome"
        };
        page.updateHandlebars('header');
        page.updateHandlebars('content', 'home', handlebarsData);
        page.updateHandlebars('footer');
    }
};
