Handlebars.registerPartial("pageHeader", Handlebars.template({"1":function(container,depth0,helpers,partials,data) {
    var helper;

  return " <small>"
    + container.escapeExpression(((helper = (helper = helpers.pageHeaderSmall || (depth0 != null ? depth0.pageHeaderSmall : depth0)) != null ? helper : helpers.helperMissing),(typeof helper === "function" ? helper.call(depth0 != null ? depth0 : {},{"name":"pageHeaderSmall","hash":{},"data":data}) : helper)))
    + "</small>";
},"compiler":[7,">= 4.0.0"],"main":function(container,depth0,helpers,partials,data) {
    var stack1, helper, alias1=depth0 != null ? depth0 : {};

  return "<div class=\"page-header\">\n    <h1>"
    + container.escapeExpression(((helper = (helper = helpers.pageHeaderH1 || (depth0 != null ? depth0.pageHeaderH1 : depth0)) != null ? helper : helpers.helperMissing),(typeof helper === "function" ? helper.call(alias1,{"name":"pageHeaderH1","hash":{},"data":data}) : helper)))
    + ((stack1 = helpers["if"].call(alias1,(depth0 != null ? depth0.pageHeaderSmall : depth0),{"name":"if","hash":{},"fn":container.program(1, data, 0),"inverse":container.noop,"data":data})) != null ? stack1 : "")
    + "</h1>\n</div>\n";
},"useData":true}));