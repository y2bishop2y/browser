var browseRoute = require('./browse');

module.exports = function(routes) {
  routes.addRoute('/', handleIndexRoute);
};

function handleIndexRoute(state, events) {

  // Redirect to browse
  events.navigation.navigate({
    path: browseRoute.createUrl('/aghassemi0.mtv.corp.google.com:5167')
  });
}