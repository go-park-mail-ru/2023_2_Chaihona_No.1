(function() {
  var template = Handlebars.template, templates = Handlebars.templates = Handlebars.templates || {};
templates['comment'] = template({"compiler":[8,">= 4.3.0"],"main":function(container,depth0,helpers,partials,data) {
    var helper, alias1=depth0 != null ? depth0 : (container.nullContext || {}), alias2=container.hooks.helperMissing, alias3="function", alias4=container.escapeExpression, lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return "<link rel=\"stylesheet\" href=\"components/Comment/comment.css\">\r\n<div class=\"comment\">\r\n    <div class=\"comment-ava\"><img src=\"img/"
    + alias4(((helper = (helper = lookupProperty(helpers,"picname") || (depth0 != null ? lookupProperty(depth0,"picname") : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"picname","hash":{},"data":data,"loc":{"start":{"line":3,"column":43},"end":{"line":3,"column":54}}}) : helper)))
    + "\"></div>\r\n    <div class=\"comment-data\">\r\n        <div class = \"comment-nickname\">"
    + alias4(((helper = (helper = lookupProperty(helpers,"nickname") || (depth0 != null ? lookupProperty(depth0,"nickname") : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"nickname","hash":{},"data":data,"loc":{"start":{"line":5,"column":40},"end":{"line":5,"column":52}}}) : helper)))
    + "</div>\r\n        <div class = \"comment-text\">"
    + alias4(((helper = (helper = lookupProperty(helpers,"text") || (depth0 != null ? lookupProperty(depth0,"text") : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"text","hash":{},"data":data,"loc":{"start":{"line":6,"column":36},"end":{"line":6,"column":44}}}) : helper)))
    + "</div>\r\n        <div class = \"comment-date\">"
    + alias4(((helper = (helper = lookupProperty(helpers,"datetime") || (depth0 != null ? lookupProperty(depth0,"datetime") : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"datetime","hash":{},"data":data,"loc":{"start":{"line":7,"column":36},"end":{"line":7,"column":48}}}) : helper)))
    + "</div>\r\n    </div>\r\n</div>";
},"useData":true});
})();