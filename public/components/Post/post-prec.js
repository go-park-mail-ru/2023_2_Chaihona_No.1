(function() {
  var template = Handlebars.template, templates = Handlebars.templates = Handlebars.templates || {};
templates['post'] = template({"1":function(container,depth0,helpers,partials,data) {
    var stack1, helper, alias1=depth0 != null ? depth0 : (container.nullContext || {}), alias2=container.hooks.helperMissing, alias3="function", alias4=container.escapeExpression, lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return "        <div class=\"target-separator\"></div>\r\n\r\n        <div class=\"unlocked-post\">\r\n            <div class=\"unlocked-post-topic\">\r\n                "
    + alias4(((helper = (helper = lookupProperty(helpers,"topic") || (depth0 != null ? lookupProperty(depth0,"topic") : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"topic","hash":{},"data":data,"loc":{"start":{"line":13,"column":16},"end":{"line":13,"column":25}}}) : helper)))
    + "\r\n            </div>\r\n            <div class=\"unlocked-post-content\">\r\n                "
    + alias4(((helper = (helper = lookupProperty(helpers,"text") || (depth0 != null ? lookupProperty(depth0,"text") : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"text","hash":{},"data":data,"loc":{"start":{"line":16,"column":16},"end":{"line":16,"column":24}}}) : helper)))
    + "\r\n            </div>\r\n        </div>\r\n\r\n        <div class=\"tagline\">\r\n"
    + ((stack1 = lookupProperty(helpers,"each").call(alias1,(depth0 != null ? lookupProperty(depth0,"tags") : depth0),{"name":"each","hash":{},"fn":container.program(2, data, 0),"inverse":container.noop,"data":data,"loc":{"start":{"line":21,"column":12},"end":{"line":23,"column":21}}})) != null ? stack1 : "")
    + "        </div>\r\n\r\n        <div class=\"target-separator\"></div>\r\n\r\n";
},"2":function(container,depth0,helpers,partials,data) {
    return "                <span class=\"tag\">"
    + container.escapeExpression(container.lambda(depth0, depth0))
    + "</span>\r\n";
},"4":function(container,depth0,helpers,partials,data) {
    var helper, alias1=depth0 != null ? depth0 : (container.nullContext || {}), alias2=container.hooks.helperMissing, alias3="function", alias4=container.escapeExpression, lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return "        <div class=\"locked-post\">\r\n            <div class=\"locked-post-topic\">\r\n                "
    + alias4(((helper = (helper = lookupProperty(helpers,"topic") || (depth0 != null ? lookupProperty(depth0,"topic") : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"topic","hash":{},"data":data,"loc":{"start":{"line":31,"column":16},"end":{"line":31,"column":25}}}) : helper)))
    + "\r\n            </div>\r\n            <div class=\"locked-post-price\">"
    + alias4(((helper = (helper = lookupProperty(helpers,"availability") || (depth0 != null ? lookupProperty(depth0,"availability") : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"availability","hash":{},"data":data,"loc":{"start":{"line":33,"column":43},"end":{"line":33,"column":59}}}) : helper)))
    + "</div>\r\n            <button class=\"buy-post-button\">Купить</button>\r\n        </div>\r\n";
},"6":function(container,depth0,helpers,partials,data) {
    var stack1, lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return "        <div class=\"comments-sections\">\r\n"
    + ((stack1 = lookupProperty(helpers,"each").call(depth0 != null ? depth0 : (container.nullContext || {}),(depth0 != null ? lookupProperty(depth0,"comments") : depth0),{"name":"each","hash":{},"fn":container.program(7, data, 0),"inverse":container.noop,"data":data,"loc":{"start":{"line":46,"column":12},"end":{"line":48,"column":21}}})) != null ? stack1 : "")
    + "        </div>\r\n        <div class=\"comment-input\">\r\n            <div class=\"input-ava\"><img src=\"img/default-ava.png\"></div>\r\n            <div class=\"comment-field\"><input type=\"text\" placeholder=\"Оставьте комментарий...\"></div>\r\n        </div>\r\n";
},"7":function(container,depth0,helpers,partials,data) {
    var stack1, lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return ((stack1 = container.invokePartial(lookupProperty(partials,"comment"),depth0,{"name":"comment","data":data,"indent":"                ","helpers":helpers,"partials":partials,"decorators":container.decorators})) != null ? stack1 : "");
},"compiler":[8,">= 4.3.0"],"main":function(container,depth0,helpers,partials,data) {
    var stack1, helper, alias1=depth0 != null ? depth0 : (container.nullContext || {}), alias2=container.hooks.helperMissing, alias3="function", alias4=container.escapeExpression, lookupProperty = container.lookupProperty || function(parent, propertyName) {
        if (Object.prototype.hasOwnProperty.call(parent, propertyName)) {
          return parent[propertyName];
        }
        return undefined
    };

  return "<link rel=\"stylesheet\" href=\"components/Post/post.css\">\r\n<div class=\"post\">\r\n    <div class=\"post-lock-and-date\">\r\n        <span>"
    + alias4(((helper = (helper = lookupProperty(helpers,"availability") || (depth0 != null ? lookupProperty(depth0,"availability") : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"availability","hash":{},"data":data,"loc":{"start":{"line":4,"column":14},"end":{"line":4,"column":30}}}) : helper)))
    + "</span>\r\n        <span style=\"float: right;\">"
    + alias4(((helper = (helper = lookupProperty(helpers,"datetime") || (depth0 != null ? lookupProperty(depth0,"datetime") : depth0)) != null ? helper : alias2),(typeof helper === alias3 ? helper.call(alias1,{"name":"datetime","hash":{},"data":data,"loc":{"start":{"line":5,"column":36},"end":{"line":5,"column":48}}}) : helper)))
    + "</span>\r\n    </div>\r\n\r\n"
    + ((stack1 = lookupProperty(helpers,"if").call(alias1,(depth0 != null ? lookupProperty(depth0,"unlocked") : depth0),{"name":"if","hash":{},"fn":container.program(1, data, 0),"inverse":container.program(4, data, 0),"data":data,"loc":{"start":{"line":8,"column":4},"end":{"line":36,"column":11}}})) != null ? stack1 : "")
    + "\r\n    <div class=\"like-comment-share\">\r\n        <img src=\"icons/like.svg\" style=\"margin-left: 2%\">\r\n        <img src=\"icons/comment.svg\">\r\n        <img src=\"icons/share.svg\" style=\"float:right; margin-right: 2%;\">\r\n    </div>\r\n\r\n"
    + ((stack1 = lookupProperty(helpers,"if").call(alias1,(depth0 != null ? lookupProperty(depth0,"unlocked") : depth0),{"name":"if","hash":{},"fn":container.program(6, data, 0),"inverse":container.noop,"data":data,"loc":{"start":{"line":44,"column":4},"end":{"line":54,"column":11}}})) != null ? stack1 : "")
    + "</div>";
},"usePartial":true,"useData":true});
})();