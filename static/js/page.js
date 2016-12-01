
var page = page || {

    // these needed for checking and updating the header/footer
    currentHeader : '',
    currentFooter : '',


    goToHash: function(hash) {
        if ('undefined' === typeof hash || 'home' == hash) {
            hash = 'home';
        }
        window.location.hash = '/'+hash;
    },


    parseHash: function() {

        if (!window.location.hash) {
            this.goToHash('home');
            return;
        }

        var hashElements = window.location.hash.split('/');

        if (hashElements.length <= 1 || hashElements[1] === '' ) {
            this.goToHash('home');
            return;
        }

        if (typeof handler[hashElements[1]].handle === "function") {
            handler[hashElements[1]].handle();
        }

    },


    updateHandlebars: function(handlebarsSection, templateFileName, handlebarsData) {

        if ( handlebarsSection === 'header' || handlebarsSection === 'footer' ) {

            // undeclared headers and footers are set to 'default'
            templateFileName = ( typeof templateFileName === 'undefined' ? 'default' : templateFileName );

            // if page.currentHeader is already set to this template, don't update
            if ( handlebarsSection === 'header' && this.currentHeader === templateFileName ) {
                return;
            }
            else {
                this.currentHeader = templateFileName;
            }

            // if page.currentFooter is already set to this template, don't update
            if ( handlebarsSection === 'footer' && this.currentFooter === templateFileName ) {
                return;
            }
            else {
                this.currentFooter = templateFileName;
            }
        }

        handlebarsData = ( typeof handlebarsData === 'undefined' ? {} : handlebarsData );
        $('#handlebars-'+handlebarsSection).html( JST['static/handlebars/templates/'+handlebarsSection+'/'+templateFileName+'.hbs'](handlebarsData) );
    }
};
