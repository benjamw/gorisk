
if ( typeof handler === 'undefined' ) {
    var handler = {};
}

handler['404'] = {
    handle: function() {
        handlebarsData = {
            "pageHeaderH1": "404",
            "pageHeaderSmall": "Not Found",
            "errorMessage": "The page requested was not found."
        };
        page.updateHandlebars('header');
        page.updateHandlebars('content', 'error', handlebarsData);
        page.updateHandlebars('footer');
    }
};
