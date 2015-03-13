<?php
namespace Ftl\Http;

class Router {
  const TOKEN_PATTERN = '[a-z0-9\-\_]+';

  protected $routes = array();
  protected $request;
  
  public function __construct($request){
    $this->request = $request;
  }

  public function addRoute($route, $callback){
    //Allow multiple routes per callback
    if (is_array($route)){
      foreach ($route as $actual){
        $this->addRoute($actual, $callback);
      }
      return;
    }

    //Turn the route into a valid regex
    $route = preg_replace(
      '#:('.self::TOKEN_PATTERN.')#', 
      '('.self::TOKEN_PATTERN.')', 
      $route
    );
    $this->routes["#^{$route}/?$#"] = $callback;
  }

  public function dispatch(){
    //Find which route to use
    //Execute the fn, pass in the content of placeholders in order
    $matches = array();
    foreach ($this->routes as $regex => $callback){
      if (preg_match($regex, $this->request, $matches) == 0) continue;
      array_shift($matches); //The first match is the full request
      call_user_func_array($callback, $matches);
      return true;
    }
    throw new \Exception("Page not found", 404);
  }
}
