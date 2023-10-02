(function() {
  var template = Handlebars.template, templates = Handlebars.templates = Handlebars.templates || {};
templates['sub_level'] = template({"compiler":[8,">= 4.3.0"],"main":function(container,depth0,helpers,partials,data) {
    var helper, alias1=depth0 != null ? depth0 : (container.nullContext || {}), alias2=container.hooks.helperMissing, alias3="function", alias4=container.escapeExpression, lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return "<link rel=\"stylesheet\" href=\"components/SubLevel/sub_level.css\">\r\n<div class = 'subscription-level'>\r\n    <p style=\"font-weight: 700\">"
    + alias4(((helper = (helper = lookupProperty(helpers,"name") || (depth0 != null ? lookupProperty(depth0,"name") : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"name","hash":{},"data":data,"loc":{"start":{"line":3,"column":32},"end":{"line":3,"column":40}}}) : helper)))
    + "</p>\r\n    <p style=\"font-weight: 400; color: rgba(0, 0, 0, 0.60);\">"
    + alias4(((helper = (helper = lookupProperty(helpers,"price") || (depth0 != null ? lookupProperty(depth0,"price") : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"price","hash":{},"data":data,"loc":{"start":{"line":4,"column":61},"end":{"line":4,"column":70}}}) : helper)))
    + " руб. в месяц</p>\r\n    <p style=\"font-weight: 400; color: rgba(0, 0, 0, 0.60);\">"
    + alias4(((helper = (helper = lookupProperty(helpers,"description") || (depth0 != null ? lookupProperty(depth0,"description") : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"description","hash":{},"data":data,"loc":{"start":{"line":5,"column":61},"end":{"line":5,"column":76}}}) : helper)))
    + "</p>\r\n    <button class='buy-sub'> Подписаться </button>\r\n</div>";
},"useData":true});
})();