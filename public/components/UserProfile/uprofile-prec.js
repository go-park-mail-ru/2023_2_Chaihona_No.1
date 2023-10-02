(function() {
  var template = Handlebars.template, templates = Handlebars.templates = Handlebars.templates || {};
templates['user_profile'] = template({"1":function(container,depth0,helpers,partials,data) {
    var lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return "                <div class=\"sub-element\">\r\n                    <img src=\"img/default-ava.png\">\r\n                    "
    + container.escapeExpression(container.lambda((depth0 != null ? lookupProperty(depth0,"name") : depth0), depth0))
    + "\r\n                </div>\r\n";
},"compiler":[8,">= 4.3.0"],"main":function(container,depth0,helpers,partials,data) {
    var stack1, helper, alias1=depth0 != null ? depth0 : (container.nullContext || {}), alias2=container.hooks.helperMissing, alias3="function", alias4=container.escapeExpression, lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return "<link rel=\"stylesheet\" href=\"components/UserProfile/user_profile.css\">\r\n\r\n<div class=\"user-cover\"></div>\r\n\r\n<div class=\"page-data\">\r\n\r\n    <div class=\"left-column\">\r\n        <div class=\"ava-block\">\r\n            <div class=\"ava\"></div>\r\n            <div class=\"subs-amount\">"
    + alias4(((helper = (helper = lookupProperty(helpers,"donated") || (depth0 != null ? lookupProperty(depth0,"donated") : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"donated","hash":{},"data":data,"loc":{"start":{"line":10,"column":37},"end":{"line":10,"column":48}}}) : helper)))
    + "</div>\r\n            <div class=\"subs-word\">Пожертвовано</div>\r\n        </div>\r\n    </div>\r\n\r\n    <div class=\"main-data\">\r\n        <div class=\"name-and-status\">\r\n            <div class=\"nickname\">"
    + alias4(((helper = (helper = lookupProperty(helpers,"nickname") || (depth0 != null ? lookupProperty(depth0,"nickname") : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"nickname","hash":{},"data":data,"loc":{"start":{"line":17,"column":34},"end":{"line":17,"column":46}}}) : helper)))
    + "</div>\r\n            <div class=\"status\">"
    + alias4(((helper = (helper = lookupProperty(helpers,"status") || (depth0 != null ? lookupProperty(depth0,"status") : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"status","hash":{},"data":data,"loc":{"start":{"line":18,"column":32},"end":{"line":18,"column":42}}}) : helper)))
    + "</div>\r\n        </div>\r\n    </div>\r\n    <div class=\"right-column\">\r\n        <div class=\"subscriptions\">\r\n            Подписки\r\n            <div class = 'target-separator'></div>\r\n"
    + ((stack1 = lookupProperty(helpers,"each").call(alias1,(depth0 != null ? lookupProperty(depth0,"subs") : depth0),{"name":"each","hash":{},"fn":container.program(1, data, 0),"inverse":container.noop,"data":data,"loc":{"start":{"line":25,"column":12},"end":{"line":30,"column":21}}})) != null ? stack1 : "")
    + "        </div>\r\n    </div>\r\n\r\n</div>";
},"useData":true});
})();