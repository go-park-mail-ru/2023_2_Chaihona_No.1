(function() {
  var template = Handlebars.template, templates = Handlebars.templates = Handlebars.templates || {};
templates['target'] = template({"compiler":[8,">= 4.3.0"],"main":function(container,depth0,helpers,partials,data) {
    var helper, alias1=depth0 != null ? depth0 : (container.nullContext || {}), alias2=container.hooks.helperMissing, alias3="function", alias4=container.escapeExpression, lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return "<link rel=\"stylesheet\" type=\"text/css\" href=\"components/Target/target.css\"/>\r\n<div class=\"target\">\r\n  <span style=\"font-weight: 700; position: relative; top: 5%;\">"
    + alias4(((helper = (helper = lookupProperty(helpers,"done") || (depth0 != null ? lookupProperty(depth0,"done") : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"done","hash":{},"data":data,"loc":{"start":{"line":3,"column":63},"end":{"line":3,"column":71}}}) : helper)))
    + " из "
    + alias4(((helper = (helper = lookupProperty(helpers,"target") || (depth0 != null ? lookupProperty(depth0,"target") : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"target","hash":{},"data":data,"loc":{"start":{"line":3,"column":75},"end":{"line":3,"column":85}}}) : helper)))
    + " </span> <span\r\n        style=\"font-weight: 400; position: relative; top: 5%;\">"
    + alias4(((helper = (helper = lookupProperty(helpers,"targettype") || (depth0 != null ? lookupProperty(depth0,"targettype") : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"targettype","hash":{},"data":data,"loc":{"start":{"line":4,"column":63},"end":{"line":4,"column":77}}}) : helper)))
    + "</span>\r\n  <div class=\"progress-bar\">\r\n    <div style=\"width: "
    + alias4(((helper = (helper = lookupProperty(helpers,"doneproc") || (depth0 != null ? lookupProperty(depth0,"doneproc") : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"doneproc","hash":{},"data":data,"loc":{"start":{"line":6,"column":23},"end":{"line":6,"column":35}}}) : helper)))
    + "%;\" class=\"progress-bar-done\"></div>\r\n    <div style=\"width: "
    + alias4(((helper = (helper = lookupProperty(helpers,"leftproc") || (depth0 != null ? lookupProperty(depth0,"leftproc") : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"leftproc","hash":{},"data":data,"loc":{"start":{"line":7,"column":23},"end":{"line":7,"column":35}}}) : helper)))
    + "%;\" class=\"progress-bar-remains\"></div>\r\n  </div>\r\n  <div class=\"target-text\">"
    + alias4(((helper = (helper = lookupProperty(helpers,"targettext") || (depth0 != null ? lookupProperty(depth0,"targettext") : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"targettext","hash":{},"data":data,"loc":{"start":{"line":9,"column":27},"end":{"line":9,"column":41}}}) : helper)))
    + "</div>\r\n  <button class=\"day-deneg-btn\">Пожертвовать</button>\r\n</div>";
},"useData":true});
})();